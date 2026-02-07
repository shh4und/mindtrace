<template>
  <!-- Loading state -->
  <div v-if="carregando" class="max-w-6xl mx-auto p-4 md:p-8">
    <div class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"
      ></div>
    </div>
  </div>

  <!-- Error state -->
  <div v-else-if="erro" class="max-w-6xl mx-auto p-4 md:p-8">
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 text-center">
      <div
        class="w-16 h-16 bg-red-50 rounded-full flex items-center justify-center mx-auto mb-6 text-red-400"
      >
        <font-awesome-icon :icon="faExclamationTriangle" class="w-8 h-8" />
      </div>
      <h2 class="text-xl font-bold text-gray-900 mb-2">
        Erro ao carregar respostas
      </h2>
      <p class="text-gray-500 mb-8 max-w-md mx-auto">
        {{ erro }}
      </p>
      <button
        class="inline-flex items-center px-6 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition-all shadow-md hover:shadow-lg"
        @click="voltar"
      >
        Voltar para questionários
      </button>
    </section>
  </div>

  <!-- Content -->
  <div v-else class="max-w-6xl mx-auto p-4 md:p-8">
    <!-- Header -->
    <header class="mb-10">
      <div class="flex items-center mb-6">
        <button
          class="mr-4 p-2.5 rounded-xl hover:bg-gray-100 transition-colors"
          aria-label="Voltar para questionários atribuídos"
          @click="voltar"
        >
          <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
        </button>
        <div class="flex-1">
          <div class="flex items-center gap-3 mb-2">
            <span
              class="inline-block px-3 py-1 text-xs font-mono font-bold rounded-lg"
              :class="getCodigoBadgeClass(resposta.instrumento?.codigo)"
            >
              {{ resposta.instrumento?.codigo?.toUpperCase().replace('_', '-') || 'N/A' }}
            </span>
            <span
              class="inline-flex items-center px-3 py-1 rounded-full text-xs font-bold bg-emerald-50 text-emerald-700 border border-emerald-100"
            >
              <font-awesome-icon :icon="faCheckCircle" class="w-3 h-3 mr-1" />
              Respondido
            </span>
          </div>
          <h1
            class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
          >
            {{ resposta.instrumento?.nome || 'Respostas do Questionário' }}
          </h1>
        </div>
      </div>
    </header>

    <!-- Alerta de Ideação Suicida (PHQ-9) -->
    <div
      v-if="temIdeacaoSuicida"
      class="mb-8 bg-red-50 rounded-3xl border-2 border-red-200 p-6 md:p-8"
    >
      <div class="flex items-start gap-4">
        <div class="p-3 bg-red-100 rounded-xl text-red-600 shrink-0">
          <svg class="h-6 w-6" viewBox="0 0 20 20" fill="currentColor">
            <path
              fill-rule="evenodd"
              d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
        <div>
          <p class="text-lg font-bold text-red-800">
            ⚠️ Ideação suicida presente
          </p>
          <p class="mt-1 text-sm text-red-700">
            O paciente indicou pensamentos de se ferir ou que seria melhor estar morto.
          </p>
        </div>
      </div>
    </div>

    <!-- Cards de informação -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
      <!-- Card Paciente -->
      <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-indigo-100 rounded-xl text-indigo-600">
            <font-awesome-icon :icon="faUser" class="w-6 h-6" />
          </div>
          <div>
            <p class="text-xs text-gray-500 uppercase tracking-wide font-bold">
              Paciente
            </p>
            <p class="text-lg font-extrabold text-gray-900">
              {{ resposta.paciente?.nome }}
            </p>
            <p class="text-sm text-gray-500 font-medium mt-1">
              <font-awesome-icon :icon="faEnvelope" class="w-3 h-3 mr-1" />
              {{ resposta.paciente?.email }}
            </p>
          </div>
        </div>
      </section>

      <!-- Card Datas -->
      <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow">
        <div class="space-y-4">
          <div class="flex items-center gap-3">
            <div class="p-2 bg-gray-100 rounded-lg text-gray-500">
              <font-awesome-icon :icon="faCalendarPlus" class="w-4 h-4" />
            </div>
            <div>
              <span class="text-sm text-gray-500 font-medium">Atribuído em:</span>
              <span class="ml-2 text-sm font-bold text-gray-900">{{ formatDate(resposta.data_atribuicao) }}</span>
            </div>
          </div>
          <div class="flex items-center gap-3">
            <div class="p-2 bg-emerald-100 rounded-lg text-emerald-600">
              <font-awesome-icon :icon="faCalendarCheck" class="w-4 h-4" />
            </div>
            <div>
              <span class="text-sm text-gray-500 font-medium">Respondido em:</span>
              <span class="ml-2 text-sm font-bold text-emerald-700">{{ formatDate(resposta.data_resposta) }}</span>
            </div>
          </div>
        </div>
      </section>
    </div>

    <!-- Score e Classificação -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
      <!-- Card de Pontuação -->
      <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow">
        <p class="text-xs text-gray-500 uppercase tracking-wide font-bold mb-2">
          Pontuação Total
        </p>
        <p class="text-5xl font-extrabold text-indigo-600 mb-2">
          {{ resposta.pontuacao_total?.toFixed(1) || '0' }}
        </p>
        <p class="text-sm text-gray-600 font-medium">
          <span class="font-bold text-indigo-900">Classificação:</span>
          {{ resposta.classificacao || 'Não classificado' }}
        </p>
      </section>

      <!-- Card de Resumo -->
      <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-md transition-shadow">
        <p class="text-xs text-gray-500 uppercase tracking-wide font-bold mb-4">
          Resumo da Resposta
        </p>
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600 font-medium">Perguntas respondidas:</span>
            <span class="font-bold text-gray-900">
              {{ resposta.instrumento?.perguntas?.length || 0 }}/{{ resposta.instrumento?.perguntas?.length || 0 }}
            </span>
          </div>
          <div class="flex justify-between items-center">
            <span class="text-sm text-gray-600 font-medium">Status:</span>
            <span
              class="inline-flex items-center px-3 py-1 rounded-full text-xs font-bold bg-emerald-50 text-emerald-700 border border-emerald-100"
            >
              <font-awesome-icon :icon="faCheckCircle" class="w-3 h-3 mr-1" />
              Respondido
            </span>
          </div>
        </div>
      </section>
    </div>

    <!-- Escala Comparativa (PHQ-9, GAD-7, WHO-5) -->
    <section
      v-if="['phq_9', 'gad_7', 'who_5'].includes(resposta.instrumento?.codigo)"
      class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-8 hover:shadow-md transition-shadow"
    >
      <h3 class="text-xl font-bold text-gray-800 mb-6 flex items-center gap-2">
        <span class="p-2 bg-indigo-100 rounded-lg text-indigo-600">
          <font-awesome-icon :icon="faClipboardList" class="w-5 h-5" />
        </span>
        Escala de Referência
      </h3>
      <div class="space-y-3">
        <div
          v-for="range in obterEscalaReferencia(resposta.instrumento?.codigo)"
          :key="`${range.min}-${range.max}`"
          class="flex items-center p-4 rounded-2xl border-2 transition-all"
          :class="
            isScoreInRange(resposta.pontuacao_total, range)
              ? 'border-indigo-500 bg-indigo-50'
              : 'border-gray-100 bg-gray-50'
          "
        >
          <div
            class="flex-shrink-0 w-20 font-mono text-sm font-bold"
            :class="
              isScoreInRange(resposta.pontuacao_total, range)
                ? 'text-indigo-700'
                : 'text-gray-600'
            "
          >
            {{ formatarFaixa(range) }}
          </div>
          <div
            class="flex-1 ml-4 text-sm font-bold"
            :class="
              isScoreInRange(resposta.pontuacao_total, range)
                ? 'text-indigo-900'
                : 'text-gray-700'
            "
          >
            {{ range.label }}
          </div>
          <div
            v-if="isScoreInRange(resposta.pontuacao_total, range)"
            class="flex-shrink-0 ml-2"
          >
            <span
              class="inline-flex items-center px-4 py-1.5 rounded-full text-xs font-bold bg-indigo-600 text-white shadow-sm"
            >
              Score: {{ resposta.pontuacao_total?.toFixed(1) || '0' }}
            </span>
          </div>
        </div>
      </div>
    </section>

    <!-- Domínios WHOQOL -->
    <div
      v-if="resposta.instrumento?.codigo === 'whoqol_bref'"
      class="grid grid-cols-2 md:grid-cols-4 gap-4 mb-8"
    >
      <section
        v-for="(score, dominio) in resposta.detalhes"
        :key="dominio"
        class="bg-white rounded-3xl shadow-sm border border-gray-100 p-5 text-center hover:shadow-md transition-shadow"
      >
        <p class="text-xs text-gray-500 uppercase tracking-wide font-bold mb-2">
          {{ formatarDominio(dominio) }}
        </p>
        <p class="text-3xl font-extrabold text-emerald-600 mb-2">
          {{ score?.toFixed(1) || '0' }}
        </p>
        <div class="w-full bg-gray-200 rounded-full h-2.5 overflow-hidden">
          <div
            class="bg-emerald-500 h-2.5 rounded-full transition-all"
            :style="{ width: `${(score / 100) * 100}%` }"
          />
        </div>
      </section>
    </div>

    <!-- Resumo de respostas -->
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-8">
      <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center gap-2">
        <span class="p-2 bg-purple-100 rounded-lg text-purple-600">
          <font-awesome-icon :icon="faClipboardList" class="w-5 h-5" />
        </span>
        Respostas Detalhadas
      </h2>

      <!-- Lista de perguntas com respostas -->
      <div class="space-y-4">
        <div
          v-for="(pergunta, index) in perguntasOrdenadas"
          :key="pergunta.pergunta_id"
          class="bg-gray-50 rounded-2xl border border-gray-100 p-5 hover:bg-gray-100/50 transition-colors"
        >
          <div class="mb-3">
            <span class="text-xs font-bold text-gray-400 uppercase tracking-wide">
              Pergunta {{ index + 1 }}
            </span>
            <p class="text-base font-bold text-gray-900 mt-1">
              {{ pergunta.conteudo }}
            </p>
          </div>

          <!-- Resposta selecionada -->
          <div class="mt-3 p-4 bg-indigo-50 rounded-xl border border-indigo-100">
            <div class="flex items-center">
              <span
                class="w-9 h-9 rounded-xl bg-indigo-600 text-white flex items-center justify-center mr-3 text-sm font-bold shadow-sm"
              >
                {{ getRespostaValor(pergunta.pergunta_id) }}
              </span>
              <span class="text-sm font-bold text-indigo-900">
                {{ getRespostaRotulo(pergunta.pergunta_id) }}
              </span>
            </div>
          </div>

          <!-- Todas as opções (para contexto) -->
          <div class="mt-3 pt-3 border-t border-gray-200">
            <p class="text-xs text-gray-400 mb-2 font-medium">Todas as opções:</p>
            <div class="flex flex-wrap gap-2">
              <span
                v-for="opcao in resposta.instrumento?.opcoes_escala"
                :key="opcao.valor"
                class="inline-flex items-center px-3 py-1 rounded-lg text-xs font-medium"
                :class="
                  getRespostaValor(pergunta.pergunta_id) === opcao.valor
                    ? 'bg-indigo-100 text-indigo-700 font-bold'
                    : 'bg-gray-100 text-gray-500'
                "
              >
                <span class="font-mono mr-1">{{ opcao.valor }}</span>
                {{ opcao.rotulo }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Botão voltar no final -->
    <div class="text-center">
      <button
        @click="voltar"
        class="px-8 py-3.5 bg-gray-100 text-gray-700 font-bold rounded-2xl hover:bg-gray-200 transition-all hover:shadow-md"
      >
        <font-awesome-icon :icon="faArrowLeft" class="w-4 h-4 mr-2" />
        Voltar para questionários atribuídos
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { api } from "@/services/api";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faArrowLeft,
  faCheckCircle,
  faExclamationTriangle,
  faUser,
  faEnvelope,
  faCalendarPlus,
  faCalendarCheck,
  faClipboardList,
} from "@fortawesome/free-solid-svg-icons";

