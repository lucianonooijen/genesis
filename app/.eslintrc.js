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
            jsx: true,
        },
        ecmaVersion: 2018,
        sourceType: "module",
        project: "./tsconfig.json",
    },
    settings: {
        "import/resolver": {
            node: {
                extensions: [".js", ".jsx", ".ts", ".tsx", ".json"],
            },
            typescript: {
                project: "./tsconfig.json",
            },
        },
    },
    rules: { // In alphabetical order, all rules require a comment to explain the change and necessity
        "import/extensions": 0, // The Typescript compiler will check imports extensions
        "import/prefer-default-export": 0, // Allow 'export const' exports
        "jsx-a11y/accessible-emoji": "off", // Allow use of Emojis
        "no-shadow": 0, // For allowing headacheless Redux Action in props in JS code
        "no-use-before-define": 0, // Allow sane file layouts
        "object-curly-spacing": ["error", "always"], // Force consistency for { item } instead of {item}
        "react/function-component-definition": [1, { namedComponents: "arrow-function", unnamedComponents: "arrow-function" }], // Force arrow function
        "react/jsx-filename-extension": [1, { extensions: [".tsx", ".jsx"] }], // Force files to have .jsx or .tsx
        "react/prop-types": 0, // Done by using React.FC<PropsInterface>
        "react/require-default-props": 0, // Otherwise you get false positives, as regular React ESLint rules doesn't fully get Typescript React code
        "@typescript-eslint/no-shadow": ["error"], //  For allowing headacheless Redux Action in props in TS code
        "@typescript-eslint/no-use-before-define": 0, // Allow sane file layouts
    },
    overrides: [
        {
            files: ["*.test.js", "*.test.jsx", "*.test.ts", "*.test.tsx"], // All test files
            rules: {
                "import/no-extraneous-dependencies": ["error", { devDependencies: true }], // Test files import devDependencies
                "@typescript-eslint/no-empty-function": 0, // Allow `() => {}` in test files
            },
        },
        {
            files: ["*.stories.js", "*.stories.jsx", "*.stories.ts", "*.stories.tsx"], // All Stories files
            rules: {
                "import/no-extraneous-dependencies": ["error", { devDependencies: true }], // Stories files import devDependencies
                "no-console": 0, // Allow `console` usage in Stories files
                "@typescript-eslint/no-empty-function": 0, // Allow `() => {}` in Stories files
            },
        },
    ],
}
