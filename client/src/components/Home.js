import React from 'react';
import { connect } from 'react-redux'

const Home = (props)=>{
  console.log(props)
return(
  <div>yo</div>
)
}

const mapStateToProps = state => ({
  state
})


export default connect(mapStateToProps, null)(Home)
