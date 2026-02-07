<template>
  <div class="max-w-4xl mx-auto p-4 md:p-8">
    <!-- Header -->
    <header class="mb-10">
      <h1 class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight">
        Meu Perfil
      </h1>
      <p class="text-gray-500 mt-2 text-lg">
        Gerencie suas informações pessoais e preferências de conta.
      </p>
    </header>

    <!-- Card de Informações do Perfil -->
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-6">
      <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center gap-2">
        <span class="p-2 bg-indigo-100 rounded-lg text-indigo-600">
          <font-awesome-icon :icon="faUser" class="w-5 h-5" />
        </span>
        Informações Pessoais
      </h2>
      <form @submit.prevent="saveProfile">
        <div class="space-y-5">
          <!-- Nome -->
          <div>
            <label for="name" class="block text-sm font-bold text-gray-700 mb-1.5">Nome</label>
            <input
              type="text"
              id="name"
              v-model="profile.nome"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
            />
          </div>

          <!-- Email (readonly) -->
          <div>
            <label for="email" class="block text-sm font-bold text-gray-700 mb-1.5">Email</label>
            <input
              type="email"
              id="email"
              v-model="profile.email"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-gray-50 text-gray-500 outline-none cursor-not-allowed font-medium"
              readonly
            />
          </div>

          <!-- Data de Nascimento -->
          <div>
            <label for="data_nascimento" class="block text-sm font-bold text-gray-700 mb-1.5">Data de Nascimento</label>
            <input
              type="date"
              id="data_nascimento"
              v-model="profile.data_nascimento"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
            />
          </div>

          <!-- Especialidade (Profissional) -->
          <div v-if="props.userType === TipoUsuario.Profissional">
            <label for="especialidade" class="block text-sm font-bold text-gray-700 mb-1.5">Especialidade</label>
            <input
              type="text"
              id="especialidade"
              v-model="profile.especialidade"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
            />
          </div>

          <!-- Registro Profissional (Profissional) -->
          <div v-if="props.userType === TipoUsuario.Profissional">
            <label for="registro_profissional" class="block text-sm font-bold text-gray-700 mb-1.5">Registro Profissional</label>
            <input
              type="text"
              id="registro_profissional"
              v-model="profile.registro_profissional"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
            />
          </div>

          <!-- Dependente (Paciente) -->
          <div v-if="props.userType === TipoUsuario.Paciente" class="flex items-center gap-3">
            <input
              type="checkbox"
              id="dependente"
              v-model="profile.dependente"
              class="w-5 h-5 rounded-lg border-gray-300 text-emerald-600 focus:ring-emerald-500 transition-colors"
            />
            <label for="dependente" class="text-sm font-bold text-gray-700">Dependente</label>
          </div>

          <!-- Nome Responsável (Paciente) -->
          <div v-if="props.userType === TipoUsuario.Paciente">
            <label for="nome_responsavel" class="block text-sm font-bold text-gray-700 mb-1.5">Nome Responsável</label>
            <input
              type="text"
              id="nome_responsavel"
              v-model="profile.nome_responsavel"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-emerald-100 focus:border-emerald-500 outline-none transition-all text-gray-900 font-medium"
            />
          </div>

          <!-- Contato Responsável (Paciente) -->
          <div v-if="props.userType === TipoUsuario.Paciente">
            <label for="contato_responsavel" class="block text-sm font-bold text-gray-700 mb-1.5">Contato Responsável</label>
            <input
              type="text"
              id="contato_responsavel"
              v-model="profile.contato_responsavel"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-emerald-100 focus:border-emerald-500 outline-none transition-all text-gray-900 font-medium"
            />
          </div>

          <!-- Bio -->
          <div>
            <label for="bio" class="block text-sm font-bold text-gray-700 mb-1.5">Bio</label>
            <textarea
              id="bio"
              v-model="profile.bio"
              rows="4"
              class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium resize-none"
            ></textarea>
          </div>
        </div>

        <div class="flex justify-end mt-8">
          <button
            type="submit"
            class="px-8 py-3 bg-gradient-to-r from-indigo-600 to-violet-600 text-white font-bold rounded-xl hover:from-indigo-700 hover:to-violet-700 transition-all shadow-md hover:shadow-lg hover:-translate-y-0.5"
          >
            Salvar Alterações
          </button>
        </div>
      </form>
    </section>

    <!-- Card para Alterar Senha -->
    <section class="bg-white rounded-3xl shadow-sm border border-gray-100 p-6 md:p-8 mb-6">
      <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center gap-2">
        <span class="p-2 bg-amber-100 rounded-lg text-amber-600">
          <font-awesome-icon :icon="faLock" class="w-5 h-5" />
        </span>
        Alterar Senha
      </h2>
      <form @submit.prevent="changePassword" class="space-y-5">
        <div>
          <label for="current-password" class="block text-sm font-bold text-gray-700 mb-1.5">Senha Atual</label>
          <input
            type="password"
            id="current-password"
            v-model="password.senha_atual"
            class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
          />
        </div>
        <div>
          <label for="new-password" class="block text-sm font-bold text-gray-700 mb-1.5">Nova Senha</label>
          <input
            type="password"
            id="new-password"
            v-model="password.nova_senha"
            class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
          />
        </div>
        <div>
          <label for="confirm-password" class="block text-sm font-bold text-gray-700 mb-1.5">Confirmar Nova Senha</label>
          <input
            type="password"
            id="confirm-password"
            v-model="password.nova_senha_re"
            class="w-full px-4 py-3 rounded-2xl border border-gray-200 bg-white focus:ring-4 focus:ring-indigo-100 focus:border-indigo-500 outline-none transition-all text-gray-900 font-medium"
          />
        </div>
        <div class="flex justify-end pt-2">
          <button
            type="submit"
            class="px-8 py-3 bg-gradient-to-r from-indigo-600 to-violet-600 text-white font-bold rounded-xl hover:from-indigo-700 hover:to-violet-700 transition-all shadow-md hover:shadow-lg hover:-translate-y-0.5"
          >
            Alterar Senha
          </button>
        </div>
      </form>
    </section>

    <!-- Card para Apagar Conta -->
    <section class="bg-white rounded-3xl shadow-sm border border-red-100 p-6 md:p-8">
      <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center gap-2">
        <span class="p-2 bg-red-100 rounded-lg text-red-600">
          <font-awesome-icon :icon="faTrash" class="w-5 h-5" />
        </span>
        Apagar Conta
      </h2>
      <p class="text-gray-500 mb-6 font-medium">
        Esta ação é irreversível e removerá todos os seus dados permanentemente.
      </p>
      <button
        type="button"
        @click="deleteAccount"
        class="px-6 py-3 bg-red-600 text-white font-bold rounded-xl hover:bg-red-700 transition-all shadow-md hover:shadow-lg"
      >
        <font-awesome-icon :icon="faTrash" class="w-4 h-4 mr-2" />
        Apagar Conta
      </button>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useToast } from 'vue-toastification';
import api from '../../services/api';
import { useUserStore } from '../../store/user';
import { TipoUsuario } from '../../types/usuario.js';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import {
  faUser,
  faLock,
  faTrash,
} from '@fortawesome/free-solid-svg-icons';

const props = defineProps({
  userType: {
    type: String,
    required: true,
    validator: (value) => [TipoUsuario.Paciente, TipoUsuario.Profissional].includes(value)
  }
});

const toast = useToast();
const userStore = useUserStore();

const profile = ref({
  nome: '',
  email: '',
  contato: '',
  bio: '',
  data_nascimento: '',
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

  try {
    await userStore.deleteAccount();
  } catch (error) {
    toast.error('Erro ao apagar a conta.');
    console.error("Erro ao apagar conta:", error);
  }
};
</script>

<style scoped>
/* Mantendo consistência com o estilo global */
</style>
