<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navbar publica -->
    <NavbarPublic :show-menu="false" />

    <div class="flex items-center justify-center px-4 mt-16">
      <div class="w-full max-w-md">
        <div
          class="bg-white space-y-4 rounded-2xl shadow-sm border border-gray-200 p-8"
        >
          <BaseButton variant="emerald" full-width @click="submitAtivacao"
            >Ativar Conta
          </BaseButton>

          <p v-if="erro" class="text-red-500 text-center text-sm">
            {{ erro }}
          </p>
          <p
            v-if="ativo"
            class="font-semibold text-center text-sm text-gray-900"
          >
            Conta ativada com sucesso!
          </p>
          <div class="mt-4 text-center space-y-3">
            <router-link
              to="/login"
              class="text-lg text-gray-600 hover:text-emerald-600 transition-colors"
            >
              Voltar para o Login
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter, useRoute } from "vue-router";
import { onMounted } from "vue";
import { BaseButton } from "@/components/ui";

import NavbarPublic from "@/components/layout/NavbarPublic.vue";
import api from "@/services/api";

const route = useRoute();
const token = ref(route.query.token);
const erro = ref("");
const ativo = ref(false);
const submitAtivacao = async () => {
  try {
    const tokenHash = token.value;
    await api.ativarConta(tokenHash);
    erro.value = "";
    ativo.value = true;
  } catch (error) {
    erro.value = "Código inválido ou expirado.";
    ativo.value = false;
    console.log(error);
  }
};
</script>
