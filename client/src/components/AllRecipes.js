import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import { getRecipesThunk } from '../reducers/recipes'

const AllRecipes = (props) => {
  console.log(props)
  useEffect(() => {
    props.getAllRecipes()
  }, [])
  return <div>All Recipes</div>
}

const mapStateToProps = (state) => ({
  state,
})

const mapDispatchToProps = (dispatch) => ({
  getAllRecipes: () => dispatch(getRecipesThunk()),
})
export default connect(mapStateToProps, mapDispatchToProps)(AllRecipes)
