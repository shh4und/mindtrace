<template>
  <div>
    <header class="mb-10 text-center md:text-left">
      <h1 class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight">
        Meus Pacientes
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Gerencie seus pacientes vinculados e acompanhe seus progressos.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-16">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-rose-500"></div>
    </div>

    <!-- Empty state -->
    <div v-else-if="patients.length === 0" class="text-center py-16">
      <font-awesome-icon :icon="faUsers" class="w-16 h-16 text-gray-300 mb-4" />
      <h3 class="text-lg font-medium text-gray-900 mb-2">Nenhum paciente vinculado</h3>
      <p class="text-gray-500 mb-6">Gere um convite para vincular seu primeiro paciente.</p>
      <router-link
        to="/dashboard-profissional/convite"
        class="inline-flex items-center px-4 py-2 bg-rose-500 text-white rounded-3xl hover:bg-rose-600 transition-colors"
      >
        <font-awesome-icon :icon="faEnvelope" class="mr-2" />
        Gerar Convite
      </router-link>
    </div>

    <!-- Lista de pacientes -->
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
          :title="patient.name"
          :subtitle="patient.age"
          variant="profissional"
          :aria-label="`Ver detalhes de ${patient.name}`"
          class="bg-white rounded-3xl p-6 border border-gray-100 shadow-sm hover:shadow-xl hover:border-indigo-100 hover:-translate-y-1 transition-all duration-300 cursor-pointer group flex flex-col items-center text-center relative overflow-hidden"
        >
          <!-- Background decoration -->
          <div
            class="absolute top-0 left-0 w-full h-24 bg-linear-to-b from-gray-50 to-white opacity-50"
          ></div>

          <div
            :class="[
              getAvatarColor(index),
              'w-20 h-20 rounded-2xl flex items-center justify-center text-2xl font-bold mb-4 shadow-sm z-10 transition-transform group-hover:scale-110',
            ]"
          >
            {{ patient.name.charAt(0).toUpperCase() }}
          </div>

          <h3
            class="text-lg font-bold text-gray-900 group-hover:text-indigo-600 transition-colors z-10"
          >
            {{ patient.name }}
          </h3>
          <p class="text-sm text-gray-500 font-medium mb-6 z-10">{{ patient.age }}</p>

          <button
            @click="viewPatientReport(patient.id)"
            class="mb-1 mt-auto w-full py-2.5 px-4 rounded-xl bg-indigo-50 text-indigo-600 font-semibold text-sm hover:bg-indigo-600 hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md"
          >
            <font-awesome-icon :icon="faChartLine" />
            Ver Relatório
          </button>
          <button
            @click="viewQuestFormAssign({ patientId: patient.id, patientNome: patient.name })"
            class="mb-1 mt-auto w-full py-2.5 px-4 rounded-xl bg-indigo-50 text-indigo-600 font-semibold text-sm hover:bg-indigo-600 hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md"
          >
            <font-awesome-icon :icon="faClipboardList" />
            Atribuir Questionário
          </button>
        </div>
      </div>
      <!-- Sem resultados da busca -->
      <div v-if="filteredPatients.length === 0 && searchQuery" class="text-center py-12">
        <font-awesome-icon :icon="faSearch" class="w-12 h-12 text-gray-300 mb-3" />
        <p class="text-gray-500">Nenhum paciente encontrado para "{{ searchQuery }}"</p>
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/services/api';
import { useToast } from 'vue-toastification';
import { CardListaUsuario } from '@/components/ui';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faUsers,
  faEnvelope,
  faChartLine,
  faClipboardList,
  faSearch,
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const patients = ref([]);
const isLoading = ref(true);
const searchQuery = ref('');
const toast = useToast();

// Cores para avatares
const avatarColors = [
  'bg-blue-100 text-blue-600',
  'bg-green-100 text-green-600',
  'bg-purple-100 text-purple-600',
  'bg-red-100 text-red-600',
  'bg-yellow-100 text-yellow-600',
  'bg-indigo-100 text-indigo-600',
  'bg-pink-100 text-pink-600',
  'bg-teal-100 text-teal-600',
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

// Ações do card
const cardActions = [
  { id: 'report', label: 'Ver Relatório', icon: faChartLine },
  { id: 'assign', label: 'Atribuir Questionário', icon: faClipboardList },
];

// Filtrar pacientes pela busca
const filteredPatients = computed(() => {
  if (!searchQuery.value) return patients.value;
  const query = searchQuery.value.toLowerCase();
  return patients.value.filter((p) => p.name.toLowerCase().includes(query));
});

const calculateAge = (birthdate) => {
  if (!birthdate) return '';
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
    name: 'profissional-paciente-relatorio',
    params: { patientId },
  });
};

const viewQuestFormAssign = ({ patientId, patientNome }) => {
  router.push({
    name: 'profissional-atribuir-questionario',
    params: { patientId, patientNome },
  });
};

const handleAction = (actionId, patient) => {
  if (actionId === 'report') {
    viewPatientReport(patient.id);
  } else if (actionId === 'assign') {
    viewQuestFormAssign({ patientId: patient.id, patientNome: patient.name });
  }
};

onMounted(async () => {
  try {
    const response = await api.listarPacientesDoProfissional();
    patients.value = (response.data || []).map((paciente) => ({
      id: paciente.id,
      name: paciente.usuario?.nome || 'Paciente',
      age: `${calculateAge(paciente.data_nascimento)} anos`,
      focus: paciente.usuario?.bio || 'Tratamento em andamento.',
    }));
  } catch (error) {
    console.error('Erro ao buscar pacientes:', error);
    toast.error('Erro ao carregar lista de pacientes.');
  } finally {
    isLoading.value = false;
  }
});
</script>
