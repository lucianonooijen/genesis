import React from "react";
import { mount } from "enzyme";
import App from "./App";

describe("App", () => {
  it("should render tutorial on first render", () => {
    const mountedElement = mount(<App />);
    // TODO: write test here to verify that tutorial screen renders on first init.
    expect(mountedElement);
  });
});
