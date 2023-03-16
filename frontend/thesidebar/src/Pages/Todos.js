import React, { useState, useEffect } from 'react'

export default function Todos() {
  const [todosItems, setTodosItems] = useState([])

  useEffect(() => {
    fetch('https://www.google.com/') //fetch data from API, need help on this
      .then(response => response.json()) // parse into json file
      .then(data => setTodosItems(data))
      .catch(error => console.error(error))
  }, [])

  return (
    <ul>
      {todosItems.map(item => <li key={item.id}>{item.name}</li>)}     
    </ul>
  )
}
//list output