import { create } from "zustand";
import { Meta } from "../schema/status";
import { EventType } from "../schema/common";

type State = {
  status: EventType;
  meta: Meta;
};

type Action = {
  setStatus: (status: EventType) => void;
  setMeta: (meta: Meta) => void;
  setRoomId: (id: string) => void;
};

export const useGameStatusStore = create<State & Action>()((set) => ({
  status: EventType.other,
  meta: {} as Meta,
  setStatus: (status) => set(() => ({ status: status })),
  setMeta: (meta) => set(() => ({ meta: meta })),
  setRoomId: (id) => set(() => ({ meta: { roomId: id } })),
}));
