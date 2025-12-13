import pluginVue from "eslint-plugin-vue";

export default [
  ...pluginVue.configs["flat/recommended"],
  {
    rules: {
      "vue/no-unused-vars": "error",
      "vue/multi-word-component-names": "off",
      "vue/require-default-prop": "off",
    },
  },
];
