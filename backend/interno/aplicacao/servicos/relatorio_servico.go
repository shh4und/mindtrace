package servicos

import (
	"errors"
	"mindtrace/backend/interno/persistencia/postgres"
	"time"

	"gorm.io/gorm"
)

type PontoDeDadosDTO struct {
	Data     time.Time `json:"data"`
	Valor    int16     `json:"valor"`
	Humor    int16     `json:"humor"`
	Anotacao string    `json:"anotacao,omitempty"`
}

type RelatorioPacienteDTO struct {
	GraficoSono    []PontoDeDadosDTO `json:"grafico_sono"`
	GraficoEnergia []PontoDeDadosDTO `json:"grafico_energia"`
	GraficoStress  []PontoDeDadosDTO `json:"grafico_stress"`
	MediaSono      float64           `json:"media_sono"`
	MediaEnergia   float64           `json:"media_energia"`
	MediaStress    float64           `json:"media_stress"`
}

type RelatorioServico interface {
	GerarRelatorioPaciente(pacienteID uint, filtroPeriodo int64) (*RelatorioPacienteDTO, error)
}

type relatorioServico struct {
	registroHumorRepositorio postgres.RegistroHumorRepositorio
}

func NovoRelatorioServico(registroHumorRepo postgres.RegistroHumorRepositorio) *relatorioServico {
	return &relatorioServico{registroHumorRepositorio: registroHumorRepo}
}

func (rs *relatorioServico) GerarRelatorioPaciente(pacienteID uint, filtroPeriodo int64) (*RelatorioPacienteDTO, error) {
	relatorioPacienteFeito := &RelatorioPacienteDTO{}
	periodoDeTempoDias := time.Duration(filtroPeriodo) * 24 * time.Hour
	now := time.Now()
	dataFim := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	dataInicio := dataFim.AddDate(0, 0, -int(periodoDeTempoDias))

	registrosHumor, err := rs.registroHumorRepositorio.BuscarPorPacienteEPeriodo(pacienteID, dataInicio, dataFim)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
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
