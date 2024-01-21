import { create } from "zustand";
import { EventType, Protocol } from "../schema/common";

type State = {
  schema: Protocol;
};

type Action = {
  setSchema: (schema: Protocol) => void;
};

export const useSchemaStore = create<State & Action>((set) => ({
  schema: { eventType: EventType.other },
  setSchema: (schema) => set(() => ({ schema: schema })),
}));
