import React from 'react';

const Button = (props) => {
    return (
        <button
            onClick={props.handleClick}
        >
            {props.value}
        </button>
    )
}
export default Button;