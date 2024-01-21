import { create } from "zustand";
import { Chat } from "../schema/status";

type State = {
  chats: Chat[];
};

type Action = {
  addChats: (chat: Chat) => void;
};

export const useChatStore = create<State & Action>((set, get) => ({
  chats: [],
  addChats: (chat) => set(() => ({ chats: [...get().chats, chat] })),
}));
