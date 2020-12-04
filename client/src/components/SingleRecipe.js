import React, { useEffect, useState } from 'react'
import { withRouter } from 'react-router'
import { Modal } from './Modals/RecipePayment'

const SingleRecipe = (props) => {
  const [show, setShow] = useState(false)
  console.log(props)

  const handleClose = () => {
    setShow(false)
  }

  return (
    <div>
      <Modal show={show} handleClose={handleClose}>
        <input
          type="tel"
          inputmode="numeric"
          pattern="[0-9\s]{13,19}"
          autocomplete="cc-number"
          maxlength="19"
          placeholder="xxxx xxxx xxxx xxxx"
        ></input>
      </Modal>
      <button
        type="button"
        onClick={() => {
          setShow(true)
        }}
      >
        Order Recipe
      </button>
      <p1>Reviews</p1>
    </div>
  )
}

export default withRouter(SingleRecipe)
