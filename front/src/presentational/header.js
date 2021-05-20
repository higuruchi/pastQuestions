import React, { Component } from 'react';
import Text from '../atoms/text'
import SearchButton from '../atoms/searchButton'


class Header extends Component {
    render() {
        return (
            <header id="header">
                <div class="search">
                    <Text handleChange={this.props.handleChange}/>
                    <SearchButton
                    handleClick={this.props.handleClick}
                    />
                <label for="login">
                    <i class="fas fa-sign-in-alt"></i>
                </label>
                </div>
            </header>
        )
    };
}
export default Header;