// Escalas de referência hardcoded
const escalasReferencia = {
  phq_9: [
    { min: 0, max: 5, label: "Ausência ou sintomas depressivos mínimos" },
    { min: 5, max: 10, label: "Sintomas depressivos leves" },
    { min: 10, max: 15, label: "Depressão moderada" },
    { min: 15, max: 20, label: "Depressão moderadamente grave" },
    { min: 20, max: 27, label: "Depressão grave" },
  ],
  gad_7: [
    { min: 0, max: 5, label: "Ansiedade mínima" },
    { min: 5, max: 10, label: "Ansiedade leve" },
    { min: 10, max: 15, label: "Ansiedade moderada" },
    { min: 15, max: 21, label: "Ansiedade grave" },
  ],
  who_5: [
    { min: 0, max: 29, label: "Bem-estar muito baixo" },
    { min: 29, max: 50, label: "Bem-estar reduzido" },
    { min: 50, max: 100, label: "Bem-estar preservado" },
  ],
};

// Mapeamento de domínios WHOQOL para label em português
const dominiosWHOQOL = {
  fisico: "Físico",
  psicologico: "Psicológico",
  relacoes_sociais: "Relações Sociais",
  ambiente: "Ambiente",
};

const props = defineProps({
  atribuicaoId: {
    type: String,
    required: true,
  },
});

