import { FC, useState } from "react";
import styles from "./index.module.css";

export const Form: FC = () => {
  const [message, setMessage] = useState<string>("");
  return (
    <>
      <div className={styles.container}>
        <input
          className={styles.form}
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="入力してね！"
        />
        <input type="button">送信</input>
      </div>
    </>
  );
};
