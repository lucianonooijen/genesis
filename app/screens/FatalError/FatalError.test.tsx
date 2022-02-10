import React from "react";
import { render } from "@testing-library/react-native";
import FatalError from "./FatalError";

describe("Home", () => {
    it("should render without throwing", () => {
        render(<FatalError title="test" description="test" />);
    });

    it("should render the title and description", () => {
        const title = "TestTitle";
        const descr = "TestDescr";
        const r = render(<FatalError title={title} description={descr} />);

        r.getByText(title);
        r.getByText(descr);
    });
});
