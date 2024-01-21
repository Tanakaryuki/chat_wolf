import type { FC, ReactNode } from "react";
import styles from "./index.module.css";

type Props = {
  children: ReactNode;
};

export const PlayingContainer: FC<Props> = ({ children }) => {
  return (
    <main className={styles.container}>
      <div className={styles.cloud} />
      {children}
    </main>
  );
};
