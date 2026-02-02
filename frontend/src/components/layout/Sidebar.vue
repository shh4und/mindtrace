<template>
  <aside
    :class="[
      'fixed md:static inset-y-0 left-0 transform transition-transform duration-300 ease-in-out z-50',
      isOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0',
      isCollapsed ? 'md:w-20' : 'md:w-64',
      'w-64 bg-white border-r border-gray-200 flex flex-col shadow-lg md:shadow-none md:relative',
    ]"
    role="navigation"
    :aria-label="`Menu de navegação do ${variantLabel}`"
  >
    <button
      @click="toggleCollapse"
      :class="[
        'absolute -right-6 top-10 z-40 w-6 h-12 hidden md:flex items-center justify-center border border-l-0 border-gray-200 rounded-r-xl shadow-sm transition-all',
        variant === 'paciente'
          ? 'bg-emerald-50 text-emerald-600'
          : 'bg-rose-50 text-rose-700',
      ]"
      :aria-label="isCollapsed ? 'Expandir menu' : 'Recolher menu'"
    >
      <font-awesome-icon
        :icon="dynamicIcon"
        class="w-3 h-3 transition-transform group-hover:scale-110"
      />
    </button>

    <nav class="flex-1 px-3 pb-3 overflow-y-auto pt-6">
      <ul class="space-y-1" role="menubar">
        <li v-for="item in menuItems" :key="item.name" role="none">
          <button
            @click="handleNavigate(item.view)"
            role="menuitem"
            :class="[
              'w-full flex items-center px-3 py-3 rounded-lg font-medium transition-all duration-200',
              activeView === item.view
                ? activeItemClasses
                : 'text-gray-600 hover:bg-gray-50',
            ]"
            :title="isCollapsed ? item.label : ''"
          >
            <font-awesome-icon
              :icon="item.icon"
              :class="['w-5 h-5 shrink-0', isCollapsed ? 'mx-auto' : 'mr-3']"
            />
            <transition name="fade">
              <span v-if="!isCollapsed" class="truncate">{{ item.label }}</span>
            </transition>
          </button>
        </li>
      </ul>
    </nav>

    <div v-if="!isCollapsed" class="p-4 border-t border-gray-200 bg-gray-50">
      <div class="text-xs text-gray-500 text-center">
        <p class="font-medium">🧠 MindTrace</p>
        <p>{{ variantLabel }}</p>
      </div>
    </div>
  </aside>

  <Teleport to="body">
    <Transition name="overlay">
      <div
        v-if="isOpen"
        @click="closeSidebar"
        class="fixed inset-0 bg-black/50 z-40 md:hidden"
      ></div>
    </Transition>
  </Teleport>
</template>

<script setup>
import { ref, computed } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import {
  faChevronLeft,
  faChevronRight,
  faCheck,
} from "@fortawesome/free-solid-svg-icons";

const props = defineProps({
  menuItems: { type: Array, required: true },
  activeView: { type: String, default: "" },
  variant: { type: String, default: "paciente" },
  isOpen: { type: Boolean, default: false }, // Controla o estado de abertura no mobile
});

const emit = defineEmits(["navigate", "close"]);

const isCollapsed = ref(false); // Controla o estado de recolhimento no desktop

const dynamicIcon = computed(() => {
  return isCollapsed.value ? faChevronRight : faChevronLeft;
});

const toggleCollapse = () => {
  isCollapsed.value = !isCollapsed.value;
};

const closeSidebar = () => {
  emit("close");
};

const handleNavigate = (view) => {
  emit("navigate", view);
  closeSidebar();
};

const activeItemClasses = computed(() =>
  props.variant === "paciente"
    ? "bg-emerald-50 text-emerald-700"
    : "bg-rose-50 text-rose-700"
);

const variantLabel = computed(() =>
  props.variant === "paciente" ? "Paciente" : "Profissional"
);
</script>