const router = useRouter();
const toast = useToast();

const carregando = ref(true);
const erro = ref(null);
const resposta = ref({});

// Mapa de respostas para acesso rápido: pergunta_id -> valor
const respostasMap = computed(() => {
  const map = {};
  if (
    resposta.value.perguntas_respostas &&
    Array.isArray(resposta.value.perguntas_respostas)
  ) {
    for (const r of resposta.value.perguntas_respostas) {
      map[r.pergunta_id] = r.valor;
    }
  }
  return map;
});

// Perguntas ordenadas por ordem_item
const perguntasOrdenadas = computed(() => {
  if (!resposta.value.instrumento?.perguntas) return [];
  return [...resposta.value.instrumento.perguntas].sort(
    (a, b) => a.ordem_item - b.ordem_item
  );
});

const getRespostaValor = (perguntaId) => {
  return respostasMap.value[perguntaId] ?? "-";
};

const getRespostaRotulo = (perguntaId) => {
  const valor = respostasMap.value[perguntaId];
  if (valor === undefined || !resposta.value.instrumento?.opcoes_escala)
    return "Não respondida";
  const opcao = resposta.value.instrumento.opcoes_escala.find(
    (o) => o.valor === valor
  );
  return opcao?.rotulo || `Valor: ${valor}`;
};

