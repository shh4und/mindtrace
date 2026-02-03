<template>
  <div>
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Meus Profissionais</h1>
      <p class="text-gray-600 mt-1">Profissionais de saúde vinculados ao seu acompanhamento.</p>
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
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="(profissional, index) in profissionais"
        :key="profissional.id"
        :title="profissional.nome"
        :subtitle="profissional.especialidade"
        variant="profissional"
        :aria-label="`Ver perfil de ${profissional.nome}`"
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
          {{ profissional.nome.charAt(0).toUpperCase() }}
        </div>

        <h3
          class="text-lg font-bold text-gray-900 group-hover:text-indigo-600 transition-colors z-10"
        >
          {{ profissional.nome }}
        </h3>
        <p class="text-sm text-gray-500 font-medium mb-6 z-10">
          {{ profissional.especialidade }}
        </p>

        <button
          @click="viewProfissionalProfile(profissional.id)"
          class="mb-1 mt-auto w-full py-2.5 px-4 rounded-xl bg-indigo-50 text-indigo-600 font-semibold text-sm hover:bg-indigo-600 hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md"
        >
          <font-awesome-icon :icon="faIdCard" />
          Ver perfil de {{ profissional.nome }}
        </button>
      </div>
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
import { faUserDoctor, faLink, faIdCard, faEnvelope } from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const toast = useToast();

const profissionais = ref([]);
const isLoading = ref(true);

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
const cardActions = [{ id: 'profile', label: 'Ver Perfil', icon: faIdCard }];

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
    profissionais.value = (response.data || []).map((prof) => ({
      id: prof.id,
      nome: prof.usuario?.nome || 'Profissional',
      especialidade: prof.especialidade || 'Especialidade não informada',
      email: prof.usuario?.email,
    }));
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
    email: 'maria.silva@email.com',
  },
  {
    id: 2,
    nome: 'Dr. João Santos',
    especialidade: 'Psiquiatra - CRM 54321',
    email: 'joao.santos@email.com',
  },
];

onMounted(fetchProfissionais);
</script>
