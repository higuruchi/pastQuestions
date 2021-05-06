import GetComments from './comment.js'
import GetPastQuestion from './pastQuestion.js'

function GetPastQuestionAndComment(event) {
    // let commentsWrapper = document.createElement('div');
    // let pastQuestionWrapper = document.createElement('div');
    const classId = event.target.getAttribute('data-classid');
    let main = document.getElementById('main');

    // commentsWrapper.setAttribute('id', 'commentsWrapper');
    // pastQuestionWrapper.setAttribute('id', 'pastQuestionWrapper');
    removeleftContent();
    removeRightContent();

    // main.appendChild(pastQuestionWrapper);
    // main.appendChild(commentsWrapper);
    GetPastQuestion(classId);
    GetComments(classId);

    // let main = document.getElementById('main');

    // main.appendChild(pastQuestionWrapper);
    // main.appendChild(commentsWrapper);

    
    
    // function GetPastQuestion() {
    //     httpRequest = new XMLHttpRequest();
    //     httpRequest.onreadystatechange = function() {
    //         if (httpRequest.readyState === XMLHttpRequest.DONE) {
    //             if (httpRequest.status === 200) {
    //                 pastQuestionInfo = JSON.parse(httpRequest.responseText);
    //                 SetPastQuestion();
    //             }
    //         }
    //     }
    //     httpRequest.open('GET', '/pastQuestion/?className='+searchCondition, true);
    //     httpRequest.send();
    //     removeMainContent();
    // }

    // function SetPastQuestion() {

    // }

    // function GetComments() {
    //     let httpRequest = new XMLHttpRequest();
    //     httpRequest.onreadystatechange = function() {
    //         if (httpRequest.readyState === XMLHttpRequest.DONE) {
    //             if (httpRequest.status === 200) {
    //                 let commentInfo = JSON.parse(httpRequest.responseText);
    //                 SetComments(commentInfo);
    //             }
    //         }
    //     }
    //     httpRequest.open('GET', '/comments/main/?classId='+event.target.getAttribute('data-classid'), true);
    //     httpRequest.send();
    //     removeMainContent();
    // }

    function removeleftContent() {
        let leftWrapper = document.getElementById('leftWrapper');
        leftWrapper.innerHTML = '';
    }
    function removeRightContent() {
        let rightWrapper = document.getElementById('rightWrapper');
        rightWrapper.innerHTML = '';
    }

}
export default GetPastQuestionAndComment