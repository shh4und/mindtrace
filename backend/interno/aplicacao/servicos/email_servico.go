package servicos

import (
	"bytes"
	"fmt"
	"log"
	"mindtrace/backend/interno/persistencia/repositorios"
	"net/smtp"
	"os"
	"path/filepath"
	"text/template"

	"gorm.io/gorm"
)

// EmailServico define os metodos para gerenciamento de emails
type EmailServico interface {
	EnviarEmailAtivacao(toEmail string, emailVerifHash string) error
	VerificarHashToken(tokenHash string) error
	EnviarEmail(toEmails []string, subject, body string) error
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
	subject := "Go SMTP | MindTrace | Ativacao de Email"

	frontendURL := os.Getenv("FRONTEND_ORIGINS")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	link := fmt.Sprintf("%s/ativacao?token=%s", frontendURL, emailVerifHash)

	ativacaoAbsPathAuto, err := filepath.Abs("interno/aplicacao/servicos/templates/ativacao.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFiles(ativacaoAbsPathAuto)
	if err != nil {
		return err
	}

	var bodyBuffer bytes.Buffer
	err = tmpl.Execute(&bodyBuffer, struct{ Link string }{Link: link})
	if err != nil {
		return err
	}

	if err := es.EnviarEmail([]string{toEmail}, subject, bodyBuffer.String()); err != nil {
		return err
	}

	return nil
}

func (es *emailServico) EnviarEmail(toEmails []string, subject, body string) error {

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	address := smtpHost + ":" + smtpPort
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "\n"
	body = "<body>" + body + "</body>"
	message := []byte(subject + mime + body)

	fmt.Println("message:", string(message))
	err := smtp.SendMail(address, nil, "no-reply@mindtrace.services", toEmails, message)
	if err != nil {
		log.Printf("u.Sendmail() err: %v", err)
		return err
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
