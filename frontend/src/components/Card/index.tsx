import { FC } from 'react'
import { Wolf } from './Wolf'
import { Citizen } from './Citizen'

type Props = {
  isWolf: boolean
}

export const Card: FC<Props> = ({ isWolf }) => {
  return <>{isWolf ? <Wolf /> : <Citizen />}</>
}
