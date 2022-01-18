# API package

## How to add a new POST request

1. Add your request body and (optionally) response body types to the `types` directory
1. Run `yarn run aegis`
1. Add your endpoint to `internal/endpoints.ts`
1. Define your call in `index.ts` with the correct endpoint, types and decoders
1. Import your call and use it
