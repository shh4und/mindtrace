<template>
  <aside class="w-full lg:w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200 flex flex-col">
    <!-- Profile Section -->
    <div class="p-6 border-b border-gray-200 relative">
      <button @click="toggleProfileCard"
        class="flex items-center space-x-3 w-full text-left p-2 rounded-lg hover:bg-gray-100 transition-colors">
        <font-awesome-icon :icon="['fas', 'user-doctor']" class="w-8 h-8 rounded-full p-2 bg-gray-200 text-gray-600" />
        <div v-if="userStore.user" class="flex-1">
          <h2 class="font-semibold text-gray-900 truncate">{{ userStore.user.nome || 'Profissional'}}</h2>
          <p class="text-sm text-gray-500">{{ userStore.user.registro_profissional || 'Profissional'}}</p>
        </div>
        <div v-else class="flex-1">
          <div class="h-4 bg-gray-200 rounded w-3/4 mb-1"></div>
          <div class="h-3 bg-gray-200 rounded w-1/2"></div>
        </div>
      </button>

      <!-- Profile Pop-up Card -->
      <div v-if="isProfileCardVisible && userStore.user"
        @click.stop
        class="absolute z-50 top-full mt-2 w-72 bg-white rounded-lg shadow-xl border border-gray-200 p-4">
        <div class="flex items-center mb-4">
          <font-awesome-icon :icon="['fas', 'user-doctor']"
            class="w-10 h-10 rounded-full p-2 mr-3 bg-gray-200 text-gray-600" />
          <div class="flex-1">
            <h3 class="font-bold text-lg">{{ userStore.user.nome }}</h3>
            <p class="text-sm text-gray-600">{{ userStore.user.email }}</p>
          </div>
          <button @click="closeProfileCard" class="text-gray-400 hover:text-gray-600">
            <i class="fa-solid fa-times"></i>
          </button>
        </div>
        <div class="space-y-2 text-sm mb-4">
          <p><strong class="font-medium">Especialidade:</strong> {{ userStore.user.especialidade || 'Não informado' }}</p>
          <p><strong class="font-medium">Registro:</strong> {{ userStore.user.registro_profissional || 'Não informado' }}</p>
          <!-- Novo Campo: Contato -->
          <p><strong class="font-medium">Contato:</strong> {{ userStore.user.contato || 'Não informado' }}</p>
        </div>
        <button @click="editProfile" class="w-full bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-2 px-4 rounded-lg transition-colors duration-200">
          <font-awesome-icon :icon="['fas', 'pen-to-square']" />
          <span>Editar Perfil</span>
        </button>
      </div>
    </div>

    <nav class="p-4 flex-1">
      <ul class="space-y-1">
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'pacientes')" class="flex items-center py-3 px-4 rounded-lg text-gray-600 font-medium transition-colors hover:bg-gray-100 hover:text-gray-800">
            <i class="fa-solid fa-users fa-fw mr-3"></i>
            <span>Pacientes</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'relatorios')" class="flex items-center py-3 px-4 rounded-lg text-gray-600 font-medium transition-colors hover:bg-gray-100 hover:text-gray-800">
            <i class="fa-solid fa-users fa-fw mr-3"></i>
            <span>Relatórios</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'convite')" class="flex items-center py-3 px-4 rounded-lg text-gray-600 font-medium transition-colors hover:bg-gray-100 hover:text-gray-800">
            <i class="fa-solid fa-ticket fa-fw mr-3"></i>
            <span>Gerar Convite</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'editar-perfil')" class="flex items-center py-3 px-4 rounded-lg text-gray-600 font-medium transition-colors hover:bg-gray-100 hover:text-gray-800">
            <i class="fa-solid fa-user-pen fa-fw mr-3"></i>
            <span>Editar Perfil</span>
          </a>
        </li>
      </ul>
    </nav>

    <div class="p-4 border-t border-gray-200">
      <a href="#" @click.prevent="performLogout" class="sidebar-item">
        <i class="fa-solid fa-right-from-bracket fa-fw mr-3"></i>
        <span>Sair</span>
      </a>
    </div>
  </aside>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '../../store/user';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUserDoctor, faPenToSquare, faTimes } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';

library.add(faUserDoctor, faPenToSquare, faTimes);

const emit = defineEmits(['navigate']);
const userStore = useUserStore();
const isProfileCardVisible = ref(false);

onMounted(() => {
  userStore.fetchUser('profissional');
});

const toggleProfileCard = () => {
  isProfileCardVisible.value = !isProfileCardVisible.value;
};

const closeProfileCard = () => {
  isProfileCardVisible.value = false;
};

const performLogout = () => {
  if (confirm('Tem certeza que deseja sair?')) {
    userStore.logout();
  }
};

const editProfile = () => {
  emit('navigate', 'editar-perfil');
  closeProfileCard();
};
</script>
