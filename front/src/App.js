import React, { Component } from 'react';
import Header from './presentational/header'
// import MainContent from './presentational/mainContent';
import SideBar from './presentational/sideBar';
import NavIcon from './molecules/navIcon';
import LoginForm from './presentational/loginForm'

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            mainContentState: "home",
            searchText: "",
            homeData: {},
            questionBoardData: {},
            studentData: {},
            loginForm: {
                flg: false,
                studentId: "",
                password: "",
            }
        }
        this.handleHome = this.handleHome.bind(this);
        this.handleQuestionBoard = this.handleQuestionBoard.bind(this);
        this.handleStudent = this.handleStudent.bind(this);
        this.handleHeaderTextChange = this.handleHeaderTextChange.bind(this);
        this.handleHeaderButtonClick = this.handleHeaderButtonClick.bind(this);
        this.handleStudentIdChange = this.handleStudentIdChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleLoginClick = this.handleLoginClick.bind(this);
    }
    // sideBar----------------------------------------------------------

    handleHome() {
        this.setState({
            mainContentState: "home"
        })
    }

    handleQuestionBoard() {
        this.setState({
            mainContentState: "questionBoard"
        })
    }

    handleStudent() {
        this.setState({
            mainContentState: "student"
        })
    }
    
    // loginForm--------------------------------------------------------
    handleStudentIdChange(event) {
        console.log(this.state.loginForm.studentId)
        this.setState({
            loginForm: {
                studentId: event.target.value,
                password: this.state.loginForm.password
            }
        });
    }
    
    handlePasswordChange(event) {
        console.log(this.state.loginForm.password)
        this.setState({
            loginForm: {
                studentId: this.state.loginForm.studentId,
                password: event.target.value
            }
        });
    }

    handleLoginClick() {
        console.log(`stuid=${this.state.loginForm.studentId} pass=${this.state.loginForm.password}`);
        
        let studentId = this.state.loginForm.studentId;
        let password = this.state.loginForm.password;
        let httpRequest = new XMLHttpRequest();
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let ret = JSON.parse(httpRequest.responseText);
                    console.log(ret);
                }
            }
        }
        httpRequest.open('POST', 'http://172.28.0.3:8080/login', true);
        httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        httpRequest.send(`studentId=${studentId}&password=${password}`);
    }
    // header-------------------------------------------------------------

    handleHeaderTextChange(event) {
        this.setState({
            searchText: event.target.value
        })
    }
    handleHeaderButtonClick() {
        // apiサーバからデータを取得する処理

        if (this.mainContentState === "home") {
            console.log("send packet to get class information");
        } else if (this.mainContentState === "questionBoard") {
            console.log("send packet to get question board information");
        } else if (this.mainContentState === "student") {
            console.log("send packet to get student information");
        }
    }
    // --------------------------------------------------------------------

    render() {
        return (
            <div>
                 <input id="menu" type="checkbox"/>
                <label for="menu" class="back"></label>
                <NavIcon/>
                <SideBar
                    handleHome={this.handleHome}
                    handleQuestionBoard={this.handleQuestionBoard}
                    handleStudent={this.handleStudent}
                />
                <input type="checkbox" id="login"/>
                <LoginForm
                    handleStudentIdChange={this.handleStudentIdChange}
                    handlePasswordChange={this.handlePasswordChange}
                    handleClick={this.handleLoginClick}
                />
                <Header
                    handleChange={this.handleHeaderTextChange}
                    handleClick={this.handleHeaderButtonClick}
                />
                
            </div>
        )
    };
}


export default App;
