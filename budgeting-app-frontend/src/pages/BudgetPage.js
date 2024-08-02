import React, { useEffect, useState } from 'react';
import axios from 'axios';
import BudgetCategory from '../components/BudgetCategory';
import AddBudgetItemModal from '../components/AddBudgetItemModal';

const BudgetPage = () => {
  const [budgetItems, setBudgetItems] = useState([]);
  const [categories, setCategories] = useState([]);
  const [showAddModal, setShowAddModal] = useState(false);

  useEffect(() => {
    // Fetch budget items and categories
    axios.get('/api/budget-items')
      .then(response => {
        setBudgetItems(response.data);
        setCategories(getCategories(response.data));
      })
      .catch(error => console.log(error));
  }, []);

  const getCategories = (items) => {
    // Function to extract categories from budget items
    const categories = [...new Set(items.map(item => item.category))];
    return categories;
  };

  return (
    <div>
      <h1>Budget Items</h1>
      <button onClick={() => setShowAddModal(true)}>Add Budget Item</button>
      {categories.map(category => (
        <BudgetCategory key={category} category={category} items={budgetItems.filter(item => item.category === category)} />
      ))}
      {showAddModal && <AddBudgetItemModal closeModal={() => setShowAddModal(false)} />}
    </div>
  );
};

export default BudgetPage;
