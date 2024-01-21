import { EventType } from "../../schema/common";
import { PlayingContainer } from "./playing";
import { WaitingContainer } from "./waiting";
import { ResultContainer } from "./result";
import { CreateRoom } from "./event/CreateRoom";
import { EnterRoom } from "./event/EnterRoom";

export const SwitchingHub = (eventType: EventType) => {
  switch (eventType) {
    case EventType.askQuestion:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    case EventType.changeRoomOwner:
      return (
        <WaitingContainer>
          <></>
        </WaitingContainer>
      );
    case EventType.createRoom:
      return (
        <WaitingContainer>
          <CreateRoom />
        </WaitingContainer>
      );
    case EventType.endQandA:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    case EventType.enterRoom:
      return (
        <WaitingContainer>
          <EnterRoom />
        </WaitingContainer>
      );
    case EventType.exitRoom:
      return (
        <WaitingContainer>
          <></>
        </WaitingContainer>
      );
    case EventType.gameResult:
      return (
        <ResultContainer>
          <></>
        </ResultContainer>
      );
    case EventType.giveAnswer:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    case EventType.prepareCompletion:
      return (
        <WaitingContainer>
          <></>
        </WaitingContainer>
      );
    case EventType.sendChat:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    case EventType.sendTime:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    case EventType.setOption:
      return (
        <WaitingContainer>
          <></>
        </WaitingContainer>
      );
    case EventType.startGame:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    case EventType.vote:
      return (
        <PlayingContainer>
          <></>
        </PlayingContainer>
      );
    default:
      return (
        <></>
        // <WaitingContainer>
        //   <></>
        // </WaitingContainer>
      );
  }
};
