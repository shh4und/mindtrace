<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased flex flex-col">
  <!-- Navbar superior -->
    <TopNavbar 
      :user-type="TipoUsuario.Profissional" 
      @edit-profile="handleNavigation('editar-perfil')"
      @logout="handleLogout"
    />

  <!-- Conteudo principal com sidebar -->
    <div class="flex flex-1 overflow-hidden">
  <!-- Barra lateral -->
      <SidebarProfissional 
        :active-view="activeView"
        @navigate="handleNavigation" 
      />

  <!-- Area de conteudo principal -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
  <!-- Lista de pacientes ou relatorio de um paciente especifico -->
        <div v-if="activeView === 'pacientes'">
          <ListaPacientes v-if="!selectedPatientId" @view-patient="showPatientReport" />
          <div v-else>
            <button @click="showPatientList" class="mb-6 flex items-center text-sm font-medium text-rose-600 hover:text-rose-800 transition-colors">
              <i class="fa-solid fa-arrow-left mr-2"></i>
              Voltar para a lista de pacientes
            </button>
            <Relatorio @view-relatorios="showPatientReport" :user-type="TipoUsuario.Profissional" :patient-id="selectedPatientId" />
          </div>
        </div>

  <!-- Outras views disponiveis -->
        <GerarConvite v-if="activeView === 'convite'" />
        <EditarPerfil v-if="activeView === 'editar-perfil'" :user-type="TipoUsuario.Profissional" />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '../../store/user';
import { TipoUsuario } from '../../types/usuario.js';
import TopNavbar from '../../components/layout/TopNavbar.vue';
import SidebarProfissional from '../../components/layout/SidebarProfissional.vue';
import ListaPacientes from './ListaPacientes.vue';
import Relatorio from '../shared/Relatorio.vue';
import GerarConvite from './GerarConvite.vue';
import EditarPerfil from '../shared/EditarPerfil.vue';

const userStore = useUserStore();
const activeView = ref('pacientes'); // Visao inicial
const selectedPatientId = ref(null);

const handleNavigation = (view) => {
  activeView.value = view;
  selectedPatientId.value = null; // Reseta a selecao de paciente ao navegar
};

const showPatientReport = (patientId) => {
  selectedPatientId.value = patientId;
  // A view ja e pacientes entao apenas a condicao do vif velse muda
};

const showPatientList = () => {
  selectedPatientId.value = null;
};

const handleLogout = () => {
  userStore.logout();
};

onMounted(async () => {
  // Busca dados do usuario se ainda nao estiverem carregados
  if (!userStore.user) {
    await userStore.fetchUser(TipoUsuario.Profissional);
  }
});
</script>

