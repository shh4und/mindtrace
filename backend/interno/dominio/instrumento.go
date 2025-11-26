package dominio

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Erros de validação - Instrumento
var (
	ErrCodigoInstrumentoVazio     = errors.New("codigo do instrumento nao pode estar vazio")
	ErrCodigoInstrumentoInvalido  = errors.New("codigo deve ter entre 3 e 50 caracteres alfanumericos")
	ErrNomeInstrumentoVazio       = errors.New("nome do instrumento nao pode estar vazio")
	ErrNomeInstrumentoInvalido    = errors.New("nome deve ter entre 3 e 255 caracteres")
	ErrAlgoritmoPontuacaoVazio    = errors.New("algoritmo de pontuacao nao pode estar vazio")
	ErrAlgoritmoPontuacaoInvalido = errors.New("algoritmo de pontuacao invalido")
	ErrVersaoInvalida             = errors.New("versao deve ser maior que zero")
	ErrInstrumentoSemPerguntas    = errors.New("instrumento deve ter ao menos uma pergunta")
	ErrInstrumentoSemOpcoesEscala = errors.New("instrumento deve ter opcoes de escala definidas")
	ErrInstrumentoPadraoImutavel  = errors.New("instrumentos padronizados nao podem ser editados")
	ErrInstrumentoJaRespondido    = errors.New("instrumento com respostas nao pode ser editado")
	ErrCodigoInstrumentoJaExiste  = errors.New("codigo do instrumento ja existe")
)

// Erros de validação - Pergunta
var (
	ErrPerguntaConteudoVazio   = errors.New("conteudo da pergunta nao pode estar vazio")
	ErrPerguntaOrdemInvalida   = errors.New("ordem do item deve ser maior que zero")
	ErrPerguntaDominioInvalido = errors.New("dominio invalido para este instrumento")
)

// Erros de validação - OpcaoEscala
var (
	ErrOpcaoEscalaRotuloVazio   = errors.New("rotulo da opcao nao pode estar vazio")
	ErrOpcaoEscalaValorInvalido = errors.New("valor da opcao deve estar no intervalo permitido")
	ErrOpcaoEscalaDuplicada     = errors.New("valor ja existe para este instrumento")
)

// Constantes de instrumentos padronizados (IMUTÁVEIS)
var InstrumentosPadronizados = map[string]bool{
	"phq_9":       true,
	"gad_7":       true,
	"whoqol_bref": true,
	"who_5":       true,
}

// Algoritmos de pontuação válidos
var AlgoritmosValidos = map[string]bool{
	"phq_9":       true, // Soma simples (0-27)
	"gad_7":       true, // Soma simples (0-21)
	"whoqol_bref": true, // Média por domínio
	"who_5":       true, // Soma convertida (0-100)
}

// Domínios válidos para WHOQOL-BREF
var DominiosWHOQOL = map[string]bool{
	"Geral":            true,
	"Físico":           true,
	"Psicológico":      true,
	"Relações Sociais": true,
	"Meio Ambiente":    true,
}

