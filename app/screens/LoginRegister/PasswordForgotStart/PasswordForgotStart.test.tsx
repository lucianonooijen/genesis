import React from "react";
import { render } from "@testing-library/react-native";
import useMockNavigation from "test/mockNavigation";
import PasswordForgotStart from "./PasswordForgotStart";

describe("PasswordForgotStart", () => {
    it("should render without throwing", () => {
        render(<PasswordForgotStart navigation={useMockNavigation()} />);
    });
});
