import React from 'react';
import Button from '../atoms/button'
import Text from '../atoms/text'

let MainRight = (props) => {
    if (props.mainContentState === "home") {
        let mainCommentData = "";
        let replyCommentData = "";

        if (props.mainComments.length !== 0) {
            mainCommentData = props.mainComments.map(function(data) {
                return (
                        <li
                            key={data.commentId}
                            data-classid={data.classId}
                            data-commentid={data.commentId}
                            onClick={props.handleGetReplyComment}
                        >
                            <div class="studentId">
                                <i class="fas fa-user-circle"></i>
                                {data.studentId}
                            </div>
                            <div class="mainComment">
                                {data.comment}
                            </div>
                        </li>
                )
            });
        }

        if (props.replyComments.length !== 0) {
            replyCommentData = props.replyComments.map(function(data) {
                return (
                    <li
                        // key={data.commentReplyId}
                        data-classid={data.classId}
                        data-commentid={data.commentId}
                        data-commentreplyid={data.commentReplyId}
                    >
                        <div class="studentId">
                            <i class="fas fa-user-circle"></i>
                            {data.studentId}
                        </div>
                        <div class="replyComment">
                            {data.comment}
                        </div>
                    </li>
                )
            });
        }

        if (props.mainComment.flg === true && props.replyComment.flg === true) {
            return (
                <div id="mainRight">
                    <div class="comment">
                        <Text placeholder="コメントする" handleChange={props.handleChangeMainComment}/>
                        <Button handleClick={props.handlePostMainComment}/>
                        <ul>
                            {mainCommentData}
                        </ul>
                    </div>
                    <div class="replyComment">
                        <Text placeholder="返信する" handleChange={props.handleChangeReplyComment}/>
                        <Button handleClick={props.handlePostReplyComment}/>
                        <ul>
                            {replyCommentData}
                        </ul>
                    </div>
                </div>
            )
        } else if (props.mainComment.flg === true && props.replyComment.flg === false) {
            return (
                <div id="mainRight">
                    <div class="comment">
                        <Text placeholder="コメントする" handleChange={props.handleChangeMainComment}/>
                        <Button handleClick={props.handlePostMainComment}/>
                        <ul>
                            {mainCommentData}
                        </ul>
                    </div>
                </div>
            )
        } else if (props.mainComment.flg === false && props.replyComment.flg === false) {
            return (
                <div id="mainRight">
                </div>
            )
        }
       
    } else if (props.mainContentState === "questionBoard") {
        return (
            <div>
                <Text placeholder="授業名" handleChange={props.handleChangeQuestionBoardClassId}/>
                <Text placeholder="質問" handleChange={props.handleChangeQuestionBoardQuestion}/>
                <Button handleClick={props.handlePostQuestionBoard}/>
            </div>
        )
    } else {
        return <div></div>
    }
}
export default MainRight;