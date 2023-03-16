import React, { useState, useEffect } from 'react'

export default function Home() {
  const [username, password] = useState([])

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('http://localhost:4000/list', {
        method: 'GET',
        mode: 'cors',
        credentials: 'include',
        headers: {
          'Content-type': 'application/json; charset=UTF-8',
        },
      })
      const data = await response.json()
      const names = data.map(item => item.name)
      password(names)
    }
    fetchData().catch(error => console.error(error))
  }, [])

  return (
    <ul>
      {username.map((name, index) => <li key={index}>{name}</li>)}     
    </ul>
  )
}