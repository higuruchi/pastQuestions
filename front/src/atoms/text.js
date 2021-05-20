import React from 'react'

const Text = (props) => {
    return (
        <input
            type="text"
            placeholder={props.placeholder}
            onChange={props.handleChange}
        />
    )
}
export default Text;