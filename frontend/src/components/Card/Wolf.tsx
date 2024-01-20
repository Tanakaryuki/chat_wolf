import { FC } from 'react'
import styles from './wolf.module.css'

export const Wolf: FC = () => {
  return (
    <div className={styles['card-container']}>
      <div className={styles.card}>
        <div className={styles['card-front']}></div>
        <div className={styles['card-back']}></div>
      </div>
    </div>
  )
}
