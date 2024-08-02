import React from 'react';
import BudgetItem from './BudgetItem';

const BudgetCategory = ({ category, items }) => {
  return (
    <div>
      <h2>{category}</h2>
      {items.map(item => (
        <BudgetItem key={item.id} item={item} />
      ))}
    </div>
  );
};

export default BudgetCategory;
