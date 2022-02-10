import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import useMockNavigation from "__mocks__/mockNavigation";
import LoginRegisterLanding from "./LoginRegisterLanding";
import { LoginRegisterScreens } from "../../../router/types";

describe("LoginRegisterLanding", () => {
    it("should render without throwing", () => {
        render(<LoginRegisterLanding navigation={useMockNavigation()} />);
    });

    it("should navigate to register when corresponding button is pressed", () => {
        const nav = useMockNavigation();
        const r = render(<LoginRegisterLanding navigation={nav} />);

        const button = r.getByA11yLabel("Register");
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(button);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
        expect(nav.navigate).toHaveBeenCalledWith(
            LoginRegisterScreens.Register,
        );
    });

    it("should navigate to login when corresponding button is pressed", () => {
        const nav = useMockNavigation();
        const r = render(<LoginRegisterLanding navigation={nav} />);

        const button = r.getByA11yLabel("Login");
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(button);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
        expect(nav.navigate).toHaveBeenCalledWith(LoginRegisterScreens.Login);
    });

    it("should navigate to password reset when corresponding button is pressed", () => {
        const nav = useMockNavigation();
        const r = render(<LoginRegisterLanding navigation={nav} />);

        const button = r.getByA11yLabel("Forgot password");
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(button);
        expect(nav.navigate).toHaveBeenCalledTimes(1);
        expect(nav.navigate).toHaveBeenCalledWith(
            LoginRegisterScreens.PasswordForgotStart,
        );
    });
});