const getCodigoBadgeClass = (codigo) => {
  const classes = {
    phq_9: "bg-blue-100 text-blue-700",
    gad_7: "bg-amber-100 text-amber-700",
    whoqol_bref: "bg-emerald-100 text-emerald-700",
    who_5: "bg-purple-100 text-purple-700",
  };
  return classes[codigo] || "bg-gray-100 text-gray-700";
};

const formatDate = (dateString) => {
  if (!dateString) return "-";
  const date = new Date(dateString);
  return date.toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
};

const obterEscalaReferencia = (codigo) => {
  return escalasReferencia[codigo] || [];
};

const isScoreInRange = (score, range) => {
  if (score === null || score === undefined) return false;
  if (range.max === null) {
    return score >= range.min;
  }
  return score >= range.min && score < range.max;
};

const formatarFaixa = (range) => {
  if (range.max === null) {
    return `≥ ${Math.floor(range.min)}`;
  }
  return `${Math.floor(range.min)}-${Math.floor(range.max - 0.1)}`;
};

const formatarDominio = (dominio) => {
  return dominiosWHOQOL[dominio] || dominio;
};

const perguntaIdeacaoSuicida = computed(() => {
  if (resposta.value.instrumento?.codigo !== "phq_9") return null;
  return resposta.value.perguntas?.find((p) => p.ordem_item === 8);
});

const respostaIdeacaoSuicida = computed(() => {
  const pergunta = perguntaIdeacaoSuicida.value;
  if (!pergunta) return null;
  const respostaPergunta = resposta.value.perguntas_respostas?.find(
    (r) => r.pergunta_id === pergunta.id
  );
  return respostaPergunta ? respostaPergunta.valor : 0;
});

const temIdeacaoSuicida = computed(() => {
  return respostaIdeacaoSuicida.value > 0;
});

const voltar = () => {
  router.push({ name: "profissional-questionarios-atribuidos" });
};

onMounted(async () => {
  try {
    carregando.value = true;
    const response = await api.visualizarRespostas(props.atribuicaoId);

    if (!response.data) {
      throw new Error("Dados da resposta inválidos");
    }

    resposta.value = response.data;
  } catch (error) {
    console.error("Erro ao carregar respostas:", error);
    erro.value =
      error.response?.data?.erro ||
      "Não foi possível carregar as respostas do questionário";
    toast.error(erro.value);
  } finally {
    carregando.value = false;
  }
});
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
