import React, { Component } from 'react';
import Header from './presentational/header'
import MainContent from './presentational/mainContent';
import SideBar from './presentational/sideBar';
import NavIcon from './molecules/navIcon';
import LoginForm from './presentational/loginForm';
// import {handleHome, handleQuestionBoard, handleStudent} from './lib/sideBar';


class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            mainContentState: "home",
            searchText: "",
            classesData: [],
            mainComments: [],
            replyComments: [],
            mainComment: {
                classId: "",
                comment: ""
            },
            replyComment: {
                classId: "",
                mainCommentId: null,
                comment: "",
            },
            questionBoardData: {},
            studentData: {
                flg: false,
                studentId: "",
                studentName: "",
            },
            loginForm: {
                flg: false,
                studentId: "",
                password: "",
            },
            hoge: "hoge"
        }
        
        this.handleHome = this.handleHome.bind(this);
        this.handleQuestionBoard = this.handleQuestionBoard.bind(this);
        this.handleStudent = this.handleStudent.bind(this);
        
        this.handleHeaderTextChange = this.handleHeaderTextChange.bind(this);
        this.handleHeaderButtonClick = this.handleHeaderButtonClick.bind(this);

        this.handleStudentIdChange = this.handleStudentIdChange.bind(this);
        this.handlePasswordChange = this.handlePasswordChange.bind(this);
        this.handleLoginClick = this.handleLoginClick.bind(this);

        this.handleGetMainComment = this.handleGetMainComment.bind(this);
        this.handleGetReplyComment = this.handleGetReplyComment.bind(this);

        this.handleChangeMainComment = this.handleChangeMainComment.bind(this);
        this.handleChangeReplyComment = this.handleChangeReplyComment.bind(this);

        this.handlePostMainComment = this.handlePostMainComment.bind(this);

        this.handleClass = this.handleClass.bind(this);
        // this.render = this.render.bind(this);
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
        this.setState({
            loginForm: {
                flg: this.state.loginForm.flg,
                studentId: event.target.value,
                password: this.state.loginForm.password
            }
        });
    }
    
    handlePasswordChange(event) {
        this.setState({
            loginForm: {
                flg: this.state.loginForm.flg,
                studentId: this.state.loginForm.studentId,
                password: event.target.value
            }
        });
    }
    handleLoginClick() {
        
        let studentId = this.state.loginForm.studentId;
        let password = this.state.loginForm.password;
        let httpRequest = new XMLHttpRequest();

        httpRequest.onreadystatechange = function () {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let ret = JSON.parse(httpRequest.responseText);
                    console.log(ret);
                    this.setState({
                        loginForm: {
                            flg: true,
                            studentId: this.state.loginForm.studentId,
                            password: this.state.loginForm.password
                        },
                        studentData: {
                            flg: true,
                            studentId: ret.body.studentId,
                            studentName: ret.body.studentName
                        }
                    });
                }
            }
        }.bind(this);

        // 上のbind(this) なぜできたのかが分からない
        // bindについて勉強する必要がある

        httpRequest.open('POST', 'http://localhost/login', true);
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
    // class-------------------------------------------------------------------

    handleClass() {
        let httpRequest = new XMLHttpRequest();
        let classInfo;
        
        if (!httpRequest) {
                return false;
        }
        httpRequest.onreadystatechange = function () {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    classInfo = JSON.parse(httpRequest.responseText);
                    if (classInfo.body.length !== 0) {
                        this.setState({
                            classesData: classInfo.body.slice(0)
                        });
                    }
                }
            }  
        }.bind(this);

        httpRequest.open('GET', 'http://localhost/classes/', true);
        httpRequest.send();
    }
    // mainComment------------------------------------------------------------

    handleGetMainComment(event) {
        let httpRequest = new XMLHttpRequest();
        let classId = event.currentTarget.dataset.classid
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let commentInfo = JSON.parse(httpRequest.responseText);
    
                    if (commentInfo.body.length !== 0) {
                        this.setState({
                            mainComments: commentInfo.body.slice(0),
                            mainComment: {
                                classId: classId,
                                comment: ""
                            }
                        });
                    }
                    // console.log(commentInfo.body.slice(0));
                }
            }
        }.bind(this);

        console.log(event.currentTarget.dataset.classid);
    
        httpRequest.open('GET', `http://localhost/comments/main/?classId=${classId}&commentId=0`, true);
        httpRequest.send();
    }

    handleChangeMainComment(event) {
        this.setState({
            mainComment: {
                classId: this.state.mainComment.classId,
                comment: event.target.value
            }
        })
    }

    handlePostMainComment() {
        if (this.state.loginForm.flg) {
            let httpRequest = new XMLHttpRequest();
            httpRequest.onreadystatechange = function() {
                if (httpRequest.readyState === XMLHttpRequest.DONE) {
                    if (httpRequest.status === 200) {
                        let comment = JSON.parse(httpRequest.responseText);
                        console.log(comment);
                    }
                }
            }
            httpRequest.open('POST', 'http://localhost/comments/main/', true);
            httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
            httpRequest.send(`classId=${this.state.mainComment.classId}&studentId=${this.state.studentData.studentId}&comment=${this.state.mainComment.comment}`);
        } else {
            alert("ログインをしてください");
        }
    }
    // replyComment-----------------------------------------------------------
    handleGetReplyComment(event) {
        let httpRequest = new XMLHttpRequest();
        let classId = event.currentTarget.dataset.classid;
        let commentId = event.currentTarget.dataset.commentid;
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let replyComments = JSON.parse(httpRequest.responseText);
                    
                    if (replyComments.body.length !== 0) {
                        this.setState({
                            replyComments: replyComments.body.slice(0),
                            replyComment: {
                                classId: classId,
                                commentId: commentId,
                                comment: ""
                            }
                        })
                    }
                }
            }
        }.bind(this);
        httpRequest.open('GET', `http://localhost/comments/reply/?classId=${classId}&commentId=${commentId}`, true);
        httpRequest.send();
    }

    handleChangeReplyComment(event) {
        this.setState({
            replyComment: {
                classId: this.state.replyComment.classId,
                commentId: this.state.replyComment.commentId,
                comment: event.target.value
            }
        })
    }

    handlePostReplyComment(event) {

    }
    // -----------------------------------------------------------------------    

    // -----------------------------------------------------------------------    
    // -----------------------------------------------------------------------    
    // -----------------------------------------------------------------------
    // -----------------------------------------------------------------------
    componentWillMount() {
        this.handleClass();
    }
    
    render() {
        // let returnJSX;

        if (this.state.loginForm.flg) {
            return  (
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
                    <Header
                        handleChange={this.handleHeaderTextChange}
                        handleClick={this.handleHeaderButtonClick}
                        studentData={this.state.studentData}
                    />
                    <MainContent 
                       mainContentState={this.state.mainContentState}
                       classesData={this.state.classesData}
                       handleGetMainComment={this.handleGetMainComment}
                       handleGetReplyComment={this.handleGetReplyComment}
                       handleChangeMainComment={this.handleChangeMainComment}
                       handleChangeReplyComment={this.handleChangeReplyComment}
                       handlePostMainComment={this.handlePostMainComment}
                       mainComments={this.state.mainComments}
                       replyComments={this.state.replyComments}
                       mainComment={this.state.mainComment}
                       replyComment={this.state.replyComment}
                    />
                </div>
            )
        } else {
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
                        studentData={this.state.studentData}
                    />
                    <MainContent 
                        mainContentState={this.state.mainContentState}
                        classesData={this.state.classesData}
                        handleGetMainComment={this.handleGetMainComment}
                        handleGetReplyComment={this.handleGetReplyComment}
                        handleChangeMainComment={this.handleChangeMainComment}
                        handleChangeReplyComment={this.handleChangeReplyComment}
                        handlePostMainComment={this.handlePostMainComment}
                        mainComments={this.state.mainComments}
                        replyComments={this.state.replyComments}
                        mainComment={this.state.mainComment}
                        replyComment={this.state.replyComment}
                    />
                    
                </div>
            )
        }
    };
}


export default App;
