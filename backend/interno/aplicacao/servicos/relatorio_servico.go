package servicos

import (
	"errors"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

// PontoDeDadosDTO representa um ponto de dados para graficos
type PontoDeDadosDTO struct {
	Data     time.Time `json:"data"`
	Valor    int16     `json:"valor"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

// RelatorioPacienteDTO representa o relatorio de um paciente
type RelatorioPacienteDTO struct {
	GraficoSono    []PontoDeDadosDTO `json:"grafico_sono"`
	GraficoEnergia []PontoDeDadosDTO `json:"grafico_energia"`
	GraficoStress  []PontoDeDadosDTO `json:"grafico_stress"`
	MediaSono      float64           `json:"media_sono"`
	MediaEnergia   float64           `json:"media_energia"`
	MediaStress    float64           `json:"media_stress"`
}

// RelatorioServico define os metodos para gerar relatorios
type RelatorioServico interface {
	GerarRelatorioPaciente(userID uint, filtroPeriodo int64) (*RelatorioPacienteDTO, error)
	GerarRelatorioPacienteDoProfissional(pacienteID uint, filtroPeriodo int64) (*RelatorioPacienteDTO, error)
}

// relatorioServico implementa a interface RelatorioServico
type relatorioServico struct {
	db                       *gorm.DB
	registroHumorRepositorio repositorios.RegistroHumorRepositorio
	usuarioRepositorio       repositorios.UsuarioRepositorio
}

// NovoRelatorioServico cria uma nova instancia de RelatorioServico
func NovoRelatorioServico(db *gorm.DB, registroHumorRepo repositorios.RegistroHumorRepositorio, usuarioRepo repositorios.UsuarioRepositorio) RelatorioServico {
	return &relatorioServico{
		db:                       db,
		registroHumorRepositorio: registroHumorRepo,
		usuarioRepositorio:       usuarioRepo,
	}
}

// GerarRelatorioPaciente gera um relatorio para o paciente autenticado
func (rs *relatorioServico) GerarRelatorioPaciente(userID uint, filtroPeriodo int64) (*RelatorioPacienteDTO, error) {
	relatorioPacienteFeito := &RelatorioPacienteDTO{
		GraficoSono:    make([]PontoDeDadosDTO, 0),
		GraficoEnergia: make([]PontoDeDadosDTO, 0),
		GraficoStress:  make([]PontoDeDadosDTO, 0),
	}

	paciente, err := rs.usuarioRepositorio.BuscarPacientePorUsuarioID(rs.db, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return relatorioPacienteFeito, nil
		}
		return nil, err
	}

	now := time.Now()
	dataFim := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	dataInicio := dataFim.AddDate(0, 0, -int(filtroPeriodo))

	registrosHumor, err := rs.registroHumorRepositorio.BuscarPorPacienteEPeriodo(paciente.ID, dataInicio, dataFim)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return relatorioPacienteFeito, nil
		}
		return nil, err
	}

	totalSono := 0
	totalEnergia := 0
	totalStress := 0

	for _, registro := range registrosHumor {

		totalSono += int(registro.HorasSono)
		totalEnergia += int(registro.NivelEnergia)
		totalStress += int(registro.NivelStress)

		pontoDeDadosSono := PontoDeDadosDTO{
			Data:     registro.DataHoraRegistro,
			Valor:    registro.HorasSono,
			Humor:    registro.NivelHumor,
			Anotacao: registro.Observacoes,
		}
		relatorioPacienteFeito.GraficoSono = append(relatorioPacienteFeito.GraficoSono, pontoDeDadosSono)

		pontoDeDadosEnergia := PontoDeDadosDTO{
			Data:     registro.DataHoraRegistro,
			Valor:    registro.NivelEnergia,
			Humor:    registro.NivelHumor,
			Anotacao: registro.Observacoes,
		}
		relatorioPacienteFeito.GraficoEnergia = append(relatorioPacienteFeito.GraficoEnergia, pontoDeDadosEnergia)

		pontoDeDadosStress := PontoDeDadosDTO{
			Data:     registro.DataHoraRegistro,
			Valor:    registro.NivelStress,
			Humor:    registro.NivelHumor,
			Anotacao: registro.Observacoes,
		}
		relatorioPacienteFeito.GraficoStress = append(relatorioPacienteFeito.GraficoStress, pontoDeDadosStress)

	}
	if len(registrosHumor) > 0 {
		relatorioPacienteFeito.MediaSono = float64(totalSono) / float64(len(registrosHumor))
		relatorioPacienteFeito.MediaEnergia = float64(totalEnergia) / float64(len(registrosHumor))
		relatorioPacienteFeito.MediaStress = float64(totalStress) / float64(len(registrosHumor))

	}

	return relatorioPacienteFeito, nil
}

// GerarRelatorioPacienteDoProfissional gera um relatorio para um paciente especifico pelo profissional
func (rs *relatorioServico) GerarRelatorioPacienteDoProfissional(pacienteID uint, filtroPeriodo int64) (*RelatorioPacienteDTO, error) {
	relatorioPacienteFeito := &RelatorioPacienteDTO{
		GraficoSono:    make([]PontoDeDadosDTO, 0),
		GraficoEnergia: make([]PontoDeDadosDTO, 0),
		GraficoStress:  make([]PontoDeDadosDTO, 0),
	}

	now := time.Now()
	dataFim := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	dataInicio := dataFim.AddDate(0, 0, -int(filtroPeriodo))

	registrosHumor, err := rs.registroHumorRepositorio.BuscarPorPacienteEPeriodo(pacienteID, dataInicio, dataFim)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return relatorioPacienteFeito, nil
		}
		return nil, err
	}

	totalSono := 0
	totalEnergia := 0
	totalStress := 0

	for _, registro := range registrosHumor {

		totalSono += int(registro.HorasSono)
		totalEnergia += int(registro.NivelEnergia)
		totalStress += int(registro.NivelStress)

		pontoDeDadosSono := PontoDeDadosDTO{
			Data:     registro.DataHoraRegistro,
			Valor:    registro.HorasSono,
			Humor:    registro.NivelHumor,
			Anotacao: registro.Observacoes,
		}
		relatorioPacienteFeito.GraficoSono = append(relatorioPacienteFeito.GraficoSono, pontoDeDadosSono)

		pontoDeDadosEnergia := PontoDeDadosDTO{
			Data:     registro.DataHoraRegistro,
			Valor:    registro.NivelEnergia,
			Humor:    registro.NivelHumor,
			Anotacao: registro.Observacoes,
		}
		relatorioPacienteFeito.GraficoEnergia = append(relatorioPacienteFeito.GraficoEnergia, pontoDeDadosEnergia)

		pontoDeDadosStress := PontoDeDadosDTO{
			Data:     registro.DataHoraRegistro,
			Valor:    registro.NivelStress,
			Humor:    registro.NivelHumor,
			Anotacao: registro.Observacoes,
		}
		relatorioPacienteFeito.GraficoStress = append(relatorioPacienteFeito.GraficoStress, pontoDeDadosStress)

	}
	if len(registrosHumor) > 0 {
		relatorioPacienteFeito.MediaSono = float64(totalSono) / float64(len(registrosHumor))
		relatorioPacienteFeito.MediaEnergia = float64(totalEnergia) / float64(len(registrosHumor))
		relatorioPacienteFeito.MediaStress = float64(totalStress) / float64(len(registrosHumor))

	}

	return relatorioPacienteFeito, nil
}
