import React from "react";
import { render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import PasswordForgotComplete from "./PasswordForgotComplete";

describe("PasswordForgotComplete", () => {
    it("should render without throwing", () => {
        render(<PasswordForgotComplete navigation={useMockNavigation()} />);
    });
});
