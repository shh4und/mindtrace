/**
 * Tipos e constantes relacionados a usuários
 * Versão JavaScript pura - compatível com todos os arquivos .js e .vue
 */

/**
 * Enum para tipos de usuário
 * Valores devem corresponder aos retornados pela API
 */
export const TipoUsuario = {
  Profissional: 'profissional',
  Paciente: 'paciente'
};

/**
 * Verifica se é profissional
 * @param {Object} usuario - objeto do usuário
 * @returns {boolean}
 */
export function isProfissional(usuario) {
  return usuario?.tipo_usuario === TipoUsuario.Profissional;
}

/**
 * Verifica se é paciente
 * @param {Object} usuario - objeto do usuário
 * @returns {boolean}
 */
export function isPaciente(usuario) {
  return usuario?.tipo_usuario === TipoUsuario.Paciente;
}

/**
 * Valida se o tipo de usuário é válido
 * @param {string} tipo - tipo a ser validado
 * @returns {boolean}
 */
export function isValidTipoUsuario(tipo) {
  return Object.values(TipoUsuario).includes(tipo);
}

/**
 * Retorna o label formatado do tipo de usuário
 * @param {string} tipo - tipo do usuário
 * @returns {string}
 */
export function getTipoUsuarioLabel(tipo) {
  const labels = {
    [TipoUsuario.Profissional]: 'Profissional',
    [TipoUsuario.Paciente]: 'Paciente'
  };
  return labels[tipo] || 'Desconhecido';
}