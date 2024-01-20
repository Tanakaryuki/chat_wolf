import { Protocol } from '../schema/common'
import { transferCommonSchema } from '../schema/dto/common'

export const transferMessage = (event: MessageEvent<string>): Protocol => {
  const data = transferCommonSchema(event.data)
  return data
}
