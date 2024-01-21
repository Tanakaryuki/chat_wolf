import { create } from 'zustand'
import { persist, devtools } from 'zustand/middleware'
import { generateUUID } from '../libs/uuid'

type State = {
  uuid: string
}

type Action = {
  updateUUID: () => void
}

export const useUUIDStore = create<State & Action>()(
  devtools(
    persist(
      set => ({
        uuid: generateUUID(),
        updateUUID: () => set(() => ({ uuid: generateUUID() })),
      }),
      {
        name: 'user-uuid',
      },
    ),
  ),
)
