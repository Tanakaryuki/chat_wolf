import { create } from "zustand";
import { Participant } from "../schema/status";

type State = {
  participants: Participant[];
};

type Action = {
  setParticipants: (participants: Participant[]) => void;
  addParticipant: (Participant: Participant) => void;
  removeParticipant: (id: string) => void;
  resetParticipant: () => void;
};

export const useParticipantsStore = create<State & Action>()((set, get) => ({
  participants: [],
  setParticipants: (participants) =>
    set(() => ({ participants: participants })),
  addParticipant: (participant) =>
    set(() => ({ participants: [...get().participants, participant] })),
  removeParticipant: (id) =>
    set(() => ({
      participants: get().participants.filter((e) => {
        return e.id !== id;
      }),
    })),
  resetParticipant: () => set(() => ({ participants: [] })),
}));
