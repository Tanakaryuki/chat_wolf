import type { FC } from 'react'
import { Button } from '../../components/Button/Button'
import styles from './index.module.css'
import { Title } from '../../components/Title'
import { useUUIDStore } from '../../store/useUUIDStore'
import { Avatar } from '../../components/Avatar'

export const Container: FC = () => {
  const userIconId = useUUIDStore(state => state.uuid)
  const generateUUID = useUUIDStore(state => state.updateUUID)
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
        <Button>ルーム作成</Button>
        <Button>ルーム入室</Button>
        <button onClick={() => generateUUID()}>hoge</button>
        {userIconId}
        <Avatar size={100} color='red' name={userIconId} />
      </div>
    </main>
  )
}
