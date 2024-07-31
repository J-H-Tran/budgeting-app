// HomePage.test.js
import { render, screen } from '@testing-library/react';
import HomePage from './HomePage';

test('renders HomePage component', () => {
  render(<HomePage />);
  const linkElement = screen.getByText(/welcome to the budgeting app/i);
  expect(linkElement).toBeInTheDocument();
});
