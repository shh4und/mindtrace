package mappers

import (
	"encoding/json"
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/dominio"
)

// ===== MAPEADORES PARA SAÍDA =====

func UsuarioParaDTOOut(usuario *dominio.Usuario) *dtos.UsuarioDTOOut {
	return &dtos.UsuarioDTOOut{
		ID:          usuario.ID,
		Email:       usuario.Email,
		Nome:        usuario.Nome,
		TipoUsuario: dominio.TipoUsuarioParaString(usuario.TipoUsuario),
		Contato:     usuario.Contato,
		Bio:         usuario.Bio,
		CreatedAt:   usuario.CreatedAt,
		UpdatedAt:   usuario.UpdatedAt,
	}
}

func ProfissionalParaDTOOut(prof *dominio.Profissional) *dtos.ProfissionalDTOOut {
	if prof == nil {
		return nil
	}
	usuarioDTO := UsuarioParaDTOOut(&prof.Usuario)

	return &dtos.ProfissionalDTOOut{
		ID:                   prof.ID,
		Usuario:              *usuarioDTO,
		DataNascimento:       prof.DataNascimento,
		Especialidade:        prof.Especialidade,
		RegistroProfissional: prof.RegistroProfissional,
		CreatedAt:            prof.CreatedAt,
		UpdatedAt:            prof.UpdatedAt,
	}
}

func PacienteParaDTOOut(pac *dominio.Paciente) *dtos.PacienteDTOOut {
	if pac == nil {
		return nil
	}
	usuarioDTO := UsuarioParaDTOOut(&pac.Usuario)

	profissionalDTOs := make([]dtos.ProfissionalDTOOut, len(pac.Profissionais))
	for i, prof := range pac.Profissionais {
		profissionalDTOs[i] = *ProfissionalParaDTOOut(&prof)
	}

	return &dtos.PacienteDTOOut{
		ID:             pac.ID,
		Usuario:        *usuarioDTO,
		DataNascimento: pac.DataNascimento,
		Dependente:     &pac.Dependente,
		Profissionais:  profissionalDTOs,
		CreatedAt:      pac.CreatedAt,
		UpdatedAt:      pac.UpdatedAt,
	}
}

func RegistroHumorParaDTOOut(reg *dominio.RegistroHumor) *dtos.RegistroHumorDTOOut {
	return &dtos.RegistroHumorDTOOut{
		ID:               reg.ID,
		PacienteID:       reg.PacienteID,
		NivelHumor:       reg.NivelHumor,
		HorasSono:        reg.HorasSono,
		NivelEnergia:     reg.NivelEnergia,
		NivelStress:      reg.NivelStress,
		AutoCuidado:      reg.AutoCuidado,
		Observacoes:      reg.Observacoes,
		DataHoraRegistro: reg.DataHoraRegistro,
		CreatedAt:        reg.CreatedAt,
	}
}

func ResumoPacienteParaDTOOut(reg *dominio.RegistroHumor) *dtos.ResumoPacienteDTOOut {
	return &dtos.ResumoPacienteDTOOut{
		Data:     reg.DataHoraRegistro,
		Humor:    reg.NivelHumor,
		Anotacao: reg.Observacoes,
	}
}

// ===== MAPEADORES PARA SAÍDA SIMPLIFICADA (Para APIs) =====

// PacientesParaDTOOut converte um slice de Pacientes para DTOs de saída
func PacientesParaDTOOut(pacientes []dominio.Paciente) []dtos.PacienteDTOOut {
	dtos := make([]dtos.PacienteDTOOut, len(pacientes))
	for i, pac := range pacientes {
		dtos[i] = *PacienteParaDTOOut(&pac)
	}
	return dtos
}

// ProfissionaisParaDTOOut converte um slice de Profissionais para DTOs de saída
func ProfissionaisParaDTOOut(profissionais []dominio.Profissional) []dtos.ProfissionalDTOOut {
	dtos := make([]dtos.ProfissionalDTOOut, len(profissionais))
	for i, prof := range profissionais {
		dtos[i] = *ProfissionalParaDTOOut(&prof)
	}
	return dtos
}

// ===== MAPEADORES PARA ENTRADA =====

func RegistrarUsuarioDTOInParaEntidade(dto *dtos.RegistrarUsuarioDTOIn) *dominio.Usuario {
	return &dominio.Usuario{
		Email: dto.Email,
		Senha: dto.Senha, // Será hasheado no serviço!
		Nome:  dto.Nome,
	}
}

func RegistrarProfissionalDTOInParaEntidade(dto *dtos.RegistrarProfissionalDTOIn) (*dominio.Usuario, *dominio.Profissional) {
	usuario := &dominio.Usuario{
		Nome:        dto.Nome,
		Email:       dto.Email,
		Senha:       dto.Senha, // Será hasheado no serviço!
		TipoUsuario: 2,
		CPF:         dto.CPF,
		Contato:     dto.Contato,
	}

	profissional := &dominio.Profissional{
		DataNascimento:       dto.DataNascimento,
		Especialidade:        dto.Especialidade,
		RegistroProfissional: dto.RegistroProfissional,
	}

	return usuario, profissional
}

