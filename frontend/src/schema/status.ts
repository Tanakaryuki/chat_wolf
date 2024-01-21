import { EventType } from "./common";

export type Status = {
  eventType: EventType;
};

export type Meta = {
  roomId: string;
};

export type Participant = {
  id: string;
  name: string;
  icon: string;
  score: number;
};
