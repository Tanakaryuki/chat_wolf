import { create } from 'zustand'
import ReconnectingWebSocket from 'reconnecting-websocket'
import { MutableRefObject, useRef } from 'react'

type State = {
  socketRef: MutableRefObject<ReconnectingWebSocket | undefined>
}

type Action = {
  setRef: (ref: MutableRefObject<ReconnectingWebSocket | undefined>) => void
}

export const useRefStore = create<State & Action>()(set => ({
  socketRef: useRef(),
  setRef: ref => set(() => ({ socketRef: ref })),
}))
