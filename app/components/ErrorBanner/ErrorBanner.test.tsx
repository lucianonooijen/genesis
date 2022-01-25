import React from "react";
import { render } from "@testing-library/react-native";
import ErrorBanner from "./ErrorBanner";

describe("ErrorBanner", () => {
    it("should return null when error is empty string", () => {
        const r = render(<ErrorBanner error="" />);
        expect(r.toJSON()).toBeFalsy();
    });

    it("should return null when error is null", () => {
        const r = render(<ErrorBanner error={null} />);
        expect(r.toJSON()).toBeFalsy();
    });

    it("should return null when error is not set", () => {
        const r = render(<ErrorBanner />);
        expect(r.toJSON()).toBeFalsy();
    });

    it("should return the banner containing the text when error is truthy", () => {
        const errorString = "this is an error";
        const r = render(<ErrorBanner error={errorString} />);
        expect(r.toJSON()).toBeTruthy();
    });
});