// Instrumento representa os metadados de um questionário
type Instrumento struct {
	ID                 uint   `gorm:"primaryKey"`
	Codigo             string `gorm:"uniqueIndex;not null;column:codigo"`
	Nome               string `gorm:"not null;column:nome"`
	Descricao          string `gorm:"type:text;column:descricao"`
	AlgoritmoPontuacao string `gorm:"not null;column:algoritmo_pontuacao"`
	Versao             int    `gorm:"default:1;column:versao"`
	EstaAtivo          bool   `gorm:"default:true;column:esta_ativo"`

	// Relacionamentos
	Perguntas    []Pergunta    `gorm:"foreignKey:InstrumentoID;constraint:OnDelete:CASCADE"`
	OpcoesEscala []OpcaoEscala `gorm:"foreignKey:InstrumentoID;constraint:OnDelete:CASCADE"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Instrumento) TableName() string {
	return "instrumentos"
}

// EhPadronizado verifica se o instrumento é um template do sistema (IMUTÁVEL)
func (i *Instrumento) EhPadronizado() bool {
	return InstrumentosPadronizados[i.Codigo]
}

// ValidarCodigo valida o código do instrumento
func (i *Instrumento) ValidarCodigo() error {
	if i.Codigo == "" {
		return ErrCodigoInstrumentoVazio
	}
	if len(i.Codigo) < 3 || len(i.Codigo) > 50 {
		return ErrCodigoInstrumentoInvalido
	}
	// Validar formato alfanumérico com underscore
	for _, c := range i.Codigo {
		if !((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_') {
			return ErrCodigoInstrumentoInvalido
		}
	}
	return nil
}

// ValidarNome valida o nome do instrumento
func (i *Instrumento) ValidarNome() error {
	if i.Nome == "" {
		return ErrNomeInstrumentoVazio
	}
	if len(i.Nome) < 3 || len(i.Nome) > 255 {
		return ErrNomeInstrumentoInvalido
	}
	return nil
}

// ValidarAlgoritmoPontuacao valida o algoritmo de pontuação
func (i *Instrumento) ValidarAlgoritmoPontuacao() error {
	if i.AlgoritmoPontuacao == "" {
		return ErrAlgoritmoPontuacaoVazio
	}
	if !AlgoritmosValidos[i.AlgoritmoPontuacao] {
		return ErrAlgoritmoPontuacaoInvalido
	}
	return nil
}

// ValidarVersao valida a versão do instrumento
func (i *Instrumento) ValidarVersao() error {
	if i.Versao < 1 {
		return ErrVersaoInvalida
	}
	return nil
}

// ValidarPerguntas valida se o instrumento tem perguntas
func (i *Instrumento) ValidarPerguntas() error {
	if len(i.Perguntas) == 0 {
		return ErrInstrumentoSemPerguntas
	}

	// Validar cada pergunta
	for _, p := range i.Perguntas {
		if err := p.Validar(i.Codigo); err != nil {
			return err
		}
	}

	// Validar ordem sequencial (não pode ter lacunas)
	ordens := make(map[int]bool)
	for _, p := range i.Perguntas {
		if ordens[p.OrdemItem] {
			return errors.New("ordem de item duplicada")
		}
		ordens[p.OrdemItem] = true
	}

	return nil
}

// ValidarOpcoesEscala valida as opções de escala
func (i *Instrumento) ValidarOpcoesEscala() error {
	if len(i.OpcoesEscala) == 0 {
		return ErrInstrumentoSemOpcoesEscala
	}

	// Validar cada opção
	valores := make(map[int]bool)
	for _, op := range i.OpcoesEscala {
		if err := op.Validar(); err != nil {
			return err
		}
		if valores[op.Valor] {
			return ErrOpcaoEscalaDuplicada
		}
		valores[op.Valor] = true
	}

	return nil
}

// PodeSerEditado verifica se o instrumento pode ser modificado
func (i *Instrumento) PodeSerEditado() error {
	// Instrumentos padronizados são IMUTÁVEIS
	if i.EhPadronizado() {
		return ErrInstrumentoPadraoImutavel
	}
	return nil
}

// Validar executa todas as validações do instrumento
func (i *Instrumento) Validar() error {
	if err := i.ValidarCodigo(); err != nil {
		return err
	}
	if err := i.ValidarNome(); err != nil {
		return err
	}
	if err := i.ValidarAlgoritmoPontuacao(); err != nil {
		return err
	}
	if err := i.ValidarVersao(); err != nil {
		return err
	}
	if err := i.ValidarPerguntas(); err != nil {
		return err
	}
	if err := i.ValidarOpcoesEscala(); err != nil {
		return err
	}
	return nil
}

// Pergunta representa um item individual do instrumento
type Pergunta struct {
	ID                   uint   `gorm:"primaryKey"`
	InstrumentoID        uint   `gorm:"not null;index;column:instrumento_id;uniqueIndex:idx_pergunta_unica,priority:1"`
	OrdemItem            int    `gorm:"not null;column:ordem_item;uniqueIndex:idx_pergunta_unica,priority:2"`
	Dominio              string `gorm:"size:100;column:dominio;uniqueIndex:idx_pergunta_unica,priority:3"`
	Conteudo             string `gorm:"type:text;not null;column:conteudo;uniqueIndex:idx_pergunta_unica,priority:4"`
	EhPontuacaoInvertida bool   `gorm:"default:false;column:eh_pontuacao_invertida"`
}

func (Pergunta) TableName() string {
	return "perguntas"
}

// ValidarConteudo valida o conteúdo da pergunta
func (p *Pergunta) ValidarConteudo() error {
	if p.Conteudo == "" {
		return ErrPerguntaConteudoVazio
	}
	return nil
}

// ValidarOrdem valida a ordem do item
func (p *Pergunta) ValidarOrdem() error {
	if p.OrdemItem < 1 {
		return ErrPerguntaOrdemInvalida
	}
	return nil
}

// ValidarDominio valida o domínio (específico para WHOQOL)
func (p *Pergunta) ValidarDominio(codigoInstrumento string) error {
	// WHOQOL-BREF exige domínios válidos
	if codigoInstrumento == "whoqol_bref" {
		if p.Dominio == "" {
			return errors.New("pergunta do WHOQOL-BREF deve ter dominio definido")
		}
		if !DominiosWHOQOL[p.Dominio] {
			return ErrPerguntaDominioInvalido
		}
	}
	return nil
}

// Validar executa todas as validações da pergunta
func (p *Pergunta) Validar(codigoInstrumento string) error {
	if err := p.ValidarConteudo(); err != nil {
		return err
	}
	if err := p.ValidarOrdem(); err != nil {
		return err
	}
	if err := p.ValidarDominio(codigoInstrumento); err != nil {
		return err
	}
	return nil
}

// OpcaoEscala define as respostas possíveis (Likert)
type OpcaoEscala struct {
	ID            uint   `gorm:"primaryKey"`
	InstrumentoID uint   `gorm:"not null;index;column:instrumento_id;uniqueIndex:idx_opcao_escala_unica,priority:1"`
	Valor         int    `gorm:"not null;column:valor;uniqueIndex:idx_opcao_escala_unica,priority:2"`
	Rotulo        string `gorm:"not null;column:rotulo;uniqueIndex:idx_opcao_escala_unica,priority:3"`
}

func (OpcaoEscala) TableName() string {
	return "opcoes_escala"
}

// ValidarRotulo valida o rótulo da opção
func (o *OpcaoEscala) ValidarRotulo() error {
	if o.Rotulo == "" {
		return ErrOpcaoEscalaRotuloVazio
	}
	return nil
}

// ValidarValor valida o valor da opção (específico por escala)
func (o *OpcaoEscala) ValidarValor() error {
	// Valores típicos: 0-3 (PHQ-9, GAD-7) ou 0-5 (WHO-5) ou 1-5 (WHOQOL)
	if o.Valor < 0 || o.Valor > 5 {
		return ErrOpcaoEscalaValorInvalido
	}
	return nil
}

// Validar executa todas as validações da opção de escala
func (o *OpcaoEscala) Validar() error {
	if err := o.ValidarRotulo(); err != nil {
		return err
	}
	if err := o.ValidarValor(); err != nil {
		return err
	}
	return nil
}
