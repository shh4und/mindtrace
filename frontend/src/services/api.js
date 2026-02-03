import axios from "axios";

// Determina a base da API:
// 1. Valor definido em build: import.meta.env.VITE_API_BASE_URL
// 2. Fallback: same-origin + "/api/v1" (ajuste se seu backend nao tiver prefixo)
const buildTimeBase = import.meta.env?.VITE_API_BASE_URL;
const fallbackBase = `${window.location.origin}/api/v1`;
const baseURL = buildTimeBase || fallbackBase;

// instancia um cliente axios padrao para consumir a API
const apiClient = axios.create({
  baseURL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true, // caso cookies venham a ser usados
});

// Interceptor: Adiciona o token JWT a todas as requisicoes protegidas
apiClient.interceptors.request.use(
  (config) => {
    // Rotas que nao precisam de token
    const publicRoutes = [
      "/entrar/login",
      "/profissionais/registrar",
      "/pacientes/registrar",
    ];
    if (publicRoutes.includes(config.url)) {
      return config;
    }

    const token = localStorage.getItem("access_token");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);
apiClient.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    // Check if error is 401, not a retry, and NOT the refresh endpoint itself (to avoid loops)
    if (
      error.response &&
      error.response.status === 401 &&
      !originalRequest._retry &&
      !originalRequest.url.includes("/entrar/refresh")
    ) {
      originalRequest._retry = true;
      try {
        const refreshToken = localStorage.getItem("refresh_token");
        
        if (!refreshToken) {
            return Promise.reject(error);
        }

        // Call refresh endpoint with object payload
        const response = await api.refreshTokenRotation({ refresh_token: refreshToken });
        const { access_token, refresh_token: new_refresh_token } = response.data;

        // Update local storage
        localStorage.setItem("access_token", access_token);
        if (new_refresh_token) {
            localStorage.setItem("refresh_token", new_refresh_token);
        }

        // Update Authorization header
        apiClient.defaults.headers.common["Authorization"] = `Bearer ${access_token}`;
        originalRequest.headers["Authorization"] = `Bearer ${access_token}`;

        // Retry the original request using apiClient
        return apiClient(originalRequest);
      } catch (refreshError) {
        // If refresh fails, clear tokens and reject
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");
        // Optionally redirect to login here if not handled by router guards
        // window.location.href = '/login'; 
        return Promise.reject(refreshError);
      }
    }
    return Promise.reject(error);
  }
);
// Funções de API que nossos componentes irao usar
const api = {
  // autenticacao e sessao
  login(credentials) {
    return apiClient.post("/entrar/login", credentials);
  },
  refreshTokenRotation(refreshToken) {
    return apiClient.post("/entrar/refresh", refreshToken);
  },
  registrarPaciente(data) {
    return apiClient.post("/pacientes/registrar", data);
  },
  registrarProfissional(data) {
    return apiClient.post("/profissionais/registrar", data);
  },
  ativarConta(token) {
    return apiClient.get(`/entrar/ativar?token=${token}`);
  },
  reenvioAtivacao(payload) {
    return apiClient.post("/entrar/ativar/reenviar", payload);
  },
  // --- Usuario ---
  buscarPerfil() {
    return apiClient.get("/usuarios/");
  },
  proprioPerfilPaciente() {
    return apiClient.get("/usuarios/paciente");
  },
  proprioPerfilProfissional() {
    return apiClient.get("/usuarios/profissional");
  },
  listarPacientesDoProfissional() {
    return apiClient.get("/usuarios/profissional/pacientes");
  },
  listarProfissionaisDoPaciente() {
    return apiClient.get("/usuarios/paciente/profissionais");
  },
  atualizarPerfil(data) {
    return apiClient.put("/usuarios/perfil", data);
  },
  alterarSenha(passwords) {
    return apiClient.put("/usuarios/perfil/alterar-senha", passwords);
  },
  deletarConta() {
    return apiClient.delete("/usuarios/perfil/anonimizar");
  },

  // --- Registro de Humor ---
  registrarHumor(data) {
    return apiClient.post("/registro-humor/", data);
  },

  // --- Relatorios ---
  buscarRelatorio(periodo) {
    return apiClient.get(`/relatorios/?periodo=${periodo}`);
  },
  buscarRelatorioPacienteDoProfissional(periodo, pacienteID) {
    return apiClient.get(
      `/relatorios/paciente-lista?periodo=${periodo}&pacienteID=${pacienteID}`
    );
  },

  // --- Resumo ---
  buscarResumo() {
    return apiClient.get("/resumo/");
  },

  // --- Convites ---
  gerarConvite(payload) {
    return apiClient.post("/convites/gerar", payload);
  },
  vincularComToken(token) {
    return apiClient.post("/convites/vincular", { token });
  },
  consultarToken(tokenValue) {
    return apiClient.get(`/convites/info?token=${tokenValue}`);
  },

  listarQuestionarios() {
    return apiClient.get("/instrumentos/listar-instrumentos");
  },

  atribuirQuestionario(pacienteId, instrumentoId, instrumentoCodigo) {
    return apiClient.post(
      `/instrumentos/atribuir-instrumento?pacienteID=${pacienteId}&instrumentoID=${instrumentoId}&instrumentoCodigo=${instrumentoCodigo}`
    );
  },

  listarAtribuicoesPaciente() {
    return apiClient.get("/instrumentos/listar-atribuicoes-paciente");
  },

  listarAtribuicoesProfissional() {
    return apiClient.get("/instrumentos/listar-atribuicoes-profissional");
  },

  buscarAtribuicao(atribuicaoId) {
    return apiClient.get(
      `/instrumentos/atribuicao?atribuicaoID=${atribuicaoId}`
    );
  },

  enviarResposta(payload) {
    return apiClient.post("/instrumentos/registrar-respostas", payload);
  },

  visualizarRespostas(atribuicaoId) {
    return apiClient.get(
      `/instrumentos/visualizar-respostas?atribuicaoID=${atribuicaoId}`
    );
  },
};

// exporta o cliente configurado e funcoes auxiliares
export { apiClient, api };
export default api;
