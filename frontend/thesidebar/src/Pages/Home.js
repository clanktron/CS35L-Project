import React, { useState, useEffect } from 'react'

export default function Home() {
  const [homeItems, setHomeItems] = useState([])

  useEffect(() => {
    fetch('https://www.google.com/') //fetch data from API, need help on this
      .then(response => response.json()) // parse into json file
      .then(data => setHomeItems(data))
      .catch(error => console.error(error))
  }, [])

  return (
    <ul>
      {homeItems.map(item => <li key={item.id}>{item.name}</li>)}     
    </ul>
  )
}
//list output