import React, { useEffect, useState } from 'react'
import axios from 'axios'

const FilteredRecipes = ({ displayRecipes }) =>{
  return (
    <div>
      {displayRecipes? displayRecipes.map((cur,ind)=>{
        return (
          <div key={ind}>
          <p>{cur.ethnicity}
          {cur.method}
          {cur.recipename}
          </p>
          </div>
        )
      }):<div>No recipes with this combo</div>}
    </div>
  )
}

export default FilteredRecipes
