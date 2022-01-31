import React, { useContext } from "react";
import { render } from "@testing-library/react-native";
import { View, Text } from "react-native";
import UserProfileStateContext, {
    UserProfileState,
    UserProfileStateContextProvider,
    UserProfileStateContextProviderTest,
    initialUserProfileState,
} from "./UserProfileState";

describe("UserProfileStateContextProvider", () => {
    it("should render the children", () => {
        const Child = () => (
            <View>
                <Text>Example</Text>
            </View>
        );
        const r = render(
            <UserProfileStateContextProvider>
                <Child />
            </UserProfileStateContextProvider>,
        );
        const child = r.getByText(/Example/);
        expect(child).toBeTruthy();
    });
});

describe("UserProfileStateContextProviderTest", () => {
    it("should render the children", () => {
        const Child = () => (
            <View>
                <Text>Example</Text>
            </View>
        );
        const r = render(
            <UserProfileStateContextProviderTest
                userProfileState={initialUserProfileState}
            >
                <Child />
            </UserProfileStateContextProviderTest>,
        );
        const child = r.getByText(/Example/);
        expect(child).toBeTruthy();
    });

    it("should pass the appState without modifying anything with custom data", () => {
        const userProfileState: UserProfileState = {
            profile: {
                firstName: "Ted",
            },
            hasLoaded: true,
            setProfile: jest.fn(),
        };

        const Child = () => {
            const contextData = useContext(UserProfileStateContext);
            expect(contextData).toEqual(userProfileState);
            return (
                <View>
                    <Text>Example</Text>
                </View>
            );
        };

        render(
            <UserProfileStateContextProviderTest
                userProfileState={userProfileState}
            >
                <Child />
            </UserProfileStateContextProviderTest>,
        );
    });
});
