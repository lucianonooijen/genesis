import React from "react";
import { render } from "@testing-library/react-native";
import Router from "./Router";

describe("router", () => {
    it("should render", () => {
        render(
            <Router appState={{ hasSeenTutorial: false, isLoggedIn: false }} />,
        );
    });
});
