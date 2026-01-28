package servicos

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"net/smtp"
	"os"
	"text/template"
	"time"

	"gorm.io/gorm"
)

//go:embed templates/ativacao.html
var templateAtivacaoString string
var tmplAtivacao = template.Must(template.New("ativacao").Parse(templateAtivacaoString))

//go:embed templates/convite_profissional.html
var templateConviteString string
var tmplConvite = template.Must(template.New("convite").Parse(templateConviteString))

//go:embed templates/atribuicao.html
var templateAtribuicaoString string
var tmplAtribuicao = template.Must(template.New("atribuicao").Parse(templateAtribuicaoString))

//go:embed templates/monitoramento.html
var templateMonitoramentoString string
var tmplMonitoramento = template.Must(template.New("monitoramento").Parse(templateMonitoramentoString))

// EmailServico define os metodos para gerenciamento de emails
type EmailServico interface {
	EnviarEmailAtivacao(toEmail string, emailVerifHash string) error
	VerificarHashToken(tokenHash string) error
	EnviarEmail(toEmails []string, subject, body string) error
	EnviarEmailConvite(toEmail string, emailVerifHash string) error
	EnviarEmailAtribuicao(toEmail, nomePaciente, nomeProfissional, nomeInstrumento string) error
	EnviarEmailAlertaMonitoramento(toEmail, nomeProfissional, nomePaciente, status string) error
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

	var bodyBuffer bytes.Buffer
	if err := tmplAtivacao.Execute(&bodyBuffer, struct{ Link string }{Link: link}); err != nil {
		return err
	}

	if err := es.EnviarEmail([]string{toEmail}, subject, bodyBuffer.String()); err != nil {
		return err
	}

	return nil
}

func (es *emailServico) EnviarEmailConvite(toEmail string, tokenVinculo string) error {
	// Configuração da mensagem
	subject := "Go SMTP | MindTrace | Convite de Vínculo"

	frontendURL := os.Getenv("FRONTEND_ORIGINS")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}

	link := fmt.Sprintf("%s/dashboard-paciente/vincular?token=%s", frontendURL, tokenVinculo)

	var bodyBuffer bytes.Buffer
	if err := tmplConvite.Execute(&bodyBuffer, struct {
		Link         string
		TokenVinculo string
	}{Link: link, TokenVinculo: tokenVinculo}); err != nil {
		return err
	}

	if err := es.EnviarEmail([]string{toEmail}, subject, bodyBuffer.String()); err != nil {
		return err
	}

	return nil
}

func (es *emailServico) EnviarEmailAtribuicao(toEmail, nomePaciente, nomeProfissional, nomeInstrumento string) error {
	subject := "MindTrace | Nova Atividade Atribuída"

	frontendURL := os.Getenv("FRONTEND_ORIGINS")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	link := fmt.Sprintf("%s/dashboard-paciente/questionarios", frontendURL)

	var bodyBuffer bytes.Buffer
	if err := tmplAtribuicao.Execute(&bodyBuffer, struct {
		NomePaciente     string
		NomeProfissional string
		NomeInstrumento  string
		Link             string
	}{
		NomePaciente:     nomePaciente,
		NomeProfissional: nomeProfissional,
		NomeInstrumento:  nomeInstrumento,
		Link:             link,
	}); err != nil {
		return err
	}

	return es.EnviarEmail([]string{toEmail}, subject, bodyBuffer.String())
}

func (es *emailServico) EnviarEmailAlertaMonitoramento(toEmail, nomeProfissional, nomePaciente, status string) error {
	subject := fmt.Sprintf("MindTrace | Alerta de Monitoramento - %s", status)

	frontendURL := os.Getenv("FRONTEND_ORIGINS")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	// Redireciona para lista de pacientes ou detalhe específico (se houver rota)
	link := fmt.Sprintf("%s/dashboard-profissional/pacientes", frontendURL)

	var bodyBuffer bytes.Buffer
	if err := tmplMonitoramento.Execute(&bodyBuffer, struct {
		NomeProfissional string
		NomePaciente     string
		Status           string
		Link             string
	}{
		NomeProfissional: nomeProfissional,
		NomePaciente:     nomePaciente,
		Status:           status,
		Link:             link,
	}); err != nil {
		return err
	}

	return es.EnviarEmail([]string{toEmail}, subject, bodyBuffer.String())
}

func (es *emailServico) EnviarEmail(toEmails []string, subject, body string) error {

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "\n"
	body = "<body>" + body + "</body>"
	message := []byte(subject + mime + body)

	var fromEmail string
	var auth smtp.Auth

	if smtpPass == "" || smtpUser == "" {
		auth = nil
		fromEmail = "no-reply@mindtrace.services"
	} else {
		auth = smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
		fromEmail = smtpUser
	}

	address := smtpHost + ":" + smtpPort

	fmt.Println("message:", string(message))
	err := smtp.SendMail(address, auth, fromEmail, toEmails, message)
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

		if usuario.EmailVerifExpiracao == nil {
			return dominio.ErrTokenExpirado
		}

		now := time.Now()
		if now.After(*usuario.EmailVerifExpiracao) {

			return dominio.ErrTokenExpirado
		} else {
			usuario.EstaAtivo = true
			usuario.EmailVerifToken = nil
			usuario.EmailVerifExpiracao = nil
			if err := es.usuarioRepo.Atualizar(tx, usuario); err != nil {
				return err
			}
		}

		return nil
	})
	return err
}
