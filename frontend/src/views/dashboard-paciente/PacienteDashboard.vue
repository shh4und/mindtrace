<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased flex flex-col">
  <!-- Navbar superior -->
    <TopNavbar 
      :user-type="TipoUsuario.Paciente" 
      @edit-profile="handleNavigation('editar-perfil')"
      @logout="handleLogout"
    />

  <!-- Conteudo principal com sidebar -->
    <div class="flex flex-1 overflow-hidden">
  <!-- Barra lateral -->
      <SidebarPaciente 
        :active-view="activeView"
        @navigate="handleNavigation" 
      />

  <!-- Area de conteudo principal -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
        <Resumo v-if="activeView === 'resumo'" />
        <RegistroHumor v-if="activeView === 'humor'" />
        <Relatorio v-if="activeView === 'relatorios'" :user-type="TipoUsuario.Paciente"/>
        <VincularProfissional v-if="activeView === 'vincular'"/>
        <EditarPerfil v-if="activeView === 'editar-perfil'" :user-type="TipoUsuario.Paciente" />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '../../store/user';
import { TipoUsuario } from '../../types/usuario.js';
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
  // Busca dados do usuario se ainda nao estiverem carregados
  if (!userStore.user) {
    await userStore.fetchUser(TipoUsuario.Paciente);
  }
});
</script>


