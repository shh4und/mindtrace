<template>
  <div class="max-w-7xl mx-auto">
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Meus Questionários</h1>
      <p class="text-gray-500">
        Responda os questionários atribuídos pelo seu profissional
      </p>
    </div>

    <!-- Tabs -->
    <div class="flex gap-2 mb-6 border-b border-gray-200">
      <button
        v-for="tab in tabs"
        :key="tab.id"
        @click="activeTab = tab.id"
        class="px-4 py-3 font-medium text-sm transition-all relative"
        :class="[
          activeTab === tab.id
            ? 'text-indigo-600'
            : 'text-gray-500 hover:text-gray-700',
        ]"
      >
        {{ tab.label }}
        <span
          class="ml-2 inline-flex items-center justify-center w-5 h-5 rounded-full text-xs font-medium"
          :class="[
            activeTab === tab.id
              ? 'bg-indigo-100 text-indigo-600'
              : 'bg-gray-200 text-gray-600',
          ]"
        >
          {{ tab.count }}
        </span>
        <span
          v-if="activeTab === tab.id"
          class="absolute bottom-0 left-0 right-0 h-0.5 bg-indigo-600"
        ></span>
      </button>
    </div>

    <!-- Conteúdo das abas -->
    <div v-if="isLoading" class="text-center py-12">
      <div
        class="animate-spin rounded-full h-10 w-10 border-b-2 border-indigo-600 mx-auto mb-4"
      ></div>
      <p class="text-gray-500">Carregando questionários...</p>
    </div>

    <!-- Aba Pendentes -->
    <div v-else-if="activeTab === 'pendentes'">
      <div v-if="pendencias.length === 0" class="text-center py-12">
        <div
          class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4"
        >
          <font-awesome-icon
            :icon="faClipboardCheck"
            class="w-8 h-8 text-gray-400"
          />
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">
          Nenhum questionário pendente
        </h3>
        <p class="text-gray-500">
          Você não possui questionários para responder no momento.
        </p>
      </div>

      <div v-else class="space-y-4 animate-fadeIn">
        <QuestionarioCard
          v-for="pendencia in sortedPendencias"
          :key="pendencia.id"
          :pendencia="pendencia"
          @responder="responderQuestionario"
        />
      </div>
    </div>

    <!-- Aba Respondidos -->
    <div v-else-if="activeTab === 'respondidos'">
      <div v-if="respondidos.length === 0" class="text-center py-12">
        <div
          class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4"
        >
          <font-awesome-icon
            :icon="faClipboardCheck"
            class="w-8 h-8 text-gray-400"
          />
        </div>
        <h3 class="text-lg font-medium text-gray-900 mb-2">
          Nenhum questionário respondido
        </h3>
        <p class="text-gray-500">
          Você ainda não respondeu nenhum questionário.
        </p>
      </div>

      <div v-else class="space-y-4 animate-fadeIn">
        <QuestionarioCard
          v-for="respondido in sortedRespondidos"
          :key="respondido.id"
          :pendencia="respondido"
          :readonly="true"
          @responder="responderQuestionario"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { api } from "@/services/api";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import QuestionarioCard from "@/components/layout/QuestionarioCard.vue";
import {
  faClipboardCheck,
  faCalendar,
  faListOl,
  faPen,
  faCheck,
  faClock,
  faUserDoctor,
} from "@fortawesome/free-solid-svg-icons";

const router = useRouter();
const isLoading = ref(true);
const toast = useToast();
const activeTab = ref("pendentes");

// Dados
const allQuestions = ref([]);

// Computed
const pendencias = computed(() =>
  allQuestions.value.filter((q) => q.status === "PENDENTE")
);

const respondidos = computed(() =>
  allQuestions.value.filter((q) => q.status === "RESPONDIDO")
);
// Ordenar por data de atribuição (mais recentes primeiro)
const sortedPendencias = computed(() =>
  [...pendencias.value].sort(
    (a, b) => new Date(b.data_atribuicao) - new Date(a.data_atribuicao)
  )
);

const sortedRespondidos = computed(() =>
  [...respondidos.value].sort(
    (a, b) => new Date(b.data_atribuicao) - new Date(a.data_atribuicao)
  )
);

// Abas
const tabs = computed(() => [
  {
    id: "pendentes",
    label: "Pendentes",
    count: pendencias.value.length,
  },
  {
    id: "respondidos",
    label: "Respondidos",
    count: respondidos.value.length,
  },
]);

// Métodos
const responderQuestionario = (atribuicaoId) => {
  router.push({
    name: "paciente-responder-questionario",
    params: { atribuicaoId },
  });
};

const formatDate = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
};

// Lifecycle
onMounted(async () => {
  try {
    const response = await api.listarAtribuicoesPaciente();
    allQuestions.value = response.data || [];
    isLoading.value = false;
    toast.success("Seus questionários carregados com sucesso.");
  } catch (error) {
    toast.error("Erro ao carregar seus questionários.");
    console.error(error);
    isLoading.value = false;
  }
});
</script>

<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fadeIn {
  animation: fadeIn 0.3s ease-out;
}
</style>
