package servicos

import (
	"fmt"
	"mindtrace/backend/interno/persistencia/repositorios"
	"net/smtp"
	"os"

	"gorm.io/gorm"
)

// EmailServico define os metodos para gerenciamento de emails
type EmailServico interface {
	EnviarEmailAtivacao(toEmail string, emailVerifHash string) error
	VerificarHashToken(tokenHash string) error
}

// emailServico implementa a interface EmailServico
type emailServico struct {
	db          *gorm.DB
	usuarioRepo repositorios.UsuarioRepositorio
}

// NovoEmailServico cria uma nova instancia de EmailServico
func NovoEmailServico(db *gorm.DB, ur repositorios.UsuarioRepositorio) EmailServico {
	return &emailServico{db: db, usuarioRepo: ur}
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
	frontendURL = "http://localhost:9090/api/v1/entrar"

	link := fmt.Sprintf("%s/ativar?token=%s", frontendURL, emailVerifHash)
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

func (es *emailServico) VerificarHashToken(tokenHash string) error {
	err := es.db.Transaction(func(tx *gorm.DB) error {
		usuario, err := es.usuarioRepo.BuscarUsuarioPorTokenHash(tokenHash)
		if err != nil {
			return err
		}
		usuario.EstaAtivo = true
		usuario.EmailVerifHash = nil

		if err := es.usuarioRepo.Atualizar(tx, usuario); err != nil {
			return err
		}
		return nil
	})
	return err
}
