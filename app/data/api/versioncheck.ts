import { useState } from "react";
import { GenesisApiError, versionCheck } from "@genesis/api";
import config from "../../config";

type UseAppVersionCheck = () => [
    ok: boolean,
    error: null | { title: string; description: string },
    check: () => void,
];

export const useAppVersionCheck: UseAppVersionCheck = () => {
    const [hasLoaded, setHasLoaded] = useState(false);
    const [error, setError] = useState<null | {
        title: string;
        description: string;
    }>(null);

    const check = async () => {
        const apiConfig = {
            appVersion: config.appVersion,
            baseUrl: config.baseUrl,
        };

        try {
            await versionCheck(apiConfig);
            setHasLoaded(true);
        } catch (e) {
            // eslint-disable-next-line no-console
            console.warn(
                `received error during app version check: ${JSON.stringify(e)}`,
            );

            const err = e as Error;

            if (err.message === "Network Error") {
                setError({
                    title: "Server niet bereikbaar",
                    description:
                        "De Genesis server is niet bereikbaar. Probeer de app opnieuw op te starten. Lukt het niet? Neem contact op met onze support.",
                });
                setHasLoaded(true);
                return;
            }

            if (err instanceof GenesisApiError) {
                // is update required code
                if (err.err.status === 426) {
                    setError({
                        title: "App update nodig",
                        description:
                            "De app versie is te laag om met de server te kunnen communiceren. Update de app om dit probleem te verhelpen.",
                    });
                    setHasLoaded(true);
                    return;
                }
            }

            setError({
                title: "Er gaat iets goed fout",
                description:
                    "Er heeft zich een onverwachte fout voorgedaan. Neem contact op met onze support zodat we je kunnen helpen.",
            });
            setHasLoaded(true);
        }
    };

    return [hasLoaded, error, check];
};
