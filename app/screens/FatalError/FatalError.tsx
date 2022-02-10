import React from "react";
import PaddedEmptyLayout from "layouts/PaddedEmptyLayout/PaddedEmptyLayout";
import FatalErrorProps from "./FatalError.types";
import {
    ForceUpdateContainer,
    ErrorTitle,
    ErrorDescription,
} from "./FatalError.components";

const FatalError: React.FC<FatalErrorProps> = ({ title, description }) => (
    <PaddedEmptyLayout>
        <ForceUpdateContainer>
            <ErrorTitle>{title}</ErrorTitle>
            <ErrorDescription>{description}</ErrorDescription>
        </ForceUpdateContainer>
    </PaddedEmptyLayout>
);

export default FatalError;
