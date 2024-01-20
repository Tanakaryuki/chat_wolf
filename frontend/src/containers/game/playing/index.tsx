import type { FC, ReactNode } from "react";
import styles from "./index.module.css";

type Props = {
  children: ReactNode;
};

export const Container: FC<Props> = ({ children }) => {
  return <main className={styles.container}>{children}</main>;
};
