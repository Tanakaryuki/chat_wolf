import { FaRegCopy } from "react-icons/fa6";
import { IconContext } from "react-icons";
import { FC } from "react";

type Props = {
  color: string;
  size: number;
  className?: string;
};

export const Copy: FC<Props> = ({ color, size, className }) => {
  return (
    <IconContext.Provider
      value={{ color: color, size: size + "px", className: className }}
    >
      <FaRegCopy />
    </IconContext.Provider>
  );
};
