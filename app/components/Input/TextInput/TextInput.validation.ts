import { useState } from "react";

export type ValidatorHook = () => [
    error: string,
    validate: (
        email: string,
        clearError?: boolean, // a "secret" property used to clear the error when needed to avoid sending extra props
    ) => void,
];

export const useEmailValidation: ValidatorHook = () => {
    const [error, setError] = useState("");

    // From https://emailregex.com
    const emailRegex =
        // eslint-disable-next-line no-useless-escape
        /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

    const validate = (email: string, clearError?: boolean) => {
        if (clearError) {
            setError("");
            return;
        }

        const isValid = emailRegex.test(email);
        if (isValid) {
            setError("");
            return;
        }

        setError("Geen geldig emailadres");
    };

    return [error, validate];
};

export const usePasswordValidation: ValidatorHook = () => {
    const [error, setError] = useState("");

    const passwordRegex =
        /^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$/;

    const validate = (email: string, clearError?: boolean) => {
        if (clearError) {
            setError("");
            return;
        }

        const isValid = passwordRegex.test(email);
        if (isValid) {
            setError("");
            return;
        }

        setError(
            "Wachtwoord eisen zijn minimaal 8 karakters lang: letters, nummers en leestekens.",
        );
    };

    return [error, validate];
};
