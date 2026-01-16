package servicos

import (
	"fmt"
	"net/smtp"
	"os"

	"gorm.io/gorm"
)

// EmailServico define os metodos para gerenciamento de emails
type EmailServico interface {
	EnviarEmailAtivacao(toEmail string, emailVerifHash string) error
}

// emailServico implementa a interface EmailServico
type emailServico struct {
	db *gorm.DB
}

// NovoEmailServico cria uma nova instancia de EmailServico
func NovoEmailServico(db *gorm.DB) EmailServico {
	return &emailServico{db: db}
}

func (es *emailServico) EnviarEmailAtivacao(toEmail string, emailVerifHash string) error {
	// Configuração da mensagem
	subject := "Teste de Implementação Go SMTP | MindTrace"
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// Fallback para segurança se esquecer de configurar
	if smtpHost == "" {
		smtpHost = "localhost"
	}
	if smtpPort == "" {
		smtpPort = "1025"
	}

	frontendURL := os.Getenv("FRONTEND_ORIGINS")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	link := fmt.Sprintf("%s/verificar-email?token=%s", frontendURL, emailVerifHash)
	body := fmt.Sprintf("Olá!\n\nClique no link abaixo para ativar sua conta:\n%s", link)
	msg := fmt.Appendf(nil, "To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", toEmail, subject, body)

	// Endereço completo
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	err := smtp.SendMail(addr, nil, "no-reply@mindtrace.services", []string{toEmail}, msg)
	if err != nil {
		return fmt.Errorf("falha ao enviar email: %w", err)
	}
	return nil
}
