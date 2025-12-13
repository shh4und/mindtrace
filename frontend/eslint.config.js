import pluginVue from "eslint-plugin-vue";
import js from "@eslint/js";
import prettier from "eslint-config-prettier";

export default [
  js.configs.recommended,
  ...pluginVue.configs["flat/recommended"],
  prettier,
  {
    ignores: ["node_modules/", "dist/"],
    rules: {
      "vue/no-unused-vars": "warning",
      "vue/multi-word-component-names": "off",
      "vue/require-default-prop": "off",
    },
  },
];
