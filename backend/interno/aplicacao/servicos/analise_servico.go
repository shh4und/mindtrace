package servicos

import (
	"errors"
	"log"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

const (
	StatusPreocupante = "PREOCUPANTE"
	StatusAtencao     = "ATENCAO"
	StatusRegular     = "REGULAR"
)

type AnaliseServico interface {
	// GerarAnaliseHistorica: Para o frontend desenhar gráficos (substitui GerarRelatorio)
	GerarAnaliseHistorica(usuarioID, pacienteID uint, tipoUsuario string, dias int) (*dtos.AnalisePacienteDTOOut, error)

	// ExecutarMonitoramento: Chamado automaticamente após novos registros ou via cron job
	ExecutarMonitoramento(pacienteID uint) error
}

type analiseServico struct {
	db           *gorm.DB
	registroRepo repositorios.RegistroHumorRepositorio
	usuarioRepo  repositorios.UsuarioRepositorio
	emailServico EmailServico
	// alertaRepo  repositorios.AlertaRepositorio // Futuro: para persistir o alerta
}

func NovoAnaliseServico(db *gorm.DB, regRepo repositorios.RegistroHumorRepositorio, userRepo repositorios.UsuarioRepositorio, es EmailServico) AnaliseServico {
	return &analiseServico{
		db:           db,
		registroRepo: regRepo,
		usuarioRepo:  userRepo,
		emailServico: es,
	}
}

func (s *analiseServico) GerarAnaliseHistorica(usuarioID, pacienteID uint, tipoUsuario string, dias int) (*dtos.AnalisePacienteDTOOut, error) {
	if dias <= 0 || dias > 90 {
		return nil, errors.New("periodo invalido")
	}

	now := time.Now()
	dataInicio := now.AddDate(0, 0, -dias)

	if dominio.StringParaTipoUsuario(tipoUsuario) == dominio.TipoUsuarioPaciente && pacienteID == 0 {
		pacienteInfo, err := s.usuarioRepo.BuscarPacientePorUsuarioID(s.db, usuarioID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, dominio.ErrUsuarioNaoEncontrado
			}
			return nil, err
		}

		pacienteID = pacienteInfo.ID
	}

	registros, err := s.registroRepo.BuscarPorPacienteEPeriodo(pacienteID, dataInicio, now)
	if err != nil {
		return nil, err
	}

	analise := &dtos.AnalisePacienteDTOOut{
		GraficoSono:    make([]dtos.PontoDeDadosDTOOut, 0),
		GraficoEnergia: make([]dtos.PontoDeDadosDTOOut, 0),
		GraficoStress:  make([]dtos.PontoDeDadosDTOOut, 0),
		StatusAtual:    StatusRegular, // Default
	}

	var somaSono, somaEnergia, somaStress, somaHumor int

	for _, reg := range registros {
		// Popula gráficos
		analise.GraficoSono = append(analise.GraficoSono, dtos.PontoDeDadosDTOOut{Data: reg.DataHoraRegistro, Valor: reg.HorasSono, Humor: reg.NivelHumor})
		analise.GraficoEnergia = append(analise.GraficoEnergia, dtos.PontoDeDadosDTOOut{Data: reg.DataHoraRegistro, Valor: reg.NivelEnergia, Humor: reg.NivelHumor})
		analise.GraficoStress = append(analise.GraficoStress, dtos.PontoDeDadosDTOOut{Data: reg.DataHoraRegistro, Valor: reg.NivelStress, Humor: reg.NivelHumor})

		// Acumula para médias
		somaSono += int(reg.HorasSono)
		somaEnergia += int(reg.NivelEnergia)
		somaStress += int(reg.NivelStress)
		somaHumor += int(reg.NivelHumor)
	}

	if len(registros) > 0 {
		count := float64(len(registros))
		analise.MediaSono = float64(somaSono) / count
		analise.MediaEnergia = float64(somaEnergia) / count
		analise.MediaStress = float64(somaStress) / count
		analise.MediaHumor = float64(somaHumor) / count

		// Recalcula o status baseado nos dados carregados
		analise.ValorIBG = s.calcularIBG(analise.MediaSono, analise.MediaHumor, analise.MediaStress, analise.MediaEnergia)
		analise.StatusAtual = s.calcularStatus(analise.ValorIBG)
	}

	return analise, nil
}

// ExecutarMonitoramento é o método "Trigger"
func (s *analiseServico) ExecutarMonitoramento(pacienteID uint) error {
	// 1. Busca os últimos X registros (ex: 7 dias ou 5 registros)
	registros, err := s.registroRepo.BuscarPorNUltimosRegistros(pacienteID, 5)
	if err != nil {
		return err
	}
	if len(registros) == 0 {
		return nil
	}

	// 2. Calcula médias rápidas
	var somaHumor, somaStress, somaSono, somaEnergia int
	for _, r := range registros {
		somaHumor += int(r.NivelHumor)
		somaStress += int(r.NivelStress)
		somaSono += int(r.HorasSono)
		somaEnergia += int(r.NivelEnergia)
	}
	mediaHumor := float64(somaHumor) / float64(len(registros))
	mediaStress := float64(somaStress) / float64(len(registros))
	mediaSono := float64(somaSono) / float64(len(registros))
	mediaEnergia := float64(somaEnergia) / float64(len(registros))

	// 3. Verifica Padrão
	ibg := s.calcularIBG(mediaSono, mediaHumor, mediaStress, mediaEnergia) // Simplificado para exemplo
	status := s.calcularStatus(ibg)
	if status == StatusPreocupante {
		// Busca paciente e seus profissionais para notificar
		var paciente *dominio.Paciente
		paciente, err := s.usuarioRepo.BuscarProfissionaisDoPaciente(s.db, pacienteID)
		if err == nil && paciente != nil {
			for _, prof := range paciente.Profissionais {
				if prof.Usuario.Email != "" {
					// Envia email de alerta (assincrono/goroutine opcional)
					go func(emailProf, nomeProf, nomePac, st string) {
						defer func() {
							if r := recover(); r != nil {
								log.Printf("ERRO CRITICO: Panic recuperado ao enviar email de alerta: %v", r)
							}
						}()
						s.emailServico.EnviarEmailAlertaMonitoramento(emailProf, nomeProf, nomePac, st)
					}(prof.Usuario.Email, prof.Usuario.Nome, paciente.Usuario.Nome, status)
				}
			}
		}
	}
	log.Printf(
		"Monitoramento realizado as: %v\nPaciente ID: %d\nStatus Calculado: %s",
		time.Now(), pacienteID, status)
	return nil
}

func (s *analiseServico) calcularStatus(ibg float64) string {
	switch {
	case ibg >= 0.70:
		return StatusRegular // Verde (> 70%)
	case ibg >= 0.40:
		return StatusAtencao // Amarelo (40% - 69%)
	default:
		return StatusPreocupante // Vermelho (< 40%)
	}
}

func (s *analiseServico) calcularIBG(sono, humor, stress, energia float64) float64 {
	// Delega ao metodo do dominio para manter logica centralizada
	rh := &dominio.RegistroHumor{
		NivelHumor:   int16(humor),
		HorasSono:    int16(sono),
		NivelStress:  int16(stress),
		NivelEnergia: int16(energia),
	}
	return rh.CalcularIBG()
}
