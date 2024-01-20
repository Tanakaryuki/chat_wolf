import { FaUser } from 'react-icons/fa6'
import { IconContext } from 'react-icons'
import { FC } from 'react'

type Props = {
  color: string
  size: number
  className?: string
}

export const Human: FC<Props> = ({ color, size, className }) => {
  return (
    <IconContext.Provider value={{ color: color, size: size + 'px', className: className }}>
      <FaUser />
    </IconContext.Provider>
  )
}
