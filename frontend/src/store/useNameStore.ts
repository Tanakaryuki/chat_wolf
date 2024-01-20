import { create } from "zustand";
import { persist, devtools } from "zustand/middleware";

type State = {
  name: string;
};

type Action = {
  setName: (name: string) => void;
};

export const useNameStore = create<State & Action>()(
  devtools(
    persist(
      (set) => ({
        name: "simesaba",
        setName: (name) => set(() => ({ name: name })),
      }),
      {
        name: "user-name",
      }
    )
  )
);
