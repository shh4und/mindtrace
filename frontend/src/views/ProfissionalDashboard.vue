<template>
  <div class="min-h-screen bg-gray-100 font-sans antialiased flex flex-col lg:flex-row">
    <NavbarProfissional @navigate="handleNavigation" />

    <main class="flex-1 p-4 sm:p-6 md:p-8 overflow-y-auto">
      <!-- Mostra a lista de pacientes ou o relatório de um paciente específico -->
      <div v-if="activeView === 'pacientes'">
        <ListaPacientes v-if="!selectedPatientId" @view-patient="showPatientReport" />
        <div v-else>
          <button @click="showPatientList" class="mb-6 flex items-center text-sm font-medium text-indigo-600 hover:text-indigo-800">
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
</template>

<script setup>
import { ref } from 'vue';
import NavbarProfissional from '../components/NavbarProfissional.vue';
import ListaPacientes from '../components/ListaPacientes.vue';
import Relatorio from '../components/Relatorio.vue';
import GerarConvite from '../components/GerarConvite.vue';
import EditarPerfil from '../components/EditarPerfil.vue';

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

</script>

