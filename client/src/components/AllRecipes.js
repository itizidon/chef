import React, { useEffect, useState } from 'react'
import { connect } from 'react-redux'
import { getRecipesThunk } from '../reducers/recipes'
import FilterForm from './FilterForm'
import axios from 'axios'

const AllRecipes = (props) => {
  const [recipes, setRecipes] = useState({ RecipeKey: 'get all' })
  const [allData, setAllData] = useState([])

  const updateRecipes = (filters) => {
    setRecipes(filters)
  }

  useEffect(() => {
    async function fetchingData() {
      const { data } = await axios.post(
        'http://localhost:8080/getRecipes',
        recipes
      )
      setAllData(data)
    }
    fetchingData()

  }, [])
  console.log(allData)
  return (
    <div>
      <div>All Recipes</div>
      <FilterForm updateRecipes={updateRecipes}></FilterForm>
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
