// import { FC } from "react";
// import { EventType, Protocol } from "../schema/common";
// import { Chat, Colors, Participant } from "../schema/status";
// import {
//   useChatStore,
//   useGameStatusStore,
//   useParticipantsStore,
// } from "../store";

// type Props = {
//   schema: Protocol;
// };

// export const Listener: FC<Props> = ({ schema }) => {
//   const setRoomId = useGameStatusStore((state) => state.setRoomId);
//   const addParticipant = useParticipantsStore((state) => state.addParticipant);
//   const addChats = useChatStore((state) => state.addChats);
//   switch (schema.eventType) {
//     case EventType.createRoom:
//       setRoomId(schema.room?.roomId!);
//       break;
//     case EventType.enterRoom:
//       const p: Participant = {
//         id: schema.user?.id!,
//         icon: schema.user?.icon!,
//         name: schema.user?.displayName!,
//         score: 0,
//         color: Colors.green,
//       };
//       addParticipant(p);
//       break;
//     case EventType.sendChat:
//       const chat: Chat = {
//         userId: schema.user?.id!,
//         icon: schema.user?.icon!,
//         name: schema.user?.displayName!,
//         message: schema.chatText!,
//         color: Colors.green,
//       };
//       addChats(chat);
//       break;
//     default:
//       console.log(schema);
//       break;
//   }
// };
