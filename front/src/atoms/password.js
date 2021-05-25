import React from 'react'

let Password = (props) => {
    return (
        <input type="password"
        placeholder="password"
        onChange={props.handleChange}
        />
    )
}
export default Password;