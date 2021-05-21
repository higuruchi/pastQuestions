import React, { Component } from 'react';
import Text from '../atoms/text'
import SearchButton from '../atoms/searchButton'


let Header = (props) => {
    if (props.studentData.flg) {
        return (
            <header id="header">
                <div class="search">
                    <Text handleChange={props.handleChange}/>
                    <SearchButton
                    handleClick={props.handleClick}
                    />
                <label>
                    <i class="fas fa-user-circle"></i>
                    <span>{props.studentData.studentName}</span>
                </label>
                </div>
            </header>
        );
    } else {
        return (
            <header id="header">
                <div class="search">
                    <Text handleChange={props.handleChange}/>
                    <SearchButton
                    handleClick={props.handleClick}
                    />
                <label for="login">
                    <i class="fas fa-sign-in-alt"></i>
                </label>
                </div>
            </header>
        );
    }
}
export default Header;