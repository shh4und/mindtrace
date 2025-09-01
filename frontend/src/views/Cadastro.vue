<template>
  <div class="min-h-screen bg-gray-50">
    <header class="w-full p-6">
      <router-link to="/" class="items-center space-x-2 border-0 border-gray-200 rounded-lg p-2">
        <font-awesome-icon :icon="faBrain" class="text-rose-300 text-2xl" />
        <span class="text-3xl font-bold text-emerald-600 whitespace-nowrap">MindTrace</span>
      </router-link>
    </header>

    <div class="flex items-center justify-center px-4 py-8">
      <div class="w-full max-w-2xl">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-2xl font-semibold text-center text-gray-900 mb-6">Criar Conta</h2>

          <!-- Seletor de Tipo de Conta -->
          <div class="grid grid-cols-2 gap-4 mb-8">
            <button @click="form.userType = 'paciente'"
              :class="['p-4 rounded-lg border-2 text-center transition', form.userType === 'paciente' ? 'border-emerald-500 bg-emerald-50' : 'border-gray-300 bg-white hover:border-emerald-400']">
              <span class="text-lg font-medium">Sou Paciente</span>
            </button>
            <button @click="form.userType = 'profissional'"
              :class="['p-4 rounded-lg border-2 text-center transition', form.userType === 'profissional' ? 'border-rose-500 bg-rose-50' : 'border-gray-300 bg-white hover:border-rose-400']">
              <span class="text-lg font-medium">Sou Profissional</span>
            </button>
          </div>

          <form v-if="form.userType" @submit.prevent="handleRegister" class="space-y-6">
            <!-- Campos Comuns -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label for="nome" class="block text-base font-medium text-gray-700 mb-1">Nome Completo</label>
                <input type="text" id="nome" v-model="form.nome" class="w-full input-style" required />
              </div>
              <div>
                <label for="cpf" class="block text-base font-medium text-gray-700 mb-1">CPF</label>
                <input type="text" id="cpf" v-model="form.cpf" class="w-full input-style" required />
              </div>
              <div>
                <label for="email" class="block text-base font-medium text-gray-700 mb-1">Email</label>
                <input type="email" id="email" v-model="form.email" class="w-full input-style" required />
              </div>
              <div>
                <label for="senha" class="block text-base font-medium text-gray-700 mb-1">Senha</label>
                <input type="password" id="senha" v-model="form.senha" @input="validatePassword"
                  class="w-full input-style" required />
                <p v-if="passwordError" class="text-sm text-red-600 mt-1">{{ passwordError }}</p>
                <p class="text-sm text-gray-500 mt-1">Use 8+ caracteres com letras, números e símbolos como !@#$%^&*</p>
              </div>
              <div>
                <label for="confirm_password" class="block text-base font-medium text-gray-700 mb-1">Confirme sua
                  Senha</label>
                <input type="password" id="confirm_password" v-model="form.confirm_password" class="w-full input-style"
                  required />
              </div>
              <div>
                <label for="contato" class="block text-base font-medium text-gray-700 mb-1">Número para Contato
                  (Opcional)</label>
                <input type="tel" id="contato" v-model="form.contato" class="w-full input-style" />
              </div>
            </div>

            <!-- Campos do Profissional -->
            <div v-if="form.userType === 'profissional'"
              class="grid grid-cols-1 md:grid-cols-2 gap-6 pt-4 border-t  border-gray-300">
              <div>
                <label for="especialidade" class="block text-base font-medium text-gray-700 mb-1">Especialidade</label>
                <input type="text" id="especialidade" v-model="form.especialidade" class="w-full input-style"
                  required />
              </div>
              <div>
                <label for="registro_profissional" class="block text-base font-medium text-gray-700 mb-1">Nº Registro
                  Profissional (CRP, etc)</label>
                <input type="text" id="registro_profissional" v-model="form.registro_profissional"
                  class="w-full input-style" required />
              </div>
            </div>

            <!-- Campos do Paciente -->
            <div v-if="form.userType === 'paciente'"
              class="grid grid-cols-1 md:grid-cols-2 gap-6 pt-4 border-t border-gray-300">
              <div>
                <label for="idade" class="block text-base font-medium text-gray-700 mb-1">Idade</label>
                <input type="number" id="idade" v-model="form.idade" class="w-full input-style" required />
              </div>
              <div class="flex flex-row items-center gap-2">
                <label for="dependente" class="text-base font-medium text-gray-700 mb-1">É dependente?</label>
                <input id="dependente" v-model="form.dependente" type="checkbox"
                  class="h-5 w-5 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500" />
              </div>

              <div v-if="form.dependente === true">
                <label for="nome_responsavel" class="block text-base font-medium text-gray-700 mb-1">Nome do
                  Responsável</label>
                <input type="text" id="nome_responsavel" v-model="form.nome_responsavel" class="w-full input-style"
                  required />
              </div>
              <div v-if="form.dependente === true">
                <label for="contato_responsavel" class="block text-base font-medium text-gray-700 mb-1">Contato do
                  Responsável</label>
                <input id="contato_responsavel" v-model="form.contato_responsavel" type="tel" class="w-full input-style"
                  required />
              </div>

            </div>



            <button type="submit" :disabled="isSubmitDisabled"
              class="w-full font-medium py-3 px-4 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-offset-2 outline-none"
              :class="isSubmitDisabled ? 'bg-gray-400 cursor-not-allowed' : 'bg-emerald-600 hover:bg-emerald-700 text-white'">
              Criar Conta de {{ form.userType === 'paciente' ? 'Paciente' : 'Profissional' }}
            </button>
          </form>

          <div class="mt-6 text-center">
            <router-link to="/login" class="text-sm text-gray-600 hover:text-emerald-600 transition-colors">
              Já tem uma conta? Faça login
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useUserStore } from '../store/user';
import { faBrain, faCaretDown, faCaretUp } from '@fortawesome/free-solid-svg-icons';

