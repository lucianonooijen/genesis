import ApiConfig from "@genesis/api/types/Config";

export type GetApiCall<ResPayload> = (
    apiConfig: ApiConfig,
) => Promise<ResPayload>;

export type PostApiCall<ReqPayload, ResPayload> = (
    apiConfig: ApiConfig,
    req: ReqPayload,
) => Promise<ResPayload>;

// ReqPayload and ResPayload should be equal in PUT, but declaring them separately is cleaner in case there will be exceptions
export type PutApiCall<ReqPayload, ResPayload> = (
    apiConfig: ApiConfig,
    req: ReqPayload,
) => Promise<ResPayload>;

export type DeleteApiCall<ReqPayload, ResPayload> = (
    apiConfig: ApiConfig,
    req: ReqPayload,
) => Promise<ResPayload>;
