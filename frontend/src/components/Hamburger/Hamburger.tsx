import styles from "./index.module.css";
import { useState } from "react";
import Popup from "reactjs-popup";
import { Menu } from "./Menu";

export const Hamburger = () => {
  const [isOpen, setIsOpen] = useState(false);
  return (
    <div className={styles.container}>
      <Popup
        trigger={
          <div
            className={`${styles.openbtn8} ${isOpen ? styles.active : ""}`}
            // onClick={() => setIsOpen(!isOpen)}
          >
            <div className={styles.openbtnArea}>
              <span></span>
              <span></span>
              <span></span>
            </div>
          </div>
        }
        onOpen={() => setIsOpen(true)}
        onClose={() => setIsOpen(false)}
        position="bottom right"

        // modal
        // nested
        //   open
        //   lockScroll={true}
      >
        <>
          <Menu />
        </>
      </Popup>
    </div>
  );
};
