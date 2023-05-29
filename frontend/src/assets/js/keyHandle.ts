export const keyHandle = (key: KeyboardEvent, func: Function) => {
    if (key.key === "Enter") {
        func();
    }
}