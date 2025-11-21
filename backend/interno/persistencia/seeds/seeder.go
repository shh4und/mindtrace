package seeds

import (
	_ "embed"
	"log"

	"gorm.io/gorm"
)

//go:embed instrumentos_padrao.sql
var sqlInstrumentos string

func ExecutarSeeds(db *gorm.DB) {
	log.Println("Iniciando seeds de instrumentos padrao...")

	err := db.Exec(sqlInstrumentos).Error
	if err != nil {
		log.Fatalf("Error ao executar script de seed de instrumentos: %v", err)
	}

	log.Println("Seed de instrumentos realizado com sucesso.")
}
