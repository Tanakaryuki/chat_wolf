import { FC, ReactNode } from 'react'
import styles from './index.module.css'

type Props = {
  children: ReactNode
}

export const Button: FC<Props> = ({ children }) => {
  return (
    <button type='button' className={styles.button}>
      {children}
    </button>
  )
}
