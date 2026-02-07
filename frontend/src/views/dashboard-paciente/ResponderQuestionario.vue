<template>
  <!-- Loading state -->
  <div v-if="carregando" class="max-w-4xl mx-auto p-4 md:p-8">
    <div class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-600"
      ></div>
    </div>
  </div>

  <!-- Error state -->
  <div v-else-if="erro" class="max-w-4xl mx-auto p-4 md:p-8">
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 text-center">
      <div
        class="w-16 h-16 bg-red-50 rounded-full flex items-center justify-center mx-auto mb-6 text-red-400"
      >
        <font-awesome-icon :icon="faExclamationTriangle" class="w-8 h-8" />
      </div>
      <h2 class="text-xl font-bold text-gray-900 mb-2">
        Erro ao carregar questionário
      </h2>
      <p class="text-gray-500 mb-8 max-w-md mx-auto">
        {{ erro }}
      </p>
      <button
        class="inline-flex items-center px-6 py-3 bg-emerald-600 text-white font-bold rounded-xl hover:bg-emerald-700 transition-all shadow-md hover:shadow-lg"
        @click="router.push({ name: 'paciente-questionarios' })"
      >
        Voltar para questionários
      </button>
    </section>
  </div>

  <!-- Content -->
  <div v-else class="max-w-4xl mx-auto p-4 md:p-8">
    <!-- Header fixo -->
    <section
      class="sticky top-0 z-10 bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-8"
    >
      <div class="flex items-center mb-5">
        <button
          class="mr-4 p-2.5 rounded-xl hover:bg-gray-100 transition-colors"
          aria-label="Voltar para questionários"
          @click="confirmarSaida"
        >
          <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
        </button>
        <div class="flex-1">
          <span
            class="inline-block px-3 py-1 text-xs font-mono font-bold rounded-lg mb-1"
            :class="getCodigoBadgeClass(instrumento.codigo)"
          >
            {{ instrumento.codigo?.toUpperCase().replace('_', '-') || 'N/A' }}
          </span>
          <h1
            class="text-2xl md:text-3xl font-extrabold text-gray-900 tracking-tight"
          >
            {{ instrumento.nome }}
          </h1>
        </div>
      </div>

      <!-- Barra de progresso -->
      <div class="bg-gray-100 rounded-full h-2.5 overflow-hidden">
        <div
          class="bg-emerald-500 h-2.5 rounded-full transition-all duration-300"
          :style="{ width: `${progresso}%` }"
        />
      </div>
      <p class="text-sm text-gray-500 mt-3 text-center font-medium">
        {{ Object.keys(respostas).length }} de {{ instrumento.perguntas?.length || 0 }} perguntas
        respondidas
      </p>
    </section>

    <!-- Todas as perguntas em scroll -->
    <div class="space-y-6">
      <section
        v-for="(pergunta, index) in instrumento.perguntas"
        :id="`pergunta-${pergunta.pergunta_id}`"
        :key="pergunta.pergunta_id"
        class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 transition-all hover:shadow-md"
        :class="
          respostas[pergunta.pergunta_id] !== undefined ? 'border-emerald-200 bg-emerald-50/30' : ''
        "
      >
        <div class="mb-5">
          <div class="flex items-center gap-2 mb-2">
            <span class="text-xs font-bold text-gray-400 uppercase tracking-wide">
              Pergunta {{ index + 1 }}
            </span>
            <span
              v-if="respostas[pergunta.pergunta_id] !== undefined"
              class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-bold bg-emerald-100 text-emerald-700"
            >
              <font-awesome-icon :icon="faCheckCircle" class="w-3 h-3 mr-1" />
              Respondida
            </span>
          </div>
          <p class="text-lg font-bold text-gray-900">
            {{ pergunta.conteudo }}
          </p>
        </div>

        <!-- Opções de resposta (Likert) -->
        <div class="space-y-3">
          <div
            v-for="opcao in instrumento.opcoes_escala"
            :key="`${pergunta.pergunta_id}-${opcao.valor}`"
            class="flex items-center p-4 border-2 rounded-2xl cursor-pointer transition-all duration-200"
            :class="
              respostas[pergunta.pergunta_id]?.[0] === opcao.valor
                ? 'border-emerald-500 bg-emerald-50 shadow-sm'
                : 'border-gray-100 hover:border-emerald-200 hover:bg-gray-50'
            "
            @click="selecionarResposta(pergunta.pergunta_id, opcao.valor, pergunta.dominio)"
          >
            <span
              class="w-9 h-9 rounded-xl border-2 flex items-center justify-center mr-4 transition-all shrink-0 text-sm font-bold"
              :class="
                respostas[pergunta.pergunta_id]?.[0] === opcao.valor
                  ? 'border-emerald-500 bg-emerald-500 text-white'
                  : 'border-gray-300 text-gray-500'
              "
            >
              {{ opcao.valor }}
            </span>
            <span class="flex-1 text-sm font-medium text-gray-700">{{ opcao.rotulo }}</span>
            <font-awesome-icon
              v-if="respostas[pergunta.pergunta_id]?.[0] === opcao.valor"
              :icon="faCheckCircle"
              class="w-5 h-5 text-emerald-500 shrink-0"
            />
          </div>
        </div>
      </section>
    </div>

    <!-- Botão fixo de envio -->
    <div
      class="fixed bottom-0 left-0 right-0 p-4 bg-white/80 backdrop-blur-md border-t border-gray-200 md:static md:bg-transparent md:border-none md:p-0 md:mt-8 z-10 md:flex md:justify-end"
    >
      <button
        :disabled="!todasRespondidas"
        class="w-full md:w-auto bg-gradient-to-r from-emerald-600 to-teal-600 hover:from-emerald-700 hover:to-teal-700 text-white font-bold py-4 px-10 rounded-2xl shadow-lg hover:shadow-xl hover:-translate-y-1 transition-all duration-300 disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-y-0 disabled:hover:shadow-lg flex items-center justify-center space-x-2"
        @click="enviarRespostas"
      >
        <font-awesome-icon :icon="faPaperPlane" class="w-5 h-5 mr-2" />
        <span v-if="todasRespondidas">Enviar Respostas</span>
        <span v-else>Responda todas as perguntas para enviar</span>
      </button>
    </div>
    <!-- Spacer for mobile fixed bottom button -->
    <div class="h-24 md:h-0"></div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { api } from '@/services/api';
