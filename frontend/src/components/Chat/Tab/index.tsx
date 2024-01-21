import { FC, useState } from "react";
import styles from "./index.module.css";
import { Contents } from "../Contents";

export const Tab: FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div
      className={`${styles.container} ${
        isOpen ? styles.active : styles.deactive
      }`}
    >
      <div onClick={() => setIsOpen(!isOpen)}>
        <div className={styles.chat}>チャット</div>
      </div>
      <div className={styles.contentsContainer}>
        <Contents />
      </div>
    </div>
  );
};
