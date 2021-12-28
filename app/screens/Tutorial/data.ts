import { TutorialPageData } from "../../layouts/TutorialLayout/TutorialLayout.types";
import { TutorialScreens } from "../../router/types";
import { imageOne, imageThree, imageTwo } from "./images/images";

const tutorialData: TutorialPageData[] = [
    {
        id: "tutorial_1",
        image: imageOne,
        title: "Example page one",
        text: "The Industrial Revolution and its consequences have been a disaster for the human race.",
        nextScreen: TutorialScreens.ScreenTwo,
    },
    {
        id: "tutorial_2",
        image: imageTwo,
        title: "Example page two",
        text: "The Industrial Revolution and its consequences have been a disaster for the human race.",
        nextScreen: TutorialScreens.ScreenThree,
    },
    {
        id: "tutorial_1",
        image: imageThree,
        title: "Example page three",
        text: "The Industrial Revolution and its consequences have been a disaster for the human race.",
        nextScreen: TutorialScreens.ScreenOne, // FIXME: edit, or implement isLastPage for example, and load hasSeenTutorial to storage
    },
];

export default tutorialData;
