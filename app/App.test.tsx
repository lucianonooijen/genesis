import React from "react";
import { render } from "@testing-library/react-native";
import { App } from "./App";
import { AppStateContextProviderTest } from "./data/AppState/AppState";
import { UserProfileStateContextProviderTest } from "./data/UserProfileState/UserProfileState";

describe("app", () => {
    it("should render without throwing", () => {
        render(
            <AppStateContextProviderTest>
                <UserProfileStateContextProviderTest>
                    <App />
                </UserProfileStateContextProviderTest>
            </AppStateContextProviderTest>,
        );
    });
});
