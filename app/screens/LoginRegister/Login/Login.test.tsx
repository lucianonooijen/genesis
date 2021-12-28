import React from "react";
import { render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import Login from "./Login";

describe("Login", () => {
    it("should render without throwing", () => {
        render(<Login navigation={useMockNavigation()} />);
    });
});
