import React, { useState } from 'react';
import axios from 'axios';

const AddBudgetItemModal = ({ closeModal }) => {
  const [name, setName] = useState('');
  const [amount, setAmount] = useState('');
  const [category, setCategory] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    const newItem = { name, amount, category };
    axios.post('/api/budget-items', newItem)
      .then(response => {
        closeModal();
        window.location.reload(); // Reload to see new item
      })
      .catch(error => console.log(error));
  };

  return (
    <div className="modal">
      <h2>Add Budget Item</h2>
      <form onSubmit={handleSubmit}>
        <input type="text" value={name} onChange={(e) => setName(e.target.value)} placeholder="Name" required />
        <input type="number" value={amount} onChange={(e) => setAmount(e.target.value)} placeholder="Amount" required />
        <input type="text" value={category} onChange={(e) => setCategory(e.target.value)} placeholder="Category" required />
        <button type="submit">Add</button>
        <button onClick={closeModal}>Cancel</button>
      </form>
    </div>
  );
};

export default AddBudgetItemModal;
