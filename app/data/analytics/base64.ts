export interface GenesisJWT {
    aud: string;
    exp: number;
    jti: string;
    iss: string;
    sub: string;
}

// from https://snack.expo.dev/BktW0xdje
// inspired by https://github.com/davidchambers/Base64.js/blob/master/base64.js

export const parseJwt = (jwt: string): GenesisJWT => {
    const base64Url = jwt.split(".")[1];
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
        Base64.atob(base64)
            .split("")
            .map(c => `%${`00${c.charCodeAt(0).toString(16)}`.slice(-2)}`)
            .join(""),
    );

    return JSON.parse(jsonPayload);
};

/* eslint-disable */

const chars =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
const Base64 = {
    btoa: (input = "") => {
        const str = input;
        let output = "";

        for (
            let block = 0, charCode, i = 0, map = chars;
            str.charAt(i | 0) || ((map = "="), i % 1);
            output += map.charAt(63 & (block >> (8 - (i % 1) * 8)))
        ) {
            charCode = str.charCodeAt((i += 3 / 4));

            if (charCode > 0xff) {
                throw new Error(
                    "'btoa' failed: The string to be encoded contains characters outside of the Latin1 range.",
                );
            }

            block = (block << 8) | charCode;
        }

        return output;
    },

    atob: (input = "") => {
        const str = input.replace(/=+$/, "");
        let output = "";

        if (str.length % 4 == 1) {
            throw new Error(
                "'atob' failed: The string to be decoded is not correctly encoded.",
            );
        }
        for (
            let bc = 0, bs = 0, buffer, i = 0;
            (buffer = str.charAt(i++));
            ~buffer && ((bs = bc % 4 ? bs * 64 + buffer : buffer), bc++ % 4)
                ? (output += String.fromCharCode(255 & (bs >> ((-2 * bc) & 6))))
                : 0
        ) {
            buffer = chars.indexOf(buffer);
        }

        return output;
    },
};
