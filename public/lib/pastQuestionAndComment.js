import GetComments from './comment.js'
import GetPastQuestion from './pastQuestion.js'

function GetPastQuestionAndComment(event) {
    const classId = event.target.getAttribute('data-classid');
    console.log(document.getElementById("main"))
    removeMainContent();
    
    GetPastQuestion(classId);
    GetComments(classId);
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
export default GetPastQuestionAndComment