import { useEffect, useRef, useState } from "react";
import { Hamburger } from "../components/Hamburger/Hamburger";
import { GameContainer } from "../containers/game";
import ReconnectingWebSocket from "reconnecting-websocket";
import { useGameStatusStore } from "../store";
import { EventType } from "../schema/common";
import { transferMessage } from "../libs/webSocket";
import { Protocol } from "../schema/common";
// const ReconnectingWebSocket = lazy(() => import("reconnecting-websocket"));

const Game = () => {
  const [message, setMessage] = useState<string>();
  const [_schema, setSchema] = useState<Protocol>();
  const socketRef = useRef<ReconnectingWebSocket>();
  const query = new URLSearchParams(location.search);
  const roomId = query.get("id");
  const eventype = query.get("type");
  //   const status = useGameStatusStore((state) => state.status);
  const setStatus = useGameStatusStore((state) => state.setStatus);
  const setRoomId = useGameStatusStore((state) => state.setRoomId);

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
      import.meta.env.VITE_WS_URL ?? "ws://localhost:8080/ws"
    );
    socketRef.current = websocket;

    const onMessage = (event: MessageEvent<string>) => {
      setSchema(transferMessage(event));
      setMessage(event.data);
    };

    console.log(message);
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
