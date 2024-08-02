import React from 'react';
import axios from 'axios';

const RemoveBudgetItemModal = ({ closeModal, itemId }) => {
  const handleRemove = () => {
    axios.delete(`/api/budget-items/${itemId}`)
      .then(response => {
        closeModal();
        window.location.reload(); // Reload to see the item removed
      })
      .catch(error => console.log(error));
  };

  return (
    <div className="modal">
      <h2>Remove Budget Item</h2>
      <p>Are you sure you want to remove this budget item?</p>
      <button onClick={handleRemove}>Remove</button>
      <button onClick={closeModal}>Cancel</button>
    </div>
  );
};

export default RemoveBudgetItemModal;
