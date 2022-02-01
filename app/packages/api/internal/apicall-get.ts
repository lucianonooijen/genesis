import axios from "axios";
import { Decoder } from "decoders";
import { GetApiCall } from "@genesis/api/types";
import { generateApiHeaders, jsonCheck, throwIfResponseError } from "./helpers";

const generateGetApiCall =
    <Res>(
        endpoint: string,
        isAuthenticated: boolean,
        resDecoder?: Decoder<Res>,
    ): GetApiCall<Res> =>
    async apiConfig => {
        // Check JWT and request body
        if (isAuthenticated && !apiConfig.jwt) {
            throw new Error(
                "Endpoint requires JWT to be present in API config",
            );
        }

        // Make API call
        const callUrl = `${apiConfig.baseUrl}/${endpoint}`;
        console.log(`[API] GET call to ${callUrl}`); // eslint-disable-line no-console
        const res = await axios.get(callUrl, {
            headers: generateApiHeaders(apiConfig),
            validateStatus: () => true, // this prevents Axios from throwing errors on non-2xx codes, we handle it ourselves
        });

        // Check if the request errored (non-2xx), if so, throw formatted error
        throwIfResponseError(res);

        // Validate response data type (successful response)
        const resData = res.data as Res;
        if (resDecoder) {
            jsonCheck(resData, resDecoder);
        }

        // Everything is successfully validated
        return resData;
    };

export default generateGetApiCall;
