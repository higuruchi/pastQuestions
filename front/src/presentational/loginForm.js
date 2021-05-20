import React from 'react'
import Button from '../atoms/button'
import Text from '../atoms/text'
import Password from '../atoms/password'

let LoginForm = (props) => {
    return (
        <div class="loginForm">
            <Text placeholder="studentId" handleChange={props.handleStudentIdChange}/>
            <Password handleChange={props.handlePasswordChange}/>
            <Button handleClick={props.handleClick} value="sign in"/>
        </div>
    )
}

export default LoginForm;