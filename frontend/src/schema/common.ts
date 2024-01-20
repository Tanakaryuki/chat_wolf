export type Protocol = {
  eventType: EventType
  user?: User
  room?: Room
  chatText?: string
  option?: Option
  timeNow?: string
  win?: Role
  users?: User[]
}

export enum EventType {
  createRoom = 'create_room',
  enterRoom = 'enter_room',
  changeRoomOwner = 'change_room_owner',
  exitRoom = 'exit_room',
  sendChat = 'send_chat',
  setOption = 'set_option',
  startGame = 'start_game',
  sendTime = 'send_time',
  askQuestion = 'ask_question',
  endQandA = 'end_Q_and_A',
  giveAnswer = 'give_answer',
  vote = 'vote',
  gameResult = 'game_result',
  prepareCompletion = 'prepare_completion',
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
