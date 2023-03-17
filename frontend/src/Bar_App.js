import React, { useState, useCallback, useEffect } from 'react';
import {parse, stringify} from 'lossless-json'
import RemainderBar from './Bar_list.js';

//create a list remote at the beginning (if needed)



const LISTS_INITIAL_STATE = [];

function To_do_bar() {
  let bar_path = 'http://localhost:4000/list/'
  let title = 'To-do-bar';
  const [lists, updateLists] = useState(LISTS_INITIAL_STATE);
  const [if_bar_exist, set_if_bar_exist] = useState(true);

  const create_list = async (listname) => {
   await fetch('http://localhost:4000/list/', {
      method: 'POST',
      mode: 'cors',
      credentials: 'include',
      headers: {
         'Content-type': 'application/json; charset=UTF-8',
      },
      body: JSON.stringify({
         name:listname
      }),
   })
      .then((data) => console.log(data));
      getBars();
 };

//create a bar for test
  useState(()=>{
   if(if_bar_exist){
       create_list("badexample");
   }
   set_if_bar_exist('false');
},[])

//get a list from the remote list
const getBars = async () => {
  await fetch(`${bar_path}`, {
     method: 'GET',
     mode: 'cors',
     credentials: 'include',
     headers: {
        'Content-type': 'application/json',
     },
  })
      .then(response => { return response.text();})
      .then(responseData => {console.log(parse(responseData)); return parse(responseData);})
      .then(response => {response && updateLists((response));})
      .catch(err => console.error(err));
};

//add
//the same function as create_list

//delete item remotely
  const deletebar = async(name)=> {
    await fetch(`${bar_path}/${name}`, {
       method: 'DELETE',
       mode: 'cors',
       credentials: 'include',
       headers: {'Content-type': 'application/json'}
    })
    .then((res) => console.log(`${bar_path}/${name}`));
     getBars();
  };

//get the Note once to start
   useEffect(() => {
     getBars();
   },[])

  
  return (
    <div className="Bar_container">
      <div className="Bar_row">
      <RemainderBar title={title} items={lists} addNewItem={create_list} deleteItem={deletebar}/> 
      </div>
    </div>
  );
}


export default To_do_bar;
