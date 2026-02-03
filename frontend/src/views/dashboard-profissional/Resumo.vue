<template>
  <div class="max-w-6xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center md:text-left">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-indigo-50 text-indigo-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-4 w-4"
          viewBox="0 0 20 20"
          fill="currentColor"
        >
          <path
            fill-rule="evenodd"
            d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
            clip-rule="evenodd"
          />
        </svg>
        <span>{{ currentDay }}, {{ currentDate }}</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Painel do Profissional
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Visão geral dos seus pacientes e atividades recentes.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"
      ></div>
    </div>

    <template v-else>
      <!-- Cards de Estatísticas -->
      <section class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <div
          v-for="stat in estatisticas"
          :key="stat.id"
          class="bg-white p-6 md:p-8 rounded-3xl shadow-sm border border-gray-100 hover:shadow-lg transition-all duration-300 cursor-pointer group"
          @click="stat.route && navigateTo(stat.route)"
        >
          <div class="flex items-center justify-between mb-4">
            <div
              :class="[
                stat.bgColor,
                stat.iconColor,
                'p-4 rounded-2xl transition-transform group-hover:scale-110',
              ]"
            >
              <font-awesome-icon :icon="stat.icon" class="w-8 h-8" />
            </div>
            <span
              v-if="stat.trend"
              :class="[
                'text-xs font-bold px-3 py-1 rounded-full border',
                stat.trend > 0
                  ? 'bg-emerald-50 text-emerald-700 border-emerald-100'
                  : 'bg-gray-50 text-gray-600 border-gray-100',
              ]"
            >
              {{ stat.trend > 0 ? "+" : "" }}{{ stat.trend }}% este mês
            </span>
          </div>
          <div>
            <p class="text-4xl font-extrabold text-gray-900 mb-1">
              {{ stat.value }}
            </p>
            <p class="text-gray-500 font-medium">{{ stat.label }}</p>
          </div>
          <div
            v-if="stat.details"
            class="mt-6 pt-4 border-t border-gray-100 flex gap-4"
          >
            <div
              v-for="(detail, idx) in stat.details"
              :key="idx"
              class="flex items-center gap-2"
            >
              <span
                class="w-3 h-3 rounded-full shadow-sm"
                :class="detail.colorClass"
              ></span>
              <span class="text-sm font-semibold text-gray-600">{{
                detail.label
              }}</span>
            </div>
          </div>
        </div>
      </section>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-8">
        <!-- Coluna Esquerda: Ações Rápidas e Pacientes (2/3 da largura em telas grandes) -->
        <div class="lg:col-span-2 space-y-8">
          <!-- Ações Rápidas -->
          <section
            class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 relative overflow-hidden"
          >
            <div
              class="absolute top-0 left-0 w-full h-1.5 bg-gradient-to-r from-amber-400 via-orange-400 to-rose-400"
            ></div>
            <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
              <span class="p-2 bg-amber-100 rounded-lg text-amber-600 mr-3">
                <font-awesome-icon :icon="faBolt" class="w-5 h-5" />
              </span>
              Ações Rápidas
            </h2>
            <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
              <button
                v-for="action in acoesRapidas"
                :key="action.id"
                @click="navigateTo(action.route)"
                class="group flex flex-col items-center justify-center p-4 rounded-2xl border-2 border-dashed border-gray-200 hover:border-indigo-400 hover:bg-indigo-50 transition-all duration-300"
              >
                <div
                  class="w-10 h-10 rounded-full bg-indigo-100 text-indigo-600 flex items-center justify-center mb-3 group-hover:scale-110 transition-transform"
                >
                  <font-awesome-icon :icon="action.icon" class="w-5 h-5" />
                </div>
                <span
                  class="text-xs font-bold text-gray-700 group-hover:text-indigo-700 text-center"
                  >{{ action.label }}</span
                >
              </button>
            </div>
          </section>

          <!-- Questionários Pendentes (Tabela) -->
          <section
            class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8"
          >
            <div class="flex items-center justify-between mb-6">
              <h2 class="text-xl font-bold text-gray-800 flex items-center">
                <span class="p-2 bg-purple-100 rounded-lg text-purple-600 mr-3">
                  <font-awesome-icon :icon="faClipboardList" class="w-5 h-5" />
                </span>
                Questionários Recentes
              </h2>
              <button
                @click="navigateTo('profissional-questionarios-atribuidos')"
                class="text-sm font-bold text-indigo-600 hover:text-indigo-800 transition-colors"
              >
                Ver todos
              </button>
            </div>

            <div
              v-if="questionariosRecentes.length === 0"
              class="text-center py-10"
            >
              <div
                class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4"
              >
                <font-awesome-icon
                  :icon="faClipboardList"
                  class="w-8 h-8 text-gray-300"
                />
              </div>
              <p class="text-gray-500 font-medium">
                Nenhum questionário atribuído recentemente.
              </p>
            </div>

            <div v-else class="overflow-hidden">
              <table class="w-full">
                <thead>
                  <tr
                    class="text-left text-xs font-bold text-gray-400 uppercase tracking-wider border-b border-gray-100"
                  >
                    <th class="pb-4 pl-2">Paciente</th>
                    <th class="pb-4">Questionário</th>
                    <th class="pb-4">Status</th>
                    <th class="pb-4 text-right pr-2">Data</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-50">
                  <tr
                    v-for="q in questionariosRecentes"
                    :key="q.id"
                    class="hover:bg-gray-50 cursor-pointer transition-colors group"
                    @click="q.respondido && viewResponses(q.id)"
                  >
                    <td class="py-4 pl-2">
                      <span class="font-bold text-gray-800 group-hover:text-indigo-600 transition-colors">{{
                        q.pacienteNome
                      }}</span>
                    </td>
                    <td class="py-4 text-gray-600 text-sm font-medium">
                      {{ q.questionarioNome }}
                    </td>
                    <td class="py-4">
                      <span
                        :class="[
                          'px-3 py-1 text-xs font-bold rounded-full border',
                          q.respondido
                            ? 'bg-emerald-50 text-emerald-700 border-emerald-100'
                            : 'bg-amber-50 text-amber-700 border-amber-100',
                        ]"
                      >
                        {{ q.respondido ? "Respondido" : "Pendente" }}
                      </span>
                    </td>
                    <td
                      class="py-4 text-gray-400 text-xs font-semibold text-right pr-2"
                    >
                      {{ q.dataFormatada }}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- Coluna Direita: Últimos Pacientes (1/3 da largura) -->
        <div class="space-y-8">
          <section
            class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 h-full"
          >
            <div class="flex items-center justify-between mb-6">
              <h2 class="text-xl font-bold text-gray-800 flex items-center">
                <span class="p-2 bg-blue-100 rounded-lg text-blue-600 mr-3">
                  <font-awesome-icon :icon="faUsers" class="w-5 h-5" />
                </span>
                Pacientes
              </h2>
              <button
                v-if="ultimosPacientes.length > 0"
                @click="navigateTo('profissional-pacientes')"
                class="w-8 h-8 flex items-center justify-center rounded-full bg-gray-50 text-gray-400 hover:bg-indigo-50 hover:text-indigo-600 transition-colors"
              >
                <font-awesome-icon :icon="faChevronRight" class="w-4 h-4" />
              </button>
            </div>

            <div
              v-if="ultimosPacientes.length === 0"
              class="text-center py-8 flex flex-col items-center justify-center h-64"
            >
              <div
                class="w-16 h-16 bg-blue-50 rounded-full flex items-center justify-center mb-4 text-blue-300"
              >
                <font-awesome-icon :icon="faUserPlus" class="w-8 h-8" />
              </div>
              <p class="text-gray-900 font-bold mb-1">Comece agora</p>
              <p class="text-gray-500 text-sm mb-6 max-w-[200px]">
                Adicione pacientes para começar a monitorar o progresso deles.
              </p>
              <button
                @click="navigateTo('profissional-convite')"
                class="px-6 py-3 bg-indigo-600 hover:bg-indigo-700 text-white font-bold rounded-xl shadow-md transition-all hover:shadow-lg text-sm"
              >
                Gerar Convite
              </button>
            </div>

            <ul v-else class="space-y-4">
              <li
                v-for="(paciente, index) in ultimosPacientes"
                :key="paciente.id"
                @click="viewPatientReport(paciente.id)"
                class="flex items-center p-3 rounded-2xl border border-gray-100 hover:border-blue-200 hover:shadow-md hover:bg-blue-50/30 cursor-pointer transition-all duration-300"
              >
                <div
                  :class="[
                    getAvatarColor(index),
                    'w-12 h-12 rounded-xl flex items-center justify-center shadow-sm font-bold text-lg',
                  ]"
                >
                  {{ paciente.name.charAt(0).toUpperCase() }}
                </div>
                <div class="flex-1 min-w-0 ml-4">
                  <p class="font-bold text-gray-900 truncate">
                    {{ paciente.name }}
                  </p>
                  <p class="text-xs font-semibold text-gray-500">
                    {{ paciente.age }}
                  </p>
                </div>
              </li>
            </ul>

            <button
              v-if="ultimosPacientes.length > 0"
               @click="navigateTo('profissional-convite')"
              class="w-full mt-6 py-3 border-2 border-dashed border-gray-200 rounded-xl text-gray-500 font-bold hover:border-indigo-400 hover:text-indigo-600 hover:bg-indigo-50 transition-all duration-300 flex items-center justify-center gap-2"
            >
              <font-awesome-icon :icon="faUserPlus" />
              Novo Paciente
            </button>
          </section>
        </div>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import api from "@/services/api";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faUsers,
  faUser,
  faUserPlus,
  faClipboardList,
  faEnvelope,
  faChartLine,
  faBolt,
  faChevronRight,
} from "@fortawesome/free-solid-svg-icons";

