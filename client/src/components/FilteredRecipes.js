import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'

const FilteredRecipes = ({ displayRecipes }) => {
  return (
    <div>
      {displayRecipes ? (
        displayRecipes.map((cur, ind) => {
          return (
            <div key={ind}>
              <Link
                to={`/user/${cur.userid}`}
              >
                <h6> {cur.ethnicity}</h6>
                <h6> {cur.method}</h6>
                <h6>{cur.recipename}</h6>
              </Link>
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
