<template>
  <div class="min-h-screen bg-gray-50">
    <header class="w-full p-6">
      <router-link to="/" class="text-3xl font-bold text-emerald-600">MindTrace</router-link>
    </header>

    <div class="flex items-center justify-center px-4 py-8">
      <div class="w-full max-w-2xl">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-2xl font-semibold text-center text-gray-900 mb-8">Criar Conta</h2>

          <form @submit.prevent="handleRegister" class="space-y-6">
            <div class="space-y-6">
              <h3 class="text-lg font-medium text-gray-900 border-b pb-2">Informações Pessoais</h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label for="fullName" class="block text-sm font-medium text-gray-700 mb-2">Nome Completo</label>
                  <input type="text" id="fullName" v-model="form.fullName" placeholder="Nome Completo" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="birthDate" class="block text-sm font-medium text-gray-700 mb-2">Data de Nascimento</label>
                  <input type="date" id="birthDate" v-model="form.birthDate" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors text-gray-900" required />
                </div>
                <div>
                  <label for="cpf" class="block text-sm font-medium text-gray-700 mb-2">CPF</label>
                  <input type="text" id="cpf" v-model="form.cpf" placeholder="000.000.000-00" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="phone" class="block text-sm font-medium text-gray-700 mb-2">Telefone</label>
                  <input type="tel" id="phone" v-model="form.phone" placeholder="(XX) XXXXX-XXXX" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
              </div>
            </div>

            <div class="space-y-6">
              <h3 class="text-lg font-medium text-gray-900 border-b pb-2">Informações de Endereço</h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label for="cep" class="block text-sm font-medium text-gray-700 mb-2">CEP</label>
                  <input type="text" id="cep" v-model="form.cep" placeholder="00000-000" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="street" class="block text-sm font-medium text-gray-700 mb-2">Rua</label>
                  <input type="text" id="street" v-model="form.street" placeholder="Rua Exemplo" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="number" class="block text-sm font-medium text-gray-700 mb-2">Número</label>
                  <input type="text" id="number" v-model="form.number" placeholder="123" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="neighborhood" class="block text-sm font-medium text-gray-700 mb-2">Bairro</label>
                  <input type="text" id="neighborhood" v-model="form.neighborhood" placeholder="Bairro Exemplo" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="city" class="block text-sm font-medium text-gray-700 mb-2">Cidade</label>
                  <input type="text" id="city" v-model="form.city" placeholder="Cidade Exemplo" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="state" class="block text-sm font-medium text-gray-700 mb-2">Estado</label>
                  <input type="text" id="state" v-model="form.state" placeholder="Estado Exemplo" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
              </div>
            </div>

            <div class="space-y-6">
              <h3 class="text-lg font-medium text-gray-900 border-b pb-2">Informações da Conta</h3>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label for="email" class="block text-sm font-medium text-gray-700 mb-2">Email</label>
                  <input type="email" id="email" v-model="form.email" placeholder="Digite seu e-mail" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
                <div>
                  <label for="password" class="block text-sm font-medium text-gray-700 mb-2">Senha</label>
                  <div class="relative">
                    <input :type="passwordFieldType" id="password" v-model="form.password" placeholder="Crie sua senha" class="w-full px-4 py-3 pr-12 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                    <button type="button" @click="togglePasswordVisibility" class="absolute right-4 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors">
                      <svg v-if="!showPassword" class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.878 9.878L3 3m6.878 6.878L21 21"></path>
                      </svg>
                      <svg v-else class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path>
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"></path>
                      </svg>
                    </button>
                  </div>
                </div>
                <div>
                  <label for="confirmPassword" class="block text-sm font-medium text-gray-700 mb-2">Confirme sua Senha</label>
                  <input type="password" id="confirmPassword" v-model="form.confirmPassword" placeholder="Confirme sua senha" class="w-full px-4 py-3 rounded-lg border border-gray-300 focus:ring-2 focus:ring-emerald-500 focus:border-emerald-500 outline-none transition-colors" required />
                </div>
              </div>
            </div>

            <button type="submit" class="w-full bg-emerald-600 hover:bg-emerald-700 text-white font-medium py-3 px-4 rounded-lg transition-colors duration-200 focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 outline-none">
              Criar Conta
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
import { ref, reactive, computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const form = reactive({
  fullName: '',
  birthDate: '',
  cpf: '',
  phone: '',
  cep: '',
  street: '',
  number: '',
  neighborhood: '',
  city: '',
  state: '',
  email: '',
  password: '',
  confirmPassword: ''
});

const showPassword = ref(false);

const passwordFieldType = computed(() => (showPassword.value ? 'text' : 'password'));

const togglePasswordVisibility = () => {
  showPassword.value = !showPassword.value;
};

const handleRegister = () => {
  if (form.password !== form.confirmPassword) {
    alert('As senhas não coincidem!');
    return;
  }

  console.log('Dados do formulário de cadastro:', form);
  alert('Cadastro realizado com sucesso! Bem-vindo(a), ' + form.fullName + '!');
  router.push('/login');
};
</script>
