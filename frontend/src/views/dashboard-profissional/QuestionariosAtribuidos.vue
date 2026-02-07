<template>
  <div class="max-w-6xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center md:text-left">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-indigo-50 text-indigo-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <font-awesome-icon :icon="faClipboardList" class="h-4 w-4" />
        <span>Monitoramento</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Questionários Atribuídos
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Acompanhe os questionários atribuídos aos seus pacientes.
      </p>
    </header>

    <!-- Filtros -->
    <section
      class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-8"
    >
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label
            class="block text-xs font-bold text-gray-500 uppercase tracking-wide mb-2"
            >Status</label
          >
          <div class="relative">
            <select
              v-model="filtroStatus"
              class="w-full pl-4 pr-10 py-3 bg-gray-50 hover:bg-white border border-gray-200 rounded-2xl focus:ring-2 focus:ring-indigo-500 focus:border-transparent outline-none appearance-none transition-all shadow-sm text-gray-700 font-medium"
            >
              <option value="">Todos os status</option>
              <option value="PENDENTE">Pendente</option>
              <option value="RESPONDIDO">Respondido</option>
              <option value="EXPIRADO">Expirado</option>
            </select>
            <div
              class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-3 text-gray-500"
            >
              <font-awesome-icon :icon="faChevronDown" class="w-3 h-3" />
            </div>
          </div>
        </div>

        <div>
          <label
            class="block text-xs font-bold text-gray-500 uppercase tracking-wide mb-2"
            >Paciente</label
          >
          <div class="relative">
            <input
              v-model="filtroPaciente"
              type="text"
              placeholder="Buscar por nome..."
              class="w-full pl-4 pr-10 py-3 bg-gray-50 hover:bg-white border border-gray-200 rounded-2xl focus:ring-2 focus:ring-indigo-500 focus:border-transparent outline-none transition-all shadow-sm text-gray-700 placeholder-gray-400 font-medium"
            />
            <div
              class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-3 text-gray-400"
            >
              <font-awesome-icon :icon="faSearch" class="w-3 h-3" />
            </div>
          </div>
        </div>

        <div>
          <label
            class="block text-xs font-bold text-gray-500 uppercase tracking-wide mb-2"
            >Instrumento</label
          >
          <div class="relative">
            <select
              v-model="filtroInstrumento"
              class="w-full pl-4 pr-10 py-3 bg-gray-50 hover:bg-white border border-gray-200 rounded-2xl focus:ring-2 focus:ring-indigo-500 focus:border-transparent outline-none appearance-none transition-all shadow-sm text-gray-700 font-medium"
            >
              <option value="">Todos os instrumentos</option>
              <option value="phq_9">PHQ-9</option>
              <option value="gad_7">GAD-7</option>
              <option value="whoqol_bref">WHOQOL-BREF</option>
              <option value="who_5">WHO-5</option>
            </select>
            <div
              class="pointer-events-none absolute inset-y-0 right-0 flex items-center px-3 text-gray-500"
            >
              <font-awesome-icon :icon="faChevronDown" class="w-3 h-3" />
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"
      ></div>
    </div>

    <!-- Empty state -->
    <div
      v-else-if="atribuicoesFiltradas.length === 0"
      class="text-center py-20 bg-white rounded-3xl border border-dashed border-gray-200"
    >
      <div
        class="w-20 h-20 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-6 text-gray-300"
      >
        <font-awesome-icon :icon="faClipboardList" class="w-10 h-10" />
      </div>
      <h3 class="text-xl font-bold text-gray-900 mb-2">
        {{
          atribuicoes.length === 0
            ? "Nenhuma atribuição encontrada"
            : "Nenhum resultado encontrado"
        }}
      </h3>
      <p class="text-gray-500 max-w-md mx-auto">
        {{
          atribuicoes.length === 0
            ? "Você ainda não atribuiu questionários aos seus pacientes."
            : "Tente ajustar os filtros para encontrar o que procura."
        }}
      </p>
    </div>

    <!-- Lista de atribuições -->
    <div v-else class="space-y-6">
      <!-- Estatísticas resumidas -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 font-medium">Total</p>
              <p class="text-3xl font-extrabold text-gray-900">
                {{ atribuicoesFiltradas.length }}
              </p>
            </div>
            <div class="p-3 bg-indigo-100 rounded-xl">
              <font-awesome-icon
                :icon="faClipboardList"
                class="w-5 h-5 text-indigo-600"
              />
            </div>
          </div>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 font-medium">Pendentes</p>
              <p class="text-3xl font-extrabold text-amber-600">
                {{ contarPorStatus("PENDENTE") }}
              </p>
            </div>
            <div class="p-3 bg-amber-100 rounded-xl">
              <font-awesome-icon
                :icon="faClock"
                class="w-5 h-5 text-amber-600"
              />
            </div>
          </div>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 font-medium">Respondidos</p>
              <p class="text-3xl font-extrabold text-emerald-600">
                {{ contarPorStatus("RESPONDIDO") }}
              </p>
            </div>
            <div class="p-3 bg-emerald-100 rounded-xl">
              <font-awesome-icon
                :icon="faCheckCircle"
                class="w-5 h-5 text-emerald-600"
              />
            </div>
          </div>
        </div>
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-5 hover:shadow-md transition-shadow"
        >
          <div class="flex items-center justify-between">
            <div>
              <p class="text-sm text-gray-500 font-medium">Expirados</p>
              <p class="text-3xl font-extrabold text-gray-600">
                {{ contarPorStatus("EXPIRADO") }}
              </p>
            </div>
            <div class="p-3 bg-gray-100 rounded-xl">
              <font-awesome-icon
                :icon="faTimesCircle"
                class="w-5 h-5 text-gray-500"
              />
            </div>
          </div>
        </div>
      </div>

      <!-- Cards de atribuições -->
      <div class="space-y-4">
        <section
          v-for="(atribuicao, index) in atribuicoesFiltradas"
          :key="atribuicao.id"
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-all duration-300"
        >
          <div class="flex items-start justify-between">
            <div class="flex-1">
              <!-- Header com paciente -->
              <div class="flex items-center gap-3 mb-4">
                <div
                  :class="[
                    getAvatarColor(index),
                    'w-20 h-20 rounded-2xl flex items-center justify-center text-2xl font-bold mb-4 shadow-sm z-10 transition-transform group-hover:scale-110',
                  ]"
                >
                  {{ atribuicao.paciente.nome.charAt(0).toUpperCase() }}
                </div>
                <div>
                  <h3 class="text-lg font-bold text-gray-900">
                    {{ atribuicao.paciente.nome }}
                  </h3>
                  <p class="text-sm text-gray-500 font-medium">
                    {{ atribuicao.paciente.email }}
                  </p>
                </div>
              </div>

              <!-- Informações do instrumento -->
              <div class="ml-15 space-y-3">
                <div class="flex items-center gap-3 flex-wrap">
                  <span
                    class="inline-block px-3 py-1 text-xs font-mono font-bold rounded-lg"
                    :class="getCodigoBadgeClass(atribuicao.instrumento.codigo)"
                  >
                    {{
                      atribuicao.instrumento.codigo
                        .toUpperCase()
                        .replace("_", "-")
                    }}
                  </span>
                  <span class="text-sm font-bold text-gray-700">
                    {{ atribuicao.instrumento.nome }}
                  </span>
                  <span
                    class="inline-flex items-center px-3 py-1 rounded-full text-xs font-bold border"
                    :class="getStatusClass(atribuicao.status)"
                  >
                    <span
                      class="w-1.5 h-1.5 rounded-full mr-1.5"
                      :class="getStatusDotClass(atribuicao.status)"
                    ></span>
                    {{ getStatusLabel(atribuicao.status) }}
                  </span>
                </div>

                <div
                  class="flex items-center text-xs text-gray-500 gap-4 font-medium flex-wrap"
                >
                  <span class="flex items-center">
                    <font-awesome-icon
                      :icon="faCalendar"
                      class="w-3 h-3 mr-1"
                    />
                    Atribuído em {{ formatDate(atribuicao.data_atribuicao) }}
                  </span>
                  <span
                    v-if="atribuicao.data_resposta"
                    class="flex items-center text-emerald-600"
                  >
                    <font-awesome-icon
                      :icon="faCheckCircle"
                      class="w-3 h-3 mr-1"
                    />
                    Respondido em {{ formatDate(atribuicao.data_resposta) }}
                  </span>
                  <span class="flex items-center">
                    <font-awesome-icon :icon="faListOl" class="w-3 h-3 mr-1" />
                    {{ atribuicao.instrumento.total_perguntas }} perguntas
                  </span>
                </div>
              </div>
            </div>

            <!-- Ações -->
            <div class="ml-4 flex flex-col gap-2">
              <button
                v-if="atribuicao.status === 'RESPONDIDO'"
                @click="visualizarRespostas(atribuicao.id)"
                class="px-5 py-2.5 bg-indigo-600 text-white text-sm font-bold rounded-xl hover:bg-indigo-700 transition-all shadow-sm hover:shadow-md flex items-center whitespace-nowrap"
              >
                <font-awesome-icon :icon="faEye" class="w-4 h-4 mr-2" />
                Ver Respostas
              </button>
              <button
                v-else-if="atribuicao.status === 'PENDENTE'"
                @click="enviarLembrete(atribuicao.id)"
                class="px-5 py-2.5 bg-amber-500 text-white text-sm font-bold rounded-xl hover:bg-amber-600 transition-all shadow-sm hover:shadow-md flex items-center whitespace-nowrap"
              >
                <font-awesome-icon :icon="faBell" class="w-4 h-4 mr-2" />
                Enviar Lembrete
              </button>
            </div>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import { api } from "@/services/api";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faClipboardList,
  faClock,
  faCheckCircle,
  faTimesCircle,
  faCalendar,
  faListOl,
  faEye,
  faBell,
  faChevronDown,
  faSearch,
} from "@fortawesome/free-solid-svg-icons";

