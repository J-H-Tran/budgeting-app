import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/HomePage';
import BudgetPage from './pages/BudgetPage';
import AccountsPage from './pages/AccountsPage';
import TransactionsPage from './pages/TransactionsPage';
import Navbar from './components/Navbar';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/budget" element={<BudgetPage />} />
        <Route path="/accounts" element={<AccountsPage />} />
        <Route path="/transactions" element={<TransactionsPage />} />
      </Routes>
    </Router>
  );
}

export default App;
