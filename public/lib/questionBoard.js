let leftWrapper = document.getElementById('leftWrapper');
let rightWrapper = document.getElementById('rightWrapper');
let studentId = document.getElementById('studentId').innerText;

function QuestionBoard(classId) {
    removeMainContent();
    GetQuestionBoard();
}

function GetQuestionBoard() {
    let httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                let questionBoardInfo = JSON.parse(httpRequest.responseText);
                SetQuestionBoard(questionBoardInfo);
            }
        }
    }
    httpRequest.open('GET', '/questionBoards/', true);
    httpRequest.send();
}

function SetQuestionBoard(questionBoardInfo) {
    let questionBoards = questionBoardInfo.body;

    if (questionBoards !== null) {
        questionBoards.map((questionBoard)=>{
            let questionBoardWrapper = document.createElement('details');
            let questionBoardWrapperSummary = document.createElement('summary');
            let questionBoardWrapperClassId = document.createElement('div');
            let questionBoardWrapperYear = document.createElement('div');
            let questionBoardWrapperStudentId = document.createElement('div');
            let questionBoardWrapperQuestion = document.createElement('div');
            let questionBoardWrapperQuestionReply = document.createElement('div');
            let questionBoardReplyText = document.createElement('input');
            let questionBoardReplyButton = document.createElement('input');
            let questionBoardReplyForm = document.createElement('form');

            questionBoardWrapperClassId.innerText = `classId : ${questionBoard.classId}`;
            questionBoardWrapperYear.innerText = `year : ${questionBoard.year}`;
            questionBoardWrapperStudentId.innerText = `studentId : ${questionBoard.studentId}`;
            questionBoardWrapperQuestion.innerText = `question : ${questionBoard.question}`;

            questionBoardWrapperQuestionReply.setAttribute('id', `qb${questionBoard.classId}-${questionBoard.questionBoardId}`);

            questionBoardReplyText.setAttribute('type', 'text');
            questionBoardReplyButton.setAttribute('type', 'button');
            questionBoardReplyButton.setAttribute('value', 'send');
            questionBoardReplyButton.setAttribute('data-classid', questionBoard.classId);
            questionBoardReplyButton.setAttribute('data-studentid', questionBoard.studentId);
            questionBoardReplyButton.setAttribute('data-questionboardid', questionBoard.questionBoardId);
            questionBoardReplyButton.addEventListener('click', PostQuestionBoardReply);
            questionBoardReplyForm.appendChild(questionBoardReplyText);
            questionBoardReplyForm.appendChild(questionBoardReplyButton);

            if (questionBoard.questionBoardReply !== null) {
                questionBoard.questionBoardReply.map((questionBoardReply)=>{
                    let questionBoardReplytext = document.createElement('div');
                    questionBoardReplytext.innerText = questionBoardReply.reply;
                    questionBoardWrapperQuestionReply.appendChild(questionBoardReplytext);
                });
            }
            questionBoardWrapperSummary.appendChild(questionBoardWrapperClassId);
            questionBoardWrapperSummary.appendChild(questionBoardWrapperYear);
            questionBoardWrapperSummary.appendChild(questionBoardWrapperStudentId);
            questionBoardWrapperSummary.appendChild(questionBoardWrapperQuestion);
            questionBoardWrapper.appendChild(questionBoardWrapperQuestionReply);
            questionBoardWrapper.appendChild(questionBoardWrapperSummary);
            questionBoardWrapper.appendChild(questionBoardReplyForm);
            leftWrapper.appendChild(questionBoardWrapper);
        });

    }
}

function PostQuestionBoardReply(event) {
    let classId = event.target.getAttribute('data-classid');
    let studentId = event.target.getAttribute('data-studentid');
    let questionBoardId = event.target.getAttribute('data-questionboardid');
    let reply = event.target.previousElementSibling.value;
    event.target.previousElementSibling.value = '';

    console.log(reply);

    let httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                let questionBoardReplyInfo = JSON.parse(httpRequest.responseText);
                let questionBoardReply = questionBoardReplyInfo.body.pop();
                let newQuestionBoardReply = document.createElement('div');
                newQuestionBoardReply.innerText = questionBoardReply.reply;
                document.getElementById(`qb${questionBoardReply.classId}-${questionBoardReply.questionBoardId}`).appendChild(newQuestionBoardReply);
            }
        }
    }
    httpRequest.open('POST', `/questionBoardsReply/${classId}/${questionBoardId}/${studentId}/`, true);
    httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    httpRequest.send('reply='+reply);
}

function removeMainContent() {
    removeRightContent();
    removeleftContent();
}

function removeleftContent() {
    let leftWrapper = document.getElementById('leftWrapper');
    leftWrapper.innerHTML = '';
}
function removeRightContent() {
    let rightWrapper = document.getElementById('rightWrapper');
    rightWrapper.innerHTML = '';
}
export default QuestionBoard;