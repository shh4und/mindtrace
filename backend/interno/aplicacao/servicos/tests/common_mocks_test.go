package tests

import (
	"github.com/stretchr/testify/mock"
)

// MockEmailServico simula o servico de email
type MockEmailServico struct {
	mock.Mock
}

func (m *MockEmailServico) EnviarEmailAtivacao(toEmail, token string) error {
	args := m.Called(toEmail, token)
	return args.Error(0)
}

func (m *MockEmailServico) VerificarHashToken(tokenHash string) error {
	args := m.Called(tokenHash)
	return args.Error(0)
}

func (m *MockEmailServico) EnviarEmail(toEmails []string, subject, body string) error {
	args := m.Called(toEmails, subject, body)
	return args.Error(0)
}

func (m *MockEmailServico) EnviarEmailConvite(toEmail, token string) error {
	args := m.Called(toEmail, token)
	return args.Error(0)
}

func (m *MockEmailServico) EnviarEmailAtribuicao(toEmail, nomePaciente, nomeProfissional, nomeInstrumento string) error {
	args := m.Called(toEmail, nomePaciente, nomeProfissional, nomeInstrumento)
	return args.Error(0)
}

func (m *MockEmailServico) EnviarEmailAlertaMonitoramento(toEmail, nomeProfissional, nomePaciente, status string) error {
	args := m.Called(toEmail, nomeProfissional, nomePaciente, status)
	return args.Error(0)
}
