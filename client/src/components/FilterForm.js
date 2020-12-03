import React, { useEffect, useState } from 'react'
import axios from 'axios'

const FilterForm = ({ updateRecipes }) => {
  const [filter, setFilter] = useState({ RecipeKey: 'get all' })

  const [checkedEthnicity, setCheckedEthnicity] = useState({})
  const [checkedRecipename, setCheckedRecipename] = useState({})
  const [checkedTime, setCheckedTime] = useState({})

  const [tags, setTags] = useState([])

  useEffect(() => {
    async function fetchingTags() {
      const { data } = await axios.get('http://localhost:8080/getTags')
      setTags(data)
    }
    fetchingTags()
  }, [])

  const submitHandler = (e) => {
    e.preventDefault()

    let recipes = Object.keys(checkedRecipename)
    let ethnicities = Object.keys(checkedEthnicity)
    let times = Object.keys(checkedTime)
    let finalQuery = {RecipeKey: 'get all'}

    if(recipes.length!==0){
      finalQuery.Recipename = recipes
      delete finalQuery.RecipeKey
    }
    if(ethnicities.length!==0){
      finalQuery.Ethnicity = ethnicities
      delete finalQuery.RecipeKey
    }
    if(times.length!==0){
      finalQuery.Time = times
      delete finalQuery.RecipeKey
    }

    updateRecipes(finalQuery)
  }

  return (
    <form onSubmit={(e) => submitHandler(e)}>
      {tags[0] ? (
        <div>
          <h6 className="filter">Ethnicity</h6>
          {tags[0].ethnicity.map((cur, inx) => {
            return (
              <label key={inx}>
                {cur}
                <input
                  name="Ethnicity"
                  onChange={() => {
                    if (checkedEthnicity[cur]) {
                      setCheckedEthnicity((ethnicities) => {
                        const cloneEthnicities = { ...ethnicities }
                        delete cloneEthnicities[cur]
                        return cloneEthnicities
                      })
                    } else {
                      setCheckedEthnicity((ethnicities) => {
                        const cloneEthnicities = { ...ethnicities }
                        cloneEthnicities[cur] = cur
                        return cloneEthnicities
                      })
                    }
                  }}
                  type="checkbox"
                />
              </label>
            )
          })}
          <h6 className="filter">Recipe</h6>
          {tags[0].recipename.map((cur, inx) => {
            return (
              <label key={inx}>
                {cur}
                <input
                  name="Recipename"
                  onChange={() => {
                    if (checkedRecipename[cur]) {
                      setCheckedRecipename((recipename) => {
                        const cloneRecipenames = { ...recipename }
                        delete cloneRecipenames[cur]
                        return cloneRecipenames
                      })
                    } else {
                      console.log('this is hit')
                      setCheckedRecipename((recipename) => {
                        const cloneRecipename = { ...recipename }
                        cloneRecipename[cur] = cur
                        return cloneRecipename
                      })
                    }
                  }}
                  type="checkbox"
                />
              </label>
            )
          })}
          <h6 className="filter">Time</h6>
          {tags[0].time.map((cur, inx) => {
            return (
              <label key={inx}>
                {cur}
                <input
                  name="time"
                  onChange={() => {
                    if (checkedTime[cur]) {
                      setCheckedTime((times) => {
                        const cloneTime = { ...times }
                        delete cloneTime[cur]
                        return cloneTime
                      })
                    } else {
                      console.log('this is hit')
                      setCheckedTime((times) => {
                        const cloneTime = { ...times }
                        cloneTime[cur] = cur
                        return cloneTime
                      })
                    }
                  }}
                  type="checkbox"
                />
              </label>
            )
          })}
        </div>
      ) : null}
      <input type="submit" value="Submit" />
    </form>
  )
}

export default FilterForm
