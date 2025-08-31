<template>
  <aside class="w-full lg:w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200 flex flex-col">
    <!-- Profile Section -->
    <div class="p-6 border-b border-gray-200 relative">
      <button @click="isProfileCardVisible = !isProfileCardVisible"
        class="flex items-center space-x-3 w-full text-left p-2 rounded-lg hover:bg-gray-100 transition-colors">
        <font-awesome-icon :icon="['fas', 'user-doctor']" class="w-8 h-8 rounded-full p-2 bg-gray-200 text-gray-600" />
        <div v-if="userStore.user" class="flex-1">
          <h2 class="font-semibold text-gray-900 truncate">{{ userStore.user.nome }}</h2>
          <p class="text-sm text-gray-500">{{ userStore.user.profissional.especialidade || 'Profissional' }}</p>
        </div>
        <div v-else class="flex-1">
          <div class="h-4 bg-gray-200 rounded w-3/4 mb-1"></div>
          <div class="h-3 bg-gray-200 rounded w-1/2"></div>
        </div>
      </button>

      <!-- Profile Pop-up Card -->
      <div v-if="userStore.user" class="flex-1">
        <h2 class="font-semibold text-gray-900 truncate">{{ userStore.user.nome }}</h2>
        <p class="text-sm text-gray-500">{{ userStore.user.profissional.especialidade || 'Profissional' }}</p>
      </div>
      <div v-if="isProfileCardVisible && userStore.user"
        class="absolute z-20 bottom-full mb-2 w-72 bg-white rounded-lg shadow-xl border border-gray-200 p-4">
        <div class="flex items-center mb-4">
          <font-awesome-icon :icon="['fas', 'user-doctor']"
            class="w-10 h-10 rounded-full p-2 mr-3 bg-gray-200 text-gray-600" />
          <div>
            <h3 class="font-bold text-lg">{{ userStore.user.nome }}</h3>
            <p class="text-sm text-gray-600">{{ userStore.user.email }}</p>
          </div>
        </div>
        <div class="space-y-2 text-sm mb-4">
          <p><strong class="font-medium">Especialidade:</strong> {{ userStore.user.profissional.especialidade || 'Não informado' }}</p>
          <p><strong class="font-medium">Registro:</strong> {{ userStore.user.profissional.registro_profissional || 'Não informado' }}</p>
        </div>
      </div>
    </div>

    <nav class="p-4 flex-1">
      <ul class="space-y-1">
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'pacientes')" class="sidebar-item">
            <i class="fa-solid fa-users fa-fw mr-3"></i>
            <span>Pacientes</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'convite')" class="sidebar-item">
            <i class="fa-solid fa-ticket fa-fw mr-3"></i>
            <span>Gerar Convite</span>
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
import { useUserStore } from '../store/user'; // 1. Importar o store Pinia
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUserDoctor, faPenToSquare } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';

library.add(faUserDoctor, faPenToSquare);

const emit = defineEmits(['navigate']);
const userStore = useUserStore(); // 2. Inicializar o store
const isProfileCardVisible = ref(false);

onMounted(() => {
  // 3. Chamar a ação do store com o tipo correto
  userStore.fetchUser('profissional');
});

const performLogout = () => {
  // 4. Chamar a ação de logout
  userStore.logout();
};

const editProfile = () => {
  emit('navigate', 'editar-perfil');
  isProfileCardVisible.value = false;
};
</script>

<style scoped>
.sidebar-item {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  border-radius: 8px;
  color: #4B5563;
  font-weight: 500;
  transition: background-color 0.2s, color 0.2s;
  cursor: pointer;
}

.sidebar-item:hover {
  background-color: #F3F4F6;
  color: #1F2937;
}

.sidebar-item.active {
  background-color: #EEF2FF;
  color: #4338CA;
}

.truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>