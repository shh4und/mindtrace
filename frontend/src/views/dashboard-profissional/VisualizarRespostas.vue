<template>
  <!-- Loading state -->
  <div v-if="carregando" class="max-w-3xl mx-auto pb-8 pt-8">
    <div class="flex flex-col items-center justify-center min-h-[60vh]">
      <div
        class="animate-spin rounded-full h-16 w-16 border-b-2 border-indigo-600 mb-4"
      />
      <p class="text-gray-600">Carregando respostas...</p>
    </div>
  </div>

  <!-- Error state -->
  <div v-else-if="erro" class="max-w-3xl mx-auto pb-8 pt-8">
    <div class="bg-red-50 border border-red-200 rounded-lg p-6 text-center">
      <font-awesome-icon
        :icon="faExclamationTriangle"
        class="w-12 h-12 text-red-500 mb-4"
      />
      <h2 class="text-xl font-semibold text-red-800 mb-2">
        Erro ao carregar respostas
      </h2>
      <p class="text-red-600 mb-4">
        {{ erro }}
      </p>
      <button
        class="px-6 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors"
        @click="voltar"
      >
        Voltar para questionários
      </button>
    </div>
  </div>

  <!-- Content -->
  <div v-else class="max-w-3xl mx-auto pb-8">
    <!-- Header -->
    <div class="pb-4 mb-6 border-b border-gray-200">
      <div class="flex items-center mb-4 pt-4">
        <button
          class="mr-4 p-2 rounded-lg hover:bg-gray-100 transition-colors"
          aria-label="Voltar para questionários atribuídos"
          @click="voltar"
        >
          <font-awesome-icon
            :icon="faArrowLeft"
            class="w-5 h-5 text-gray-600"
          />
        </button>
        <div class="flex-1">
          <span
            class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md mb-1"
            :class="getCodigoBadgeClass(resposta.instrumento?.codigo)"
          >
            {{
              resposta.instrumento?.codigo?.toUpperCase().replace("_", "-") ||
              "N/A"
            }}
          </span>
          <h1 class="text-2xl font-bold text-gray-900">
            {{ resposta.instrumento?.nome || "Respostas do Questionário" }}
          </h1>
        </div>
        <span
          class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium bg-emerald-100 text-emerald-700"
        >
          <font-awesome-icon :icon="faCheckCircle" class="w-4 h-4 mr-2" />
          Respondido
        </span>
      </div>
    </div>

    <!-- Cards de informação -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
      <!-- Card Paciente -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-5">
        <div class="flex items-center mb-3">
          <div
            class="w-10 h-10 bg-indigo-100 rounded-full flex items-center justify-center mr-3"
          >
            <font-awesome-icon :icon="faUser" class="w-5 h-5 text-indigo-600" />
          </div>
          <div>
            <p class="text-xs text-gray-500 uppercase tracking-wide">
              Paciente
            </p>
            <p class="font-semibold text-gray-900">
              {{ resposta.paciente?.nome }}
            </p>
          </div>
        </div>
        <p class="text-sm text-gray-500 ml-13">
          <font-awesome-icon :icon="faEnvelope" class="w-4 h-4 mr-1" />
          {{ resposta.paciente?.email }}
        </p>
      </div>

      <!-- Card Datas -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-5">
        <div class="space-y-3">
          <div class="flex items-center">
            <font-awesome-icon
              :icon="faCalendarPlus"
              class="w-4 h-4 text-gray-400 mr-2"
            />
            <span class="text-sm text-gray-500">Atribuído em:</span>
            <span class="ml-2 text-sm font-medium text-gray-900">{{
              formatDate(resposta.data_atribuicao)
            }}</span>
          </div>
          <div class="flex items-center">
            <font-awesome-icon
              :icon="faCalendarCheck"
              class="w-4 h-4 text-emerald-500 mr-2"
            />
            <span class="text-sm text-gray-500">Respondido em:</span>
            <span class="ml-2 text-sm font-medium text-emerald-700">{{
              formatDate(resposta.data_resposta)
            }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Resumo de respostas -->
    <div class="bg-indigo-50 border border-indigo-200 rounded-lg p-4 mb-6">
      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <font-awesome-icon
            :icon="faClipboardList"
            class="w-5 h-5 text-indigo-600 mr-2"
          />
          <span class="font-medium text-indigo-900"
            >Total de perguntas respondidas</span
          >
        </div>
        <span class="text-lg font-bold text-indigo-700">
          {{ resposta.instrumento?.perguntas?.length || 0 }}
        </span>
      </div>
    </div>

    <!-- Lista de perguntas com respostas -->
    <h2 class="text-lg font-semibold text-gray-900 mb-4">
      Respostas detalhadas
    </h2>
    <div class="space-y-4">
      <div
        v-for="(pergunta, index) in perguntasOrdenadas"
        :key="pergunta.pergunta_id"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-5"
      >
        <div class="mb-3">
          <span
            class="text-xs font-medium text-gray-400 uppercase tracking-wide"
          >
            Pergunta {{ index + 1 }}
          </span>
          <p class="text-base font-medium text-gray-900 mt-1">
            {{ pergunta.conteudo }}
          </p>
        </div>

        <!-- Resposta selecionada -->
        <div class="mt-3 p-3 bg-indigo-50 rounded-lg border border-indigo-100">
          <div class="flex items-center">
            <span
              class="w-8 h-8 rounded-full bg-indigo-600 text-white flex items-center justify-center mr-3 text-sm font-semibold"
            >
              {{ getRespostaValor(pergunta.pergunta_id) }}
            </span>
            <span class="text-sm font-medium text-indigo-900">
              {{ getRespostaRotulo(pergunta.pergunta_id) }}
            </span>
          </div>
        </div>

        <!-- Todas as opções (para contexto) -->
        <div class="mt-3 pt-3 border-t border-gray-100">
          <p class="text-xs text-gray-400 mb-2">Todas as opções:</p>
          <div class="flex flex-wrap gap-2">
            <span
              v-for="opcao in resposta.instrumento?.opcoes_escala"
              :key="opcao.valor"
              class="inline-flex items-center px-2 py-1 rounded text-xs"
              :class="
                getRespostaValor(pergunta.pergunta_id) === opcao.valor
                  ? 'bg-indigo-100 text-indigo-700 font-medium'
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

    <!-- Botão voltar no final -->
    <div class="mt-8 text-center">
      <button
        @click="voltar"
        class="px-6 py-3 bg-gray-100 text-gray-700 font-medium rounded-lg hover:bg-gray-200 transition-colors"
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
