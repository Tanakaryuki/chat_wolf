import { FC, useEffect } from "react";
import styles from "./index.module.css";
import { useParticipantsStore } from "../../../../store/useParticipantsStore";
import { useNameStore, useRefStore, useUUIDStore } from "../../../../store";
import { EventType, Protocol } from "../../../../schema/common";
import { commonSchemaToJSON } from "../../../../schema/dto/common";
import { Colors } from "../../../../schema/status";

export const CreateRoom: FC = () => {
  const addParticipant = useParticipantsStore((state) => state.addParticipant);
  const name = useNameStore((state) => state.name);
  const uuid = useUUIDStore((state) => state.uuid);
  const ref = useRefStore((state) => state.socketRef);
  useEffect(() => {
    addParticipant({
      name: name,
      icon: uuid,
      id: uuid,
      score: 0,
      color: Colors.green,
    });
    const data: Protocol = {
      eventType: EventType.createRoom,
      user: {
        id: uuid,
        displayName: name,
        icon: uuid,
      },
      room: {
        roomOwnerId: uuid,
      },
      option: {
        turnNum: 5,
        discussTime: 180,
        voteTime: 60,
        participantsNum: 6,
      },
    };
    const dataString = commonSchemaToJSON(data);
    ref?.current?.send(dataString);
  }, []);
  return <div className={styles.container}></div>;
};
