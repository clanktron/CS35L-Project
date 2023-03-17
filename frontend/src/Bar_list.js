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
    className="Bar_InputBar" 
    value={value}
    onKeyUp={handleKeyUpEvent}
    onChange={handleChangeEvent}
    placeholder="Plan your work and work your plan." />;
}


function RemainderBar(props) {
  
  const {items,title, addNewItem, deleteItem} = props;

  const count = items?.length;

  return (
    <div className="Bar_remainderlist">

        <header>
          <h1>{title}</h1>
          <InputBox addNewItem={addNewItem} />
        </header>

        <ul className="Bar_list-group list-group-flush">

          {items && Object.keys(items).map((item,i) => (
              <li key={i} className="Bar_list-group-item">

                  <div>
                  <label className="Bar_Remaindertext" htmlFor={`remainder-item-check-${i}`}>
                    {items[item].Name}
                  </label>
                  </div>

                  <div>
                  <button className="Bar_Deletebutton" onClick={()=>deleteItem(items[item].Name)}>
                    Finish!
                  </button>
                  </div>

              </li>
            ))}

        </ul>

        <div className="Bar_remainder-footer">
          <span className="Bar_count-remainders">{count}</span>
          {' tasks left'}
        </div>

    </div>
  );
}

export default RemainderBar;