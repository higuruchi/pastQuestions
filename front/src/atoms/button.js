import React from 'react';

const Button = (props) => {
    return (
        <button
            onClick={props.handleClick}
        >
            <i class="fas fa-paper-plane"></i>
        </button>
    )
}
export default Button;