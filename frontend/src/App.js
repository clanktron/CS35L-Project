import React, { useState, useCallback } from 'react';
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
  const [items, updateItems] = useState(ITEMS_INITIAL_STATE);

  const addNewItem = useCallback(
    text => {
      updateItems(items => {
        const nextId = items.length + 1;
        const newItem = {
          id: nextId,
          text: text
        };

        return [...items, newItem];
      });
    },
    [updateItems]
  );
  



  return (
    <div className="container">
      <div className="row">
        <RemainderList title={title} items={items} addNewItem={addNewItem}/>
      </div>
    </div>
  );
}

export default App;
