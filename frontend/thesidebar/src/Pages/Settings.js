import React, { useState, useEffect } from 'react'

export default function Settings() {
  const [settingsItems, setSettingsItems] = useState([])

  useEffect(() => {
    fetch('https://www.google.com/') //fetch data from API, need help on this
      .then(response => response.json()) // parse into json file
      .then(data => setSettingsItems(data))
      .catch(error => console.error(error))
  }, [])

  return (
    <ul>
      {settingsItems.map(item => <li key={item.id}>{item.name}</li>)}     
    </ul>
  )
}
//list output