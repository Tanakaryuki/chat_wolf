import type { EventType, Protocol } from '../common'

export const transferCommonSchema = (jsonString: string): Protocol => {
  const rawData: RawProtocol = JSON.parse(jsonString)
  const data: Protocol = {
    eventType: rawData.event_type,
    user: {
      id: rawData.user?.id,
      displayName: rawData.user?.display_name,
      icon: rawData.user?.icon,
      isWolf: rawData.user?.is_wolf,
      score: rawData.user?.score,
      word: rawData.user?.word,
      isParticipant: rawData.user?.is_participant,
    },
  }
  return data
}

type RawProtocol = {
  event_type: EventType
  user?: User
  room?: Room
  chat_text?: string
  option?: Option
  time_now?: string
  win?: Role
  users?: User[]
}

type User = {
  id?: string
  display_name?: string
  icon?: string
  is_wolf?: boolean
  score?: number
  word?: string
  is_participant?: string
  vote?: {
    id?: string
    display_name?: string
  }
}

type Room = {
  room_id?: string
  room_owner_id?: string
  vote_ended?: boolean
}

type Option = {
  turn_num?: string
  discuss_time?: string
  vote_time?: string
  participants_num?: string
}

enum Role {
  wolf = 'wolf',
  citizen = 'citizen',
}
