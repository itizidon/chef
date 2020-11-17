import axios from 'axios'
const GET_RECIPES = 'GET_RECIPES'

const getRecipes = (userInfo) => ({
  type: GET_RECIPES,
  userInfo,
})

export const getRecipesThunk = (recipeInfo) => {
  return async (dispatch) => {
    const { data } = await axios.post('http://localhost:8080/getRecipes', recipeInfo)
    dispatch(getRecipes(data))
    console.log(data)
  }
}

export default (state = {}, action) => {
  switch (action.type) {
    case GET_RECIPES:
      return {
        result: action.payload,
      }
    default:
      return state
  }
}
