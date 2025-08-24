<template>
  <div class="min-h-screen bg-gray-100 font-sans antialiased flex flex-col lg:flex-row">
    <Navbar @navigate="handleNavigation" />

    <main class="flex-1 p-8 overflow-y-auto">
      <Resumo v-if="activeView === 'resumo'" />
      <RegistroHumor v-if="activeView === 'humor'" />
      <Relatorio v-if="activeView === 'relatorios'" />
      <Perfil v-if="activeView === 'perfil'" user-type="paciente" />

    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Navbar from '../components/Navbar.vue';
import RegistroHumor from '../components/RegistroHumor.vue';
import Resumo from '../components/Resumo.vue';
import Relatorio from '../components/Relatorio.vue';
import Perfil from '../components/Perfil.vue';
import '../assets/paciente_dashboard.css';

const activeView = ref('resumo');

const handleNavigation = (view) => {
  activeView.value = view;
};
</script>
