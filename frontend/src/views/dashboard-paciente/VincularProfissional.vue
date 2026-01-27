<template>
  <div class="flex justify-center items-start pt-10">
    <div
      class="w-full space-y-4 max-w-lg bg-white rounded-xl shadow-sm border border-gray-200 p-8"
    >
      <h1 class="text-3xl font-bold text-gray-900 mb-2 text-center">
        Vincular a um Profissional
      </h1>
      <p class="text-gray-500 mb-8 text-center">
        Insira o token de convite fornecido pelo seu profissional clique no link
        recebido por e-mail.
      </p>

      <BaseInput
        v-model="token"
        label="Código do Convite"
        placeholder="Cole o token aqui..."
        :disabled="loading"
        @input="verificarTokenManualmente"
      />
      <div
        v-if="dadosConvite"
        class="bg-emerald-50 p-4 rounded-lg border border-emerald-200 flex items-center gap-3"
      >
        <font-awesome-icon
          :icon="faUserDoctor"
          class="w-3.5 h-3.5 text-emerald-600 shrink-0"
        />
        <div>
          <p class="text-sm text-emerald-800 font-medium">
            Convite encontrado:
          </p>
          <p class="text-lg font-bold text-emerald-900">
            {{ dadosConvite.nome_profissional }}
          </p>
          <p class="text-xs text-emerald-700">
            {{ dadosConvite.especialidade }}
          </p>
        </div>
      </div>

      <p v-if="erro" class="text-red-500 text-sm mt-1">{{ erro }}</p>
      <BaseButton
        variant="emerald"
        full-width
        :loading="loadingVinculo"
        :disabled="!dadosConvite"
        @click="bindWithToken"
        >Confirmar Vínculo
      </BaseButton>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import { useRoute, useRouter } from "vue-router";
import api from "@/services/api";
import { useToast } from "vue-toastification";
import { BaseInput, BaseButton } from "@/components/ui";
import { debounce } from "lodash";
import { faUserDoctor } from "@fortawesome/free-solid-svg-icons";

const route = useRoute();
const router = useRouter();
const toast = useToast();

const token = ref("");
const dadosConvite = ref({
  nome_profissional: "",
  especialidade: "",
  valido: false,
});
const erro = ref("");
const loading = ref(false);
const loadingVinculo = ref(false);

onMounted(() => {
  const tokenUrl = route.query.token;
  if (tokenUrl) {
    token.value = tokenUrl;
    // router.replace({ query: null });
  }
});
/**
 * TODO:
 * verificar por possivel melhoria no uso do watch
 * // router.replace({ query: null }); e limpagem de url
 */

watch(token, async (newToken, oldToken) => {
  verificarTokenManualmente();
});

// Função que consulta a API para ver de quem é o convite
const buscarInfoToken = async (tokenValue) => {
  if (!tokenValue || tokenValue.length < 10) return; // Tamanho mínimo para evitar chamadas inúteis

  loading.value = true;
  erro.value = "";
  dadosConvite.value = null;

  try {
    const response = await api.consultarToken(tokenValue);
    dadosConvite.value = response.data;
  } catch (err) {
    erro.value = "Código inválido ou expirado.";
    dadosConvite.value = null;
  } finally {
    loading.value = false;
  }
};
// Debounce para não chamar a API a cada letra digitada (se for digitação manual)
const verificarTokenManualmente = debounce(() => {
  buscarInfoToken(token.value);
}, 500);

const bindWithToken = async () => {
  if (!token.value.trim()) {
    return;
  }

  loadingVinculo.value = true;
  try {
    await api.vincularComToken(token.value.trim());
    toast.success(
      `Você agora está vinculado a ${dadosConvite.value.nome_profissional}`
    );
    router.push({ name: "paciente-profissionais" });
  } catch (error) {
    toast.error(error.response?.data?.erro || "Erro ao vincular.");
    toast.error(errorMessage);
    console.error("Erro ao vincular com token:", error);
  } finally {
    loadingVinculo.value = false;
  }
};
</script>
