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
  color: Colors;
};

export type Chat = {
  userId: string;
  icon: string;
  name: string;
  message: string;
  color: Colors;
};

export enum Colors {
  red = "#FF0000",
  orange = "#FFA500",
  yellowGreen = "#95DC4E",
  green = "#009944",
  lightBlue = "#2986E8",
  purple = "#954EDC",
}
