import { Routes, Route } from "react-router-dom";
import { Top } from "./pages/Top";
// import { Game } from './pages/Game'
import { Suspense, lazy } from "react";
const Game = lazy(() => import("./pages/Game"));

// import About from './routes/about';
// import Contact from './routes/contact';

function Routing() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Routes>
        <Route path="/" element={<Top />} />
        <Route path="/game" element={<Game />} />
      </Routes>
    </Suspense>
  );
}

export default Routing;
