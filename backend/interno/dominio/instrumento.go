package dominio

import (
	"time"

	"gorm.io/gorm"
)

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

// TableName define o nome da tabela no banco de dados
func (Instrumento) TableName() string {
	return "instrumentos"
}

// Pergunta representa um item individual do instrumento
type Pergunta struct {
	ID                   uint   `gorm:"primaryKey"`
	InstrumentoID        uint   `gorm:"not null;index;column:instrumento_id"`
	OrdemItem            int    `gorm:"not null;column:ordem_item"`
	Dominio              string `gorm:"size:100;column:dominio"` // Para WHOQOL
	Conteudo             string `gorm:"type:text;not null;column:conteudo"`
	EhPontuacaoInvertida bool   `gorm:"default:false;column:eh_pontuacao_invertida"`
}

func (Pergunta) TableName() string {
	return "perguntas"
}

// OpcaoEscala define as respostas possíveis (Likert)
type OpcaoEscala struct {
	ID            uint   `gorm:"primaryKey"`
	InstrumentoID uint   `gorm:"not null;index;column:instrumento_id"`
	Valor         int    `gorm:"not null;column:valor"`  // 0, 1, 2...
	Rotulo        string `gorm:"not null;column:rotulo"` // "Vários dias"
}

func (OpcaoEscala) TableName() string {
	return "opcoes_escala"
}
