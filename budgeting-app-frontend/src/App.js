import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import HomePage from './pages/HomePage';
import BudgetPage from './pages/BudgetPage';
import AccountsPage from './pages/AccountsPage';
import TransactionsPage from './pages/TransactionsPage';

function App() {
  return (
    <Router>
      <Switch>
        <Route path="/" exact component={HomePage} />
        <Route path="/budget" component={BudgetPage} />
        <Route path="/accounts" component={AccountsPage} />
        <Route path="/transactions" component={TransactionsPage} />
      </Switch>
    </Router>
  );
}

export default App;
