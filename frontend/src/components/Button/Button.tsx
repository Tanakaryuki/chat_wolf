import { FC, MouseEventHandler, ReactNode } from "react";
import styles from "./index.module.css";

type Props = {
  children: ReactNode;
  onClick?: MouseEventHandler<HTMLButtonElement>;
};

export const Button: FC<Props> = ({ children, onClick = undefined }) => {
  return (
    <button type="button" className={styles.button} onClick={onClick}>
      {children}
    </button>
  );
};
