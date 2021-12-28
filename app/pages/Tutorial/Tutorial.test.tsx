import React from "react";
import {render} from "@testing-library/react-native";
import Tutorial from "./Tutorial";
import useMockNavigation from "../../test/mockNavigation";

describe("Landing", () => {
    it("should render without throwing", () => {
        render(<Tutorial navigation={useMockNavigation()} />);
    });
});
