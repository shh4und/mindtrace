/**
 * Utilitários para manipulação de tokens JWT
 * Centraliza funções de parsing e extração de informações do token
 */

/**
 * Decodifica o payload de um token JWT
 * @param {string} token - Token JWT completo
 * @returns {Object|null} - Payload decodificado ou null se inválido
 */
export function parseJwt(token) {
  if (!token) return null;
  
  try {
    const base64Payload = token.split('.')[1];
    if (!base64Payload) return null;
    
    return JSON.parse(atob(base64Payload));
  } catch (error) {
    console.error('Erro ao decodificar token JWT:', error);
    return null;
  }
}

/**
 * Obtém o token JWT armazenado no localStorage
 * @returns {string|null} - Token ou null se não existir
 */
export function getStoredToken() {
  return localStorage.getItem('token');
}

/**
 * Extrai o role do usuário do token JWT armazenado
 * @returns {string|null} - Role do usuário ('profissional' ou 'paciente') ou null
 */
export function getUserRoleFromToken() {
  const token = getStoredToken();
  if (!token) return null;
  
  const payload = parseJwt(token);
  return payload?.role || null;
}

/**
 * Extrai o ID do usuário do token JWT armazenado
 * @returns {string|null} - ID do usuário ou null
 */
export function getUserIdFromToken() {
  const token = getStoredToken();
  if (!token) return null;
  
  const payload = parseJwt(token);
  return payload?.sub || payload?.user_id || null;
}

/**
 * Verifica se o token está expirado
 * @param {string} token - Token JWT (opcional, usa o armazenado se não fornecido)
 * @returns {boolean} - true se expirado ou inválido, false se válido
 */
export function isTokenExpired(token = null) {
  const targetToken = token || getStoredToken();
  if (!targetToken) return true;
  
  const payload = parseJwt(targetToken);
  if (!payload?.exp) return true;
  
  // exp é em segundos, Date.now() é em milissegundos
  return Date.now() >= payload.exp * 1000;
}

/**
 * Verifica se o usuário está autenticado (tem token válido)
 * @returns {boolean}
 */
export function isAuthenticated() {
  const token = getStoredToken();
  return !!token && !isTokenExpired(token);
}

/**
 * Remove o token do localStorage
 */
export function clearToken() {
  localStorage.removeItem('token');
}

/**
 * Armazena um token no localStorage
 * @param {string} token - Token JWT a ser armazenado
 */
export function setToken(token) {
  localStorage.setItem('token', token);
}
