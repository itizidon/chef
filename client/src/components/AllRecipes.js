import React from 'react';
import { connect } from 'react-redux'

const AllRecipes = (props)=>{
  console.log(props)
return(
  <div>All Recipes</div>
)
}

const mapStateToProps = state => ({
  state
})


export default connect(mapStateToProps, null)(AllRecipes)
