module.exports = {
    root: true,
    extends: [
        "airbnb",
        "airbnb/hooks",
        "plugin:@typescript-eslint/recommended",
        "prettier",
        "plugin:prettier/recommended",
    ],
    plugins: ["@typescript-eslint", "react", "prettier", "security"],
    parser: "@typescript-eslint/parser",
    parserOptions: {
        ecmaFeatures: {
            jsx: true
        },
        ecmaVersion: 2018,
        sourceType: "module",
        project: "./tsconfig.json",
    },
    settings: {
        "import/resolver": {
            node: {
                paths: ["src"],
                extensions: [".js", ".jsx", ".ts", ".tsx", ".json"],
            },
        },
    },
    rules: { // In alphabetical order
        "import/extensions": ["error", "never"],
        "import/no-unresolved": 0,
        "import/prefer-default-export": 0,
        "jsx-a11y/accessible-emoji": "off", // Allow use of Emojis
        "no-shadow": 0, // For allowing headacheless Redux Action in props
        "no-use-before-define": 0, // Allow sane file layouts
        "react/function-component-definition": [1, { namedComponents: "arrow-function", unnamedComponents: "arrow-function" }],
        "react/jsx-filename-extension": [1, { extensions: [".tsx", ".jsx"] }],
        "react/prop-types": 0, // Done by using React.FC<PropsInterface>
        "@typescript-eslint/no-shadow": ["error"],
        "@typescript-eslint/no-use-before-define": 0, // Allow sane file layouts
    }
}
