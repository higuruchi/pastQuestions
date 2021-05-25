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
            />
            <MainRight
                mainContentState={props.mainContentState}
                mainComments={props.mainComments}
                replyComments={props.replyComments}
                handleGetReplyComment={props.handleGetReplyComment}
                handleChangeMainComment={props.handleChangeMainComment}
                handleChangeReplyComment={props.handleChangeReplyComment}
                handlePostMainComment={props.handlePostMainComment}
                mainComment={props.mainComment}
                replyComment={props.replyComment}
            />
        </div>
    )

}

export default MainContent;