import React from "react";
import { render } from "@testing-library/react-native";
import Router from "./Router";
import { initialAppState } from "../data/AppState/AppState";

describe("router", () => {
    it("should render", () => {
        render(<Router appState={initialAppState} />);
    });
});
