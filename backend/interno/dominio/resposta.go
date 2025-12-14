package dominio

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Resposta armazena o resultado processado e o dado bruto (JSON)
type Resposta struct {
	ID           uint       `gorm:"primaryKey"`
	AtribuicaoID uint       `gorm:"uniqueIndex;not null;column:atribuicao_id"`
	Atribuicao   Atribuicao `gorm:"foreignKey:AtribuicaoID"`
	// Metadados Relacionais (Para Relatórios)
	PontuacaoTotal float64 `gorm:"type:decimal(10,2);column:pontuacao_total"`
	Classificacao  string  `gorm:"size:255;column:classificacao"` // Ex: "Depressão Moderada"

	// Armazenamento Híbrido (JSONB)
	// Guarda exatamente o que o front enviou: { "q1": 2, "q2": 0 ... }
	DadosBrutos datatypes.JSON `gorm:"type:jsonb;column:dados_brutos"`

	DataResposta time.Time `gorm:"autoCreateTime;column:data_resposta"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (Resposta) TableName() string {
	return "respostas"
}
