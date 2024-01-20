import { create } from 'zustand'
import { Status } from '../schema/status'

type State = {
  status: Status
}

type Action = {
  setStatus: (status: Status) => void
}

export const useGameStatusStore = create<State & Action>()(set => ({
  status: {} as Status,
  setStatus: status => set(() => ({ status: status })),
}))
