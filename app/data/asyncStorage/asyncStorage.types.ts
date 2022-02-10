export interface DataStorage<T> {
    get: () => Promise<T>;
    set: (data: T) => Promise<void>;
    remove: () => Promise<void>;
}

export type AppStateStorage = DataStorage<{
    jwt: string | null;
    hasSeenTutorial: boolean;
}>;
