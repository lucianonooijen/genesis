import React from "react";
import { render } from "@testing-library/react-native";
import Account from "./Account";
import {
    initialUserProfileState,
    UserProfileState,
    UserProfileStateContextProviderTest,
} from "../../data/UserProfileState/UserProfileState";

describe("Account", () => {
    it("should render without throwing", () => {
        const userProfileState: UserProfileState = {
            ...initialUserProfileState,
            profile: { firstName: "Ted" },
            hasLoaded: true,
        };
        render(
            <UserProfileStateContextProviderTest
                userProfileState={userProfileState}
            >
                <Account />
            </UserProfileStateContextProviderTest>,
        );
    });
});
