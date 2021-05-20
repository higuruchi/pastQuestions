import React, { Component } from 'react';
import Comment from '../atoms/comment';
import Text from '../atoms/text';
import Button from '../atoms/button';

class Comments extends Component {
    constructor (props) {
        super(props);
        this.state = {
            mainComment: "",
            mainComments: this.getMainComments()
        }
        this.sendMainComment = this.sendMainComment.bind(this);
        this.getMainComments = this.getMainComments.bind(this);
        this.setMainComments = this.setMainComments.bind(this);
        this.handleChangeText = this.handleChangeText.bind(this);
    }

    sendMainComment() {
        // コメントを送信する処理
    }

    getMainComments() {
        // コメントを取得する処理
    }

    setMainComments() {

    }

    handleChangeText(event) {
        // textに変化があるび、stateを変える
        this.setState({
            mainComment: event.target.value
        })
    }

    render() {
        let mainCommentLi = this.state.mainComments.map((comment) => {
            return <li><Comment
                        comment={comment}
                    /></li>
        })
        return(
            <div>
                <ul>]
                    {mainCommentLi}
                </ul>
                <Text handleChange={this.handleChange} placeholder="コメント"/>
                <Button handleClick={this.sendMainComment} value="送信"/>
            </div>
        );
    }
}
export default Comments;