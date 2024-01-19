import type { FC } from 'react'
import { Button } from '../../components/Button/Button'
import styles from './index.module.css'
import { Title } from '../../components/Title'

export const Container: FC = () => {
  return (
    <main className={styles.container}>
        <div className={styles.titleContainer}>

      <Title>chat<br/>wolf</Title>
        </div>
      <div className={styles.buttonContainer}>
        <Button>ルーム作成</Button>
        <Button>ルーム入室</Button>
      </div>
    </main>
  )
}
