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
                flg: false,
                classId: "",
                comment: ""
            },
            replyComment: {
                flg: false,
                classId: "",
                mainCommentId: null,
                comment: "",
            },
            questionBoards: [],
            questionBoard: {
                classId: "",
                question: ""
            },
            questionBoardComment: {
                classId: "",
                questionBoardId: "",
                comment: ""
            },
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
        this.handlePostReplyComment = this.handlePostReplyComment.bind(this);
        
        this.handleGetQuestionBoard = this. handleGetQuestionBoard.bind(this);
        this.handleChangeQuestionBoardQuestion = this.handleChangeQuestionBoardQuestion.bind(this);
        this.handleChangeQuestionBoardClassId = this.handleChangeQuestionBoardClassId.bind(this);
        this.handlePostQuestionBoard = this.handlePostQuestionBoard.bind(this);


        this.handleClass = this.handleClass.bind(this);
        // this.render = this.render.bind(this);
    }
    // sideBar----------------------------------------------------------

    handleHome() {
        this.setState({
            mainContentState: "home",
            classesData: [],
            mainComment: {
                flg: false,
                classId: "",
                comment: ""
            },
            replyComment: {
                flg: false,
                classId: "",
                mainCommentId: null,
                comment: ""
            }
        });
        this.handleClass();
    }

    handleQuestionBoard() {
        this.setState({
            mainContentState: "questionBoard"
        });
        this.handleGetQuestionBoard();
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

                    if (commentInfo.body !== null) {
                        this.setState({
                            mainComments: commentInfo.body.slice(0),
                            mainComment: {
                                flg: true,
                                classId: classId,
                                comment: ""
                            }
                        });
                    }
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
                flg: this.state.mainComment.flg,
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
                        let tmpArr = this.state.mainComments;
                        tmpArr.unshift(comment.body[0]);

                        this.setState({

                            // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                            mainComments: tmpArr.slice(0),
                            // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
                            mainComment: {
                                flg: this.state.mainComment.flg,
                                classId: this.state.mainComment.classId,
                                // mainCommentId: this.state.mainComment.mainCommentId,
                                comment: ""
                            }
                        });
                    }
                }
            }.bind(this);
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


        console.log(classId, commentId);

        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let replyComments = JSON.parse(httpRequest.responseText);
    
                    if (replyComments.body.length !== 0) {
                        this.setState({
                            replyComments: replyComments.body.slice(0),
                            replyComment: {
                                flg: true,
                                classId: classId,
                                mainCommentId: commentId,
                                comment: ""
                            }
                        });
                    } 
                    // else {
                    //     this.setState({
                    //         replyComment: {
                    //             classId: classId,
                    //             mainCommentId: commentId,
                    //             comment: ""
                    //         }
                    //     })
                    // }
                }
            }
        }.bind(this);
        httpRequest.open('GET', `http://localhost/comments/reply/?classId=${classId}&commentId=${commentId}`, true);
        httpRequest.send();
    }

    handleChangeReplyComment(event) {
        this.setState({
            replyComment: {
                flg: this.state.replyComment.flg,
                classId: this.state.replyComment.classId,
                mainCommentId: this.state.replyComment.mainCommentId,
                comment: event.target.value
            }
        })
    }

    handlePostReplyComment(event) {
        let httpRequest = new XMLHttpRequest();

        if (this.state.loginForm.flg) {
            httpRequest.onreadystatechange = function() {
                if (httpRequest.readyState === XMLHttpRequest.DONE) {
                    if (httpRequest.status === 200) {
                        let comment = JSON.parse(httpRequest.responseText);
                        let tmpArr = this.state.replyComments;
                        tmpArr.unshift(comment.body[0]);

                        this.setState({
                            replyComments: tmpArr.slice(0),
                            replyComment: {
                                flg: this.state.replyComment.flg,
                                classId: this.state.replyComment.classId,
                                mainCommentId: this.state.replyComment.mainCommentId,
                                comment: ""
                            }
                        });
                    }
                }
            }.bind(this)
            httpRequest.open('POST', 'http://localhost/comments/reply/', true);
            httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
            httpRequest.send(`classId=${this.state.replyComment.classId}&studentId=${this.state.studentData.studentId}&comment=${this.state.replyComment.comment}&commentId=${this.state.replyComment.mainCommentId}`);
        } else {
            alert("ログインをしてください");
        }
    }
    // questionBoard------------------------------------------------------------------    

    handleGetQuestionBoard() {
        let httpRequest = new XMLHttpRequest();
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let questionBoardInfo = JSON.parse(httpRequest.responseText);
                    console.log(questionBoardInfo);

                    this.setState({
                        questionBoards: questionBoardInfo.body.slice(0)
                    })
                }
            }
        }.bind(this)
        httpRequest.open('GET', 'http://localhost/questionBoards/', true);
        httpRequest.send();
    }

    handleChangeQuestionBoardQuestion(event) {
        let question = event.target.value;

        console.log(question);

        this.setState({
            questionBoard: {
                classId: this.state.questionBoard.classId,
                question: question
            }
        });
    }

    handleChangeQuestionBoardClassId(event) {
        let httpRequest = new XMLHttpRequest();
        let searchCondition = event.target.value;
        console.log(searchCondition);
    
        if (!httpRequest) {
            return false;
        }
        httpRequest.onreadystatechange = function () {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let classInfo = JSON.parse(httpRequest.responseText);
                    if (classInfo.body !== null) {

                        this.setState({
                            questionBoard: {
                                classId: classInfo.body[0].classId,
                                question: this.state.questionBoard.question
                            }
                        });
                    } 
                }
            }
        }.bind(this);

        httpRequest.open('GET', 'http://localhost/classes/?className='+searchCondition, true);
        httpRequest.send();
    }

    handlePostQuestionBoard() {
        let classId = this.state.questionBoard.classId;
        let question = this.state.questionBoard.question;
        let studentId = this.state.studentData.studentId;

        if (this.state.studentData.flg) {
            if (question !== "" && classId !== "") {
                let httpRequest = new XMLHttpRequest();
                httpRequest.onreadystatechange = function() {
                    if (httpRequest.readyState === XMLHttpRequest.DONE) {
                        if (httpRequest.status === 200) {
                            let questionBoardInfo = JSON.parse(httpRequest.responseText);
                            // questionBoardInfo.body.pop();
                            console.log(questionBoardInfo);
                        }
                    }
                }
                httpRequest.open('POST', `http://localhost/questionBoards/${classId}/2020/${studentId}/`, true);
                httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
                httpRequest.send('question='+question);
            } else {
                alert(`classid=${classId}`);
            }
        } else {
            alert("ログインをしてください");
        }
    }

    // handlePostQuestionBoardReply() {

    // }
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
                        handleGetMainComment={this.handleGetMainComment}
                        handleGetReplyComment={this.handleGetReplyComment}
                        
                        handleChangeMainComment={this.handleChangeMainComment}
                        handleChangeReplyComment={this.handleChangeReplyComment}
                        
                        handlePostMainComment={this.handlePostMainComment}
                        handlePostReplyComment={this.handlePostReplyComment}

                        handleChangeQuestionBoardQuestion={this.handleChangeQuestionBoardQuestion}
                        handleChangeQuestionBoardClassId={this.handleChangeQuestionBoardClassId}
                        handlePostQuestionBoard={this.handlePostQuestionBoard}
                        
                        mainContentState={this.state.mainContentState}
                        classesData={this.state.classesData}
                        mainComments={this.state.mainComments}
                        replyComments={this.state.replyComments}
                        questionBoards={this.state.questionBoards}
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
                        handleGetMainComment={this.handleGetMainComment}
                        handleGetReplyComment={this.handleGetReplyComment}
                        
                        handleChangeMainComment={this.handleChangeMainComment}
                        handleChangeReplyComment={this.handleChangeReplyComment}
                        
                        handlePostMainComment={this.handlePostMainComment}
                        handlePostReplyComment={this.handlePostReplyComment}

                        handleChangeQuestionBoardQuestion={this.handleChangeQuestionBoardQuestion}
                        handleChangeQuestionBoardClassId={this.handleChangeQuestionBoardClassId}
                        handlePostQuestionBoard={this.handlePostQuestionBoard}
                        
                        mainContentState={this.state.mainContentState}
                        classesData={this.state.classesData}
                        
                        mainComments={this.state.mainComments}
                        replyComments={this.state.replyComments}
                        
                        questionBoards={this.state.questionBoards}
                        mainComment={this.state.mainComment}
                        replyComment={this.state.replyComment}
                    />
                    
                </div>
            )
        }
    };
}


export default App;
