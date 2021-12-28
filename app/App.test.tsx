import React from "react";
import { render } from "@testing-library/react-native";
import App from "./App";

describe("app", () => {
    it("should render without throwing", () => {
        render(<App />);
    });
});
