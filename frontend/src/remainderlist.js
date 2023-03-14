import React from 'react';

function RemainderList(props) {
  
  const { items, title } = props;

  const count = items.length;

  return (
    <div className="remainderlist">

          <h1>{title}</h1>

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