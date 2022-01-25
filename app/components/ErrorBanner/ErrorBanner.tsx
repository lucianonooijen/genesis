import React, { useEffect, useState } from "react";
import ErrorBannerProps from "./ErrorBanner.types";
import { ErrorBannerContainer, ErrorBannerText } from "./ErrorBanner.styles";

const ErrorBanner: React.FC<ErrorBannerProps> = ({ error }) => {
    const [errString, setErrString] = useState("");

    useEffect(() => {
        setErrString(errStringExtractor(error));
    }, [error]);

    if (!error) {
        return null;
    }

    return (
        <ErrorBannerContainer>
            <ErrorBannerText>{errString}</ErrorBannerText>
        </ErrorBannerContainer>
    );
};

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const errStringExtractor = (error: any): string => {
    if (!error) {
        return "";
    }

    if (error instanceof Error) {
        return error.message;
    }

    if (typeof error === "string") {
        return error;
    }

    if (typeof error === "object") {
        return "Something went wrong";
    }

    return "";
};

export default ErrorBanner;
