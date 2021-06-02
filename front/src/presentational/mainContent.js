import React from 'react';
import MainLeft from '../molecules/mainLeft';
import MainRight from '../molecules/mainRight';

let MainContent = (props) => {
    return (
        <div id="main">
            <MainLeft
                classesData={props.classesData}
                mainContentState={props.mainContentState}
                handleGetMainComment={props.handleGetMainComment}
                questionBoards={props.questionBoards}
            />
            <MainRight
                mainContentState={props.mainContentState}
                
                mainComments={props.mainComments}
                replyComments={props.replyComments}

                handleGetMainComment={props.handleGetMainComment}
                handleGetReplyComment={props.handleGetReplyComment}
                
                handleChangeMainComment={props.handleChangeMainComment}
                handleChangeReplyComment={props.handleChangeReplyComment}
                
                handlePostMainComment={props.handlePostMainComment}
                handlePostReplyComment={props.handlePostReplyComment}

                handleChangeQuestionBoardQuestion={props.handleChangeQuestionBoardQuestion}
                handleChangeQuestionBoardClassId={props.handleChangeQuestionBoardClassId}
                handlePostQuestionBoard={props.handlePostQuestionBoard}

                mainComment={props.mainComment}
                replyComment={props.replyComment}
            />
        </div>
    )

}

export default MainContent;