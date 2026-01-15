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
              <BaseInput
                v-model="form.nome"
                label="Nome Completo"
                autocomplete="name"
                required
              />
              <BaseInput
                v-model="form.cpf"
                label="CPF"
                inputmode="numeric"
                required
              />
              <BaseInput
                v-model="form.email"
                type="email"
                label="Email"
                autocomplete="email"
                required
              />
              <BaseInput
                v-model="form.senha"
                type="password"
                label="Senha"
                autocomplete="new-password"
                :error="passwordError"
                hint="Use 8+ caracteres com letras, números e símbolos como !@#$%^&*"
                required
                @blur="validatePassword"
              />
              <BaseInput
                v-model="form.confirm_password"
                type="password"
                label="Confirme sua Senha"
                autocomplete="new-password"
                :error="form.confirm_password && form.senha !== form.confirm_password ? 'As senhas não coincidem' : ''"
                required
              />
              <BaseInput
                v-model="form.contato"
                type="tel"
                label="Número para Contato (Opcional)"
                autocomplete="tel"
              />
              <BaseInput
                v-model="form.data_nascimento"
                type="date"
                label="Data de Nascimento"
                autocomplete="bday"
                required
              />
            </div>

            <!-- Campos profissional -->
            <fieldset v-if="form.userType === TipoUsuario.Profissional" class="pt-4 border-t border-gray-300">
              <legend class="sr-only">Informações profissionais</legend>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <BaseInput
                  v-model="form.especialidade"
                  label="Especialidade"
                  required
                />
                <BaseInput
                  v-model="form.registro_profissional"
                  label="Nº Registro Profissional (CRP, etc)"
                  required
                />
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

                <BaseInput
                  v-if="form.dependente === true"
                  v-model="form.nome_responsavel"
                  label="Nome do Responsável"
                  required
                />
                <BaseInput
                  v-if="form.dependente === true"
                  v-model="form.contato_responsavel"
                  type="tel"
                  label="Contato do Responsável"
                  required
                />
              </div>
            </fieldset>

            <BaseButton
              type="submit"
              :variant="form.userType === TipoUsuario.Profissional ? 'rose' : 'emerald'"
              size="lg"
              full-width
              :disabled="isSubmitDisabled"
            >
              Criar Conta de {{ form.userType === TipoUsuario.Paciente ? 'Paciente' : 'Profissional' }}
            </BaseButton>
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
import { BaseInput, BaseButton } from '@/components/ui';

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