package seeds

import (
	_ "embed"
	"log"
	"os"
	"strings"

	"gorm.io/gorm"
)

//go:embed instrumentos_padrao.sql
var sqlInstrumentos string

//go:embed dados_mock.sql
var sqlDadosMock string

func ExecutarSeeds(db *gorm.DB) {
	log.Println("Iniciando seeds de instrumentos padrao...")

	err := db.Exec(sqlInstrumentos).Error
	if err != nil {
		log.Fatalf("Error ao executar script de seed de instrumentos: %v", err)
	}

	log.Println("Seed de instrumentos realizado com sucesso.")
}

// ExecutarSeedsMock executa seeds de dados mockados apenas em ambiente de desenvolvimento
// Verificar a variável de ambiente GO_ENV ou APP_ENV para determinar o ambiente
func ExecutarSeedsMock(db *gorm.DB) {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = os.Getenv("APP_ENV")
	}

	// Apenas executa em ambiente de desenvolvimento
	if !isDevEnvironment(env) {
		log.Println("Ambiente de producao detectado. Pulando seeds de dados mock.")
		return
	}

	log.Println("Ambiente de desenvolvimento detectado. Iniciando seeds de dados mock...")

	err := db.Exec(sqlDadosMock).Error
	if err != nil {
		log.Printf("Aviso: Erro ao executar script de seed de dados mock (pode ser duplicata): %v", err)
		return
	}

	log.Println("Seed de dados mock realizado com sucesso.")
}

// isDevEnvironment verifica se o ambiente é de desenvolvimento
func isDevEnvironment(env string) bool {
	env = strings.ToLower(env)
	return env == "" || env == "development" || env == "dev" || env == "local"
}
