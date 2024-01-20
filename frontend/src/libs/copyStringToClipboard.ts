export const copyStringToClipboard = (text: string) => {
  navigator.clipboard.writeText(text).then(
    () => {
      console.log("Async: Copying to clipboard was successful!");
    },
    (err) => {
      console.error("Async: Could not copy text: ", err);
    }
  );
};
