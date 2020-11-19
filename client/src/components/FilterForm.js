import React, { useEffect, useState } from 'react'
import axios from 'axios'


const FilterForm = ({ updateRecipes }) => {
  const [filter, setFilter] = useState({ RecipeKey: 'get all' })

  const [tags, setTags] = useState([])

  useEffect(()=>{
    async function fetchingTags() {
      const { data } = await axios.get(
        'http://localhost:8080/getTags')
      setTags(data)
    }
    fetchingTags()
  }, [])

  console.log(tags)
  return (
    <form onSubmit={(event) => updateRecipes(filter)}>

    </form>
  )
}

export default FilterForm
