import React from "react";
import {render} from "@testing-library/react-native";
import Landing from "./Landing";

describe("Landing", () => {
    it("should render without throwing", () => {
        render(<Landing />);
    });
});
