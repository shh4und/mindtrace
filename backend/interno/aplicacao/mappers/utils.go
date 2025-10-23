package mappers

import (
	"mindtrace/backend/interno/aplicacao/dtos"
	"mindtrace/backend/interno/dominio"
)

// ===== MAPEADORES PARA SAÍDA =====

func UsuarioParaDTOOut(usuario *dominio.Usuario) *dtos.UsuarioDTOOut {
	return &dtos.UsuarioDTOOut{
		ID:        usuario.ID,
		Email:     usuario.Email,
		Nome:      usuario.Nome,
		Contato:   usuario.Contato,
		Bio:       usuario.Bio,
		CreatedAt: usuario.CreatedAt,
		UpdatedAt: usuario.UpdatedAt,
	}
}

func ProfissionalParaDTOOut(prof *dominio.Profissional) *dtos.ProfissionalDTOOut {
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
		TipoUsuario: "profissional",
		CPF:         dto.CPF,
		Contato:     dto.Contato,
	}

	profissional := &dominio.Profissional{
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
		TipoUsuario: "paciente",
		CPF:         dto.CPF,
		Contato:     dto.Contato,
	}

	paciente := &dominio.Paciente{
		DataNascimento:       dto.DataNascimento,
		Dependente:           dto.Dependente,
		NomeResponsavel:      dto.NomeResponsavel,
		ContatoResponsavel:   dto.ContatoResponsavel,
		DataInicioTratamento: dto.DataInicioTratamento,
	}

	return usuario, paciente
}

func CriarRegistroHumorDTOInParaEntidade(dto *dtos.CriarRegistroHumorDTOIn, pacienteID uint) *dominio.RegistroHumor {
	return &dominio.RegistroHumor{
		PacienteID:       pacienteID,
		NivelHumor:       dto.NivelHumor,
		HorasSono:        dto.HorasSono,
		NivelEnergia:     dto.NivelEnergia,
		NivelStress:      dto.NivelStress,
		AutoCuidado:      dto.AutoCuidado,
		Observacoes:      dto.Observacoes,
		DataHoraRegistro: dto.DataHoraRegistro,
	}
}
