module.exports = {
    preset: "@testing-library/react-native",
    setupFilesAfterEnv: ["@testing-library/jest-native/extend-expect"],
    setupFiles: ["<rootDir>/__mocks__/setupMock.js"],
    moduleFileExtensions: ["ts", "tsx", "js", "jsx", "json", "node"],
    moduleNameMapper: {
        "\\.(gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$": "<rootDir>/__mocks__/fileMock.js",
        "\\.(jpg|jpeg|png)$": "<rootDir>/__mocks__/imageMock.js",
    },
};
