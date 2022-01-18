import ApiConfig from "@genesis/api/types/Config";

export type PostApiCall<ReqPayload, ResPayload> = (
    apiConfig: ApiConfig,
    req: ReqPayload,
) => Promise<ResPayload>;
