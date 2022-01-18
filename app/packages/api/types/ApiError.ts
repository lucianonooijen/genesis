export default interface ApiError {
    title: string;
    status: number;
    detail?: string;
    rawError?: string;
}
