<template>
  <div>
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Meus Profissionais</h1>
      <p class="text-gray-600 mt-1">
        Profissionais de saúde vinculados ao seu acompanhamento.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-16">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-600"></div>
    </div>

    <!-- Empty state -->
    <div v-else-if="profissionais.length === 0" class="text-center py-16">
      <font-awesome-icon :icon="faUserDoctor" class="w-16 h-16 text-gray-300 mb-4" />
      <h3 class="text-lg font-medium text-gray-900 mb-2">Nenhum profissional vinculado</h3>
      <p class="text-gray-500 mb-6">
        Você ainda não está vinculado a nenhum profissional de saúde.
      </p>
      <router-link 
        to="/dashboard-paciente/vincular"
        class="inline-flex items-center px-4 py-2 bg-emerald-600 text-white rounded-lg hover:bg-emerald-700 transition-colors"
      >
        <font-awesome-icon :icon="faLink" class="mr-2" />
        Vincular Profissional
      </router-link>
    </div>

    <!-- Lista de profissionais -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <CardListaUsuario
        v-for="(profissional, index) in profissionais"
        :key="profissional.id"
        :title="profissional.nome"
        :subtitle="profissional.especialidade"
        variant="paciente"
        :avatar-color="getAvatarColor(index)"
        :actions="cardActions"
        :aria-label="`Ver perfil de ${profissional.nome}`"
        @click="viewProfissionalProfile(profissional.id)"
        @action="handleAction($event, profissional)"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import api from '@/services/api';
import { CardListaUsuario } from '@/components/ui';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faUserDoctor, 
  faLink, 
  faIdCard,
  faEnvelope
} from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const toast = useToast();

const profissionais = ref([]);
const isLoading = ref(true);

// Cores para avatares
const avatarColors = [
  'bg-emerald-500',
  'bg-teal-500',
  'bg-cyan-500',
  'bg-blue-500',
  'bg-indigo-500',
  'bg-violet-500',
  'bg-purple-500',
  'bg-fuchsia-500',
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

// Ações do card
const cardActions = [
  { id: 'profile', label: 'Ver Perfil', icon: faIdCard }
];

const viewProfissionalProfile = (profissionalId) => {
  // Por enquanto, exibe toast - pode ser expandido para modal ou página
  toast.info('Visualização de perfil em desenvolvimento.');
};

const handleAction = (actionId, profissional) => {
  if (actionId === 'profile') {
    viewProfissionalProfile(profissional.id);
  }
};

// Buscar profissionais vinculados
const fetchProfissionais = async () => {
  isLoading.value = true;
  try {
    const response = await api.listarProfissionaisDoPaciente();
    profissionais.value = response.data || [];
  } catch (error) {
    console.error('Erro ao buscar profissionais:', error);
    // Se o endpoint não existir, usa mock
    if (error.response?.status === 404 || error.response?.status === 501) {
      toast.warning('Usando dados de demonstração.');
      profissionais.value = getMockProfissionais();
    } else {
      toast.error('Erro ao carregar lista de profissionais.');
    }
  } finally {
    isLoading.value = false;
  }
};

// Mock de profissionais para desenvolvimento
const getMockProfissionais = () => [
  {
    id: 1,
    nome: 'Dra. Maria Silva',
    especialidade: 'Psicóloga - CRP 06/12345',
    email: 'maria.silva@email.com'
  },
  {
    id: 2,
    nome: 'Dr. João Santos',
    especialidade: 'Psiquiatra - CRM 54321',
    email: 'joao.santos@email.com'
  }
];

onMounted(fetchProfissionais);
</script>
