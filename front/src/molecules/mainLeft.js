import React from 'react';

let MainLeft = (props) => {
    if (props.mainContentState === "home") {
        let classData = "";

        classData = props.classesData.map(function(data) {
            return (
                    <li
                        key={data.classId}
                        data-classid={data.classId}
                        onClick={props.handleGetMainComment}
                    >
                        <i class="fas fa-book-reader"></i>
                        {data.className}
                    </li>
                    )
        });
        return (
            <div id="mainLeft">
                <ul class="classes">
                    {classData}
                </ul>
            </div>
        )
    } else if (props.mainContentState === "questionBoard") {
        let questionBoard = "";

        questionBoard = props.questionBoards.map(function(data) {
            let questionBoardReply = "";

            if (data.questionBoardReply !== null) {
                questionBoardReply = data.questionBoardReply.map(function(rep) {
                    return (
                        <li key={rep.questionBoardReplyId}>
                            <div>
                                <i class="fas fa-user-circle"></i>
                                {rep.studentId}
                            </div>
                            <div>
                                {rep.reply}
                            </div>
                        </li>
                    )
                });    
            }

            return (
                <details
                    key={data.questionBoardId}
                    data-classid={data.classId}
                    data-questionboardid={data.questionBoardId}
                >
                    <summary>
                        <i class="fas fa-question"></i>
                        <div class="className">{data.className}</div>
                        <div class="question">{data.question}</div>
                    </summary>
                    <ul class="questionBoardReply">{questionBoardReply}</ul>
                </details>
            )
        });

        return (
            <div id="mainLeft">
                <div class="questionBoard">
                    {questionBoard}
                </div>
            </div>
        )
    }
}

export default MainLeft;