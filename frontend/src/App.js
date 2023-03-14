import React from 'react'
import RemainderList from './remainderlist';

function App() {
  let items = [
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
  let title = 'Remainder';

  return (
    <div className="container">
      <div className="row">
        <RemainderList title={title} items={items} />
      </div>
    </div>
  );
}

export default App;
