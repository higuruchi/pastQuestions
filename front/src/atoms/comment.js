import React, { Component } from 'react';
import List from './list'
import Text from './text'
import Button from './button'

class Comment extends Component {
    constructor(props) {
        super(props);
        this.state = {
            classId: props.comment.classId,
            commentId: props.comment.commentId,
            mainComment: props.comment.mainComment,
            goodFlg: props.comment.goodFlg,
            badFlg: props.comment.BadFlg,
            studentId: props.comment.studentId,
            replyComments: [],
            replyComment: ""
        }
        this.getReplyComments = this.getReplyComments.bind(this);
        this.sendReplyComment = this.sendReplyComment.bind(this);
        this.checkText = this.checkText.bind(this);
    }

    getReplyComments() {
        // 返信コメントをサーバから取得する処理
        // httpRequest = new XMLHttpRequest();

        // httpRequest.onreadystatechange = function() {
        //     if(httpRequest.readyState === XMLHttpRequest.DONE) {
        //         if (httpRequest.status === 200) {

        //         }
        //     }
        // }

        // this.setState({
        //     replyComments: 
        // });
    }
    
    sendReplyComment() {

    }

    handleChangeText(event) {
        this.setState({
            replyComment: event.target.value
        })
    }
    render() {
        return (
            <div>
                <details onClick={this.getReplyComments}>
                    <summary>
                        <hidden>{this.state.classId}</hidden>
                        <hidden>{this.state.commentId}</hidden>
                        <div>{this.state.studentId}</div>
                        <div>{this.state.comment}</div>
                        <div><spn>👍{this.state.goodFlg}</spn>
                        <spn>⤵{this.state.badFlg}</spn></div>
                    </summary>
                    <List
                    list={this.props.replyComment}
                    />
                    <Text placeholder="コメント" handleChange={this.handleChangeText}/>
                    <Button handleClick={this.sendReplyComment} value="送信"/>
                </details>

            </div>
        );
    }
}
export default Comment;