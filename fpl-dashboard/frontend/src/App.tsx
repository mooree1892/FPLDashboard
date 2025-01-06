import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Dashboard from './pages/Dashboard.tsx';
import TransfersPage from './pages/TransfersPage.tsx';
import PlayersPage from './pages/PlayersPage.tsx';

const App: React.FC = () => {
  return (
    <Router>
      <div className="App">
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/transfers" element={<TransfersPage />} />
          <Route path="/players" element={<PlayersPage />} />
        </Routes>
      </div>
    </Router>
  );
};

export default App;