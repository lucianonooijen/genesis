/* eslint-disable no-unused-vars, @typescript-eslint/no-unused-vars */
// DO NOT EDIT - Generated using Aegis
import {
    Decoder,
    constant,
    optional,
    string,
    boolean,
    number,
    null_,
    undefined_,
    object,
    array,
    either,
    either3,
    either4,
    either5,
    either6,
    either7,
    either8,
    either9,
} from "decoders";

export const apiErrorDecoder = object({
    title: string,
    status: number,
    detail: optional(string),
    rawError: optional(string),
});
export const apiConfigDecoder = object({
    baseUrl: string,
    jwt: optional(string),
});
export const jwtResponseDecoder = object({
    jwt: string,
});
export const loginRequestDecoder = object({
    email: string,
    password: string,
});
export const passwordResetStartRequestDecoder = object({
    email: string,
});
export const passwordResetCompleteRequestDecoder = object({
    resetToken: string,
    password: string,
});
export const userProfileDecoder = object({
    firstName: string,
});
export const deleteRequestDecoder = object({
    password: string,
});
export const registerRequestDecoder = object({
    email: string,
    password: string,
    firstName: string,
});
