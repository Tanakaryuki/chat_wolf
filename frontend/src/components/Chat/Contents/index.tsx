import { ComponentProps, FC } from 'react'
import { Message } from './Message'
import { generateUUID } from '../../../libs/uuid'
import styles from './index.module.css'
import { Form } from './Form'

export const Contents: FC = () => {
  const messages: Array<ComponentProps<typeof Message>> = [
    {
      name: 'hoge',
      icon: generateUUID(),
      message: 'メッセージのコンポーネントはこんな感じでいいと思うかな？',
      color: 'red',
    },
    {
      name: 'fuga',
      icon: generateUUID(),
      message: 'メッセージのコンポーネントはこんな感じでいいと思うかな？',
      color: '#aaaa77',
    },

    {
      name: 'simesabaaaaa',
      icon: generateUUID(),
      message: 'メッセージのコンポーネントはこんな感じでいいと思うかな？',
      color: 'black',
    },
    {
      name: 'niwaota',
      icon: generateUUID(),
      message: 'メッセージのコンポーネントはこんな感じでいいと思うかな？',
      color: 'green',
    },
  ]
  return (
    <div className={styles.container}>
        <div className={styles.messagesContainer}>
      {messages.map((e,i) => {
        return <Message key={i} message={e.message} icon={e.icon} name={e.name} color={e.color} />
      })}
      </div>
      <div className={styles.formContainer}>
        <Form />
      </div>
    </div>
  )
}
