<template>
  <div class="flex-1 p-4 md:p-8 w-full">
    <h1 class="text-3xl font-bold text-gray-900 mb-6">Meu Perfil</h1>

    <!-- Card de Informações do Perfil -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 md:p-8 mb-8">
      <h2 class="text-xl font-semibold text-gray-900 mb-6">Informações Pessoais</h2>
      <form @submit.prevent="saveProfile">
        <div class="space-y-6">
          <!-- Nome -->
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700">Nome</label>
            <input type="text" id="name" v-model="profile.name" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
          </div>

          <!-- Email -->
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input type="email" id="email" v-model="profile.email" class="mt-1 block w-full px-3 py-2 bg-gray-100 border border-gray-300 rounded-md shadow-sm sm:text-sm" readonly>
          </div>

          <!-- Especialidade (Apenas para Profissional) -->
          <div v-if="props.userType === 'profissional'">
            <label for="especialidade" class="block text-sm font-medium text-gray-700">Especialidade</label>
            <input type="text" id="especialidade" v-model="profile.especialidade" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
          </div>

          <!-- Bio -->
          <div>
            <label for="bio" class="block text-sm font-medium text-gray-700">Bio</label>
            <textarea id="bio" v-model="profile.bio" rows="4" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"></textarea>
          </div>
        </div>
        <div class="flex justify-end mt-6">
          <button type="submit" class="bg-indigo-600 text-white font-medium px-6 py-2 rounded-lg hover:bg-indigo-700 transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Salvar Alterações
          </button>
        </div>
      </form>
    </div>

    <!-- Card para Alterar Senha -->
    <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 md:p-8">
      <h2 class="text-xl font-semibold text-gray-900 mb-6">Alterar Senha</h2>
      <form @submit.prevent="changePassword" class="space-y-4">
        <div>
          <label for="current-password" class="block text-sm font-medium text-gray-700">Senha Atual</label>
          <input type="password" id="current-password" v-model="password.senha_atual" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
        </div>
        <div>
          <label for="new-password" class="block text-sm font-medium text-gray-700">Nova Senha</label>
          <input type="password" id="new-password" v-model="password.nova_senha" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
        </div>
        <div>
          <label for="confirm-password" class="block text-sm font-medium text-gray-700">Confirmar Nova Senha</label>
          <input type="password" id="confirm-password" v-model="password.nova_senha_re" class="mt-1 block w-full px-3 py-2 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm">
        </div>
        <div class="flex justify-end pt-2">
          <button type="submit" class="bg-indigo-500 text-white font-medium px-6 py-2 rounded-lg hover:bg-indigo-600 transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            Alterar Senha
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useToast } from 'vue-toastification';
import api from '../services/api';

const props = defineProps({
  userType: {
    type: String,
    required: true, // 'paciente' ou 'profissional'
  }
});

const toast = useToast();

// Padronizado para corresponder ao DTO do backend
const profile = ref({
  nome: '',
  email: '', // O email não é editável, mas será exibido
  contato: '',
  bio: '',
  // Campo específico do profissional
  especialidade: '' 
});

const password = ref({
  senha_atual: '',
  nova_senha: '',
  nova_senha_re: ''
});

// Carrega os dados do perfil quando o componente é montado
onMounted(async () => {
  try {
    const response = await api.buscarPerfil();
    const userData = response.data;
    profile.value.nome = userData.nome;
    profile.value.email = userData.email;
    profile.value.contato = userData.contato;
    profile.value.bio = userData.bio;

    // TODO: Adicionar lógica para buscar dados específicos do profissional se necessário
    // if (props.userType === 'profissional') { ... }

  } catch (error) {
    toast.error('Não foi possível carregar os dados do perfil.');
    console.error("Erro ao buscar perfil:", error);
  }
});

const saveProfile = async () => {
  try {
    const profileData = {
      nome: profile.value.nome,
      contato: profile.value.contato,
      bio: profile.value.bio,
    };
    await api.atualizarPerfil(profileData);
    toast.success('Perfil atualizado com sucesso!');
  } catch (error) {
    toast.error('Erro ao atualizar o perfil.');
    console.error("Erro ao salvar perfil:", error);
  }
};

const changePassword = async () => {
  if (password.value.nova_senha !== password.value.nova_senha_re) {
    toast.error('As novas senhas não coincidem.');
    return;
  }
  try {
    await api.alterarSenha(password.value);
    toast.success('Senha alterada com sucesso!');
    password.value = { senha_atual: '', nova_senha: '', nova_senha_re: '' };
  } catch (error) {
    toast.error('Não foi possível alterar a senha. Verifique sua senha atual.');
    console.error("Erro ao alterar senha:", error);
  }
};
</script>
