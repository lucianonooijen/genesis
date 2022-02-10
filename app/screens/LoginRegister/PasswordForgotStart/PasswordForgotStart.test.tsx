import React from "react";
import { fireEvent, render, waitFor } from "@testing-library/react-native";
import useMockNavigation from "__mocks__/mockNavigation";
import PasswordForgotStart from "./PasswordForgotStart";

describe("PasswordForgotStart", () => {
    it("should render without throwing", () => {
        render(<PasswordForgotStart navigation={useMockNavigation()} />);
    });

    it("should only password reset start after filling in account details", async () => {
        const nav = useMockNavigation();
        const r = render(
            <PasswordForgotStart
                navigation={nav}
                passwordResetStartApiCall={jest.fn()}
            />,
        );

        // Expect button to be disabled and not call navigate
        const resetButton = r.getByA11yLabel("Reset my password");
        expect(resetButton.props.accessibilityState.disabled).toBeTruthy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(resetButton);
        expect(nav.navigate).toHaveBeenCalledTimes(0);

        // Change data
        fireEvent.changeText(r.getByA11yLabel("Email"), "email@example.com");

        // Expect login to be possible now
        expect(resetButton.props.accessibilityState.disabled).toBeFalsy();
        expect(nav.navigate).toHaveBeenCalledTimes(0);
        fireEvent.press(resetButton);
        await waitFor(() => expect(nav.navigate).toHaveBeenCalledTimes(1));
    });
});
