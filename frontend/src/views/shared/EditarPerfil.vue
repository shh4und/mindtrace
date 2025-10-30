<template>
  <div class="flex flex-col items-center justify-center p-4 md:p-8 w-full">
    <h1 class="text-3xl font-bold text-gray-900 mb-6">Meu Perfil</h1>

    <!-- Card de Informações do Perfil -->
    <div class="w-full max-w-2xl bg-white rounded-xl shadow-sm border border-gray-200 p-6 md:p-8 mb-8">
      <h2 class="text-xl font-semibold text-gray-900 mb-6">Informações Pessoais</h2>
      <form @submit.prevent="saveProfile">
        <div class="space-y-6">
          <!-- Nome -->
          <div>
            <label for="name" class="block text-sm font-medium text-gray-700">Nome</label>
            <input type="text" id="name" v-model="profile.nome"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>

          <!-- Email -->
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input type="email" id="email" v-model="profile.email"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 bg-gray-100 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900"
              readonly>
          </div>

          <div>
            <label for="data_nascimento" class="block text-sm font-medium text-gray-700">Data de Nascimento</label>
            <input type="date" id="data_nascimento" v-model="profile.data_nascimento"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>


          <!-- (Apenas para Profissional) -->
          <div v-if="props.userType === TipoUsuario.Profissional">
            <label for="especialidade" class="block text-sm font-medium text-gray-700">Especialidade</label>
            <input type="text" id="especialidade" v-model="profile.especialidade"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>
          <div v-if="props.userType === TipoUsuario.Profissional">
            <label for="registro_profissional" class="block text-sm font-medium text-gray-700">Registro
              Profissional</label>
            <input type="text" id="registro_profissional" v-model="profile.registro_profissional"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>

          <!-- (Apenas para Paciente) -->
          <div v-if="props.userType === TipoUsuario.Paciente">
            <label for="dependente" class="block text-sm font-medium text-gray-700">Dependente</label>
            <input type="checkbox" id="dependente" v-model="profile.dependente"
              class="w-5 h-5 mt-3 px-4 py-5 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>
          <div v-if="props.userType === TipoUsuario.Paciente">
            <label for="nome_responsavel" class="block text-sm font-medium text-gray-700">Nome Responsável</label>
            <input type="text" id="nome_responsavel" v-model="profile.nome_responsavel"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>
          <div v-if="props.userType === TipoUsuario.Paciente">
            <label for="contato_responsavel" class="block text-sm font-medium text-gray-700">Contato Responsável</label>
            <input type="text" id="contato_responsavel" v-model="profile.contato_responsavel"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
          </div>

          <!-- Bio -->
          <div>
            <label for="bio" class="block text-sm font-medium text-gray-700">Bio</label>
            <textarea id="bio" v-model="profile.bio" rows="4"
              class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900"></textarea>
          </div>
        </div>
        <div class="flex justify-end mt-6">
          <button type="submit"
            class="bg-emerald-600 text-white font-medium px-6 py-2 rounded-lg hover:bg-emerald-700 transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-emerald-500">
            Salvar Alterações
          </button>
        </div>
      </form>
    </div>

    <!-- Card para Alterar Senha -->
    <div class="w-full max-w-2xl bg-white rounded-xl shadow-sm border border-gray-200 p-6 md:p-8">
      <h2 class="text-xl font-semibold text-gray-900 mb-6">Alterar Senha</h2>
      <form @submit.prevent="changePassword" class="space-y-4">
        <div>
          <label for="current-password" class="block text-sm font-medium text-gray-700">Senha Atual</label>
          <input type="password" id="current-password" v-model="password.senha_atual"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
        </div>
        <div>
          <label for="new-password" class="block text-sm font-medium text-gray-700">Nova Senha</label>
          <input type="password" id="new-password" v-model="password.nova_senha"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
        </div>
        <div>
          <label for="confirm-password" class="block text-sm font-medium text-gray-700">Confirmar Nova Senha</label>
          <input type="password" id="confirm-password" v-model="password.nova_senha_re"
            class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900">
        </div>
        <div class="flex justify-end pt-2">
          <button type="submit"
            class="bg-emerald-600 text-white font-medium px-6 py-2 rounded-lg hover:bg-emerald-700 transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-emerald-500">
            Alterar Senha
          </button>
        </div>
      </form>
    </div>
    <div class="w-full max-w-2xl bg-white rounded-xl shadow-sm border border-red-200 p-6 md:p-8 mt-8">
      <h2 class="text-xl font-semibold text-red-900 mb-6">Apagar Conta</h2>
      <p class="text-gray-700 mb-4">Esta ação é irreversível e removerá todos os seus dados.</p>
      <button type="button" @click="deleteAccount"
        class="bg-red-600 text-white font-medium px-6 py-2 rounded-lg hover:bg-red-700 transition-colors focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500">
        Apagar Conta
      </button>
    </div>
  </div>
  <!-- Card para Apagar Conta -->

