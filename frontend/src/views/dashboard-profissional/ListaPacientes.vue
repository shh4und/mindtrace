<template>
  <div>
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Meus Pacientes</h1>
      <p class="text-gray-600 mt-1">
        Gerencie seus pacientes vinculados e acompanhe seu progresso.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-16">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-rose-500"
      ></div>
    </div>

    <!-- Empty state -->
    <div v-else-if="patients.length === 0" class="text-center py-16">
      <font-awesome-icon :icon="faUsers" class="w-16 h-16 text-gray-300 mb-4" />
      <h3 class="text-lg font-medium text-gray-900 mb-2">
        Nenhum paciente vinculado
      </h3>
      <p class="text-gray-500 mb-6">
        Gere um convite para vincular seu primeiro paciente.
      </p>
      <router-link
        to="/dashboard-profissional/convite"
        class="inline-flex items-center px-4 py-2 bg-rose-500 text-white rounded-lg hover:bg-rose-600 transition-colors"
      >
        <font-awesome-icon :icon="faEnvelope" class="mr-2" />
        Gerar Convite
      </router-link>
    </div>

    <!-- Lista de pacientes -->
    <template v-else>
      <!-- Barra de busca -->
      <div class="mb-6">
        <div class="relative">
          <font-awesome-icon
            :icon="faSearch"
            class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-400 w-5 h-5"
          />
          <input
            v-model="searchQuery"
            type="text"
            placeholder="Buscar paciente por nome..."
            class="w-full pl-12 pr-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-rose-500 focus:border-rose-500 outline-none transition-colors"
          />
        </div>
      </div>

      <!-- Grid de pacientes -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
        <CardListaUsuario
          v-for="(patient, index) in filteredPatients"
          :key="patient.id"
          :title="patient.name"
          :subtitle="patient.age"
          variant="profissional"
          :avatar-color="getAvatarColor(index)"
          :actions="cardActions"
          :aria-label="`Ver detalhes de ${patient.name}`"
          @click="viewPatientReport(patient.id)"
          @action="handleAction($event, patient)"
        />
      </div>

      <!-- Sem resultados da busca -->
      <div
        v-if="filteredPatients.length === 0 && searchQuery"
        class="text-center py-12"
      >
        <font-awesome-icon
          :icon="faSearch"
          class="w-12 h-12 text-gray-300 mb-3"
        />
        <p class="text-gray-500">
          Nenhum paciente encontrado para "{{ searchQuery }}"
        </p>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import { useRouter } from "vue-router";
import api from "@/services/api";
import { useToast } from "vue-toastification";
import { CardListaUsuario } from "@/components/ui";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faUsers,
  faEnvelope,
  faChartLine,
  faClipboardList,
  faSearch,
} from "@fortawesome/free-solid-svg-icons";

const router = useRouter();
const patients = ref([]);
const isLoading = ref(true);
const searchQuery = ref("");
const toast = useToast();

// Cores para avatares
const avatarColors = [
  "bg-blue-100 text-blue-600",
  "bg-green-100 text-green-600",
  "bg-purple-100 text-purple-600",
  "bg-red-100 text-red-600",
  "bg-yellow-100 text-yellow-600",
  "bg-indigo-100 text-indigo-600",
  "bg-pink-100 text-pink-600",
  "bg-teal-100 text-teal-600",
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

// Ações do card
const cardActions = [
  { id: "report", label: "Ver Relatório", icon: faChartLine },
  { id: "assign", label: "Atribuir Questionário", icon: faClipboardList },
];

// Filtrar pacientes pela busca
const filteredPatients = computed(() => {
  if (!searchQuery.value) return patients.value;
  const query = searchQuery.value.toLowerCase();
  return patients.value.filter((p) => p.name.toLowerCase().includes(query));
});

const calculateAge = (birthdate) => {
  if (!birthdate) return "";
  const birthDate = new Date(birthdate);
  const today = new Date();
  let age = today.getFullYear() - birthDate.getFullYear();
  const m = today.getMonth() - birthDate.getMonth();
  if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
    age--;
  }
  return age;
};

const viewPatientReport = (patientId) => {
  router.push({
    name: "profissional-paciente-relatorio",
    params: { patientId },
  });
};

const viewQuestFormAssign = ({ patientId, patientNome }) => {
  router.push({
    name: "profissional-atribuir-questionario",
    params: { patientId, patientNome },
  });
};

const handleAction = (actionId, patient) => {
  if (actionId === "report") {
    viewPatientReport(patient.id);
  } else if (actionId === "assign") {
    viewQuestFormAssign({ patientId: patient.id, patientNome: patient.name });
  }
};

onMounted(async () => {
  try {
    const response = await api.listarPacientesDoProfissional();
    patients.value = (response.data || []).map((paciente) => ({
      id: paciente.id,
      name: paciente.usuario?.nome || "Paciente",
      age: `${calculateAge(paciente.data_nascimento)} anos`,
      focus: paciente.usuario?.bio || "Tratamento em andamento.",
    }));
  } catch (error) {
    console.error("Erro ao buscar pacientes:", error);
    toast.error("Erro ao carregar lista de pacientes.");
  } finally {
    isLoading.value = false;
  }
});
</script>
