import React, { useEffect, useState } from 'react'

const FilterForm = ({ updateRecipes }) => {
  const [filter, setFilter] = useState({ RecipeKey: 'get all' })

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
