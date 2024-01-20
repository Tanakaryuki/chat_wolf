import { FC } from "react";
import { Avatar } from "../../Avatar";
import { useUUIDStore } from "../../../store";
import { useNameStore } from "../../../store";
import styles from "./index.module.css";
import { Rotate } from "../../icons/Rotate";

export const Menu: FC = () => {
  const userIconId = useUUIDStore((state) => state.uuid);
  const generateUUID = useUUIDStore((state) => state.updateUUID);
  const userName = useNameStore((state) => state.name);
  const setName = useNameStore((state) => state.setName);
  return (
    <div className={styles.container}>
      <div>
        <div className={styles.avatar}>
          <Avatar name={userIconId} color="#00000000" size={50} />
        </div>
        <button onClick={() => generateUUID()} className={styles.button}>
          <Rotate color="black" size={15} />
        </button>
      </div>
      <div>
        <input
          value={userName}
          onChange={(e) => setName(e.target.value)}
          className={styles.nameInput}
        />
      </div>
    </div>
  );
};
