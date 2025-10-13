<template>
  <div class="bg-gray-50 text-gray-800">
    <!-- Cabecalho -->
    <NavbarPublic />

    <main>
      <!-- Secao Hero -->
      <section id="home" class="relative py-24 md:py-40 bg-gradient-to-br from-emerald-100 to-rose-100">
        <div class="container mx-auto px-6 text-center">
          <h1 class="text-4xl md:text-5xl lg:text-6xl font-extrabold leading-tight tracking-wide text-gray-900">
            Sua mente é um universo. <span class="text-emerald-600">Explore-o.</span>
          </h1>
          <p class="mt-6 md:mt-8 text-lg md:text-xl max-w-3xl mx-auto text-gray-700">
            Um diário emocional para te ajudar a se entender melhor e a se conectar com quem pode te ajudar.
          </p>
          <a href="#features"
            class="mt-10 inline-block bg-emerald-600 text-white font-bold py-4 px-10 rounded-full shadow-lg hover:bg-emerald-700 transition duration-300 transform hover:scale-105">
            Comece sua Jornada
          </a>
        </div>
      </section>

  <!-- Secao recursos -->
      <section id="features" class="py-24 bg-white">
        <div class="container mx-auto px-6 text-center">
          <h2 class="text-3xl md:text-4xl font-bold text-gray-900">Como o MindTrace te Ajuda</h2>
          <p class="mt-4 text-lg text-gray-600 max-w-2xl mx-auto">
            Ferramentas intuitivas e divertidas para o seu bem-estar emocional.
          </p>
          <div class="mt-16 grid grid-cols-1 md:grid-cols-3 gap-10">
            <!-- Recurso 1 -->
            <div
              class="feature-card bg-gray-50 p-8 rounded-2xl shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl">
              <div class="w-16 h-16 bg-emerald-100 rounded-full flex items-center justify-center mx-auto mb-6">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-emerald-600" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round"
                    d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Monitoramento de Humor</h3>
              <p class="mt-2 text-gray-600">
                Anote como você se sente, de forma rápida e fácil. Com o tempo, você vai descobrir padrões e gatilhos
                emocionais.
              </p>
            </div>
            <!-- Recurso 2 -->
            <div
              class="feature-card bg-gray-50 p-8 rounded-2xl shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl">
              <div class="w-16 h-16 bg-rose-100 rounded-full flex items-center justify-center mx-auto mb-6">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-rose-600" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Análises e Relatórios</h3>
              <p class="mt-2 text-gray-600">
                Veja sua jornada em gráficos intuitivos. Informação valiosa para você e para seu terapeuta.
              </p>
            </div>
            <!-- Recurso 3 -->
            <div
              class="feature-card bg-gray-50 p-8 rounded-2xl shadow-lg transform transition duration-500 hover:scale-105 hover:shadow-2xl">
              <div class="w-16 h-16 bg-amber-100 rounded-full flex items-center justify-center mx-auto mb-6">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-amber-600" fill="none" viewBox="0 0 24 24"
                  stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round"
                    d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a2 2 0 01-2-2V7a2 2 0 012-2h2.586a1 1 0 01.707.293l2.414 2.414a1 1 0 00.707.293H17z" />
                </svg>
              </div>
              <h3 class="text-xl font-bold text-gray-900">Conexão com Profissionais</h3>
              <p class="mt-2 text-gray-600">
                Conecte-se com profissionais da psicologia de forma segura e descomplicada.
              </p>
            </div>
          </div>
        </div>
      </section>

  <!-- Secao sobre -->
      <section id="about" class="py-24 bg-gray-100">
        <div class="container mx-auto px-6">
          <div class="flex flex-col md:flex-row items-center justify-between">
            <div class="md:w-1/2 mb-10 md:mb-0 text-center md:text-left">
              <h2 class="text-3xl md:text-4xl font-bold text-gray-900">Sua mente, seu universo.</h2>
              <p class="mt-4 text-lg text-gray-600">
                MindTrace nasceu da ideia de que a tecnologia pode ser uma grande aliada no caminho do autoconhecimento.
                Nossa missão é oferecer um espaço seguro e acolhedor para que você possa explorar seus sentimentos,
                entender suas emoções e cuidar da sua saúde mental com leveza e curiosidade.
              </p>
            </div>
            <div class="md:w-1/2 flex justify-center">
              <div
                class="w-80 h-80 bg-gradient-to-br from-emerald-300 to-rose-300 rounded-full flex items-center justify-center shadow-2xl">
                <div class="w-72 h-72 bg-gray-50 rounded-full flex items-center justify-center">
                  <font-awesome-icon :icon="faBrain" class="text-[1000%] text-rose-300" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

  <!-- Secao contato -->
      <section id="contact" class="bg-gradient-to-r from-emerald-600 to-green-500 text-white py-24">
        <div class="container mx-auto px-6 text-center">
          <h2 class="text-3xl md:text-4xl font-bold">Pronto para começar sua jornada?</h2>
          <p class="mt-4 text-lg max-w-2xl mx-auto opacity-90">
            Crie sua conta e comece a trilhar um caminho de autoconhecimento e bem-estar. É gratuito para começar.
          </p>
          <router-link to="/cadastro"
            class="mt-10 inline-block bg-white text-emerald-600 font-bold py-4 px-10 rounded-full shadow-lg hover:bg-gray-100 transition duration-300 transform hover:scale-105">
            Crie sua conta gratuita
          </router-link>
        </div>
      </section>
    </main>

    <footer class="bg-gray-800 text-gray-300 py-8">
      <div class="container mx-auto px-6 text-center">
        <p>&copy; 2025 MindTrace. Todos os direitos reservados.</p>
        <div class="mt-4 space-x-4">
          <a href="#" class="hover:text-white transition duration-300">Política de Privacidade</a>
          <a href="#" class="hover:text-white transition duration-300">Termos de Serviço</a>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import NavbarPublic from '../components/layout/NavbarPublic.vue';
import { faBrain, } from '@fortawesome/free-solid-svg-icons';

</script>

<style>
/* Estilos para o menu de navegacao responsivo */
#mobile-menu {
  transition: all 0.3s ease-in-out;
  transform: translateY(-100%);
  opacity: 0;
}

/* Estado ativo do menu mobile */
#mobile-menu.active {
  transform: translateY(0);
  opacity: 1;
}

/* Animacao para os cards de feature */
.feature-card {
  animation: fadeInUp 0.6s ease-in-out forwards;
  opacity: 0;
}

.feature-card:nth-child(1) {
  animation-delay: 0.2s;
}

.feature-card:nth-child(2) {
  animation-delay: 0.4s;
}

.feature-card:nth-child(3) {
  animation-delay: 0.6s;
}

@keyframes fadeInUp {
  from {
    transform: translateY(30px);
    opacity: 0;
  }

  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>
