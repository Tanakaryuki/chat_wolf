import { useEffect, useRef } from "react";
import { Hamburger } from "../components/Hamburger/Hamburger";
import { GameContainer } from "../containers/game";
import ReconnectingWebSocket from "reconnecting-websocket";
import {
  useChatStore,
  useGameStatusStore,
  useParticipantsStore,
  useRefStore,
  useSchemaStore,
} from "../store";
import { EventType } from "../schema/common";
import { transferMessage } from "../libs/webSocket";
import { Chat, Colors, Participant } from "../schema/status";
// import { Protocol } from "../schema/common";
// const ReconnectingWebSocket = lazy(() => import("reconnecting-websocket"));

const Game = () => {
  const socketRef = useRef<ReconnectingWebSocket>();
  const setRef = useRefStore((state) => state.setRef);
  const query = new URLSearchParams(location.search);
  const roomId = query.get("id");
  const eventype = query.get("type");
  //   const status = useGameStatusStore((state) => state.status);
  const setStatus = useGameStatusStore((state) => state.setStatus);
  const setRoomId = useGameStatusStore((state) => state.setRoomId);
  const setSchema = useSchemaStore((state) => state.setSchema);
  const schema = useSchemaStore((state) => state.schema);
  const addParticipant = useParticipantsStore((state) => state.addParticipant);
  const addChats = useChatStore((state) => state.addChats);

  useEffect(() => {
    if (eventype === EventType.createRoom) {
      setStatus(EventType.createRoom);
      if (roomId) {
        setRoomId(roomId);
      }
    } else if (eventype === EventType.enterRoom) {
      setStatus(EventType.enterRoom);
      if (roomId) {
        setRoomId(roomId);
      }
    }
    const websocket = new ReconnectingWebSocket(
      import.meta.env.VITE_WS_URL ?? "ws://localhost:80/ws"
    );
    socketRef.current = websocket;
    setRef(socketRef);

    const onMessage = (event: MessageEvent<string>) => {
      const s = transferMessage(event);
      setSchema(s);

      switch (s.eventType) {
        case EventType.createRoom:
          setRoomId(s.room?.roomId!);
          break;
        case EventType.enterRoom:
          const p: Participant = {
            id: s.user?.id!,
            icon: s.user?.icon!,
            name: s.user?.displayName!,
            score: 0,
            color: Colors.green,
          };
          addParticipant(p);
          break;
        case EventType.sendChat:
          const chat: Chat = {
            userId: s.user?.id!,
            icon: s.user?.icon!,
            name: s.user?.displayName!,
            message: s.chatText!,
            color: Colors.green,
          };
          addChats(chat);
          break;
        default:
          console.log(schema);
          break;
      }
    };

    websocket.addEventListener("message", onMessage);

    return () => {
      websocket.close();
      websocket.removeEventListener("message", onMessage);
    };
  }, []);

  return (
    <>
      <Hamburger />
      {/* {SwitchingHub(schema)} */}
      {/* <button
        type="button"
        onClick={() => {
          socketRef.current?.send("gggo");
        }}
      >
        送信
      </button>
      <div>{message}</div> */}
      <GameContainer />
    </>
  );
};

export default Game;
