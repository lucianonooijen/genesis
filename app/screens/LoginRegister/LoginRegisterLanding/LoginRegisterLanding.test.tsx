import React from "react";
import { render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import LoginRegisterLanding from "./LoginRegisterLanding";

describe("LoginRegisterLanding", () => {
    it("should render without throwing", () => {
        render(<LoginRegisterLanding navigation={useMockNavigation()} />);
    });
});
