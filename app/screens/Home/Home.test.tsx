import React from "react";
import { render } from "@testing-library/react-native";
import Home from "./Home";
import useMockNavigation from "../../test/mockNavigation";

describe("Home", () => {
    it("should render without throwing", () => {
        render(<Home navigation={useMockNavigation()} />);
    });
});
