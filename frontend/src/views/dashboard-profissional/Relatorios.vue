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
        Relatórios de Pacientes
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Selecione um paciente para analisar o progresso detalhado.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"
      ></div>
    </div>

    <!-- Empty state -->
    <div v-else-if="patients.length === 0" class="text-center py-20 bg-white rounded-3xl border border-dashed border-gray-200">
      <div class="w-20 h-20 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-6 text-gray-300">
         <font-awesome-icon :icon="faUsers" class="w-10 h-10" />
      </div>
      <h3 class="text-xl font-bold text-gray-900 mb-2">
        Nenhum paciente vinculado
      </h3>
      <p class="text-gray-500 mb-8 max-w-md mx-auto">
        Você precisa vincular pacientes à sua conta para visualizar seus relatórios de acompanhamento.
      </p>
      <router-link
        to="/dashboard-profissional/convite"
        class="inline-flex items-center px-6 py-3 bg-indigo-600 text-white font-bold rounded-xl hover:bg-indigo-700 transition-all shadow-md hover:shadow-lg hover:-translate-y-0.5"
      >
        <font-awesome-icon :icon="faEnvelope" class="mr-2" />
        Convidar Paciente
      </router-link>
    </div>

    <!-- Lista de pacientes para seleção -->
    <template v-else>
      <!-- Barra de busca -->
      <div class="mb-8 relative max-w-2xl">
        <font-awesome-icon
          :icon="faSearch"
          class="absolute left-5 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5"
        />
        <input
          v-model="searchQuery"
          type="text"
          placeholder="Buscar paciente por nome..."
          class="w-full pl-14 pr-6 py-4 rounded-2xl border border-gray-200 bg-white shadow-sm focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-700 placeholder-gray-400 font-medium"
        />
      </div>

      <!-- Grid de pacientes -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <div
          v-for="(patient, index) in filteredPatients"
          :key="patient.id"
          @click="viewPatientReport(patient.id)"
          class="bg-white rounded-3xl p-6 border border-gray-100 shadow-sm hover:shadow-xl hover:border-indigo-100 hover:-translate-y-1 transition-all duration-300 cursor-pointer group flex flex-col items-center text-center relative overflow-hidden"
        >
          <!-- Background decoration -->
          <div class="absolute top-0 left-0 w-full h-24 bg-gradient-to-b from-gray-50 to-white opacity-50"></div>
          
          <div
            :class="[
              getAvatarColor(index),
              'w-20 h-20 rounded-2xl flex items-center justify-center text-2xl font-bold mb-4 shadow-sm z-10 transition-transform group-hover:scale-110',
            ]"
          >
             {{ patient.name.charAt(0).toUpperCase() }}
          </div>
          
          <h3 class="text-lg font-bold text-gray-900 group-hover:text-indigo-600 transition-colors z-10">
            {{ patient.name }}
          </h3>
          <p class="text-sm text-gray-500 font-medium mb-6 z-10">{{ patient.age }}</p>

          <button
             class="mt-auto w-full py-2.5 px-4 rounded-xl bg-indigo-50 text-indigo-600 font-semibold text-sm hover:bg-indigo-600 hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md"
          >
            <font-awesome-icon :icon="faChartLine" />
            Ver Relatório
          </button>
        </div>
      </div>

      <!-- Sem resultados da busca -->
      <div
        v-if="filteredPatients.length === 0 && searchQuery"
        class="text-center py-20"
      >
        <div class="w-16 h-16 bg-gray-100 rounded-full flex items-center justify-center mx-auto mb-4 text-gray-400">
           <font-awesome-icon
            :icon="faSearch"
            class="w-6 h-6"
          />
        </div>
        <p class="text-gray-900 font-medium text-lg">
          Nenhum paciente encontrado
        </p>
         <p class="text-gray-500">
          Não encontramos resultados para "{{ searchQuery }}"
        </p>
        <button 
            @click="searchQuery = ''"
            class="mt-4 text-indigo-600 font-bold hover:text-indigo-800"
        >
            Limpar busca
        </button>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import { useToast } from "vue-toastification";
import api from "@/services/api";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faUsers,
  faEnvelope,
  faChartLine,
  faSearch,
} from "@fortawesome/free-solid-svg-icons";

const router = useRouter();
const toast = useToast();

const patients = ref([]);
const isLoading = ref(true);
const searchQuery = ref("");

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


// Cores para avatares (estilo moderno)
const avatarColors = [
  "bg-blue-100 text-blue-600",
  "bg-emerald-100 text-emerald-600",
  "bg-violet-100 text-violet-600",
  "bg-rose-100 text-rose-600",
  "bg-amber-100 text-amber-600",
  "bg-cyan-100 text-cyan-600",
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

// Filtrar pacientes pela busca
const filteredPatients = computed(() => {
  if (!searchQuery.value) return patients.value;
  const query = searchQuery.value.toLowerCase();
  return patients.value.filter((p) => p.name.toLowerCase().includes(query));
});

// Calcular idade
const calculateAge = (birthdate) => {
  if (!birthdate) return "Idade não inf.";
  const birthDate = new Date(birthdate);
  const today = new Date();
  let age = today.getFullYear() - birthDate.getFullYear();
  const m = today.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }
  return `${age} anos`;
};

// Navegar para relatório do paciente
const viewPatientReport = (patientId) => {
  router.push({
    name: "profissional-paciente-relatorio",
    params: { patientId },
  });
};

// Buscar pacientes
const fetchPatients = async () => {
  isLoading.value = true;
  try {
    const response = await api.listarPacientesDoProfissional();
    patients.value = (response.data || []).map((paciente) => ({
      id: paciente.id,
      name: paciente.usuario?.nome || "Paciente",
      age: calculateAge(paciente.data_nascimento),
      // Adicionar outros campos se necessário
    }));
  } catch (error) {
    console.error("Erro ao buscar pacientes:", error);
    toast.error("Erro ao carregar lista de pacientes.");
  } finally {
    isLoading.value = false;
  }
};

onMounted(fetchPatients);
</script>

<style scoped>
/* Estilos específicos se necessário */
</style>