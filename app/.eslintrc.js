module.exports = {
  parser: "@typescript-eslint/parser",
  extends: [
    "plugin:@typescript-eslint/recommended",
    "airbnb",
    "plugin:prettier/recommended",
    "prettier",
    "prettier/react",
    "prettier/@typescript-eslint",
  ],
  plugins: ["security", "prettier"],
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: "module",
    ecmaFeatures: {
      jsx: true,
    },
  },
  settings: {
    "import/resolver": {
      node: {
        paths: ["src"],
        extensions: [".js", ".jsx", ".ts", ".tsx"],
      },
    },
  },
  globals: {
    JSX: true,
  },
  rules: {
    "no-shadow": 0, // For allowing headacheless Redux Action in props
    "no-use-before-define": 0, // Allow sane file layouts
    "import/prefer-default-export": 0,
    "import/extensions": 0,
    "prettier/prettier": ["error"],
    "react/prop-types": 0, // Done by using React.FC<PropsInterface>
    "react/jsx-filename-extension": [1, { extensions: [".tsx", ".jsx"] }],
    "react/button-has-type": 0,
    "global-require": 0,
    "@typescript-eslint/explicit-function-return-type": 0,
    "@typescript-eslint/explicit-module-boundary-types": 0,
    "@typescript-eslint/no-empty-function": 0,
    camelcase: 0,
    "@typescript-eslint/camelcase": 0,
    "no-unused-vars": "off",
    "@typescript-eslint/no-unused-vars": ["error"],
    "jsx-a11y/accessible-emoji": "off", // Allow use of Emoji's
  },
  env: {
    jest: true,
  },
};
