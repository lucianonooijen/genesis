import axios from "axios";
import { Decoder } from "decoders";
import { PutApiCall } from "@genesis/api/types";
import { generateApiHeaders, jsonCheck, throwIfResponseError } from "./helpers";

const generatePutApiCall =
    <Req, Res>(
        endpoint: string,
        isAuthenticated: boolean,
        reqDecoder: Decoder<Req>,
        resDecoder?: Decoder<Res>,
    ): PutApiCall<Req, Res> =>
    async (apiConfig, reqBody) => {
        // Check JWT and request body
        if (isAuthenticated && !apiConfig.jwt) {
            throw new Error(
                "Endpoint requires JWT to be present in API config",
            );
        }
        jsonCheck(reqBody, reqDecoder);

        // Make API call
        const callUrl = `${apiConfig.baseUrl}/${endpoint}`;
        console.log(`[API] PUT call to ${callUrl}`); // eslint-disable-line no-console
        const res = await axios.put(callUrl, reqBody, {
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

export default generatePutApiCall;
