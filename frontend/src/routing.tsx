import { Routes, Route } from 'react-router-dom'
import { Top } from './pages/Top'
import { Game } from './pages/Game'
// import About from './routes/about';
// import Contact from './routes/contact';

function Routing() {
  return (
    <Routes>
      <Route path='/' element={<Top />} />
      <Route path='/game' element={<Game />} />
    </Routes>
  )
}

export default Routing
