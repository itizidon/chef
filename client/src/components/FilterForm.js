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


  console.log(checkedEthnicity)
  return (
    <form onSubmit={(event) => updateRecipes(filter)}>
      {tags[0] ? (
        <div>
          <h6>Ethnicity</h6>
          {tags[0].ethnicity.map((cur, inx) => {
            return (
              <label key={inx}>
                {cur}
                <input
                  name="Ethnicity"
                  onChange={() => {
                    if (checkedEthnicity[cur]) {
                      setCheckedEthnicity((ethnicities) => {
                        const cloneEthnicities = {...ethnicities}
                        delete cloneEthnicities[cur]
                        return cloneEthnicities
                      })
                    }
                    else{
                      console.log('this is hit')
                      setCheckedEthnicity((ethnicities)=>{
                        const cloneEthnicities={...ethnicities}
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
          <h6>Recipe</h6>
          {tags[0].recipename.map((cur, inx) => {
            return (
              <label key={inx}>
                {cur}
                <input
                  name="Recipename"
                  onChange={() => setCheckedRecipename(inx)}
                  type="checkbox"
                />
              </label>
            )
          })}
          <h6>Time</h6>
          {tags[0].time.map((cur, inx) => {
            return (
              <label key={inx}>
                {cur}
                <input
                  name="time"
                  onChange={() => setCheckedTime(inx)}
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
