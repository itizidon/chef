import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import { getRecipesThunk } from '../reducers/recipes'

const AllRecipes = (props) => {
  useEffect(() => {
    console.log(props, 'this is props')
    props.getAllRecipes({RecipeKey: "get all"})
  }, [])
  return <div>All Recipes</div>
}

const mapStateToProps = (state) => ({
  state,
})

const mapDispatchToProps = (dispatch) => ({
  getAllRecipes: (recipeInfo) => dispatch(getRecipesThunk(recipeInfo)),
})
export default connect(mapStateToProps, mapDispatchToProps)(AllRecipes)
