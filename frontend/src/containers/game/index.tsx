import { FC } from "react";
import { SwitchingHub } from "./SwitchingHub";
import { useGameStatusStore } from "../../store";
import { useParticipantsStore } from "../../store/useParticipantsStore";
import styles from "./index.module.css";
import { Chat } from "../../components/Chat";
import { People } from "../../components/People";

// type Props = {
//   schema: Protocol;
// };

export const GameContainer: FC = () => {
  const status = useGameStatusStore((state) => state.status);
  const participants = useParticipantsStore((state) => state.participants);

  return (
    <main>
      <Chat />
      <div className={styles.peopleContainer}>
        <People participants={participants} />
      </div>
      {SwitchingHub(status)}
    </main>
  );
};
