import type { FC } from "react";
import { Button } from "../../components/Button/Button";
import styles from "./index.module.css";
import { Title } from "../../components/Title";
import { Link } from "react-router-dom";
import { generateUUID } from "../../libs/uuid";

export const Container: FC = () => {
  const uuid = generateUUID();
  return (
    <main className={styles.container}>
      <div className={styles.titleContainer}>
        <Title>
          chat
          <br />
          wolf
        </Title>
      </div>
      <div className={styles.buttonContainer}>
        <Link to={`/game/?type=create_room&id=${uuid}`}>
          <Button>ルーム作成</Button>
        </Link>
        <Link to={`/game/?type=enter_room`}>
          <Button>ルーム入室</Button>
        </Link>
      </div>
    </main>
  );
};
