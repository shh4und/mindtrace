package servicos

import (
	"errors"
	"log"
	"math"
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

const (
	PesoSono    = 0.4
	PesoHumor   = 0.2
	PesoStress  = 0.2
	PesoEnergia = 0.2
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
	// 1. Normalizar Humor (Escala 1-5, Maior é melhor)
	// (Valor - 1) / (5 - 1)
	normHumor := (humor - 1.0) / 4.0
	if normHumor < 0 {
		normHumor = 0
	}

	// 2. Normalizar Energia (Escala 1-10, Maior é melhor)
	// (Valor - 1) / (10 - 1)
	normEnergia := (energia - 1.0) / 9.0

	// 3. Normalizar Stress (Escala 1-10, MENOR é melhor - Inverso)
	// 1 - ((Valor - 1) / 9)
	normStress := 1.0 - ((stress - 1.0) / 9.0)

	// 4. Normalizar Sono (Ideal ~8h. Distância do ideal).
	// Consideramos 8h o ideal. Se afastar mais que 4h (ou seja <4 ou >12), zera.
	distancia := math.Abs(sono - 8.0)
	normSono := 1.0 - (distancia / 4.0) // Penalidade de 0.25 por hora de desvio
	if normSono < 0 {
		normSono = 0
	}

	// Calculo do Índice Ponderado
	ibg := (normHumor * PesoHumor) +
		(normStress * PesoStress) +
		(normSono * PesoSono) +
		(normEnergia * PesoEnergia)

	// Definição de Status baseada no Índice (0 a 1)
	return ibg
}
