import { FC } from "react";
import { Participant } from "../../schema/status";
import { Player } from "../Player";
import styles from "./index.module.css";

type Props = {
  participants: Participant[];
};

export const People: FC<Props> = ({ participants }) => {
  const colors = [
    "#FF0000",
    "#FFA500",
    "#95DC4E",
    "#009944",
    "#2986E8",
    "#954EDC",
  ];
  return (
    <div className={styles.container}>
      {participants.map((e, i) => {
        return (
          <>
            {i < 6 && (
              <Player
                key={e.id}
                name={e.name}
                uuid={e.icon}
                color={colors[i]}
              />
            )}
          </>
        );
      })}
    </div>
  );
};
