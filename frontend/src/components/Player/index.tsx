import { FC } from 'react'
import { Avatar } from '../Avatar'
import { Human } from '../Human'
import styles from './index.module.css'

type Props = {
  color: string
  name: string
  uuid: string
}

export const Player: FC<Props> = ({ color, name, uuid }) => {
  return (
    <div className={styles.container}>
      <Human color={color} size={130} />
      <div className={styles.nameContainer}>
        <Avatar name={uuid} color={color} size={25} />
        <div className={styles.name}>{name}</div>
      </div>
    </div>
  )
}
