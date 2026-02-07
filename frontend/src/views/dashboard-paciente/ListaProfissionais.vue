<template>
  <div class="max-w-4xl mx-auto p-4 md:p-8">
    <header class="mb-10 text-center md:text-left">
      <div
        class="inline-flex items-center justify-center space-x-2 bg-emerald-50 text-emerald-800 px-4 py-1.5 rounded-full text-sm font-medium mb-4 shadow-sm"
      >
        <font-awesome-icon :icon="faUserDoctor" class="h-4 w-4" />
        <span>Acompanhamento</span>
      </div>
      <h1
        class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
      >
        Meus Profissionais
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Profissionais de saúde vinculados ao seu acompanhamento.
      </p>
    </header>

    <!-- Loading state -->
    <div v-if="isLoading" class="flex justify-center items-center py-20">
      <div
        class="animate-spin rounded-full h-12 w-12 border-b-2 border-emerald-600"
      ></div>
    </div>

    <!-- Empty state -->
    <div
      v-else-if="profissionais.length === 0"
      class="text-center py-20 bg-white rounded-3xl border border-dashed border-gray-200"
    >
      <div
        class="w-20 h-20 bg-gray-50 rounded-full flex items-center justify-center mx-auto mb-6 text-gray-300"
      >
        <font-awesome-icon :icon="faUserDoctor" class="w-10 h-10" />
      </div>
      <h3 class="text-xl font-bold text-gray-900 mb-2">
        Nenhum profissional vinculado
      </h3>
      <p class="text-gray-500 mb-8 max-w-md mx-auto">
        Você ainda não está vinculado a nenhum profissional de saúde.
      </p>
      <router-link
        to="/dashboard-paciente/vincular"
        class="inline-flex items-center px-6 py-3 bg-emerald-600 text-white font-bold rounded-xl hover:bg-emerald-700 transition-all shadow-md hover:shadow-lg hover:-translate-y-0.5"
      >
        <font-awesome-icon :icon="faLink" class="mr-2" />
        Vincular Profissional
      </router-link>
    </div>

    <!-- Lista de profissionais -->
    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      <div
        v-for="(profissional, index) in profissionais"
        :key="profissional.id"
        class="bg-white rounded-3xl p-6 border border-gray-100 shadow-sm hover:shadow-xl hover:border-emerald-100 hover:-translate-y-1 transition-all duration-300 cursor-pointer group flex flex-col items-center text-center relative overflow-hidden"
      >
        <!-- Background decoration -->
        <div
          class="absolute top-0 left-0 w-full h-24 bg-gradient-to-b from-gray-50 to-white opacity-50"
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
          class="text-lg font-bold text-gray-900 group-hover:text-emerald-600 transition-colors z-10"
        >
          {{ profissional.nome }}
        </h3>
        <p class="text-sm text-gray-500 font-medium mb-6 z-10">
          {{ profissional.especialidade }}
        </p>

        <button
          @click="viewProfissionalProfile(profissional.id)"
          class="mb-1 mt-auto w-full py-2.5 px-4 rounded-xl bg-emerald-50 text-emerald-600 font-semibold text-sm hover:bg-emerald-600 hover:text-white transition-all duration-300 flex items-center justify-center gap-2 group-hover:shadow-md"
        >
          <font-awesome-icon :icon="faIdCard" />
          Ver Perfil
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
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUserDoctor, faLink, faIdCard } from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const toast = useToast();

const profissionais = ref([]);
const isLoading = ref(true);

// Cores para avatares (consistente com Resumo.vue)
const avatarColors = [
  'bg-blue-100 text-blue-600',
  'bg-emerald-100 text-emerald-600',
  'bg-violet-100 text-violet-600',
  'bg-rose-100 text-rose-600',
  'bg-amber-100 text-amber-600',
];

const getAvatarColor = (index) => avatarColors[index % avatarColors.length];

const viewProfissionalProfile = (profissionalId) => {
  toast.info('Visualização de perfil em desenvolvimento.');
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

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
