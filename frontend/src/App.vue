<template>
  <div id="app">
    <!-- Skip Navigation Link para acessibilidade -->
    <a 
      href="#main-content" 
      class="sr-only focus:not-sr-only focus:absolute focus:top-4 focus:left-4 focus:z-50 focus:px-4 focus:py-2 focus:bg-emerald-600 focus:text-white focus:rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500"
    >
      Pular para o conteúdo principal
    </a>
    
    <!-- Anunciador para mudanças de rota (leitores de tela) -->
    <div 
      ref="routeAnnouncer"
      role="status" 
      aria-live="polite" 
      aria-atomic="true"
      class="sr-only"
    >
      {{ routeAnnouncement }}
    </div>
    
    <router-view v-slot="{ Component }">
      <transition name="page" mode="out-in">
        <component :is="Component" />
      </transition>
    </router-view>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute();
const routeAnnouncement = ref('');

// Mapa de nomes de rota para anúncios amigáveis
const routeNames = {
  'landpage': 'Página inicial',
  'login': 'Página de login',
  'cadastro': 'Página de cadastro',
  'forgot-password': 'Recuperação de senha',
  'paciente-resumo': 'Resumo do paciente',
  'paciente-humor': 'Registro de humor',
  'paciente-relatorios': 'Relatórios',
  'paciente-vincular': 'Vincular profissional',
  'paciente-editar-perfil': 'Editar perfil',
  'profissional-pacientes': 'Lista de pacientes',
  'profissional-paciente-relatorio': 'Relatório do paciente',
  'profissional-convite': 'Gerar convite',
  'profissional-editar-perfil': 'Editar perfil'
};

// Anuncia mudanças de rota para leitores de tela
watch(
  () => route.name,
  (newRouteName) => {
    if (newRouteName) {
      const pageName = routeNames[newRouteName] || 'Nova página';
      routeAnnouncement.value = `Navegou para: ${pageName}`;
    }
  }
);
</script>

<style>
/* Transição de página suave */
.page-enter-active,
.page-leave-active {
  transition: opacity 0.2s ease;
}

.page-enter-from,
.page-leave-to {
  opacity: 0;
}

/* Classe utilitária para elementos visíveis apenas para leitores de tela */
.sr-only {
  position: absolute;
  width: 1px;
  height: 1px;
  padding: 0;
  margin: -1px;
  overflow: hidden;
  clip: rect(0, 0, 0, 0);
  white-space: nowrap;
  border-width: 0;
}

.sr-only:focus,
.focus\:not-sr-only:focus {
  position: absolute;
  width: auto;
  height: auto;
  padding: inherit;
  margin: inherit;
  overflow: visible;
  clip: auto;
  white-space: normal;
}
</style>
