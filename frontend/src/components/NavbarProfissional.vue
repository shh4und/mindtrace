<template>
  <aside class="w-full lg:w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200 flex flex-col">
    <div class="flex items-center justify-between p-6 border-b border-gray-200">
      <div class="flex items-center space-x-3">
        <img src="https://images.unsplash.com/photo-1557862921-37829c790f19?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1471&q=80"
             alt="Dra. Maria"
             class="w-12 h-12 rounded-full object-cover border-2 border-indigo-400">
        <div>
          <h2 class="font-semibold text-gray-900">Dra. Maria</h2>
          <p class="text-sm text-gray-500">Psicóloga</p>
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
        <li>
          <a href="#" @click.prevent="$emit('navigate', 'perfil')" class="sidebar-item">
            <i class="fa-solid fa-user-gear fa-fw mr-3"></i>
            <span>Meu Perfil</span>
          </a>
        </li>
      </ul>
    </nav>

    <div class="p-4 border-t border-gray-200">
        <a href="#" @click.prevent="logout" class="sidebar-item">
            <i class="fa-solid fa-right-from-bracket fa-fw mr-3"></i>
            <span>Sair</span>
        </a>
    </div>
  </aside>
</template>

<script setup>
import { useRouter } from 'vue-router';

defineEmits(['navigate']);

const router = useRouter();

const logout = () => {
  console.log('Profissional deslogado');
  router.push('/login');
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
  background-color: #EEF2FF; /* Cor de destaque para o profissional */
  color: #4338CA; /* Cor de destaque para o profissional */
}
</style>
