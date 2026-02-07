<template>
  <div class="max-w-4xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center md:text-left">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-emerald-50 text-emerald-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
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
        Resumo do Bem-Estar
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Bem-vindo(a) de volta! Acompanhe seus progressos.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-600"
      ></div>
    </div>

    <template v-else>
      <!-- Acesso Rápido -->
      <section class="mb-8">
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 relative overflow-hidden"
        >
          <div
            class="absolute top-0 left-0 w-full h-1.5 bg-gradient-to-r from-amber-400 via-orange-400 to-red-400"
          ></div>
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="p-2 bg-amber-100 rounded-lg text-amber-600 mr-3">
              <font-awesome-icon :icon="faBolt" class="w-5 h-5" />
            </span>
            Ações Rápidas
          </h2>
          <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
            <button
              v-for="action in acoesRapidas"
              :key="action.id"
              @click="navigateTo(action.route)"
              class="group flex flex-col items-center justify-center p-6 rounded-2xl border-2 border-dashed border-gray-200 hover:border-emerald-400 hover:bg-emerald-50 transition-all duration-300"
            >
              <div
                class="w-12 h-12 rounded-full bg-emerald-100 text-emerald-600 flex items-center justify-center mb-3 group-hover:scale-110 transition-transform"
              >
                <font-awesome-icon :icon="action.icon" class="w-6 h-6" />
              </div>
              <span
                class="text-sm font-semibold text-gray-700 group-hover:text-emerald-700"
                >{{ action.label }}</span
              >
            </button>
          </div>
        </div>
      </section>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
        <!-- Card: Último Registro de Humor -->
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-lg transition-all duration-300 group"
        >
          <header class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-bold text-gray-800 flex items-center gap-3">
              <span class="p-2 bg-indigo-100 rounded-lg text-indigo-600">
                <font-awesome-icon :icon="faFaceSmile" class="w-5 h-5" />
              </span>
              Último Registro
            </h2>
            <span
              v-if="userLastData.humor"
              class="text-xs font-semibold px-3 py-1 bg-gray-100 text-gray-500 rounded-full"
            >
              {{ formatDate(userLastData.data) }}
            </span>
          </header>

          <div v-if="userLastData.humor" class="flex flex-col items-center">
            <div
              class="transform transition-transform duration-300 group-hover:scale-110 mb-4"
            >
              <Emoji
                :data="emojiIndex"
                :emoji="userLastData.emoji"
                :size="64"
                set="facebook"
              />
            </div>
            <h3 class="text-2xl font-extrabold text-indigo-600 mb-2">
              {{ userLastData.label }}
            </h3>
            <div
              v-if="userLastData.anotacao"
              class="w-full bg-gray-50 rounded-xl p-4 mt-2 text-gray-600 text-sm italic text-center border border-gray-100"
            >
              "{{ userLastData.anotacao }}"
            </div>
          </div>

          <div v-else class="text-center py-8">
            <div
              class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4 text-gray-400"
            >
              <font-awesome-icon :icon="faFaceSmile" class="w-8 h-8" />
            </div>
            <p class="text-gray-500 mb-4 font-medium">
              Nenhum registro encontrado.
            </p>
            <router-link
              to="/dashboard-paciente/humor"
              class="inline-flex items-center justify-center px-6 py-2.5 border border-transparent text-sm font-medium rounded-xl text-white bg-indigo-600 hover:bg-indigo-700 shadow-md hover:shadow-lg transition-all"
            >
              Registrar Humor
            </router-link>
          </div>
        </div>

        <!-- Card: Notificações -->
        <div
          class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-lg transition-all duration-300"
        >
          <header class="flex items-center justify-between mb-6">
            <h2 class="text-xl font-bold text-gray-800 flex items-center gap-3">
              <span class="p-2 bg-rose-100 rounded-lg text-rose-600">
                <font-awesome-icon :icon="faBell" class="w-5 h-5" />
              </span>
              Notificações
            </h2>
          </header>

          <div class="flex flex-col items-center justify-center h-48 text-center">
            <div
              class="w-16 h-16 bg-rose-50 rounded-full flex items-center justify-center mb-4 text-rose-300"
            >
              <font-awesome-icon :icon="faBellSlash" class="w-8 h-8" />
            </div>
            <p class="text-gray-900 font-medium">Tudo tranquilo por aqui!</p>
            <p class="text-gray-500 text-sm mt-1 max-w-xs">
              Você não tem novas notificações pendentes no momento.
            </p>
          </div>
        </div>
      </div>

      <!-- Card: Atividades Recentes -->
      <section
        class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 relative overflow-hidden"
      >
        <div
          class="absolute left-0 top-0 bottom-0 w-1.5 bg-gradient-to-b from-emerald-400 to-teal-500"
        ></div>
        <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
          <span class="p-2 bg-emerald-100 rounded-lg text-emerald-600 mr-3">
            <font-awesome-icon :icon="faClock" class="w-5 h-5" />
          </span>
          Linha do Tempo Recente
        </h2>

        <div v-if="atividadesRecentes.length === 0" class="text-center py-10">
          <font-awesome-icon
            :icon="faListCheck"
            class="w-12 h-12 text-gray-200 mb-3"
          />
          <p class="text-gray-500 font-medium">Nenhuma atividade recente.</p>
          <p class="text-gray-400 text-sm mt-1">
            Suas atividades aparecerão aqui.
          </p>
        </div>

        <div v-else class="space-y-4">
          <div
            v-for="(atividade, index) in atividadesRecentes"
            :key="index"
            class="flex items-center p-4 rounded-2xl border border-gray-100 hover:border-emerald-200 hover:shadow-md hover:bg-emerald-50/30 transition-all duration-300 cursor-default bg-gray-50/50"
          >
            <div
              :class="[
                'shrink-0 flex items-center justify-center w-12 h-12 rounded-xl mr-4 shadow-sm',
                atividade.tipo === 'humor'
                  ? 'bg-emerald-100 text-emerald-600'
                  : 'bg-blue-100 text-blue-600',
              ]"
            >
              <font-awesome-icon
                :icon="
                  atividade.tipo === 'humor' ? faFaceSmile : faClipboardList
                "
                class="w-5 h-5"
              />
            </div>
            <div class="flex-1 min-w-0">
              <h4 class="font-bold text-gray-900 truncate">
                {{ atividade.titulo }}
              </h4>
              <p class="text-sm text-gray-500 truncate">
                {{ atividade.descricao }}
              </p>
            </div>
            <div class="flex flex-col items-end ml-4">
              <span class="text-xs font-semibold text-gray-400">{{
                atividade.dataFormatada
              }}</span>
              <div v-if="atividade.tipo === 'humor'" class="mt-1">
                <Emoji
                  :data="emojiIndex"
                  :emoji="atividade.emoji"
                  :size="20"
                  set="facebook"
                />
              </div>
            </div>
          </div>
        </div>
      </section>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import api from "@/services/api";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faFaceSmile,
  faBell,
  faBellSlash,
  faClock,
  faClipboardList,
  faListCheck,
  faBolt,
  faUserDoctor,
} from "@fortawesome/free-solid-svg-icons";

