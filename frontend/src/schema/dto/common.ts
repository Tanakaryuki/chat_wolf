import type { EventType, Protocol } from "../common";

export const transferCommonSchema = (jsonString: string): Protocol => {
  const rawData: RawProtocol = JSON.parse(jsonString);
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
    room: {
      roomId: rawData.room?.room_id,
      roomOwnerId: rawData.room?.room_owner_id,
      voteEnded: rawData.room?.vote_ended,
    },
    chatText: rawData.chat_text,
    option: {
      turnNum: rawData.option?.turn_num,
      discussTime: rawData.option?.discuss_time,
      voteTime: rawData.option?.vote_time,
      participantsNum: rawData.option?.participants_num,
    },
    timeNow: rawData.time_now,
    win: rawData.win,
    users: rawData.users,
  };
  return data;
};

export const commonSchemaToJSON = (data: Protocol): string => {
  const rawData: RawProtocol = {
    event_type: data.eventType,
    user: {
      id: data.user?.id,
      display_name: data.user?.displayName,
      icon: data.user?.icon,
      is_wolf: data.user?.isWolf,
      score: data.user?.score,
      word: data.user?.word,
      is_participant: data.user?.isParticipant,
    },
    room: {
      room_id: data.room?.roomId,
      room_owner_id: data.room?.roomOwnerId,
      vote_ended: data.room?.voteEnded,
    },
    chat_text: data.chatText,
    option: {
      turn_num: data.option?.turnNum,
      discuss_time: data.option?.discussTime,
      vote_time: data.option?.voteTime,
      participants_num: data.option?.participantsNum,
    },
    time_now: data.timeNow,
    win: data.win,
    users: data.users,
  };
  return JSON.stringify(rawData);
};

type RawProtocol = {
  event_type: EventType;
  user?: User;
  room?: Room;
  chat_text?: string;
  option?: Option;
  time_now?: string;
  win?: Role;
  users?: User[];
};

type User = {
  id?: string;
  display_name?: string;
  icon?: string;
  is_wolf?: boolean;
  score?: number;
  word?: string;
  is_participant?: string;
  vote?: {
    id?: string;
    display_name?: string;
  };
};

type Room = {
  room_id?: string;
  room_owner_id?: string;
  vote_ended?: boolean;
};

type Option = {
  turn_num?: string;
  discuss_time?: string;
  vote_time?: string;
  participants_num?: string;
};

enum Role {
  wolf = "wolf",
  citizen = "citizen",
}
