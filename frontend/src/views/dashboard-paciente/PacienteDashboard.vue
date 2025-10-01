<template>
  <div class="min-h-screen bg-gray-100 font-sans antialiased flex flex-col lg:flex-row">
    <Navbar @navigate="handleNavigation" />

    <main class="flex-1 p-8 overflow-y-auto">
      <Resumo v-if="activeView === 'resumo'" />
      <RegistroHumor v-if="activeView === 'humor'" />
      <Relatorio v-if="activeView === 'relatorios'" user-type="paciente"/>
      <VincularProfissional v-if="activeView === 'vincular'"/>
      <EditarPerfil v-if="activeView === 'editar-perfil'" user-type="paciente" />

    </main>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Navbar from '../../components/layout/NavbarPaciente.vue';
import RegistroHumor from './RegistroHumor.vue';
import Resumo from './Resumo.vue';
import Relatorio from '../shared/Relatorio.vue';
import EditarPerfil from '../shared/EditarPerfil.vue';
import VincularProfissional from './VincularProfissional.vue';

const activeView = ref('resumo');

const handleNavigation = (view) => {
  activeView.value = view;
};
</script>


