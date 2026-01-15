<template>
  <header class="bg-white border-b border-gray-200 sticky top-0 z-40 shadow-sm">
    <nav class="px-4 sm:px-6 py-3 flex justify-between items-center">
      <!-- Logo da aplicacao -->
      <router-link to="/" class="flex items-center space-x-2">
        <font-awesome-icon :icon="faBrain" class="text-rose-300 text-xl sm:text-2xl" />
        <span class="text-2xl sm:text-3xl font-bold text-emerald-600 whitespace-nowrap">MindTrace</span>
      </router-link>

      <!-- Acoes de perfil e sessao -->
      <div class="flex items-center space-x-3 sm:space-x-4">
        <!-- Botao de perfil com dropdown -->
        <div class="relative" ref="profileDropdownContainer">
          <button 
            @click="toggleProfileDropdown"
            class="flex items-center space-x-2 p-2 rounded-lg hover:bg-gray-100 transition-colors"
          >
            <font-awesome-icon 
              :icon="userType === TipoUsuario.Profissional ? faUserDoctor : faUser" 
              class="w-8 h-8 p-2 rounded-full bg-emerald-100 text-emerald-600"
            />
            <div v-if="user" class="hidden sm:block text-left">
              <p class="text-sm font-semibold text-gray-900 leading-tight">{{ userName }}</p>
              <p class="text-xs text-gray-500">{{ userType === TipoUsuario.Profissional ? 'Profissional' : 'Paciente' }}</p>
            </div>
            <font-awesome-icon :icon="faChevronDown" class="w-3 h-3 text-gray-400 hidden sm:block" />
          </button>

          <!-- Dropdown com dados do perfil -->
          <transition
            enter-active-class="transition ease-out duration-100"
            enter-from-class="transform opacity-0 scale-95"
            enter-to-class="transform opacity-100 scale-100"
            leave-active-class="transition ease-in duration-75"
            leave-from-class="transform opacity-100 scale-100"
            leave-to-class="transform opacity-0 scale-95"
          >
            <div 
              v-if="isProfileDropdownOpen && user"
              @click.stop
              class="absolute right-0 mt-2 w-72 bg-white rounded-lg shadow-xl border border-gray-200 py-2"
            >
              <!-- Dados principais do perfil -->
              <div class="px-4 py-3 border-b border-gray-100">
                <div class="flex items-center space-x-3">
                  <font-awesome-icon 
                    :icon="userType === TipoUsuario.Profissional ? faUserDoctor : faUser" 
                    class="w-10 h-10 p-2 rounded-full bg-emerald-100 text-emerald-600"
                  />
                  <div class="flex-1">
                    <h3 class="font-bold text-base text-gray-900">{{ userName }}</h3>
                    <p class="text-sm text-gray-600">{{ userEmail }}</p>
                  </div>
                </div>
              </div>

              <!-- Detalhes complementares -->
              <div class="px-4 py-3 space-y-2 text-sm border-b border-gray-100">
                <template v-if="userType === TipoUsuario.Paciente">
                  <p><strong class="font-medium text-gray-700">Idade:</strong> <span class="text-gray-600">{{ calculateAge(user.data_nascimento) }}</span></p>
                  <p><strong class="font-medium text-gray-700">Contato:</strong> <span class="text-gray-600">{{ userContato }}</span></p>
                  <p v-if="user.dependente"><strong class="font-medium text-gray-700">Responsável:</strong> <span class="text-gray-600">{{ user.nome_responsavel }}</span></p>
                </template>
                <template v-else>
                  <p><strong class="font-medium text-gray-700">Especialidade:</strong> <span class="text-gray-600">{{ user.especialidade || 'Não informado' }}</span></p>
                  <p><strong class="font-medium text-gray-700">Registro:</strong> <span class="text-gray-600">{{ user.registro_profissional || 'Não informado' }}</span></p>
                  <p><strong class="font-medium text-gray-700">Contato:</strong> <span class="text-gray-600">{{ userContato }}</span></p>
                </template>
              </div>

              <!-- Acoes rapidas -->
              <div class="py-1">
                <button 
                  @click="handleEditProfile"
                  class="w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-50 flex items-center transition-colors"
                >
                  <font-awesome-icon :icon="faPenToSquare" class="w-4 h-4 mr-3 text-gray-500" />
                  Editar Perfil
                </button>
              </div>
            </div>
          </transition>
        </div>

  <!-- Botao de sair -->
        <button 
          @click="handleLogout"
          class="flex items-center space-x-2 px-4 py-2 text-sm font-medium text-white bg-rose-500 hover:bg-rose-600 rounded-lg transition-colors shadow-sm"
        >
          <font-awesome-icon :icon="faRightFromBracket" class="w-4 h-4" />
          <span class="hidden sm:inline">Sair</span>
        </button>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { 
  faBrain, 
  faUser, 
  faUserDoctor, 
  faChevronDown, 
  faPenToSquare, 
  faRightFromBracket 
} from '@fortawesome/free-solid-svg-icons';
import { useUserStore } from '../../store/user';
import { TipoUsuario } from '../../types/usuario.js';

// Propriedades esperadas pelo layout
const props = defineProps({
  userType: {
    type: String,
    required: true,
    validator: (value) => [TipoUsuario.Paciente, TipoUsuario.Profissional].includes(value)
  }
});

// Eventos de interacao emitidos para o container
const emit = defineEmits(['edit-profile', 'logout']);

// Estado global do usuario
const userStore = useUserStore();

// Estado do dropdown de perfil
const isProfileDropdownOpen = ref(false);
const profileDropdownContainer = ref(null);

// Usuario logado obtido da store
const user = computed(() => userStore.user);

// Computed properties para acessar dados do usuário independente da estrutura
// O backend retorna dados aninhados em 'usuario' para paciente/profissional
const userName = computed(() => {
  if (!user.value) return '';
  // Verifica se existe estrutura aninhada (PacienteDTOOut/ProfissionalDTOOut)
  return user.value.usuario?.nome || user.value.nome || '';
});

const userEmail = computed(() => {
  if (!user.value) return '';
  return user.value.usuario?.email || user.value.email || '';
});

const userContato = computed(() => {
  if (!user.value) return 'Não informado';
  return user.value.usuario?.contato || user.value.contato || 'Não informado';
});

// Alterna exibicao do dropdown de perfil
const toggleProfileDropdown = () => {
  isProfileDropdownOpen.value = !isProfileDropdownOpen.value;
};

// Fecha dropdown do perfil
const closeProfileDropdown = () => {
  isProfileDropdownOpen.value = false;
};

// Dispara evento para edicao de perfil
const handleEditProfile = () => {
  closeProfileDropdown();
  emit('edit-profile');
};

// Dispara evento para encerrar sessao
const handleLogout = () => {
  closeProfileDropdown();
  emit('logout');
};

// Calcula idade estimada a partir da data de nascimento
const calculateAge = (birthDate) => {
  if (!birthDate) return 'N/A';
  const today = new Date();
  const birth = new Date(birthDate);
  let age = today.getFullYear() - birth.getFullYear();
  const monthDiff = today.getMonth() - birth.getMonth();
  if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birth.getDate())) {
    age--;
  }
  return age;
};

// Fecha dropdown quando clicar fora do container
const handleClickOutside = (event) => {
  if (profileDropdownContainer.value && !profileDropdownContainer.value.contains(event.target)) {
    closeProfileDropdown();
  }
};

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
  // Recupera dados do usuario caso estejam ausentes
  if (!user.value) {
    userStore.fetchUser(props.userType);
  }
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>