</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useToast } from 'vue-toastification';
import api from '../../services/api';
import { useUserStore } from '../../store/user';
import { TipoUsuario } from '../../types/usuario.js';

const props = defineProps({
  userType: {
    type: String,
    required: true,
    validator: (value) => [TipoUsuario.Paciente, TipoUsuario.Profissional].includes(value)
  }
});

const toast = useToast();
const userStore = useUserStore();

// Padronizado para corresponder ao DTO do backend
const profile = ref({
  nome: '',
  email: '', // O email não é editável, mas será exibido
  contato: '',
  bio: '',
  data_nascimento: '',
  // Campo específico do profissional
  especialidade: ''
});

const password = ref({
  senha_atual: '',
  nova_senha: '',
  nova_senha_re: ''
});

onMounted(async () => {
  try {
    const response = await api.buscarPerfil();
    const userData = response.data;
    profile.value.nome = userData.nome;
    profile.value.email = userData.email;
    profile.value.contato = userData.contato;
    profile.value.bio = userData.bio;

    if (props.userType === TipoUsuario.Profissional) {
      const profResponse = await api.proprioPerfilProfissional();
      profile.value.especialidade = profResponse.data.especialidade;
      profile.value.registro_profissional = profResponse.data.registro_profissional;
      if (profResponse.data.data_nascimento) {
        profile.value.data_nascimento = profResponse.data.data_nascimento.split('T')[0];
      }
    } else if (props.userType === TipoUsuario.Paciente) {
      const pacResponse = await api.proprioPerfilPaciente();
      if (pacResponse.data.data_nascimento) {
        profile.value.data_nascimento = pacResponse.data.data_nascimento.split('T')[0];
      }
      profile.value.dependente = pacResponse.data.dependente;
      profile.value.nome_responsavel = pacResponse.data.nome_responsavel;
      profile.value.contato_responsavel = pacResponse.data.contato_responsavel;
    }

    toast.success('Perfil carregado com sucesso!');
  } catch (error) {
    toast.error('Erro ao carregar perfil.');
  }
});

const saveProfile = async () => {
  const profileData = {
    nome: profile.value.nome,
    contato: profile.value.contato,
    bio: profile.value.bio,
    data_nascimento: profile.value.data_nascimento,
  };
  if (props.userType === TipoUsuario.Profissional) {
    profileData.especialidade = profile.value.especialidade;
    profileData.registro_profissional = profile.value.registro_profissional;
  } else if (props.userType === TipoUsuario.Paciente) {
    profileData.dependente = profile.value.dependente;
    profileData.nome_responsavel = profile.value.nome_responsavel;
    profileData.contato_responsavel = profile.value.contato_responsavel;
  }
  try {
    await api.atualizarPerfil(profileData);
    toast.success('Perfil atualizado com sucesso!');
  } catch (error) {
    toast.error('Erro ao atualizar perfil.');
  }
};

const changePassword = async () => {
  if (password.value.nova_senha.length < 8) {
    toast.error('A nova senha deve ter pelo menos 8 caracteres.');
    return;
  }
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

const deleteAccount = async () => {
  const confirmed = confirm('Tem certeza de que deseja apagar sua conta? Esta ação não pode ser desfeita.');
  if (!confirmed) return;

  const result = await userStore.deleteAccount();
  if (result.success) {
    toast.success('Conta deletada com sucesso.');
  } else {
    toast.error(result.error);
  }
};
</script>