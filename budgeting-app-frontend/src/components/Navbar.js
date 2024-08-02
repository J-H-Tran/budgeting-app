import React from 'react';
import { Link } from 'react-router-dom';
import '../styles/Navbar.css'; // Ensure the path is correct

const Navbar = () => {
  return (
    <nav>
      <ul className="navbar">
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/budget">Budget</Link>
        </li>
        <li>
          <Link to="/accounts">Accounts</Link>
        </li>
        <li>
          <Link to="/transactions">Transactions</Link>
        </li>
      </ul>
    </nav>
  );
};

export default Navbar;
