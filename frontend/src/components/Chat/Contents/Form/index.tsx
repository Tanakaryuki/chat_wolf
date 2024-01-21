import { FC, useState } from "react";
import styles from "./index.module.css";
import {
  useGameStatusStore,
  useNameStore,
  useRefStore,
  useUUIDStore,
} from "../../../../store";
import { commonSchemaToJSON } from "../../../../schema/dto/common";
import { EventType, Protocol } from "../../../../schema/common";

export const Form: FC = () => {
  const [message, setMessage] = useState<string>("");
  const ref = useRefStore((state) => state.socketRef);
  const id = useUUIDStore((state) => state.uuid);
  const name = useNameStore((state) => state.name);
  const roomId = useGameStatusStore((state) => state.meta.roomId);
  return (
    <>
      <div className={styles.container}>
        <input
          className={styles.form}
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="入力してね！"
        />
        <button
          className={styles.button}
          type="button"
          onClick={() => {
            const data: Protocol = {
              eventType: EventType.sendChat,
              user: {
                id: id,
                displayName: name,
                icon: id,
                isParticipant: true,
              },
              chatText: message,
              room: {
                roomId: roomId,
              },
            };
            const dataString = commonSchemaToJSON(data);
            ref?.current?.send(dataString);
          }}
        >
          送信
        </button>
      </div>
    </>
  );
};
