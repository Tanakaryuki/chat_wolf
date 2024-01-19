import { Routes, Route } from 'react-router-dom'
import { Top } from './pages/Top'
// import About from './routes/about';
// import Contact from './routes/contact';

function Routing() {
  return (
    <Routes>
      <Route path='/' element={<Top />} />
    </Routes>
  )
}

export default Routing
