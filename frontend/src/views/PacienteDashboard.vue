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
import Navbar from '../components/Navbar.vue';
import RegistroHumor from '../components/RegistroHumor.vue';
import Resumo from '../components/Resumo.vue';
import Relatorio from '../components/Relatorio.vue';
import EditarPerfil from '../components/EditarPerfil.vue';
import VincularProfissional from '@/components/VincularProfissional.vue';

const activeView = ref('resumo');

const handleNavigation = (view) => {
  activeView.value = view;
};
</script>

<style>
@import "tailwindcss";

/* Estilos personalizados para a sidebar */
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
    background-color: #DBF7E9; /* soft emerald-ish background */
    color: #166534; /* darker emerald */
}
.sidebar-item.active .sidebar-icon {
    color: #166534;
}
.sidebar-icon {
    width: 20px;
    height: 20px;
    margin-right: 12px;
    color: #6B7280;
}
</style>
