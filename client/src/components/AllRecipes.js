import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import { getRecipesThunk } from '../reducers/recipes'

const AllRecipes = (props) => {
  useEffect(() => {
    console.log(props, 'this is props')
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
