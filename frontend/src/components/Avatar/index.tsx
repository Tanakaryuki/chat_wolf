import BoringAvatar from 'boring-avatars'
import { FC } from 'react'
import styles from './index.module.css'

type Props = {
  size?: number
  name: string
  color: string
}

export const Avatar: FC<Props> = ({ size = 40, name, color }) => {
  return (
    <div
      className={styles.iconContainer}
      style={{
        width: `${size + 10}px`,
        height: `${size + 10}px`,
        backgroundColor: `${color}`,
      }}
    >
      <div style={{ margin: '5px' }}>
        <BoringAvatar size={size} name={name} variant='beam' />
      </div>
    </div>
  )
}