const router = useRouter();
const toast = useToast();

const isLoading = ref(true);
const estatisticas = ref([]);
const ultimosPacientes = ref([]);
const questionariosRecentes = ref([]);

const today = new Date();
const currentDate = computed(() =>
  today.toLocaleDateString("pt-BR", {
    year: "numeric",
    month: "long",
    day: "numeric",
  })
);
const currentDay = computed(() =>
  today.toLocaleDateString("pt-BR", { weekday: "long" })
);

// Cores para avatares (estilo pastel forte)
const avatarColors = [
  "bg-blue-100 text-blue-600",
  "bg-emerald-100 text-emerald-600",
  "bg-violet-100 text-violet-600",
  "bg-rose-100 text-rose-600",
  "bg-amber-100 text-amber-600",
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

// Ações rápidas
const acoesRapidas = [
  {
    id: "pacientes",
    label: "Pacientes",
    icon: faUsers,
    route: "profissional-pacientes",
  },
  {
    id: "convite",
    label: "Convite",
    icon: faEnvelope,
    route: "profissional-convite",
  },
  {
    id: "questionarios",
    label: "Questionários",
    icon: faClipboardList,
    route: "profissional-questionarios-atribuidos",
  },
  {
    id: "relatorios",
    label: "Relatórios",
    icon: faChartLine,
    route: "profissional-relatorios",
  },
];

// Navegação
const navigateTo = (routeName) => {
  router.push({ name: routeName });
};

const viewPatientReport = (patientId) => {
  router.push({
    name: "profissional-paciente-relatorio",
    params: { patientId },
  });
};

const viewResponses = (atribuicaoId) => {
  router.push({
    name: "profissional-visualizar-respostas",
    params: { atribuicaoId },
  });
};

// Calcular idade
const calculateAge = (birthdate) => {
  if (!birthdate) return "";
  const birthDate = new Date(birthdate);
  const today = new Date();
  let age = today.getFullYear() - birthDate.getFullYear();
  const m = today.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
};

// Formatar data
const formatDate = (dateString) => {
  if (!dateString) return "";
  return new Date(dateString).toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "short",
  });
};

