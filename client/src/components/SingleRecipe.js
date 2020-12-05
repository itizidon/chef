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
          inputMode="numeric"
          pattern="[0-9\s]{13,19}"
          autoComplete="cc-number"
          maxLength="19"
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
    </div>
  )
}

export default withRouter(SingleRecipe)
