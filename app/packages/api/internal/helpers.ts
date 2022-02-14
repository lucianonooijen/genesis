import { AxiosRequestHeaders, AxiosResponse } from "axios";
import { guard, Decoder } from "decoders";
import { apiErrorDecoder } from "@genesis/api/internal/decoders";
import ApiError from "@genesis/api/types/ApiError";
import { GenesisApiError } from "@genesis/api";
import ApiConfig from "../types/Config";

export const generateApiHeaders = (config: ApiConfig): AxiosRequestHeaders => {
    if (config.jwt) {
        return {
            "X-Genesis-Client-Version": config.appVersion,
            Authorization: `Bearer ${config.jwt}`,
        };
    }

    return {
        "X-Genesis-Client-Version": config.appVersion,
    };
};

export const jsonCheck = <T>(data: T, decoder: Decoder<T>) => {
    const decodeChecker = guard(decoder);
    // TODO: Allow underscore to be not used
    // eslint-disable-next-line no-unused-vars, @typescript-eslint/no-unused-vars
    const _ = decodeChecker(data); // Throws if it's not valid
};

export const throwIfResponseError = (res: AxiosResponse) => {
    if (res.status < 200 || res.status > 299) {
        const err = res.data as ApiError;

        // eslint-disable-next-line no-console
        console.warn(err);

        jsonCheck(err, apiErrorDecoder);

        throw new GenesisApiError(err);
    }
};