import data from "emoji-mart-vue-fast/data/all.json";
import { EmojiIndex, Emoji } from "emoji-mart-vue-fast/src";
import "emoji-mart-vue-fast/css/emoji-mart.css";

const emojiIndex = new EmojiIndex(data);
const router = useRouter();
const toast = useToast();

// --- ESTADO DO COMPONENTE ---
const isLoading = ref(true);
const userLastData = ref({});
const atividadesRecentes = ref([]);

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

const moodOptions = [
  { label: "Muito Mal", emoji: "confounded" },
  { label: "Aborrecido", emoji: "confused" },
  { label: "Neutro", emoji: "neutral_face" },
  { label: "Animado", emoji: "blush" },
  { label: "Muito Bem", emoji: "grin" },
];

// Ações rápidas
const acoesRapidas = [
  {
    id: "humor",
    label: "Registrar Humor",
    icon: faFaceSmile,
    route: "paciente-humor",
  },
  {
    id: "questionario",
    label: "Questionários",
    icon: faClipboardList,
    route: "paciente-questionarios",
  },
  {
    id: "profissionais",
    label: "Profissionais",
    icon: faUserDoctor,
    route: "paciente-profissionais",
  },
];

const navigateTo = (routeName) => {
  router.push({ name: routeName });
};

// --- FUNÇÕES AUXILIARES ---
const formatDate = (dateString) => {
  if (!dateString) return "";
  return new Date(dateString).toLocaleDateString("pt-BR", {
    weekday: "long",
    day: "numeric",
    month: "long",
  });
};

const formatDateShort = (dateString) => {
  if (!dateString) return "";
  return new Date(dateString).toLocaleDateString("pt-BR", {
    day: "numeric",
    month: "short",
    hour: "2-digit",
    minute: "2-digit",
  });
};

// --- LÓGICA DE BUSCA DE DADOS ---
const fetchSummaryData = async () => {
  isLoading.value = true;
  try {
    // Buscar resumo (último registro de humor)
    const summaryResponse = await api.buscarResumo();
    const summary = summaryResponse.data;

    if (summary && summary.humor) {
      summary.emoji = moodOptions[summary.humor - 1]?.emoji || "neutral_face";
      summary.label = moodOptions[summary.humor - 1]?.label || "Neutro";
      userLastData.value = summary;
    }

    // Buscar atividades recentes (mock + dados reais combinados)
    await fetchAtividadesRecentes(summary);
  } catch (error) {
    // Silently fail or minimal toast if critical
    console.error(error);
  } finally {
    isLoading.value = false;
  }
};

const fetchAtividadesRecentes = async (summaryData) => {
  const atividades = [];

  // Adicionar último registro de humor como atividade
  if (summaryData && summaryData.humor) {
    atividades.push({
      tipo: "humor",
      titulo: "Registro de Humor",
      descricao: `Você estava se sentindo ${
        moodOptions[summaryData.humor - 1]?.label.toLowerCase() || "neutro"
      }`,
      data: summaryData.data,
      dataFormatada: formatDateShort(summaryData.data),
      emoji: moodOptions[summaryData.humor - 1]?.emoji || "neutral_face",
    });
  }

  // Tentar buscar questionários respondidos
  try {
    const atribuicoesResponse = await api.listarAtribuicoesPaciente();
    const atribuicoes = atribuicoesResponse.data || [];

    // Filtrar apenas respondidos e pegar os mais recentes
    const respondidos = atribuicoes
      .filter((a) => a.respondido_em)
      .sort((a, b) => new Date(b.respondido_em) - new Date(a.respondido_em))
      .slice(0, 2);

    respondidos.forEach((q) => {
      atividades.push({
        tipo: "questionario",
        titulo: "Questionário Respondido",
        descricao: q.instrumento?.nome || "Avaliação Concluída",
        data: q.respondido_em,
        dataFormatada: formatDateShort(q.respondido_em),
      });
    });
  } catch (error) {
    console.log("Questionários não disponíveis:", error);
  }

  // Ordenar por data e limitar a 3
  atividadesRecentes.value = atividades
    .sort((a, b) => new Date(b.data) - new Date(a.data))
    .slice(0, 3);
};

onMounted(fetchSummaryData);
</script>

<style scoped>
/* Mantendo consistência com o estilo global se necessário */
</style>