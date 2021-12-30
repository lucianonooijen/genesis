import React from "react";
import { fireEvent, render } from "@testing-library/react-native";
import { TutorialPageData } from "./TutorialLayout.types";
import TutorialLayout from "./TutorialLayout";

const getTestData = (): TutorialPageData => ({
    id: "test",
    title: "testTitle",
    text: "testText",
    image: { uri: "test-image-source" },
    onPressContinue: jest.fn(),
});

describe("TutorialLayout", () => {
    it("should render", () => {
        const testData = getTestData();
        const r = render(<TutorialLayout pageData={testData} />);
        expect(r.getByTestId(testData.id)).toBeTruthy();
    });

    it("should show the title and text", () => {
        const testData = getTestData();
        const r = render(<TutorialLayout pageData={testData} />);
        expect(r.getByTestId("tutorial-title").children).toEqual([
            testData.title,
        ]);
        expect(r.getByTestId("tutorial-text").children).toEqual([
            testData.text,
        ]);
    });

    it("should call the navigate function with nextScreen as argument when the next button is pressed", () => {
        const testData = getTestData();
        const r = render(<TutorialLayout pageData={testData} />);
        const nextButton = r.getByTestId("tutorial-nextbutton");
        expect(nextButton).toBeTruthy();
        expect(testData.onPressContinue).toBeCalledTimes(0);
        fireEvent.press(nextButton);
        expect(testData.onPressContinue).toBeCalledTimes(1);
    });
});
