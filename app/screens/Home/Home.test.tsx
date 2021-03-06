import React from "react";
import { render } from "@testing-library/react-native";
import Home from "./Home";
import {
    initialUserProfileState,
    UserProfileState,
    UserProfileStateContextProviderTest,
} from "../../data/UserProfileState/UserProfileState";

describe("Home", () => {
    it("should render without throwing", () => {
        const userProfileState: UserProfileState = {
            ...initialUserProfileState,
            profile: { firstName: "Ted" },
            hasLoaded: true,
        };
        const r = render(
            <UserProfileStateContextProviderTest
                userProfileState={userProfileState}
            >
                <Home />
            </UserProfileStateContextProviderTest>,
        );
        r.getByText(/Ted/);
    });
});
