import React, { useState, useCallback, useEffect } from 'react';
import RemainderList from './remainderlist';

const ITEMS_INITIAL_STATE = [];

function App() {
  let list_path = 'http://localhost:4000/list/Test_list'
  let title = 'Remainder';
  const [items, updateItems] = useState(ITEMS_INITIAL_STATE);

  //create a list remote at the beginning (if needed)
  const create_list = async () => {
    await fetch('http://localhost:4000/list/', {
       method: 'POST',
       mode: 'cors',
       credentials: 'include',
       headers: {
          'Content-type': 'application/json; charset=UTF-8',
       },
       body: JSON.stringify({
          name: "Test_list"
       }),
    })
       .then((data) => console.log(data));
  };

  create_list();

//get a notes from the remote list
const getNotes = async () => {
  await fetch(`${list_path}/note`, {
     method: 'GET',
     mode: 'cors',
     credentials: 'include',
     headers: {
        'Content-type': 'application/json',
     },
  })
      .then(response => { return response.json();})
      .then(responseData => {console.log(responseData); return responseData;})
      .then(response => {updateItems((response))})
      .catch(err => console.error(err));
};


//add new item remotely
  const additem = async (text) => {
    await fetch('http://localhost:4000/list/Test_list/note', {
       method: 'POST',
       mode: 'cors',
       credentials: 'include',
       headers: {
          'Content-type': 'application/json; charset=UTF-8',
       },
       body: JSON.stringify({
          content:text
       }),
    })
       .then((data) => console.log(data));

    getNotes();
  };

//delete item remotely
  const deleteitem = async (number) => {
    await fetch(`${list_path}/note/${number}`, {
       method: 'DELETE',
       mode: 'cors',
       credentials: 'include',
       headers: {
          'Content-type': 'application/json; charset=UTF-8',
       },
    })
       .then((data) => console.log(data));
     getNotes();
  };

  

//get the Note once to start
  useEffect(() => {
    getNotes();},[])
  console.log(Object.keys(items))
  

  return (
    <div className="container">
      <div className="row">
      <RemainderList title={title} items={items} addNewItem={additem} deleteItem={deleteitem}/> 
      </div>
    </div>
  );
}


export default App;
