import { create } from "zustand";
import ReconnectingWebSocket from "reconnecting-websocket";
import { MutableRefObject } from "react";

type State = {
  socketRef: MutableRefObject<ReconnectingWebSocket | undefined> | null;
};

type Action = {
  setRef: (ref: MutableRefObject<ReconnectingWebSocket | undefined>) => void;
};

export const useRefStore = create<State & Action>()((set) => ({
  socketRef: null,
  setRef: (ref) => set(() => ({ socketRef: ref })),
}));
