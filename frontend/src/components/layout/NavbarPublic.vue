<template>
  <header class="bg-white shadow-md sticky top-0 z-50">
    <nav class="px-6 py-4 flex justify-between items-center">
      <!-- Logo -->
      <router-link to="/" class="flex items-center space-x-2">
        <font-awesome-icon :icon="faBrain" class="text-rose-300 text-2xl" />
        <span class="text-3xl font-bold text-emerald-600 whitespace-nowrap">MindTrace</span>
      </router-link>

      <!-- Desktop Navigation -->
      <div class="flex items-center space-x-4">
        <ul v-if="showMenu" class="hidden md:flex space-x-8 items-center text-lg">
          <li>
            <a href="/#home" class="hover:text-emerald-600 transition duration-300">Início</a>
          </li>
          <li>
            <a href="/#features" class="hover:text-emerald-600 transition duration-300">Recursos</a>
          </li>
          <li>
            <a href="/#about" class="hover:text-emerald-600 transition duration-300">Sobre</a>
          </li>
          <li>
            <a href="/#contact" class="hover:text-emerald-600 transition duration-300">Contato</a>
          </li>
        </ul>
        
        <router-link 
          to="/login"
          class="hidden md:block bg-emerald-600 text-white px-6 py-2 rounded-full hover:bg-emerald-700 transition duration-300 font-semibold shadow-lg"
        >
          Acessar
        </router-link>

        <!-- Mobile Menu Toggle -->
        <div class="md:hidden">
          <button @click="toggleMenu" class="text-gray-600 focus:outline-none">
            <font-awesome-icon :icon="isMenuOpen ? faCaretUp : faCaretDown" class="w-7 h-7" />
          </button>
        </div>
      </div>
    </nav>

    <!-- Mobile Menu -->
    <div 
      v-if="showMenu"
      :class="['md:hidden', { 'hidden': !isMenuOpen }]" 
      class="bg-white shadow-lg border-t border-gray-200"
    >
      <ul class="px-6 py-4 space-y-4 text-center">
        <li>
          <a href="/#home" @click="closeMenu" class="block text-gray-800 hover:bg-emerald-50 p-3 rounded-lg transition">
            Início
          </a>
        </li>
        <li>
          <a href="/#features" @click="closeMenu" class="block text-gray-800 hover:bg-emerald-50 p-3 rounded-lg transition">
            Recursos
          </a>
        </li>
        <li>
          <a href="/#about" @click="closeMenu" class="block text-gray-800 hover:bg-emerald-50 p-3 rounded-lg transition">
            Sobre
          </a>
        </li>
        <li>
          <a href="/#contact" @click="closeMenu" class="block text-gray-800 hover:bg-emerald-50 p-3 rounded-lg transition">
            Contato
          </a>
        </li>
        <li>
          <router-link 
            to="/login" 
            @click="closeMenu"
            class="block bg-emerald-600 text-white p-3 rounded-lg mt-4 font-semibold hover:bg-emerald-700 transition"
          >
            Acessar
          </router-link>
        </li>
      </ul>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faBrain, faCaretUp, faCaretDown } from '@fortawesome/free-solid-svg-icons'; // Icones usados na navbar

// Propriedades de configuracao
defineProps({
  showMenu: {
    type: Boolean,
    default: true
  }
});

// Estado reativo do menu mobile
const isMenuOpen = ref(false);

// Alterna exibicao do menu mobile
const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value;
};

// Fecha o menu mobile e mantem navegacao limpa
const closeMenu = () => {
  isMenuOpen.value = false;
};
</script>
