import React from 'react';

const SearchButton = (props) => {
    return (
        <button
            onClick={props.handleClick}
        >
            <i class="fas fa-search"></i>
        </button>
    )
}
export default SearchButton;