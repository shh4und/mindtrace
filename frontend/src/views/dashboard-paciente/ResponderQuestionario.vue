<template>
  <div class="max-w-3xl mx-auto">
    <!-- Header com progresso -->
    <div class="mb-8">
      <div class="flex items-center mb-4">
        <button 
          @click="confirmarSaida" 
          class="mr-4 p-2 rounded-lg hover:bg-gray-100 transition-colors"
          aria-label="Voltar para questionários"
        >
          <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
        </button>
        <div class="flex-1">
          <span 
            class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md mb-1"
            :class="getCodigoBadgeClass(instrumento.codigo)"
          >
            {{ instrumento.codigo.toUpperCase().replace('_', '-') }}
          </span>
          <h1 class="text-2xl font-bold text-gray-900">{{ instrumento.nome }}</h1>
        </div>
      </div>

      <!-- Barra de progresso -->
      <div class="bg-gray-100 rounded-full h-2 overflow-hidden">
        <div 
          class="bg-indigo-600 h-2 rounded-full transition-all duration-300"
          :style="{ width: `${progresso}%` }"
        ></div>
      </div>
      <p class="text-sm text-gray-500 mt-2 text-center">
        Pergunta {{ perguntaAtualIndex + 1 }} de {{ instrumento.perguntas.length }}
      </p>
    </div>

    <!-- Instrução -->
    <div class="bg-blue-50 border border-blue-200 rounded-lg p-4 mb-6">
      <p class="text-sm text-blue-800">
        <font-awesome-icon :icon="faInfoCircle" class="mr-2" />
        {{ instrumento.instrucao }}
      </p>
    </div>

    <!-- Pergunta atual -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 mb-6">
      <div class="mb-6">
        <span class="text-xs font-medium text-gray-400 uppercase tracking-wide">
          Pergunta {{ perguntaAtualIndex + 1 }}
        </span>
        <p class="text-lg font-medium text-gray-900 mt-2">
          {{ perguntaAtual.conteudo }}
        </p>
      </div>

      <!-- Opções de resposta (Likert) -->
      <div class="space-y-3">
        <label 
          v-for="opcao in instrumento.opcoesEscala" 
          :key="opcao.valor"
          class="flex items-center p-4 border rounded-lg cursor-pointer transition-all"
          :class="respostas[perguntaAtual.id] === opcao.valor 
            ? 'border-indigo-500 bg-indigo-50 ring-2 ring-indigo-500' 
            : 'border-gray-200 hover:border-gray-300 hover:bg-gray-50'"
        >
          <input 
            type="radio" 
            :name="`pergunta-${perguntaAtual.id}`"
            :value="opcao.valor"
            v-model="respostas[perguntaAtual.id]"
            class="sr-only"
          />
          <span 
            class="w-8 h-8 rounded-full border-2 flex items-center justify-center mr-4 transition-all"
            :class="respostas[perguntaAtual.id] === opcao.valor 
              ? 'border-indigo-500 bg-indigo-500 text-white' 
              : 'border-gray-300 text-gray-500'"
          >
            {{ opcao.valor }}
          </span>
          <span class="flex-1 text-gray-700">{{ opcao.rotulo }}</span>
          <font-awesome-icon 
            v-if="respostas[perguntaAtual.id] === opcao.valor"
            :icon="faCheckCircle" 
            class="w-5 h-5 text-indigo-500" 
          />
        </label>
      </div>
    </div>

    <!-- Navegação -->
    <div class="flex items-center justify-between">
      <button 
        @click="perguntaAnterior"
        :disabled="perguntaAtualIndex === 0"
        class="px-4 py-2 text-gray-600 font-medium rounded-lg hover:bg-gray-100 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
      >
        <font-awesome-icon :icon="faChevronLeft" class="w-4 h-4 mr-2" />
        Anterior
      </button>

      <button 
        v-if="perguntaAtualIndex < instrumento.perguntas.length - 1"
        @click="proximaPergunta"
        :disabled="respostas[perguntaAtual.id] === undefined"
        class="px-6 py-2 bg-indigo-600 text-white font-medium rounded-lg hover:bg-indigo-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
      >
        Próxima
        <font-awesome-icon :icon="faChevronRight" class="w-4 h-4 ml-2" />
      </button>

      <button 
        v-else
        @click="enviarRespostas"
        :disabled="!todasRespondidas"
        class="px-6 py-2 bg-emerald-600 text-white font-medium rounded-lg hover:bg-emerald-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
      >
        <font-awesome-icon :icon="faPaperPlane" class="w-4 h-4 mr-2" />
        Enviar Respostas
      </button>
    </div>

    <!-- Indicadores de perguntas -->
    <div class="flex justify-center gap-2 mt-8">
      <button
        v-for="(pergunta, index) in instrumento.perguntas"
        :key="pergunta.id"
        @click="irParaPergunta(index)"
        class="w-3 h-3 rounded-full transition-all"
        :class="getIndicadorClass(index)"
        :aria-label="`Ir para pergunta ${index + 1}`"
      ></button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faArrowLeft, 
  faInfoCircle, 
  faCheckCircle, 
  faChevronLeft, 
  faChevronRight, 
  faPaperPlane 
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const toast = useToast();

