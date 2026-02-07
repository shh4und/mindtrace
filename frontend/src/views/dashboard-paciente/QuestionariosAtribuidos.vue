<template>
  <div class="max-w-4xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center md:text-left">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-emerald-50 text-emerald-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <font-awesome-icon :icon="faClipboardList" class="h-4 w-4" />
        <span>Avaliações</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Meus Questionários
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Responda os questionários atribuídos pelo seu profissional.
      </p>
    </header>

    <!-- Tabs -->
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 overflow-hidden mb-8">
      <div class="flex border-b border-gray-100">
        <button
          v-for="tab in tabs"
          :key="tab.id"
          @click="activeTab = tab.id"
          class="flex-1 px-6 py-4 font-semibold text-sm transition-all relative text-center"
          :class="[
            activeTab === tab.id
              ? 'text-emerald-700 bg-emerald-50/50'
              : 'text-gray-500 hover:text-gray-700 hover:bg-gray-50',
          ]"
        >
          {{ tab.label }}
          <span
            class="ml-2 inline-flex items-center justify-center w-6 h-6 rounded-full text-xs font-bold"
            :class="[
              activeTab === tab.id
                ? 'bg-emerald-100 text-emerald-700'
                : 'bg-gray-100 text-gray-500',
            ]"
          >
            {{ tab.count }}
          </span>
          <span
            v-if="activeTab === tab.id"
            class="absolute bottom-0 left-0 right-0 h-0.5 bg-emerald-600"
          ></span>
        </button>
      </div>

      <!-- Conteúdo das abas -->
      <div class="p-6 md:p-8">
        <!-- Loading -->
        <div v-if="isLoading" class="flex justify-center items-center py-16">
          <div
            class="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-600"
          ></div>
        </div>

        <!-- Aba Pendentes -->
        <div v-else-if="activeTab === 'pendentes'">
          <div v-if="pendencias.length === 0" class="text-center py-16">
            <div
              class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4 text-gray-300"
            >
              <font-awesome-icon :icon="faClipboardCheck" class="w-8 h-8" />
            </div>
            <h3 class="text-lg font-bold text-gray-900 mb-2">
              Nenhum questionário pendente
            </h3>
            <p class="text-gray-500">
              Você não possui questionários para responder no momento.
            </p>
          </div>

          <div v-else class="space-y-4">
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
          <div v-if="respondidos.length === 0" class="text-center py-16">
            <div
              class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-4 text-gray-300"
            >
              <font-awesome-icon :icon="faClipboardCheck" class="w-8 h-8" />
            </div>
            <h3 class="text-lg font-bold text-gray-900 mb-2">
              Nenhum questionário respondido
            </h3>
            <p class="text-gray-500">
              Você ainda não respondeu nenhum questionário.
            </p>
          </div>

          <div v-else class="space-y-4">
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
    </section>
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
  faClipboardList,
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

// Lifecycle
onMounted(async () => {
  try {
    const response = await api.listarAtribuicoesPaciente();
    allQuestions.value = response.data || [];
    isLoading.value = false;
  } catch (error) {
    toast.error("Erro ao carregar seus questionários.");
    console.error(error);
    isLoading.value = false;
  }
});
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
