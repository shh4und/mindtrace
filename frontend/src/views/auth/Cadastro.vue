<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navbar publica -->
    <NavbarPublic :show-menu="false" />

    <div class="flex items-center justify-center px-4 py-8">
      <div class="w-full max-w-2xl">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-2xl font-semibold text-center text-gray-900 mb-6">Criar Conta</h2>

          <!-- Seletor tipo conta -->
          <fieldset class="mb-8">
            <legend class="sr-only">Tipo de conta</legend>
            <div class="grid grid-cols-2 gap-4" role="radiogroup" aria-label="Selecione o tipo de conta">
              <button 
                type="button"
                @click="form.userType = TipoUsuario.Paciente"
                :aria-pressed="form.userType === TipoUsuario.Paciente"
                :class="['p-4 rounded-lg border-2 text-center transition', form.userType === TipoUsuario.Paciente ? 'border-emerald-500 bg-emerald-50' : 'border-gray-300 bg-white hover:border-emerald-400']"
              >
                <span class="text-lg font-medium">Sou Paciente</span>
              </button>
              <button 
                type="button"
                @click="form.userType = TipoUsuario.Profissional"
                :aria-pressed="form.userType === TipoUsuario.Profissional"
                :class="['p-4 rounded-lg border-2 text-center transition', form.userType === TipoUsuario.Profissional ? 'border-rose-500 bg-rose-50' : 'border-gray-300 bg-white hover:border-rose-400']"
              >
                <span class="text-lg font-medium">Sou Profissional</span>
              </button>
            </div>
          </fieldset>

          <form v-if="form.userType" @submit.prevent="handleRegister" class="space-y-6" novalidate>
            <!-- Campos comuns -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label for="nome" class="block text-base font-medium text-gray-700 mb-1">Nome Completo</label>
                <input type="text" id="nome" v-model="form.nome" autocomplete="name" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
              </div>
              <div>
                <label for="cpf" class="block text-base font-medium text-gray-700 mb-1">CPF</label>
                <input type="text" id="cpf" v-model="form.cpf" inputmode="numeric" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
              </div>
              <div>
                <label for="email" class="block text-base font-medium text-gray-700 mb-1">Email</label>
                <input type="email" id="email" v-model="form.email" autocomplete="email" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
              </div>
              <div>
                <label for="senha" class="block text-base font-medium text-gray-700 mb-1">Senha</label>
                <input 
                  type="password" 
                  id="senha" 
                  v-model="form.senha" 
                  @input="validatePassword"
                  autocomplete="new-password"
                  :aria-invalid="!!passwordError"
                  :aria-describedby="passwordError ? 'senha-error senha-hint' : 'senha-hint'"
                  class="w-full px-4 py-3 rounded-lg border outline-none transition-colors focus:ring-2 focus:ring-emerald-500/20"
                  :class="passwordError ? 'border-red-300 focus:border-red-500' : 'border-gray-300 focus:border-emerald-500'"
                  required 
                  aria-required="true" 
                />
                <p v-if="passwordError" id="senha-error" class="text-sm text-red-600 mt-1" role="alert">{{ passwordError }}</p>
                <p id="senha-hint" class="text-sm text-gray-500 mt-1">Use 8+ caracteres com letras, números e símbolos como !@#$%^&*</p>
              </div>
              <div>
                <label for="confirm_password" class="block text-base font-medium text-gray-700 mb-1">Confirme sua Senha</label>
                <input 
                  type="password" 
                  id="confirm_password" 
                  v-model="form.confirm_password" 
                  autocomplete="new-password"
                  :aria-invalid="form.confirm_password && form.senha !== form.confirm_password"
                  :aria-describedby="form.confirm_password && form.senha !== form.confirm_password ? 'confirm-error' : undefined"
                  class="w-full px-4 py-3 rounded-lg border outline-none transition-colors focus:ring-2 focus:ring-emerald-500/20"
                  :class="form.confirm_password && form.senha !== form.confirm_password ? 'border-red-300 focus:border-red-500' : 'border-gray-300 focus:border-emerald-500'"
                  required 
                  aria-required="true"
                />
                <p v-if="form.confirm_password && form.senha !== form.confirm_password" id="confirm-error" class="text-sm text-red-600 mt-1" role="alert">As senhas não coincidem</p>
              </div>
              <div>
                <label for="contato" class="block text-base font-medium text-gray-700 mb-1">Número para Contato (Opcional)</label>
                <input type="tel" id="contato" v-model="form.contato" autocomplete="tel" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" />
              </div>
              <div>
                <label for="data_nascimento" class="block text-base font-medium text-gray-700 mb-1">Data de Nascimento</label>
                <input type="date" id="data_nascimento" v-model="form.data_nascimento" autocomplete="bday" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
              </div>
            </div>

            <!-- Campos profissional -->
            <fieldset v-if="form.userType === TipoUsuario.Profissional" class="pt-4 border-t border-gray-300">
              <legend class="sr-only">Informações profissionais</legend>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label for="especialidade" class="block text-base font-medium text-gray-700 mb-1">Especialidade</label>
                  <input type="text" id="especialidade" v-model="form.especialidade" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
                </div>
                <div>
                  <label for="registro_profissional" class="block text-base font-medium text-gray-700 mb-1">Nº Registro Profissional (CRP, etc)</label>
                  <input type="text" id="registro_profissional" v-model="form.registro_profissional" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
                </div>
              </div>
            </fieldset>

            <!-- Campos paciente -->
            <fieldset v-if="form.userType === TipoUsuario.Paciente" class="pt-4 border-t border-gray-300">
              <legend class="sr-only">Informações do paciente</legend>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="flex flex-row items-center gap-2">
                  <input id="dependente" v-model="form.dependente" type="checkbox" class="h-5 w-5 rounded border-gray-300 text-emerald-600 focus:ring-emerald-500" />
                  <label for="dependente" class="text-base font-medium text-gray-700">É dependente?</label>
                </div>

                <div v-if="form.dependente === true">
                  <label for="nome_responsavel" class="block text-base font-medium text-gray-700 mb-1">Nome do Responsável</label>
                  <input type="text" id="nome_responsavel" v-model="form.nome_responsavel" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
                </div>
                <div v-if="form.dependente === true">
                  <label for="contato_responsavel" class="block text-base font-medium text-gray-700 mb-1">Contato do Responsável</label>
                  <input id="contato_responsavel" v-model="form.contato_responsavel" type="tel" class="w-full px-4 py-3 rounded-lg border border-gray-300 outline-none transition-colors focus:border-emerald-500 focus:ring-2 focus:ring-emerald-500/20" required aria-required="true" />
                </div>
              </div>
            </fieldset>

            <button 
              type="submit" 
              :disabled="isSubmitDisabled"
              :aria-disabled="isSubmitDisabled"
              class="w-full font-medium py-3 px-4 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-offset-2 outline-none"
              :class="isSubmitDisabled ? 'bg-gray-400 cursor-not-allowed text-gray-200' : 'bg-emerald-600 hover:bg-emerald-700 text-white focus:ring-emerald-500'"
            >
              Criar Conta de {{ form.userType === TipoUsuario.Paciente ? 'Paciente' : 'Profissional' }}
            </button>
          </form>

          <div class="mt-6 text-center">
            <router-link to="/login" class="text-lg text-gray-600 hover:text-emerald-600 transition-colors">
              Já tem uma conta? Faça login
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useUserStore } from '@/store/user';
import { TipoUsuario } from '@/types/usuario.js';
import NavbarPublic from '@/components/layout/NavbarPublic.vue';

const router = useRouter();
const toast = useToast();
const userStore = useUserStore();

const form = reactive({
  userType: '', // Tipo selecionado paciente ou profissional
  nome: '',
  email: '',
  senha: '',
  confirm_password: '',
  cpf: '',
  contato: '',
  data_nascimento: new Date(),
  // Campos do profissional
  especialidade: '',
  registro_profissional: '',
  // Campos do paciente
  dependente: false,
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
      data_nascimento: new Date(form.data_nascimento).toJSON(),
    };

    let data;
    if (form.userType === TipoUsuario.Paciente) {
      data = {
        ...commonData,
        dependente: form.dependente,
        nome_responsavel: form.nome_responsavel,
        contato_responsavel: form.contato_responsavel,
      };
    } else if (form.userType === TipoUsuario.Profissional) {
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