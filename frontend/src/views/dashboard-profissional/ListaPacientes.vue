<template>
  <div>
    <h1 class="text-3xl font-bold text-gray-900 mb-8">Meus Pacientes</h1>
    <div v-if="isLoading" class="text-center py-8">
      <p class="text-gray-500">Carregando pacientes...</p>
    </div>
    <div v-else-if="patients.length === 0" class="text-center py-8">
      <p class="text-gray-500">Nenhum paciente vinculado ainda.</p>
    </div>
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="(patient, index) in patients" 
        :key="patient.id"
        @click="viewPatientReport(patient.id)"
        class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 cursor-pointer hover:shadow-lg hover:border-indigo-400 transition-all duration-200"
        role="button"
        :aria-label="`Ver relatório do paciente ${patient.name}`"
        tabindex="0"
        @keydown.enter="viewPatientReport(patient.id)"
        @keydown.space.prevent="viewPatientReport(patient.id)"
      >
        <div class="flex items-center space-x-4 mb-4">
          <div :class="getAvatarClass(index)" class="w-14 h-14 rounded-full flex items-center justify-center">
            <font-awesome-icon :icon="['fas', 'user']" class="w-8 h-8 text-white" aria-hidden="true" />
          </div>
          <div>
            <h3 class="font-semibold text-gray-900 text-lg">{{ patient.name }}</h3>
            <p class="text-sm text-gray-500">{{ patient.age }}</p>
          </div>
        </div>
        <button 
          @click.stop="viewPatientReport(patient.id)" 
          class="flex items-center text-sm font-medium text-indigo-600 hover:text-indigo-800 transition-colors"
        >
          <font-awesome-icon :icon="faChartLine" class="mr-2" aria-hidden="true" />
          Ver Relatório
        </button>
        <button 
          @click.stop="viewQuestFormAssign({patientId: patient.id, patientNome: patient.name})" 
          class="flex items-center text-sm font-medium text-indigo-600 hover:text-indigo-800 transition-colors"
        >
          <font-awesome-icon :icon="faListCheck" class="mr-2" aria-hidden="true" />
          Atribuir Questionário
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import api from '@/services/api';
import { useToast } from 'vue-toastification';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUser, faChartLine, faListCheck } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';

library.add(faUser);

const router = useRouter();
const patients = ref([]);
const isLoading = ref(true);
const toast = useToast();

// Array de cores para os avatares
const avatarColors = [
  'bg-blue-500',
  'bg-green-500',
  'bg-purple-500',
  'bg-red-500',
  'bg-yellow-500',
  'bg-indigo-500',
  'bg-pink-500',
  'bg-teal-500',
];

const getAvatarClass = (index) => {
  return avatarColors[index % avatarColors.length];
};

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
    params: { patientId } 
  });
};

const viewQuestFormAssign = ({patientId, patientNome}) => {
  router.push({ 
    name: 'profissional-atribuir-questionario', 
    params: { patientId, patientNome }  // params de rota, NÃO props de componente
  });
};

onMounted(async () => {
  try {
    const response = await api.listarPacientesDoProfissional();
    // Mapear os dados da API para o formato esperado pelo template
    patients.value = response.data.map(paciente => ({
      id: paciente.id,
      name: paciente.usuario.nome,
      age: `${calculateAge(paciente.data_nascimento)} anos`,
      focus: paciente.usuario.bio || "Tratamento em andamento.", // Usar bio ou um placeholder
    }));
  } catch (error) {
    console.error('Erro ao buscar pacientes:', error);
    toast.error('Erro ao carregar lista de pacientes.');
  } finally {
    isLoading.value = false;
  }
});
</script>
