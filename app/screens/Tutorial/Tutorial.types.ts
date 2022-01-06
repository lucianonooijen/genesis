import { TutorialPageData } from "../../layouts/TutorialLayout/TutorialLayout.types";

export type TutorialPropPageData = Omit<TutorialPageData, "onPressContinue">;

export interface TutorialProps {
    pageData: TutorialPropPageData;
}
