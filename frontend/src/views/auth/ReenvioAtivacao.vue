<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navbar publica -->
    <NavbarPublic :show-menu="false" />

    <div class="flex items-center justify-center px-4 mt-16">
      <div class="w-full max-w-md">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-3xl font-semibold text-center text-gray-900 mb-8">
            Reenvio de link de ativação
          </h2>

          <form @submit.prevent="handleReenvio" class="space-y-6">
            <!-- Campo email -->
            <BaseInput
              v-model="emailReenvio"
              type="email"
              label="E-mail"
              placeholder="Digite novamente seu e-mail para ativação."
              autocomplete="email"
              required
              size="lg"
            />

            <BaseButton type="submit" variant="emerald" size="lg" full-width>
              Reenviar link para e-mail
            </BaseButton>
          </form>

          <div class="mt-6 text-center space-y-3">
            <router-link
              to="/login"
              class="text-lg text-gray-600 hover:text-emerald-600 transition-colors"
            >
              Voltar para o Login
            </router-link>
            <h3 class="text-sm text-gray-600 mt-5">
              Verifique sua caixa de correspôndencia
            </h3>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import api from "@/services/api";
import { useToast } from "vue-toastification";
import { BaseInput, BaseButton } from "@/components/ui";

const toast = useToast();

import NavbarPublic from "@/components/layout/NavbarPublic.vue";

const emailReenvio = ref("");
const payload = {
  email: "",
};
const handleReenvio = async () => {
  try {
    payload.email = emailReenvio.value;
    await api.reenvioAtivacao(payload);

    toast.success("Email reenviado com sucesso!");
  } catch (error) {
    console.log(error);
  }
};
</script>