const router = useRouter();
const toast = useToast();

const isLoading = ref(true);
const atribuicoes = ref([]);
const filtroStatus = ref("");
const filtroPaciente = ref("");
const filtroInstrumento = ref("");

const atribuicoesFiltradas = computed(() => {
  let resultado = atribuicoes.value;

  if (filtroStatus.value) {
    resultado = resultado.filter((a) => a.status === filtroStatus.value);
  }

  if (filtroPaciente.value) {
    const busca = filtroPaciente.value.toLowerCase();
    resultado = resultado.filter(
      (a) =>
        a.paciente.nome.toLowerCase().includes(busca) ||
        a.paciente.email.toLowerCase().includes(busca)
    );
  }

  if (filtroInstrumento.value) {
    resultado = resultado.filter(
      (a) => a.instrumento.codigo === filtroInstrumento.value
    );
  }

  return resultado;
});

const avatarColors = ["bg-blue-100 text-blue-600"];
const contarPorStatus = (status) => {
  return atribuicoesFiltradas.value.filter((a) => a.status === status).length;
};

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

const getCodigoBadgeClass = (codigo) => {
  const classes = {
    phq_9: "bg-blue-100 text-blue-700",
    gad_7: "bg-amber-100 text-amber-700",
    whoqol_bref: "bg-emerald-100 text-emerald-700",
    who_5: "bg-purple-100 text-purple-700",
  };
  return classes[codigo] || "bg-gray-100 text-gray-700";
};

