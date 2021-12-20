import React from "react";
import {render} from "@testing-library/react-native";
import App from "./app";

describe("app", () => {
    it("should render without throwing", () => {
        render(<App />);
    });
});
