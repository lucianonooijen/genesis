module.exports = (api) => {
  api.cache(true);
  return {
    presets: ["babel-preset-expo"],
    plugins: [
      [
        "module-resolver",
        {
          alias: {
            config: "./src/config.ts",
            api: "./src/api",
            components: "./src/components",
            containers: "./src/containers",
            content: "./src/content",
            layouts: "./src/layouts",
            helpers: "./src/helpers",
            lang: "./src/lang",
            router: "./src/router",
            screens: "./src/screens",
            sections: "./src/sections",
            store: "./src/store",
            test: "./src/test",
          },
        },
      ],
    ],
  };
};
