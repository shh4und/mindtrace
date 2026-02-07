package servicos

import (
	"errors"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/aplicacao/mappers"
	"mindtrace/backend/interno/dominio"
	"mindtrace/backend/interno/persistencia/repositorios"
	"time"

	"gorm.io/gorm"
)

// RegistroHumorServico define os metodos para gerenciamento de registros de humor
type RegistroHumorServico interface {
	CriarRegistroHumor(dto *dtos.CriarRegistroHumorDTOIn, userID uint) (*dominio.RegistroHumor, error)
}

// registroHumorServico implementa a interface RegistroHumorServico
type registroHumorServico struct {
	db                 *gorm.DB
	repositorio        repositorios.RegistroHumorRepositorio
	usuarioRepositorio repositorios.UsuarioRepositorio
	analiseServico     AnaliseServico
}

// NovoRegistroHumorServico cria uma nova instancia de registroHumorServico
func NovoRegistroHumorServico(db *gorm.DB, repo repositorios.RegistroHumorRepositorio, userRepo repositorios.UsuarioRepositorio, analiseSvc AnaliseServico) *registroHumorServico {
	return &registroHumorServico{db: db, repositorio: repo, usuarioRepositorio: userRepo, analiseServico: analiseSvc}
}

// CriarRegistroHumor cria um novo registro de humor para o paciente
func (rhs *registroHumorServico) CriarRegistroHumor(dto *dtos.CriarRegistroHumorDTOIn, userID uint) (*dominio.RegistroHumor, error) {
	var registroHumorRealizado *dominio.RegistroHumor

	err := rhs.db.Transaction(func(tx *gorm.DB) error {
		paciente, err := rhs.usuarioRepositorio.BuscarPacientePorUsuarioID(tx, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return dominio.ErrUsuarioNaoEncontrado
			}
			return err
		}

		// Validação de taxa: Verificar último registro
		ultimoRegistro, err := rhs.repositorio.BuscarUltimoRegistroDePaciente(paciente.ID)
		if err == nil && ultimoRegistro != nil {
			// Se encontrou um registro anterior
			agora := time.Now()
			// Diferença de tempo
			diff := agora.Sub(ultimoRegistro.DataHoraRegistro)
			// Se for menor que 60 segundos (simulação de 1 dia)
			if diff < 60*time.Second {
				return dominio.ErrRegistroHumorMuitoRecente
			}
		} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			// Se houve erro de banco que não seja "não encontrado", retorna erro
			return err
		}

		novoRegistroHumor, err := mappers.CriarRegistroHumorDTOInParaEntidade(dto, paciente.ID)
		if err != nil {
			return err
		}
		// Validar o registro de humor antes de criar
		if err := novoRegistroHumor.Validar(); err != nil {
			return err
		}

		// Calcula o IBG antes de persistir
		novoRegistroHumor.IndiceBemEstarGeral = novoRegistroHumor.CalcularIBG()

		if err := rhs.repositorio.CriarRegistroHumor(tx, novoRegistroHumor); err != nil {
			return err
		}

		registroHumorRealizado = novoRegistroHumor
		return nil
	})

	if err != nil {
		return nil, err
	}

	// --- TRIGGER DE MONITORAMENTO ---
	// Executa em uma goroutine para não bloquear a resposta da API
	go func(pacID uint) {
		// TODO: tratar erros/logs aqui internamente
		_ = rhs.analiseServico.ExecutarMonitoramento(pacID)
	}(registroHumorRealizado.PacienteID)

	return registroHumorRealizado, nil
}
