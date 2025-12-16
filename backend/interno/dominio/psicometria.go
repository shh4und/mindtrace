package dominio

import "fmt"

func somarLista(listaDados []map[string]any) float64 {
	var total float64 = 0
	if len(listaDados) > 0 {
		for _, dados := range listaDados {
			total = total + dados["valor"].(float64)
		}
	}
	return total
}

// Função helper para calcular média simples de um slice
func calcularMediaSimples(valores []float64) float64 {
	if len(valores) == 0 {
		return 0
	}

	var total float64 = 0
	for _, v := range valores {
		total += v
	}
	return total / float64(len(valores))
}

type ResultadoClinico struct {
	ScoreTotal    float64
	Classificacao string
	Alertas       []string
	Detalhes      map[string]float64
}

type AvaliadorClinico interface {
	Avaliar(repostas []map[string]any) ResultadoClinico
}

type AvaliadorPHQ9 struct{}

func (av AvaliadorPHQ9) Avaliar(listaDados []map[string]any) ResultadoClinico {
	scoreTotal := somarLista(listaDados)
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
	if listaDados[8]["8"].(float64) > 0 {
		alertas = append(alertas, "Ideação suicida presente")
	}

	return ResultadoClinico{
		ScoreTotal:    scoreTotal,
		Classificacao: classificacao,
		Alertas:       alertas}
}

type AvaliadorGAD7 struct{}

func (av AvaliadorGAD7) Avaliar(listaDados []map[string]any) ResultadoClinico {
	scoreTotal := somarLista(listaDados)
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

func (av AvaliadorWHO5) Avaliar(listaDados []map[string]any) ResultadoClinico {
	scoreTotal := somarLista(listaDados) * 4
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

type AvaliadorWHOQOL struct{}

func (av AvaliadorWHOQOL) Avaliar(listaDados []map[string]any) ResultadoClinico {
	// Agrupar respostas por domínio dinamicamente
	scoresPorDominio := make(map[string][]float64)

	for _, dados := range listaDados {
		dominio := dados["dominio"].(string)
		valor := dados["valor"].(float64)

		// Adicionar o valor à lista do domínio
		scoresPorDominio[dominio] = append(scoresPorDominio[dominio], valor)
	}

	// Calcular a média de cada domínio e normalizar
	detalhes := make(map[string]float64)
	var scoreTotal float64 = 0

	for dominio, valores := range scoresPorDominio {
		// Calcular média do domínio
		media := calcularMediaSimples(valores)

		// Normalizar para escala 0-100 (escala 1-5)
		scoreNormalizado := media * 20

		detalhes[dominio] = scoreNormalizado
		scoreTotal += scoreNormalizado
	}

	// Score total é a média dos domínios
	quantidadeDominios := float64(len(detalhes))
	if quantidadeDominios > 0 {
		scoreTotal = scoreTotal / quantidadeDominios
	}

	// Classificação geral baseada no score total
	classificacao := av.classificarWHOQOL(scoreTotal)

	return ResultadoClinico{
		ScoreTotal:    scoreTotal,
		Classificacao: classificacao,
		Alertas:       []string{},
		Detalhes:      detalhes,
	}
}

func (av AvaliadorWHOQOL) classificarWHOQOL(score float64) string {
	if score < 25 {
		return "Qualidade de vida muito baixa"
	} else if score < 50 {
		return "Qualidade de vida baixa"
	} else if score < 75 {
		return "Qualidade de vida moderada"
	} else {
		return "Qualidade de vida boa"
	}
}

func CriarAvaliador(codigo string) (AvaliadorClinico, error) {
	switch codigo {
	case "phq_9":
		return AvaliadorPHQ9{}, nil
	case "gad_7":
		return AvaliadorGAD7{}, nil
	case "whoqol_bref":
		return AvaliadorWHOQOL{}, nil
	case "who_5":
		return AvaliadorWHO5{}, nil
	default:
		return nil, fmt.Errorf("instrumento nao reconhecido: %s", codigo)
	}
}
