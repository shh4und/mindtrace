<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased flex flex-col">
    <!-- Top Navbar -->
    <TopNavbar 
      user-type="paciente" 
      @edit-profile="handleNavigation('editar-perfil')"
      @logout="handleLogout"
    />

    <!-- Main Content with Sidebar -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <SidebarPaciente 
        :active-view="activeView"
        @navigate="handleNavigation" 
      />

      <!-- Main Content Area -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
        <Resumo v-if="activeView === 'resumo'" />
        <RegistroHumor v-if="activeView === 'humor'" />
        <Relatorio v-if="activeView === 'relatorios'" user-type="paciente"/>
        <VincularProfissional v-if="activeView === 'vincular'"/>
        <EditarPerfil v-if="activeView === 'editar-perfil'" user-type="paciente" />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '../../store/user';
import TopNavbar from '../../components/layout/TopNavbar.vue';
import SidebarPaciente from '../../components/layout/SidebarPaciente.vue';
import RegistroHumor from './RegistroHumor.vue';
import Resumo from './Resumo.vue';
import Relatorio from '../shared/Relatorio.vue';
import EditarPerfil from '../shared/EditarPerfil.vue';
import VincularProfissional from './VincularProfissional.vue';

const userStore = useUserStore();
const activeView = ref('resumo');

const handleNavigation = (view) => {
  activeView.value = view;
};

const handleLogout = () => {
  userStore.logout();
};

onMounted(async () => {
  // Fetch user data if not already loaded
  if (!userStore.user) {
    await userStore.fetchUser('paciente');
  }
});
</script>


