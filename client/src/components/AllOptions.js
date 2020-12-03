import React, { useEffect, useState } from 'react'
import AllRecipes from './AllRecipes'
import AllChefs from './AllChefs'

const AllOptions = () => {
  const [option, setOption] = useState(true)
  return <div>
    <button onClick={()=>setOption(false)}>See All Recipes</button>
    <button onClick={()=>setOption(true)}>See All Chefs</button>
    {option ? <AllChefs></AllChefs> : <AllRecipes></AllRecipes>}</div>
}

export default AllOptions
