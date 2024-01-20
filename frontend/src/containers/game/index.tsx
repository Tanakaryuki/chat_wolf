import { FC } from 'react'
import { Protocol } from '../../schema/common'
import { SwitchingHub } from './SwitchingHub'

type Props = {
  schema: Protocol
}

export const Game: FC<Props> = ({ schema }) => {
  return <>{SwitchingHub(schema)}</>
}
