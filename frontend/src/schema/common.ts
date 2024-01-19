export type Protocol = {
  eventType: string
  user?: User
  room?: Room
  chatText?: string
  option?: Option
  timeNow?: string
  win?: Role
  users?: User[]
}

export type User = {
  id?: string
  displayName?: string
  icon?: string
  isWolf?: boolean
  score?: number
  word?: string
  isParticipant?: string
  vote?: {
    id?: string
    displayName?: string
  }
}

export type Room = {
  roomId?: string
  roomOwnerId?: string
  voteEnded?: boolean
}

export type Option = {
  turnNum?: string
  discussTime?: string
  voteTime?: string
  participantsNum?: string
}

export enum Role {
  wolf = 'wolf',
  citizen = 'citizen',
}
