import React, { useState } from 'react'
import RemainderList from './remainderlist';

const ITEMS_INITIAL_STATE = [
  {
    id: 1,
    text: 'Finish Frontend',
    completed: false
  },
  {
    id: 2,
    text: 'Catch up recording',
    completed: false
  }
];


function App() {
  let title = 'Remainder';
  const [items] = useState(ITEMS_INITIAL_STATE);

  return (
    <div className="container">
      <div className="row">
        <RemainderList title={title} items={items} />
      </div>
    </div>
  );
}

export default App;