func RegistrarPacienteDTOInParaEntidade(dto *dtos.RegistrarPacienteDTOIn) (*dominio.Usuario, *dominio.Paciente) {
	usuario := &dominio.Usuario{
		Nome:        dto.Nome,
		Email:       dto.Email,
		Senha:       dto.Senha, // Será hasheado no serviço!
		TipoUsuario: 3,
		CPF:         dto.CPF,
		Contato:     dto.Contato,
	}

	paciente := &dominio.Paciente{
		DataNascimento:       dto.DataNascimento,
		Dependente:           *dto.Dependente,
		NomeResponsavel:      dto.NomeResponsavel,
		ContatoResponsavel:   dto.ContatoResponsavel,
		DataInicioTratamento: dto.DataInicioTratamento,
	}

	return usuario, paciente
}

func CriarRegistroHumorDTOInParaEntidade(dto *dtos.CriarRegistroHumorDTOIn, pacienteID uint) (*dominio.RegistroHumor, error) {
	autoCuidadoJSONB, err := json.Marshal(dto.AutoCuidado)
	if err != nil {
		return nil, err
	}

	return &dominio.RegistroHumor{
		PacienteID:       pacienteID,
		NivelHumor:       dto.NivelHumor,
		HorasSono:        *dto.HorasSono,
		NivelEnergia:     dto.NivelEnergia,
		NivelStress:      dto.NivelStress,
		AutoCuidado:      string(autoCuidadoJSONB),
		Observacoes:      dto.Observacoes,
		DataHoraRegistro: dto.DataHoraRegistro,
	}, nil
}

func ConviteParaDTOOut(convite *dominio.Convite) *dtos.ConviteDTOOut {
	if convite == nil {
		return nil
	}
	return &dtos.ConviteDTOOut{
		Token:         convite.Token,
		DataExpiracao: convite.DataExpiracao,
		Usado:         convite.Usado,
		CreatedAt:     convite.CreatedAt,
	}
}

func InstrumentosParaDTOOut(instrumentos []*dominio.Instrumento) []*dtos.InstrumentoDTOOut {
	instrumentoDTOs := make([]*dtos.InstrumentoDTOOut, len(instrumentos))

	for i, inst := range instrumentos {
		instrumentoDTOs[i] = &dtos.InstrumentoDTOOut{
			ID:        inst.ID,
			Codigo:    inst.Codigo,
			Nome:      inst.Nome,
			Descricao: inst.Descricao,
			Versao:    inst.Versao,
		}
	}

	return instrumentoDTOs
}

// AtribuicaoParaDTOOutPaciente converte Atribuicao para DTO (visão do paciente)
func AtribuicaoParaDTOOutPaciente(atrib *dominio.Atribuicao) *dtos.AtribuicaoDTOOut {
	if atrib == nil {
		return nil
	}

	// Contar perguntas do instrumento
	totalPerguntas := len(atrib.Instrumento.Perguntas)

	return &dtos.AtribuicaoDTOOut{
		ID:             atrib.ID,
		Status:         string(atrib.Status),
		DataAtribuicao: atrib.CreatedAt,
		DataResposta:   atrib.DataResposta,
		Instrumento: dtos.InstrumentoResumidoDTOOut{
			Codigo:         atrib.Instrumento.Codigo,
			Nome:           atrib.Instrumento.Nome,
			Descricao:      atrib.Instrumento.Descricao,
			TotalPerguntas: totalPerguntas,
		},
		Profissional: &dtos.ProfissionalResumidoDTOOut{
			ID:            atrib.Profissional.ID,
			Nome:          atrib.Profissional.Usuario.Nome,
			Email:         atrib.Profissional.Usuario.Email,
			Especialidade: atrib.Profissional.Especialidade,
		},
	}
}

// AtribuicaoParaDTOOutProfissional converte Atribuicao para DTO (visão do profissional)
func AtribuicaoParaDTOOutProfissional(atrib *dominio.Atribuicao) *dtos.AtribuicaoDTOOut {
	if atrib == nil {
		return nil
	}

	// Contar perguntas do instrumento
	totalPerguntas := len(atrib.Instrumento.Perguntas)

	return &dtos.AtribuicaoDTOOut{
		ID:             atrib.ID,
		Status:         string(atrib.Status),
		DataAtribuicao: atrib.CreatedAt,
		DataResposta:   atrib.DataResposta,
		Instrumento: dtos.InstrumentoResumidoDTOOut{
			Codigo:         atrib.Instrumento.Codigo,
			Nome:           atrib.Instrumento.Nome,
			Descricao:      atrib.Instrumento.Descricao,
			TotalPerguntas: totalPerguntas,
		},
		Paciente: &dtos.PacienteResumidoDTOOut{
			ID:    atrib.Paciente.ID,
			Nome:  atrib.Paciente.Usuario.Nome,
			Email: atrib.Paciente.Usuario.Email,
		},
	}
}

// AtribuicoesParaDTOOutPaciente converte slice de Atribuicoes para DTOs (visão paciente)
func AtribuicoesParaDTOOutPaciente(atribuicoes []*dominio.Atribuicao) []*dtos.AtribuicaoDTOOut {
	dtos := make([]*dtos.AtribuicaoDTOOut, 0)
	for _, atrib := range atribuicoes {
		dtos = append(dtos, AtribuicaoParaDTOOutPaciente(atrib))
	}
	return dtos
}

// AtribuicoesParaDTOOutProfissional converte slice de Atribuicoes para DTOs (visão profissional)
func AtribuicoesParaDTOOutProfissional(atribuicoes []*dominio.Atribuicao) []*dtos.AtribuicaoDTOOut {
	dtos := make([]*dtos.AtribuicaoDTOOut, 0)
	for _, atrib := range atribuicoes {
		dtos = append(dtos, AtribuicaoParaDTOOutProfissional(atrib))
	}
	return dtos
}
