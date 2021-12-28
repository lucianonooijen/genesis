import React from "react";
import { render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import Register from "./Register";

describe("Register", () => {
    it("should render without throwing", () => {
        render(<Register navigation={useMockNavigation()} />);
    });
});
