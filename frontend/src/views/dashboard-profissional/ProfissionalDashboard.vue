<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased flex flex-col">
    <!-- Top Navbar -->
    <TopNavbar 
      user-type="profissional" 
      @edit-profile="handleNavigation('editar-perfil')"
      @logout="handleLogout"
    />

    <!-- Main Content with Sidebar -->
    <div class="flex flex-1 overflow-hidden">
      <!-- Sidebar -->
      <SidebarProfissional 
        :active-view="activeView"
        @navigate="handleNavigation" 
      />

      <!-- Main Content Area -->
      <main class="flex-1 overflow-y-auto p-4 sm:p-6 lg:p-8">
        <!-- Mostra a lista de pacientes ou o relatório de um paciente específico -->
        <div v-if="activeView === 'pacientes'">
          <ListaPacientes v-if="!selectedPatientId" @view-patient="showPatientReport" />
          <div v-else>
            <button @click="showPatientList" class="mb-6 flex items-center text-sm font-medium text-rose-600 hover:text-rose-800 transition-colors">
              <i class="fa-solid fa-arrow-left mr-2"></i>
              Voltar para a lista de pacientes
            </button>
            <Relatorio @view-relatorios="showPatientReport" user-type="profissional" :patient-id="selectedPatientId" />
          </div>
        </div>

        <!-- Mostra outras views -->
        <GerarConvite v-if="activeView === 'convite'" />
        <EditarPerfil v-if="activeView === 'editar-perfil'" user-type="profissional" />
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useUserStore } from '../../store/user';
import TopNavbar from '../../components/layout/TopNavbar.vue';
import SidebarProfissional from '../../components/layout/SidebarProfissional.vue';
import ListaPacientes from './ListaPacientes.vue';
import Relatorio from '../shared/Relatorio.vue';
import GerarConvite from './GerarConvite.vue';
import EditarPerfil from '../shared/EditarPerfil.vue';

const userStore = useUserStore();
const activeView = ref('pacientes'); // Visão inicial
const selectedPatientId = ref(null);

const handleNavigation = (view) => {
  activeView.value = view;
  selectedPatientId.value = null; // Reseta a seleção de paciente ao navegar
};

const showPatientReport = (patientId) => {
  selectedPatientId.value = patientId;
  // A view já é 'pacientes', então apenas a condição do v-if/v-else muda
};

const showPatientList = () => {
  selectedPatientId.value = null;
};

const handleLogout = () => {
  userStore.logout();
};

onMounted(async () => {
  // Fetch user data if not already loaded
  if (!userStore.user) {
    await userStore.fetchUser('profissional');
  }
});
</script>