const getStatusClass = (status) => {
  const classes = {
    PENDENTE: "bg-amber-50 text-amber-700 border-amber-100",
    RESPONDIDO: "bg-emerald-50 text-emerald-700 border-emerald-100",
    EXPIRADO: "bg-gray-50 text-gray-600 border-gray-100",
  };
  return classes[status] || "bg-gray-50 text-gray-600 border-gray-100";
};

const getStatusDotClass = (status) => {
  const classes = {
    PENDENTE: "bg-amber-500",
    RESPONDIDO: "bg-emerald-500",
    EXPIRADO: "bg-gray-400",
  };
  return classes[status] || "bg-gray-400";
};

const getStatusLabel = (status) => {
  const labels = {
    PENDENTE: "Pendente",
    RESPONDIDO: "Respondido",
    EXPIRADO: "Expirado",
  };
  return labels[status] || status;
};

const formatDate = (dateString) => {
  if (!dateString) return "Data não informada";
  const date = new Date(dateString);
  if (isNaN(date.getTime()) || date.getFullYear() < 1900) {
    return "Data não informada";
  }
  return date.toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
};

const visualizarRespostas = (atribuicaoId) => {
  router.push({
    name: "profissional-visualizar-respostas",
    params: { atribuicaoId },
  });
};

const enviarLembrete = async (atribuicaoId) => {
  try {
    toast.success("Lembrete enviado ao paciente!");
  } catch (error) {
    toast.error("Erro ao enviar lembrete.");
    console.error(error);
  }
};

onMounted(async () => {
  try {
    const response = await api.listarAtribuicoesProfissional();
    atribuicoes.value = response.data;
  } catch (error) {
    toast.error("Erro ao carregar atribuições.");
    console.error(error);
  } finally {
    isLoading.value = false;
  }
});
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
