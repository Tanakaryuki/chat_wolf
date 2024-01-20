import { FC, useEffect } from "react";
import styles from "./index.module.css";
import { useParticipantsStore } from "../../../../store/useParticipantsStore";
import { useNameStore, useUUIDStore } from "../../../../store";

export const EnterRoom: FC = () => {
  const addParticipant = useParticipantsStore((state) => state.addParticipant);
  const name = useNameStore((state) => state.name);
  const uuid = useUUIDStore((state) => state.uuid);
  useEffect(() => {
    addParticipant({ name: name, icon: uuid, id: uuid, score: 0 });
  }, []);
  return <div className={styles.container}></div>;
};
