import axios from 'axios';

// Cria uma instância do Axios com a URL base da nossa API
const apiClient = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Interceptor: Adiciona o token JWT a todas as requisições protegidas
apiClient.interceptors.request.use(
  (config) => {
    // Rotas que não precisam de token
    const publicRoutes = ['/entrar/login', '/profissionais/registrar', '/pacientes/registrar'];
    if (publicRoutes.includes(config.url)) {
      return config;
    }

    const token = localStorage.getItem('authToken');
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
export default {
  login(credentials) {
    return apiClient.post('/entrar/login', credentials);
  },

  // --- Usuário ---
  buscarPerfil() {
    return apiClient.get('/usuarios/');
  },
  atualizarPerfil(data) {
    return apiClient.put('/usuarios/perfil', data);
  },
  alterarSenha(passwords) {
    return apiClient.put('/usuarios/perfil/alterar-senha', passwords);
  },

  // --- Registro de Humor ---
  registrarHumor(data) {
    return apiClient.post('/registro-humor/', data);
  },

  // --- Relatórios ---
  buscarRelatorio(periodo) {
    return apiClient.get(`/relatorios/?periodo=${periodo}`);
  },

  // --- Convites (Exemplo futuro) ---
  gerarConvite() {
    // return apiClient.post('/profissionais/convite');
    // Simulação por enquanto
    return Promise.resolve({
      data: {
        token: Math.random().toString(36).substring(2, 12).toUpperCase(),
        data_expiracao: new Date(new Date().getTime() + 24 * 60 * 60 * 1000),
      }
    });
  },
};
