import { FC } from 'react'
import { Avatar } from '../../../Avatar'
import styles from './index.module.css'

type Props = {
  message: string
  name: string
  icon: string
  color: string
}

export const Message: FC<Props> = ({ message, name, icon, color }) => {
  return (
    <div className={styles.container}>
      <div className={styles.userContainer}>
        <Avatar name={icon} color={color} size={20} />
        <div className={styles.name}>{name}</div>
      </div>
      <div className={styles.messageContainer}>
        <div className={styles.message}>{message}</div>
      </div>
    </div>
  )
}