const router = useRouter();
const toast = useToast();
const userStore = useUserStore();

const form = reactive({
  userType: '', // paciente ou profissional
  nome: '',
  email: '',
  senha: '',
  confirm_password: '',
  cpf: '',
  contato: '',
  // Campos do profissional
  especialidade: '',
  registro_profissional: '',
  // Campos do paciente
  dependente: false,
  idade: '',
  nome_responsavel: '',
  contato_responsavel: '',
});

const passwordError = ref('');

const isPasswordValid = computed(() => {
  if (!form.senha) return false;
  const regex = /^[a-zA-Z0-9!@#$%^&*]{8,}$/;
  return regex.test(form.senha);
});

const isSubmitDisabled = computed(() => {
  return !isPasswordValid.value || form.senha !== form.confirm_password;
});

const validatePassword = () => {
  if (!form.senha) {
    passwordError.value = '';
    return;
  }
  const regex = /^[a-zA-Z0-9!@#$%^&*]{8,}$/;
  if (!regex.test(form.senha)) {
    passwordError.value = 'A senha não atende aos critérios de segurança.';
  } else {
    passwordError.value = '';
  }
};

const handleRegister = async () => {
  if (form.senha !== form.confirm_password) {
    toast.error('As senhas não coincidem!');
    return;
  }
  if (!isPasswordValid.value) {
    toast.error('Por favor, use uma senha válida.');
    return;
  }

  try {
    const commonData = {
      nome: form.nome,
      email: form.email,
      senha: form.senha,
      cpf: form.cpf,
      contato: form.contato,
    };

    let data;
    if (form.userType === 'paciente') {
      data = {
        ...commonData,
        dependente: form.dependente,
        idade: form.idade,
        nome_responsavel: form.nome_responsavel,
        contato_responsavel: form.contato_responsavel,
      };
    } else if (form.userType === 'profissional') {
      data = {
        ...commonData,
        especialidade: form.especialidade,
        registro_profissional: form.registro_profissional,
      };
    }

    const result = await userStore.register(data, form.userType);

    if (result.success) {
      toast.success(result.message);
      router.push('/login');
    } else {
      toast.error(result.error);
    }

  } catch (error) {
    const errorMessage = error.response?.data?.erro || 'Erro desconhecido no cadastro.';
    toast.error(errorMessage);
    console.error('Falha no cadastro:', error);
  }
};
</script>

<style scoped>
.input-style {
  width: 100%;
  padding-left: 1rem;
  padding-right: 1rem;
  padding-top: 0.75rem;
  padding-bottom: 0.75rem;
  border-radius: 0.5rem;
  border: 1px solid #D1D5DB;
  /* gray-300 */
  outline: none;
  transition: color 0.2s ease, border-color 0.2s ease, box-shadow 0.2s ease;
}

.input-style:focus {
  border-color: #10B981;
  /* emerald-500 */
  box-shadow: 0 0 0 4px rgba(16, 185, 129, 0.12);
  /* subtle ring effect */
}
</style>