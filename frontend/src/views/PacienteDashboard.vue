<template>
  <div class="min-h-screen bg-gray-100 font-sans antialiased flex flex-col lg:flex-row">
    <Navbar @navigate="handleNavigation" />

    <main class="flex-1 p-8 overflow-y-auto">
      <Resumo v-if="activeView === 'resumo'" />
      <RegistroHumor v-if="activeView === 'humor'" />

      <div v-if="activeView === 'relatorios'" class="text-center mt-10">
        <h2 class="text-2xl font-bold">Página de Relatórios</h2>
        <p class="text-gray-600">Em construção...</p>
      </div>
      <div v-if="activeView === 'anotacoes'" class="text-center mt-10">
        <h2 class="text-2xl font-bold">Página de Anotações Diárias</h2>
        <p class="text-gray-600">Em construção...</p>
      </div>
      <div v-if="activeView === 'perfil'" class="text-center mt-10">
        <h2 class="text-2xl font-bold">Página de Perfil</h2>
        <p class="text-gray-600">Em construção...</p>
      </div>

    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Navbar from '../components/Navbar.vue';
import RegistroHumor from '../components/RegistroHumor.vue';
import Resumo from '../components/Resumo.vue';
import '../assets/paciente_dashboard.css';

const activeView = ref('resumo');

const handleNavigation = (view) => {
  activeView.value = view;
};
</script>
