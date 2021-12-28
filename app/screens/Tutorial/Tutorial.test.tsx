import React from "react";
import {render} from "@testing-library/react-native";
import {
    TutorialScreenOne,
    TutorialScreenTwo,
    TutorialScreenThree,
} from "./Tutorial";
import useMockNavigation from "../../test/mockNavigation";

describe("LandingPages", () => {
    it("should render TutorialScreenOne without throwing", () => {
        render(<TutorialScreenOne navigation={useMockNavigation()} />);
    });

    it("should render TutorialScreenTwo without throwing", () => {
        render(<TutorialScreenTwo navigation={useMockNavigation()} />);
    });

    it("should render TutorialScreenThree without throwing", () => {
        render(<TutorialScreenThree navigation={useMockNavigation()} />);
    });
});