const atribuicaoId = computed(() => route.params.atribuicaoId);
const perguntaAtualIndex = ref(0);
const respostas = ref({});

// Mock: dados do instrumento PHQ-9
const instrumento = ref({
  codigo: 'phq_9',
  nome: 'Patient Health Questionnaire-9',
  instrucao: 'Nas últimas 2 semanas, com que frequência você foi incomodado(a) por qualquer um dos problemas abaixo?',
  opcoesEscala: [
    { valor: 0, rotulo: 'Nenhuma vez' },
    { valor: 1, rotulo: 'Vários dias' },
    { valor: 2, rotulo: 'Mais da metade dos dias' },
    { valor: 3, rotulo: 'Quase todos os dias' }
  ],
  perguntas: [
    { id: 1, ordemItem: 1, conteudo: 'Pouco interesse ou pouco prazer em fazer as coisas' },
    { id: 2, ordemItem: 2, conteudo: 'Se sentir "para baixo", deprimido(a) ou sem perspectiva' },
    { id: 3, ordemItem: 3, conteudo: 'Dificuldade para pegar no sono ou permanecer dormindo, ou dormir mais do que de costume' },
    { id: 4, ordemItem: 4, conteudo: 'Se sentir cansado(a) ou com pouca energia' },
    { id: 5, ordemItem: 5, conteudo: 'Falta de apetite ou comendo demais' },
    { id: 6, ordemItem: 6, conteudo: 'Se sentir mal consigo mesmo(a) — ou achar que você é um fracasso ou que decepcionou sua família ou você mesmo(a)' },
    { id: 7, ordemItem: 7, conteudo: 'Dificuldade para se concentrar nas coisas, como ler o jornal ou ver televisão' },
    { id: 8, ordemItem: 8, conteudo: 'Lentidão para se movimentar ou falar, a ponto das outras pessoas perceberem? Ou o oposto — Loss inquieto(a) ou agitado(a) a ponto de você ficar andando de um lado para o outro muito mais do que de costume' },
    { id: 9, ordemItem: 9, conteudo: 'Pensar em se ferir de alguma maneira ou que seria melhor estar morto(a)' }
  ]
});

const perguntaAtual = computed(() => instrumento.value.perguntas[perguntaAtualIndex.value]);

const progresso = computed(() => {
  const respondidas = Object.keys(respostas.value).length;
  return Math.round((respondidas / instrumento.value.perguntas.length) * 100);
});

const todasRespondidas = computed(() => {
  return instrumento.value.perguntas.every(p => respostas.value[p.id] !== undefined);
});

const getCodigoBadgeClass = (codigo) => {
  const classes = {
    'phq_9': 'bg-blue-100 text-blue-700',
    'gad_7': 'bg-amber-100 text-amber-700',
    'whoqol_bref': 'bg-emerald-100 text-emerald-700',
    'who_5': 'bg-purple-100 text-purple-700'
  };
  return classes[codigo] || 'bg-gray-100 text-gray-700';
};

const getIndicadorClass = (index) => {
  const pergunta = instrumento.value.perguntas[index];
  if (index === perguntaAtualIndex.value) {
    return 'bg-indigo-600 scale-125';
  }
  if (respostas.value[pergunta.id] !== undefined) {
    return 'bg-emerald-500';
  }
  return 'bg-gray-300';
};

const proximaPergunta = () => {
  if (perguntaAtualIndex.value < instrumento.value.perguntas.length - 1) {
    perguntaAtualIndex.value++;
  }
};

const perguntaAnterior = () => {
  if (perguntaAtualIndex.value > 0) {
    perguntaAtualIndex.value--;
  }
};

const irParaPergunta = (index) => {
  perguntaAtualIndex.value = index;
};

const enviarRespostas = () => {
  // TODO: Integrar com API quando backend estiver pronto
  // const payload = {
  //   atribuicaoId: atribuicaoId.value,
  //   respostas: Object.entries(respostas.value).map(([perguntaId, valor]) => ({
  //     perguntaId: parseInt(perguntaId),
  //     valor
  //   }))
  // };
  // await api.enviarResposta(payload);

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

onMounted(() => {
  // TODO: Buscar dados do instrumento baseado na atribuição
  // const response = await api.buscarAtribuicao(atribuicaoId.value);
  // instrumento.value = response.data.instrumento;
});
</script>
