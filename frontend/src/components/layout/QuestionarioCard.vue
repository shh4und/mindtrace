<template>
  <div
    class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 hover:shadow-md transition-all hover:border-indigo-100 group"
  >
    <div class="flex items-start justify-between gap-4">
      <!-- Conteúdo Principal -->
      <div class="flex-1 min-w-0">
        <!-- Header com badges -->
        <div class="flex items-center gap-3 mb-3 flex-wrap">
          <span
            class="inline-block px-2 py-1 text-xs font-mono font-medium rounded-md shrink-0"
            :class="getCodigoBadgeClass(pendencia.instrumento.codigo)"
          >
            {{ pendencia.instrumento.codigo.toUpperCase().replace("_", "-") }}
          </span>
          <span
            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium shrink-0"
            :class="getStatusClass(pendencia.status)"
          >
            <span
              class="w-1.5 h-1.5 rounded-full mr-1.5"
              :class="getStatusDotClass(pendencia.status)"
            ></span>
            {{ getStatusLabel(pendencia.status) }}
          </span>
        </div>

        <!-- Título e Descrição -->
        <h3 class="text-lg font-semibold text-gray-900 mb-1 line-clamp-1">
          {{ pendencia.instrumento.nome }}
        </h3>
        <p class="text-sm text-gray-500 mb-4 line-clamp-2">
          {{ pendencia.instrumento.descricao }}
        </p>

        <!-- Metadados -->
        <div
          class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 text-xs text-gray-500"
        >
          <div class="flex items-center gap-2 min-w-0">
            <font-awesome-icon
              :icon="faUserDoctor"
              class="w-3.5 h-3.5 text-indigo-400 shrink-0"
            />
            <span class="truncate">
              <strong class="text-gray-900">{{
                pendencia.profissional.nome
              }}</strong>
            </span>
          </div>

          <div class="flex items-center gap-2 min-w-0">
            <font-awesome-icon
              :icon="faCalendar"
              class="w-3.5 h-3.5 text-indigo-400 shrink-0"
            />
            <span class="truncate">{{
              formatDate(pendencia.data_atribuicao)
            }}</span>
          </div>

          <div class="flex items-center gap-2 min-w-0">
            <font-awesome-icon
              :icon="faListOl"
              class="w-3.5 h-3.5 text-indigo-400 shrink-0"
            />
            <span>{{ pendencia.instrumento.total_perguntas }} perguntas</span>
          </div>

          <div v-if="daysRemaining !== null" class="flex items-center gap-2">
            <font-awesome-icon
              :icon="faClock"
              class="w-3.5 h-3.5"
              :class="[daysRemaining <= 3 ? 'text-red-400' : 'text-indigo-400']"
            />
            <span
              :class="[
                daysRemaining <= 3
                  ? 'font-semibold text-red-600'
                  : 'text-gray-500',
              ]"
            >
              {{ daysRemainingLabel }}
            </span>
          </div>
        </div>
      </div>

      <!-- Ação/Status -->
      <div class="shrink-0 ml-4">
        <button
          v-if="!readonly && pendencia.status === 'PENDENTE'"
          @click="handleResponder"
          class="px-4 py-2 bg-indigo-600 text-white text-sm font-medium rounded-lg hover:bg-indigo-700 active:scale-95 transition-all flex items-center gap-2 whitespace-nowrap shadow-sm hover:shadow-md"
        >
          <font-awesome-icon :icon="faPen" class="w-4 h-4" />
          Responder
        </button>

        <span
          v-else-if="pendencia.status === 'RESPONDIDO'"
          class="inline-flex items-center px-4 py-2 text-sm font-medium text-emerald-700 bg-emerald-50 rounded-lg gap-2 whitespace-nowrap"
        >
          <font-awesome-icon :icon="faCheck" class="w-4 h-4" />
          Concluído
        </span>

        <span
          v-else-if="pendencia.status === 'EXPIRADO'"
          class="inline-flex items-center px-4 py-2 text-sm font-medium text-gray-500 bg-gray-100 rounded-lg gap-2 whitespace-nowrap"
        >
          <font-awesome-icon :icon="faClock" class="w-4 h-4" />
          Expirado
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faUserDoctor,
  faCalendar,
  faListOl,
  faClock,
  faPen,
  faCheck,
} from "@fortawesome/free-solid-svg-icons";

const props = defineProps({
  pendencia: {
    type: Object,
    required: true,
  },
  readonly: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["responder"]);

// Computados
const daysRemaining = computed(() => {
  if (props.readonly || props.pendencia.status !== "PENDENTE") {
    return null;
  }

  if (!props.pendencia.data_expiracao) {
    return null;
  }

  const today = new Date();
  today.setHours(0, 0, 0, 0);

  const expirationDate = new Date(props.pendencia.data_expiracao);
  expirationDate.setHours(0, 0, 0, 0);

  const diffTime = expirationDate - today;
  const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

  return diffDays;
});

const daysRemainingLabel = computed(() => {
  if (daysRemaining.value === null) return "";
  if (daysRemaining.value < 0) return "Expirado";
  if (daysRemaining.value === 0) return "Hoje";
  if (daysRemaining.value === 1) return "Amanhã";
  return `${daysRemaining.value} dias`;
});

// Métodos
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
    PENDENTE: "bg-yellow-50 text-yellow-800",
    RESPONDIDO: "bg-emerald-50 text-emerald-800",
    EXPIRADO: "bg-gray-100 text-gray-600",
  };
  return classes[status] || "bg-gray-100 text-gray-600";
};

const getStatusDotClass = (status) => {
  const classes = {
    PENDENTE: "bg-yellow-500",
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
  const date = new Date(dateString);
  return date.toLocaleDateString("pt-BR", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
};

const handleResponder = () => {
  emit("responder", props.pendencia.id);
};
</script>

<style scoped>
.line-clamp-1 {
  display: -webkit-box;
  -webkit-line-clamp: 1;
  line-clamp: 1;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
