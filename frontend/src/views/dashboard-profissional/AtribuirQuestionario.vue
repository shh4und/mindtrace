<template>
  <div class="max-w-6xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center md:text-left">
      <div class="flex items-center mb-6">
        <button
          @click="voltar"
          class="mr-4 p-2.5 rounded-xl hover:bg-gray-100 transition-colors"
          aria-label="Voltar para lista de pacientes"
        >
          <font-awesome-icon :icon="faArrowLeft" class="w-5 h-5 text-gray-600" />
        </button>
        <div>
          <div
            class="inline-flex items-center justify-center space-x-2 bg-indigo-50 text-indigo-800 px-4 py-1.5 rounded-full text-sm font-medium mb-3 shadow-sm"
          >
            <font-awesome-icon :icon="faClipboardList" class="h-4 w-4" />
            <span>Atribuição</span>
          </div>
          <h1
            class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
          >
            Atribuir Questionário
          </h1>
          <p class="text-gray-500 mt-2 text-lg">
            Selecione um instrumento para atribuir ao paciente.
          </p>
        </div>
      </div>
    </header>

    <!-- Info do paciente -->
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-8">
      <div class="flex items-center gap-4">
        <div class="w-14 h-14 bg-indigo-100 rounded-2xl flex items-center justify-center shadow-sm">
          <font-awesome-icon :icon="faUser" class="w-6 h-6 text-indigo-600" />
        </div>
        <div>
          <p class="text-sm text-indigo-600 font-semibold">Atribuindo para:</p>
          <p class="text-xl font-extrabold text-gray-900">{{ pacienteNome }}</p>
        </div>
      </div>
    </section>

    <!-- Grid de instrumentos -->
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
      <section
        v-for="instrumento in instrumentos"
        :key="instrumento.codigo"
        class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 hover:shadow-lg hover:border-indigo-100 transition-all duration-300 group"
      >
        <div class="flex items-start justify-between mb-4">
          <div class="flex-1">
            <span
              class="inline-block px-3 py-1 text-xs font-mono font-bold rounded-lg mb-2"
              :class="getCodigoBadgeClass(instrumento.codigo)"
            >
              {{ instrumento.codigo.toUpperCase().replace('_', '-') }}
            </span>
            <h3 class="text-lg font-bold text-gray-900 group-hover:text-indigo-600 transition-colors">
              {{ instrumento.nome }}
            </h3>
          </div>
          <div
            class="w-12 h-12 rounded-xl flex items-center justify-center shadow-sm transition-transform group-hover:scale-110"
            :class="getIconBgClass(instrumento.codigo)"
          >
            <font-awesome-icon :icon="faClipboardList" class="w-6 h-6 text-white" />
          </div>
        </div>

        <p class="text-gray-500 text-sm mb-6 line-clamp-3 leading-relaxed">
          {{ instrumento.descricao }}
        </p>

        <button
          @click="atribuirQuestionario(instrumento)"
          class="w-full py-3 px-4 bg-indigo-50 text-indigo-600 font-bold text-sm rounded-xl hover:bg-indigo-600 hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md"
        >
          <font-awesome-icon :icon="faPaperPlane" class="w-4 h-4" />
          Atribuir
        </button>
      </section>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from '@/services/api';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faArrowLeft,
  faUser,
  faClipboardList,
  faPaperPlane,
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const route = useRoute();
const toast = useToast();

const patientId = computed(() => route.params.patientId);
const pacienteNome = computed(() => route.params.patientNome);

const instrumentos = ref([]);

onMounted(async () => {
  try {
    const response = await api.listarQuestionarios();

    if (Array.isArray(response.data) && response.data.length > 0) {
      instrumentos.value = response.data;
    } else {
      toast.warning('Houve um problema ao carregar ou não há questionários cadastrados.');
    }
  } catch (error) {
    toast.error('Erro ao carregar os questionários.');
    console.error(error);
  }
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

const getIconBgClass = (codigo) => {
  const classes = {
    phq_9: 'bg-blue-500',
    gad_7: 'bg-amber-500',
    whoqol_bref: 'bg-emerald-500',
    who_5: 'bg-purple-500',
  };
  return classes[codigo] || 'bg-gray-500';
};

const atribuirQuestionario = async (instrumento) => {
  const response = await api.atribuirQuestionario(patientId.value, instrumento.id, instrumento.codigo);
  toast.success(`Questionário "${instrumento.nome}" atribuído com sucesso!`);
};

const voltar = () => {
  router.push({ name: 'profissional-pacientes' });
};
</script>

<style scoped>
.line-clamp-3 {
  display: -webkit-box;
  -webkit-line-clamp: 3;
  line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
