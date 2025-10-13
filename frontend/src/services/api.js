import axios from 'axios';

// Determina a base da API:
// 1. Valor definido em build: import.meta.env.VITE_API_BASE_URL
// 2. Fallback: same-origin + "/api/v1" (ajuste se seu backend não tiver prefixo)
const buildTimeBase = import.meta.env?.VITE_API_BASE_URL;
const fallbackBase = `${window.location.origin}/api/v1`;
const baseURL = buildTimeBase || fallbackBase;

// instancia um cliente axios padrao para consumir a API
const apiClient = axios.create({
  baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true, // caso cookies venham a ser usados
});

// Interceptor: Adiciona o token JWT a todas as requisições protegidas
apiClient.interceptors.request.use(
  (config) => {
    // Rotas que não precisam de token
    const publicRoutes = ['/entrar/login', '/profissionais/registrar', '/pacientes/registrar'];
    if (publicRoutes.includes(config.url)) {
      return config;
    }

    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Funções de API que nossos componentes irão usar
const api = {
  // autenticacao e sessao
  login(credentials) {
    return apiClient.post('/entrar/login', credentials);
  },
  registrarPaciente(data) {
    return apiClient.post('/pacientes/registrar', data);
  },
  registrarProfissional(data) {
    return apiClient.post('/profissionais/registrar', data);
  },

  // --- Usuário ---
  buscarPerfil() {
    return apiClient.get('/usuarios/');
  },
  proprioPerfilPaciente() {
    return apiClient.get('/usuarios/paciente');
  },
  proprioPerfilProfissional() {
    return apiClient.get('/usuarios/profissional');
  },
  listarPacientesDoProfissional() {
    return apiClient.get('/usuarios/profissional/pacientes');
  },
  atualizarPerfil(data) {
    return apiClient.put('/usuarios/perfil', data);
  },
  alterarSenha(passwords) {
    return apiClient.put('/usuarios/perfil/alterar-senha', passwords);
  },
  deletarConta() {
    return apiClient.delete('/usuarios/perfil/apagar-conta');
  },

  // --- Registro de Humor ---
  registrarHumor(data) {
    return apiClient.post('/registro-humor/', data);
  },

  // --- Relatórios ---
  buscarRelatorio(periodo) {
    return apiClient.get(`/relatorios/?periodo=${periodo}`);
  },
  buscarRelatorioPacienteDoProfissional(periodo, pacienteID) {
    return apiClient.get(`/relatorios/paciente-lista?periodo=${periodo}&pacienteID=${pacienteID}`);
  },

  // --- Convites ---
  gerarConvite() {
    return apiClient.post('/convites/gerar');
  },
  vincularComToken(token) {
    return apiClient.post('/convites/vincular', { token });
  },
};

// exporta o cliente configurado e funcoes auxiliares
export { apiClient, api };
export default api;
