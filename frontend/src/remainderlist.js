import React, { useState, useCallback } from 'react';
import { KEY_RETURN } from 'keycode-js';

function InputBox(props) {
    const { addNewItem } = props;
    const [value, setValue] = useState('');
    const handleKeyUpEvent = useCallback(
      event => {
      if (event.keyCode === KEY_RETURN) {
        addNewItem(event.target.value);
        setValue('');
        console.log('KEY_RETURN pressed');
      }
    }, [addNewItem, setValue]);
    const handleChangeEvent = useCallback(
      event => {
        setValue(event.target.value);
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
  
  const { items, title, addNewItem } = props;

  const count = items.length;

  return (
    <div className="remainderlist">

        <header>
          <h1>{title}</h1>
          <InputBox addNewItem={addNewItem} />
        </header>

        <ul className="list-group list-group-flush">

          {Object.keys(items).map((item,i) => (
              <li key={i} className="list-group-item">
                  <div>
                  <label className="Remaindertext" htmlFor={`remainder-item-check-${item.Id}`}>
                    {items[item].Content}
                  </label>
                  </div>
                  <div>
                  <input className="Checkbox" type="checkbox" value="" id={`remainder-item-check-${item.Id}`} />
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