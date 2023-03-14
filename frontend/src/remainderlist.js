import React, { useState, useCallback } from 'react';
import { KEY_RETURN } from 'keycode-js';

function InputBox() {
    const [value, setValue] = useState('');
    const handleKeyUpEvent = useCallback(
      e => {
      if (e.keyCode === KEY_RETURN) {
        // Add new Todo Here
        // Clear the text box
        console.log('KEY_RETURN pressed');
      }
    }, []);
    const handleChangeEvent = useCallback(
      e => {
        setValue(e.target.value);
      },
      [setValue]
    );

    return <input 
    type="text" 
    className="InputBar" 
    value={value}
    onKeyUp={handleKeyUpEvent}
    onChange={handleChangeEvent}
    placeholder="Plan your work and work your plan." />;
}


function RemainderList(props) {
  
  const { items, title } = props;

  const count = items.length;

  return (
    <div className="remainderlist">

        <header>
        <h1>{title.toUpperCase()}</h1>
          <InputBox />
        </header>

        <ul className="list-group list-group-flush">

          {items.map(item => (
              <li key={item.id} className="list-group-item">
                <div className="form-check">
                  <input className="form-check-input" type="checkbox" value="" id={`tremainder-item-check-${item.id}`} />
                  <label className="form-check-label" htmlFor={`remainder-item-check-${item.id}`}>
                    {item.text}
                  </label>
                </div>
              </li>
            ))}

        </ul>

        <div className="remainder-footer">
          <span className="count-remainders">{count}</span>
          {' tasks left'}
        </div>

    </div>
  );
}

export default RemainderList;