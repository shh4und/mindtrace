<template>
  <aside class="w-full lg:w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200 flex flex-col">
    <!-- Profile Section -->
    <div class="p-6 border-b border-gray-200 relative">
      <button @click="toggleProfileCard"
        class="flex items-center space-x-3 w-full text-left p-2 rounded-lg hover:bg-gray-100 transition-colors">
        <font-awesome-icon :icon="['fas', 'user']" class="w-8 h-8 rounded-full p-2 bg-gray-200 text-gray-600" />
        <div v-if="userStore.user" class="flex-1">
          <h2 class="font-semibold text-gray-900 truncate">{{ userStore.user.nome }}</h2>
          <p class="text-sm text-gray-500">Paciente</p>
        </div>
        <div v-else class="flex-1">
          <div class="h-4 bg-gray-200 rounded w-3/4 mb-1"></div>
          <div class="h-3 bg-gray-200 rounded w-1/2"></div>
        </div>
      </button>

      <!-- Profile Pop-up Card -->
      <div v-if="isProfileCardVisible && userStore.user" @click.stop
        class="absolute z-50 top-full mt-2 w-72 bg-white rounded-lg shadow-xl border border-gray-200 p-4">
        <div class="flex items-center mb-4">
          <font-awesome-icon :icon="['fas', 'user']"
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
          <p><strong class="font-medium">Idade:</strong> {{ userStore.user.idade }}</p>
          <p><strong class="font-medium">Contato:</strong> {{ userStore.user.contato || 'Não informado' }}</p>
          <p><strong class="font-medium">Dependente:</strong> {{ userStore.user.dependente ? 'É dependente' : 'Não dependente'  }}</p>
          <p v-if="userStore.user.dependente"><strong class="font-medium">Responsável:</strong> {{ userStore.user.nome_responsavel }}</p>
          <p v-if="userStore.user.dependente"><strong class="font-medium">Contato do Responsável:</strong> {{ userStore.user.contato_responsavel }}</p>
        </div>
        <button @click="editProfile"
          class="w-full text-sm bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-2 px-4 rounded-lg transition-colors duration-200">
          <font-awesome-icon :icon="['fas', 'pen-to-square']" />
          <span> Editar Perfil</span>
        </button>
      </div>
    </div>

    <nav class="p-4 flex-1">
      <ul class="space-y-1">
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'resumo')" class="sidebar-item">
            <i class="fa-solid fa-home fa-fw mr-3"></i>
            <span>Resumo</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'humor')" class="sidebar-item">
            <i class="fa-regular fa-face-smile-beam fa-fw mr-3"></i>
            <span>Registro de Humor</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'relatorios')" class="sidebar-item">
            <i class="fa-solid fa-chart-line fa-fw mr-3"></i>
            <span>Relatórios</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'vincular')" class="sidebar-item">
            <i class="fa-solid fa-link fa-fw mr-3"></i>
            <span>Vincular um Profissional</span>
          </a>
        </li>
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'editar-perfil')" class="sidebar-item">
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
import { useUserStore } from '../store/user';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faUser, faPenToSquare, faTimes } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';

library.add(faUser, faPenToSquare, faTimes);

const emit = defineEmits(['navigate']);
const userStore = useUserStore();
const isProfileCardVisible = ref(false);

onMounted(() => {
  userStore.fetchUser('paciente');
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
  background-color: #DBF7E9;
  color: #166534;
}

.truncate {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}
</style>