// Buscar dados do resumo
const fetchResumoData = async () => {
  isLoading.value = true;
  try {
    // Buscar pacientes
    let totalPacientes = 0;
    try {
      const pacientesResponse = await api.listarPacientesDoProfissional();
      const pacientes = pacientesResponse.data || [];
      totalPacientes = pacientes.length;

      // Pegar últimos 3 pacientes
      ultimosPacientes.value = pacientes.slice(0, 5).map((p) => ({
        id: p.id,
        name: p.usuario?.nome || "Paciente",
        age: `${calculateAge(p.data_nascimento)} anos`,
      }));
    } catch (e) {
      console.log("Erro ao buscar pacientes:", e);
    }

    // Buscar questionários atribuídos
    let totalQuestionarios = 0;
    let pendentes = 0;
    try {
      const questionariosResponse = await api.listarAtribuicoesProfissional();
      const questionarios = questionariosResponse.data || [];
      totalQuestionarios = questionarios.length;
      pendentes = questionarios.filter((q) => !q.data_resposta).length;
      const respondidos = totalQuestionarios - pendentes;

      // Pegar últimos 5 questionários
      questionariosRecentes.value = questionarios
        .sort(
          (a, b) => new Date(b.data_atribuicao) - new Date(a.data_atribuicao)
        )
        .slice(0, 5)
        .map((q) => ({
          id: q.id,
          pacienteNome: q.paciente?.nome || "Paciente",
          questionarioNome: q.instrumento?.nome || "Questionário",
          respondido: !!q.data_resposta,
          dataFormatada: formatDate(q.data_atribuicao),
        }));

      // Montar estatísticas
      estatisticas.value = [
        {
          id: "pacientes",
          label: "Total de Pacientes",
          value: totalPacientes,
          icon: faUsers,
          bgColor: "bg-blue-100",
          iconColor: "text-blue-600",
          trend: 12, // Mock trend
          route: "profissional-pacientes",
        },
        {
          id: "questionarios",
          label: "Atividades Monitoradas",
          value: totalQuestionarios,
          icon: faClipboardList,
          bgColor: "bg-purple-100",
          iconColor: "text-purple-600",
          route: "profissional-questionarios-atribuidos",
          details: [
            { label: `${pendentes} Pendentes`, colorClass: "bg-amber-500" },
            { label: `${respondidos} Concluídas`, colorClass: "bg-emerald-500" },
          ],
        },
      ];
    } catch (e) {
      console.log("Erro ao buscar questionários:", e);
    }
  } catch (error) {
    console.error("Erro ao carregar resumo:", error);
    toast.error("Erro ao carregar dados do resumo.");
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchResumoData);
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>