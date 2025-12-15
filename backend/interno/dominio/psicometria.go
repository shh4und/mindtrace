package dominio

func somarLista(lista []int64) float64 {
	var total float64 = 0
	if len(lista) > 0 {
		for valor, _ := range lista {
			total = total + float64(valor)
		}
	}
	return total
}

type ResultadoClinico struct {
	ScoreTotal    float64
	Classificacao string
	Alertas       []string
	Detalhes      map[string]any
}

type AvaliadorClinico interface {
	Avaliar(repostas []int64) ResultadoClinico
}

type AvaliadorPHQ9 struct{}

func (av AvaliadorPHQ9) Avaliar(valorRespostas []int64) ResultadoClinico {
	scoreTotal := somarLista(valorRespostas)
	var classificacao string

	if scoreTotal >= 0 && scoreTotal < 5 {
		classificacao = "Ausência ou sintomas depressivos mínimos"
	} else if scoreTotal >= 5 && scoreTotal < 10 {
		classificacao = "Sintomas depressivos leves"
	} else if scoreTotal >= 10 && scoreTotal < 15 {
		classificacao = "Depressão moderada"
	} else if scoreTotal >= 15 && scoreTotal < 20 {
		classificacao = "Depressão moderadamente grave"
	} else if scoreTotal >= 20 && scoreTotal <= 27 {
		classificacao = "Depressão grave"
	} else {
		classificacao = "Classificação inválida"
	}

	var alertas []string
	if valorRespostas[8] > 0 {
		alertas = append(alertas, "Ideação suicida presente")
	}

	return ResultadoClinico{
		ScoreTotal:    scoreTotal,
		Classificacao: classificacao,
		Alertas:       alertas}
}

type AvaliadorGAD7 struct{}

func (av AvaliadorGAD7) Avaliar(valorRespostas []int64) ResultadoClinico {
	scoreTotal := somarLista(valorRespostas)
	var classificacao string

	if scoreTotal >= 0 && scoreTotal < 5 {
		classificacao = "Ansiedade mínima"
	} else if scoreTotal >= 5 && scoreTotal < 10 {
		classificacao = "Ansiedade leve"
	} else if scoreTotal >= 10 && scoreTotal < 15 {
		classificacao = "Ansiedade moderada"
	} else if scoreTotal >= 15 && scoreTotal <= 21 {
		classificacao = "Ansiedade grave"
	} else {
		classificacao = "Classificação inválida"
	}

	return ResultadoClinico{
		ScoreTotal:    scoreTotal,
		Classificacao: classificacao,
	}
}

type AvaliadorWHO5 struct{}

func (av AvaliadorWHO5) Avaliar(valorRespostas []int64) ResultadoClinico {
	scoreTotal := somarLista(valorRespostas) * 4
	var classificacao string

	if scoreTotal >= 0 && scoreTotal < 29 {
		classificacao = "Bem-estar muito baixo"
	} else if scoreTotal >= 29 && scoreTotal < 50 {
		classificacao = "Bem-estar reduzido"
	} else if scoreTotal >= 50 && scoreTotal <= 100 {
		classificacao = "Bem-estar preservado"
	} else {
		classificacao = "Classificação inválida"
	}

	return ResultadoClinico{
		ScoreTotal:    scoreTotal,
		Classificacao: classificacao,
	}
}

type AvaliadorWHOQOLBREF struct{}
