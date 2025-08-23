<template>
  <aside class="w-full lg:w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200 flex flex-col">
    <div class="flex items-center justify-between p-6 border-b border-gray-200">
      <div class="flex items-center space-x-3">
        <img
          src="https://images.unsplash.com/photo-1506744038136-46273834b3fb?fm=jpg&q=60&w=3000&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8M3x8bGFuZHNjYXBlfGVufDB8fDB8fHww"
          alt="Usuário Paciente" class="w-10 h-10 rounded-full object-cover">
        <div>
          <h2 class="font-semibold text-gray-900">João</h2>
          <p class="text-sm text-gray-500">Paciente</p>
        </div>
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
          <a href="#" @click.prevent="$emit('navigate', 'perfil')" class="sidebar-item">
            <i class="fa-solid fa-user fa-fw mr-3"></i>
            <span>Perfil</span>
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
  // Lógica de logout aqui
  console.log('Usuário deslogado');
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
  transition: background-color 0.2s, color 0.2s;
  cursor: pointer;
}

.sidebar-item:hover {
  background-color: #F3F4F6;
  color: #1F2937;
}

.sidebar-item.active {
  background-color: #DBF7E9;
  /* soft emerald-ish background */
  color: #166534;
  /* darker emerald */
}
</style>
