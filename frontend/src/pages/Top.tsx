import { Wolf } from "../components/Card/Wolf";
import { Contents } from "../components/Chat/Contents";
// import { Tab } from "../components/Chat/Tab";
import { Hamburger } from "../components/Hamburger/Hamburger";
import { Human } from "../components/Human";
import { Player } from "../components/Player";
import { Container } from "../containers/top";
import { Link } from "react-router-dom";

export const Top = () => {
  return (
    <>
      <Hamburger />
      <Container />
      <a href={"/game"}>hoge</a>
      {/* <Citizen></Citizen> */}
      <Wolf />
      <Human color="red" size={120} />
      <Human color="yellow" size={120} />
      <Human color="red" size={120} />
      <Player color="red" uuid="test" name="Kizuku" />
      <Contents />
      {/* <Chat /> */}
      <Link to={"/game"}>game</Link>
    </>
  );
};
