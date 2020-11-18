import React, { useEffect, useState } from 'react'

const FilterForm = ({ updateRecipes }) => {
  const [filter, setFilter] = useState({ RecipeKey: 'get all' })

  const [tags, setTags] = useState([])


  useEffect(()=>{
    async function fetchingTags() {
      const { data } = await axios.post(
        'http://localhost:8080/getRecipes',
        recipes
      )
      setTags(data)
    }
    fetchingTags()
  }, [])
  return (
    <form onSubmit={(event) => updateRecipes(filter)}>
      <label>
        Region:
        <input
          name="region"
          type="checkbox"
          // checked={this.state.isGoing}
          // onChange={this.handleInputChange}
        />
      </label>
    </form>
  )
}

export default FilterForm
