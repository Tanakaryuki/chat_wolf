import { ReactNode } from 'react'
import { EventType, Protocol } from '../../schema/common'

export const SwitchingHub = (schema: Protocol | undefined) => {
  switch (schema?.eventType) {
    case EventType.askQuestion:
      return <></>
    case EventType.changeRoomOwner:
      return <></>
    case EventType.createRoom:
      return <></>
    case EventType.endQandA:
      return <></>
    case EventType.enterRoom:
      return <></>
    case EventType.exitRoom:
      return <></>
    case EventType.gameResult:
      return <></>
    case EventType.giveAnswer:
      return <></>
    case EventType.prepareCompletion:
      return <></>
    case EventType.sendChat:
      return <></>
    case EventType.sendTime:
      return <></>
    case EventType.setOption:
      return <></>
    case EventType.startGame:
      return <></>
    case EventType.vote:
      return <></>
    default:
      return <></>
  }
}

export const SwitchingSeen = (schema: Protocol | undefined, children: ReactNode) => {
  switch (schema?.eventType) {
    case EventType.askQuestion:
      return <>{children}</>
    case EventType.changeRoomOwner:
      return <></>
    case EventType.createRoom:
      return <></>
    case EventType.endQandA:
      return <></>
    case EventType.enterRoom:
      return <></>
    case EventType.exitRoom:
      return <></>
    case EventType.gameResult:
      return <></>
    case EventType.giveAnswer:
      return <></>
    case EventType.prepareCompletion:
      return <></>
    case EventType.sendChat:
      return <></>
    case EventType.sendTime:
      return <></>
    case EventType.setOption:
      return <></>
    case EventType.startGame:
      return <></>
    case EventType.vote:
      return <></>
    default:
      return <></>
  }
}
