import React, { useEffect, useState } from 'react'

type User = {
  id: number,
  name: string,
  "e-mail": string
};

function Helloworld() {
  const [user, setUser] = useState<User>()
  useEffect(() => {
    (async() => {
      const res = await fetch("/helloworld")
      const data = await res.json()
      console.log(data)
      setUser(data)
    })()
  }, [])

  return (
    <>
      <h1>Hello world!</h1>
      <p>{user?.name}</p>
      <p>{user?.id}</p>
      <p>{user?.['e-mail']}</p>
    </>
  )
}
export default Helloworld;