import { useRouter, useRoute } from 'vue-router';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faArrowLeft,
  faCheckCircle,
  faPaperPlane,
  faExclamationTriangle,
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const toast = useToast();

const atribuicaoId = computed(() => route.params.atribuicaoId);
const respostas = ref({});
const carregando = ref(true);
const erro = ref(null);

// Dados do instrumento
const instrumento = ref({
  codigo: '',
  nome: '',
  perguntas: [],
  opcoes_escala: [],
});

const progresso = computed(() => {
  const respondidas = Object.keys(respostas.value).length;
  return Math.round((respondidas / instrumento.value.perguntas.length) * 100);
});

const todasRespondidas = computed(() => {
  return instrumento.value.perguntas.every((p) => respostas.value[p.pergunta_id] !== undefined);
});

const getCodigoBadgeClass = (codigo) => {
  const classes = {
    phq_9: 'bg-blue-100 text-blue-700',
    gad_7: 'bg-amber-100 text-amber-700',
    whoqol_bref: 'bg-emerald-100 text-emerald-700',
    who_5: 'bg-purple-100 text-purple-700',
  };
  return classes[codigo] || 'bg-gray-100 text-gray-700';
};

const selecionarResposta = (perguntaId, valor, perguntaDominio) => {
  respostas.value = {
    ...respostas.value,
    [perguntaId]: [valor, perguntaDominio],
  };
};

const enviarRespostas = async () => {
  const payload = {
    atribuicao_id: parseInt(atribuicaoId.value),
    respostas: Object.entries(respostas.value).map(([perguntaId, dadosResposta]) => ({
      pergunta_id: parseInt(perguntaId),
      valor: dadosResposta[0],
      dominio: dadosResposta[1],
    })),
  };

  await api.enviarResposta(payload);

  toast.success('Questionário enviado com sucesso! Obrigado por responder.');

  setTimeout(() => {
    router.push({ name: 'paciente-questionarios' });
  }, 1500);
};

const confirmarSaida = () => {
  const respondidas = Object.keys(respostas.value).length;
  if (respondidas > 0) {
    if (confirm('Você tem respostas não salvas. Deseja realmente sair?')) {
      router.push({ name: 'paciente-questionarios' });
    }
  } else {
    router.push({ name: 'paciente-questionarios' });
  }
};

onMounted(async () => {
  try {
    carregando.value = true;
    const response = await api.buscarAtribuicao(atribuicaoId.value);

    if (!response.data || !response.data.instrumento) {
      throw new Error('Dados do questionário inválidos');
    }

    instrumento.value = response.data.instrumento;
  } catch (error) {
    console.error('Erro ao carregar questionário:', error);
    erro.value = error.response?.data?.erro || 'Não foi possível carregar o questionário';
    toast.error(erro.value);
  } finally {
    carregando.value = false;
  }
});
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
