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
	// alertaRepo  repositorios.AlertaRepositorio // Futuro: para persistir o alerta
}

func NovoAnaliseServico(db *gorm.DB, regRepo repositorios.RegistroHumorRepositorio, userRepo repositorios.UsuarioRepositorio) AnaliseServico {
	return &analiseServico{
		db:           db,
		registroRepo: regRepo,
		usuarioRepo:  userRepo,
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
		analise.StatusAtual = s.calcularStatus(analise.MediaSono, analise.MediaHumor, analise.MediaStress, analise.MediaEnergia)
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
	status := s.calcularStatus(mediaSono, mediaHumor, mediaStress, mediaEnergia) // Simplificado para exemplo

	if status == StatusPreocupante {
		// TODO: PERSISTE O ALERTA
		// s.alertaRepo.Criar(dominio.Alerta{PacienteID: pacienteID, Tipo: status, Mensagem: "Padrão preocupante detectado"})

		// TODO: ENVIA EMAIL/NOTIFICAÇÃO
	}
	log.Printf(
		"Monitoramento realizado as: %v\nPaciente ID: %d\nDados:\n mediaHumor: %.2f, mediaStress: %.2f, mediaSono: %.2f, mediaEnergia: %.2f\nStatus: %s",
		time.Now(), pacienteID, mediaHumor, mediaStress, mediaSono, mediaEnergia, status)
	return nil
}

func (s *analiseServico) calcularStatus(sono, humor, stress, energia float64) string {
	if humor < 2.5 || stress > 8.0 || (sono < 4.0 || sono > 11.0) || energia < 2.5 {
		return StatusPreocupante
	}
	if humor < 3.5 || stress > 6.0 || (sono < 5.0 || sono > 10.0) || energia < 4.0 {
		return StatusAtencao
	}
	return StatusRegular
}
