<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navbar publica -->
    <NavbarPublic :show-menu="false" />

    <div class="flex items-center justify-center px-4 mt-16">
      <div class="w-full max-w-md">
        <div class="bg-white rounded-2xl shadow-sm border border-gray-200 p-8">
          <h2 class="text-3xl font-semibold text-center text-gray-900 mb-8">
            Conta ativada com sucesso!
          </h2>
          <div class="mt-6 text-center space-y-3">
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

import NavbarPublic from "@/components/layout/NavbarPublic.vue";
import api from "@/services/api";

const route = useRoute();
const token = ref(route.query.token);

onMounted(async () => {
  try {
    const tokenHash = token.value;
    await api.ativarConta(tokenHash);
  } catch (error) {
    console.log(error);
  }
});
</script>
