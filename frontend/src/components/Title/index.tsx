import { FC, ReactNode } from 'react'
import styles from './index.module.css'

type Props = {
  children: ReactNode
}

export const Title: FC<Props> = ({ children }) => {
  return <div className={styles.title}>{children}</div>
}
