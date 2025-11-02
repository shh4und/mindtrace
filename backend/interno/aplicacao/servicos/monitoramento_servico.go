package servicos

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

const (
	padraoPreocupante = "PREOCUPANTE"
	padraoAtencao     = "ATENCAO"
	padraoRegular     = "REGULAR"
)

// MonitoramentoServico define os metodos para gerar monitoramentos
type MonitoramentoServico interface {
	RealizarMonitoramentoPaciente(userID, pacID uint, numLimiteRegistros int) (*dtos.MonitoramentoPacienteDTOOut, error)
}

// monitoramentoServico implementa a interface MonitoramentoServico
type monitoramentoServico struct {
	db                       *gorm.DB
	registroHumorRepositorio repositorios.RegistroHumorRepositorio
	usuarioRepositorio       repositorios.UsuarioRepositorio
}

// NovoMonitoramentoServico cria uma nova instancia de MonitoramentoServico
func NovoMonitoramentoServico(db *gorm.DB, registroHumorRepo repositorios.RegistroHumorRepositorio, usuarioRepo repositorios.UsuarioRepositorio) MonitoramentoServico {
	return &monitoramentoServico{
		db:                       db,
		registroHumorRepositorio: registroHumorRepo,
		usuarioRepositorio:       usuarioRepo,
	}
}
func verificarPadrao(monitorimentoRealizado *dtos.MonitoramentoPacienteDTOOut, margemPercentual float64) string {
	mediaSonoReferencia := float64(6)
	mediaHumorReferencia := float64(3)
	mediaStressReferencia := float64(5.5)
	mediaEnergiaReferencia := float64(5.5)

	minimaSonoReferencia := mediaSonoReferencia - margemPercentual*mediaSonoReferencia
	minimaHumorReferencia := mediaHumorReferencia - margemPercentual*mediaHumorReferencia
	maximaStressReferencia := mediaStressReferencia + margemPercentual*mediaStressReferencia
	minimaEnergiaReferencia := mediaEnergiaReferencia - margemPercentual*mediaEnergiaReferencia

	// PREOCUPANTE: valores criticos (sono/humor/energia baixos OU stress alto)
	if (monitorimentoRealizado.MediaSono < minimaSonoReferencia) ||
		(monitorimentoRealizado.MediaHumor < minimaHumorReferencia) ||
		(monitorimentoRealizado.MediaStress > maximaStressReferencia) ||
		(monitorimentoRealizado.MediaEnergia < minimaEnergiaReferencia) {
		return padraoPreocupante
	}

	// ATENCAO: valores abaixo da referencia mas acima do preocupante (OU stress entre referencia e maxima)
	if (monitorimentoRealizado.MediaSono >= minimaSonoReferencia && monitorimentoRealizado.MediaSono < mediaSonoReferencia) ||
		(monitorimentoRealizado.MediaHumor >= minimaHumorReferencia && monitorimentoRealizado.MediaHumor < mediaHumorReferencia) ||
		(monitorimentoRealizado.MediaStress > mediaStressReferencia && monitorimentoRealizado.MediaStress <= maximaStressReferencia) ||
		(monitorimentoRealizado.MediaEnergia >= minimaEnergiaReferencia && monitorimentoRealizado.MediaEnergia < mediaEnergiaReferencia) {
		return padraoAtencao
	}

	// REGULAR: todos os indicadores em niveis saudaveis
	return padraoRegular

}

// GerarMonitoramentoPaciente gera um relatorio para o paciente autenticado
func (ms *monitoramentoServico) RealizarMonitoramentoPaciente(userID, pacID uint, numLimiteRegistros int) (*dtos.MonitoramentoPacienteDTOOut, error) {
	// Validar periodo
	if numLimiteRegistros < 2 {
		return nil, errors.New("periodo de filtro deve ser maior que 1")
	}
	if numLimiteRegistros > 14 {
		return nil, errors.New("periodo de filtro nao pode exceder 14 dias")
	}

	// Validar pacID
	if pacID == 0 {
		return nil, errors.New("id do paciente invalido")
	}

	monitoramentoPacienteFeito := &dtos.MonitoramentoPacienteDTOOut{
		DadosMonitoramento: make([]dtos.DadosMonitoramentoDTOOut, numLimiteRegistros),
	}

	profissional, err := ms.usuarioRepositorio.BuscarProfissionalPorUsuarioID(ms.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("paciente nao encontrado")
		}
		return nil, err
	}

	now := time.Now()
	dataAtual := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, now.Location())

	registrosHumor, err := ms.registroHumorRepositorio.BuscarPorNUltimosRegistros(pacID, numLimiteRegistros)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return monitoramentoPacienteFeito, nil
		}
		return nil, err
	}

	totalSono := 0
	totalEnergia := 0
	totalStress := 0
	totalHumor := 0

	for _, registro := range registrosHumor {

		totalSono += int(registro.HorasSono)
		totalEnergia += int(registro.NivelEnergia)
		totalStress += int(registro.NivelStress)
		totalHumor += int(registro.NivelHumor)

		dadosMonitoramento := dtos.DadosMonitoramentoDTOOut{
			Data:         registro.DataHoraRegistro,
			NivelHumor:   registro.NivelHumor,
			HorasSono:    registro.HorasSono,
			NivelEnergia: registro.NivelEnergia,
			NivelStress:  registro.NivelStress,
			Observacoes:  registro.Observacoes,
			AutoCuidado:  registro.AutoCuidado,
		}
		monitoramentoPacienteFeito.DadosMonitoramento = append(monitoramentoPacienteFeito.DadosMonitoramento, dadosMonitoramento)
	}
	monitoramentoPacienteFeito.MediaSono = float64(totalSono) / float64(numLimiteRegistros)
	monitoramentoPacienteFeito.MediaEnergia = float64(totalEnergia) / float64(numLimiteRegistros)
	monitoramentoPacienteFeito.MediaStress = float64(totalStress) / float64(numLimiteRegistros)
	monitoramentoPacienteFeito.MediaHumor = float64(totalHumor) / float64(numLimiteRegistros)
	monitoramentoPacienteFeito.TipoAlerta = string(verificarPadrao(monitoramentoPacienteFeito, 0.2))
	monitoramentoPacienteFeito.Data = dataAtual
	monitoramentoPacienteFeito.PacienteID = pacID
	monitoramentoPacienteFeito.ProfissionalID = profissional.ID
	return monitoramentoPacienteFeito, err
}
