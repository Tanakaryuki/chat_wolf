import { useEffect, useRef, useState } from "react";
import { Hamburger } from "../components/Hamburger/Hamburger";
import { Container } from "../containers/top";
import ReconnectingWebSocket from "reconnecting-websocket";
// import { transferMessage } from "../libs/webSocket";
// import { Protocol } from "../schema/common";
// import { SwitchingHub } from "../containers/game/SwitchingHub";

export const Game = () => {
  const [message, setMessage] = useState<string>();
  //   const [schema, setSchema] = useState<Protocol>();
  const socketRef = useRef<ReconnectingWebSocket>();

  useEffect(() => {
    const websocket = new ReconnectingWebSocket(
      import.meta.env.REACT_APP_WS_URL ?? "ws://localhost:8080/ws"
    );
    socketRef.current = websocket;

    const onMessage = (event: MessageEvent<string>) => {
      //   setSchema(transferMessage(event))
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
      <button
        type="button"
        onClick={() => {
          socketRef.current?.send("gggo");
        }}
      >
        送信
      </button>
      <div>{message}</div>
      <Container />
    </>
  );
};
