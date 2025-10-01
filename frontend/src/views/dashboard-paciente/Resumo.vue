<template>
  <div>
    <header class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900">Resumo</h1>
      <p class="text-gray-600 mt-1">Bem-vindo(a) de volta! Aqui está um resumo do seu bem-estar.</p>
    </header>

    <section class="grid grid-cols-1 md:grid-cols-2 gap-8 mb-8">
      <div class="bg-white p-6 rounded-lg shadow-md border border-gray-200">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Humor de Hoje</h2>
        <div class="flex items-center space-x-4">
          <span class="text-5xl">😊</span>
          <div>
            <p class="text-3xl font-bold text-green-600">Bom</p>
            <p class="text-gray-500 text-sm">Registro de hoje às 14:30h</p>
          </div>
        </div>
      </div>

      <div class="bg-white p-6 rounded-lg shadow-md border border-gray-200">
        <h2 class="text-xl font-semibold text-gray-900 mb-4">Alertas e Notificações</h2>
        <ul class="space-y-4 text-gray-700">
          <li class="flex items-center space-x-3 p-3 bg-red-50 rounded-md border border-red-200">
            <svg class="w-6 h-6 text-red-500" fill="currentColor" viewBox="0 0 20 20">
              <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
            </svg>
            <div>
              <p class="font-medium text-red-700">ALERTA: Padrão Preocupante</p>
              <p class="text-sm text-gray-600">Seu humor esteve baixo por 3 dias consecutivos.</p>
            </div>
          </li>
        </ul>
      </div>
    </section>

    <section class="bg-white p-6 rounded-lg shadow-md border border-gray-200">
      <h2 class="text-xl font-semibold text-gray-900 mb-4">Atividades Recentes</h2>
      <ul class="space-y-4">
        <li class="flex items-start space-x-4">
          <div class="flex-shrink-0 w-10 h-10 bg-gray-200 rounded-full flex items-center justify-center text-gray-600">
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-6 9l2 2 4-4"></path>
            </svg>
          </div>
          <div>
            <p class="font-medium text-gray-900">Questionário "Bem-Estar Semanal" respondido.</p>
            <p class="text-sm text-gray-500">12 de Agosto de 2025</p>
          </div>
        </li>
      </ul>
    </section>
  </div>
</template>

<script setup>
// Nenhuma lógica de script necessária para este componente estático por enquanto.
</script>
