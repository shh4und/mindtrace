package servicos

import (
	"mindtrace/backend/interno/persistencia/repositorios"

	"gorm.io/gorm"
)

type NotificacaoServico interface{}

type notificacaoServico struct {
	db                       *gorm.DB
	registroHumorRepositorio repositorios.RegistroHumorRepositorio
	usuarioRepositorio       repositorios.UsuarioRepositorio
}

func NovoNotificacaoServico(db *gorm.DB, registroHumorRepo repositorios.RegistroHumorRepositorio, usuarioRepo repositorios.UsuarioRepositorio) NotificacaoServico {
	return &notificacaoServico{
		db:                       db,
		registroHumorRepositorio: registroHumorRepo,
		usuarioRepositorio:       usuarioRepo,
	}
}
