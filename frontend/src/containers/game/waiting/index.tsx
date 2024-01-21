import { type FC, type ReactNode } from "react";
import styles from "./index.module.css";
// import { Chat } from "../../../components/Chat";
// import { People } from "../../../components/People";
// import { useParticipantsStore } from "../../../store/useParticipantsStore";
import { copyStringToClipboard } from "../../../libs/copyStringToClipboard";
import { useGameStatusStore } from "../../../store";
import { Copy } from "../../../components/icons/Copy";
// import { Colors } from "../../../schema/status";

type Props = {
  children: ReactNode;
};

export const WaitingContainer: FC<Props> = ({ children }) => {
  //   const participants = useParticipantsStore((state) => state.participants);
  const meta = useGameStatusStore((state) => state.meta);
  const url = `${location.protocol}://${location.hostname}:${location.port}/game/?type=enter_room&id=${meta.roomId}`;
  //   const addParticipant = useParticipantsStore((state) => state.addParticipant);
  //   useEffect(() => {
  //     addParticipant({
  //       name: "hoge",
  //       icon: "test",
  //       id: "888",
  //       score: 0,
  //       color: Colors.green,
  //     });
  //   }, []);
  return (
    <main className={styles.container}>
      {/* <Chat />
      <div className={styles.peopleContainer}>
        <People participants={participants} />
      </div> */}
      <div
        className={styles.linkContainer}
        onClick={() => copyStringToClipboard(url)}
      >
        <div className={styles.linkText}>招待リンクをコピー</div>
        <Copy color="black" size={35} />
      </div>
      {children}
    </main>
  );
};
