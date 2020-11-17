import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'
import { getRecipesThunk } from '../reducers/recipes'
import FilterForm from './FilterForm'

const AllRecipes = (props) => {
  const [recipes, setRecipes] = useState({ RecipeKey: 'get all' })

  const updateRecipes = (filters) => {
    setRecipes(filters)
  }

  useEffect(() => {
    props.getAllRecipes(recipes)
  }, [])
  return (
    <div>
      <div>All Recipes</div>
      <FilterForm updateRecipes={updateRecipes} ></FilterForm>
    </div>
  )
}

const mapStateToProps = (state) => ({
  state,
})

const mapDispatchToProps = (dispatch) => ({
  getAllRecipes: (recipeInfo) => dispatch(getRecipesThunk(recipeInfo)),
})
export default connect(mapStateToProps, mapDispatchToProps)(AllRecipes)
