import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'

const FilteredRecipes = ({ displayRecipes }) => {
  console.log(displayRecipes, 'fgdg')
  return (
    <div>
      {displayRecipes ? (
        displayRecipes.map((cur, ind) => {
          return (
            <div key={ind} >
              <Link to={`/user/${cur.userid}`}>YOLO</Link>
              <p>
                {cur.ethnicity}
                {cur.method}
                {cur.recipename}
              </p>
            </div>
          )
        })
      ) : (
        <div>No recipes with this combo</div>
      )}
    </div>
  )
}

export default FilteredRecipes
