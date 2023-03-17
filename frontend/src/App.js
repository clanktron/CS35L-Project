import React, { useState, useCallback, useEffect } from 'react';
import Bar_RemainderList from './remainderlist';
import { parse, stringify } from 'lossless-json';

//create a list remote at the beginning (if needed)
const Bar_create_list = async () => {
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

//Bar_create_list();

const ITEMS_INITIAL_STATE = [];

function Bar_App() {
  let list_path = 'http://localhost:4000/list/Test_list'
  let title = 'To-Dos';
  const [items, updateItems] = useState(ITEMS_INITIAL_STATE);

//get a notes from the remote list
const Bar_getNotes = async () => {
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
      .then(response => {updateItems((response));})
      .catch(err => console.error(err));
};


//add new item remotely
  const Bar_additem = async (text) => {
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
    Bar_getNotes();
  };


//delete item remotely
  const Bar_deleteitem = async (number) => {
    await fetch(`${list_path}/note/${number}`, {
       method: 'DELETE',
       headers: {'Content-type': 'application/json'}
    })
       .then((data) => console.log(`${list_path}/note/${number}`));
     Bar_getNotes();
  };
  console.log(stringify(parse()))
 


//get the Note once to start
  useEffect(() => {
   Bar_additem("sample todo");
   },[])
  
  return (
    <div className="Bar_container">
      <div className="Bar_row">
      <Bar_RemainderList title={title} items={items} addNewItem={Bar_additem} deleteItem={Bar_deleteitem}/> 
      </div>
    </div>
  );
}


export default Bar_App;

