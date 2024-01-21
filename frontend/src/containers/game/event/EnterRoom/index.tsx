import { FC, useEffect } from "react";
import styles from "./index.module.css";
import {
  useGameStatusStore,
  useNameStore,
  useRefStore,
  useUUIDStore,
} from "../../../../store";
import { commonSchemaToJSON } from "../../../../schema/dto/common";
import { EventType, Protocol } from "../../../../schema/common";
// import { Colors } from "../../../../schema/status";

export const EnterRoom: FC = () => {
  //   const addParticipant = useParticipantsStore((state) => state.addParticipant);
  const name = useNameStore((state) => state.name);
  const uuid = useUUIDStore((state) => state.uuid);
  const ref = useRefStore((state) => state.socketRef);
  const roomId = useGameStatusStore((state) => state.meta.roomId);
  useEffect(() => {
    const data: Protocol = {
      eventType: EventType.enterRoom,
      user: {
        id: uuid,
        displayName: name,
        icon: uuid,
        isParticipant: true,
      },
      room: {
        roomId: roomId,
      },
    };
    const dataString = commonSchemaToJSON(data);
    ref?.current?.send(dataString);
  }, []);

  return <div className={styles.container}></div>;
};
