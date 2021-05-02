import GetComments from './comment.js'

function GetPastQuestionAndComment(event) {
    let pastQuestioInfo;
    // let commentInfo;
    const classId = event.target.getAttribute('data-classid')
    // let pastQuestionWrapper = document.createElement('div').setAttribute('class', 'pastQuestionWrapper');
    // let commentsWrapper = document.createElement('div');
    // commentsWrapper.setAttribute('class', 'commentsWrapper');
    // document.getElementById('main').appendChild(pastQuestionWrapper);
    // document.getElementById('main').appendChild(commentsWrapper);

    // GetPastQuestion();
    removeMainContent();
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

    function removeMainContent() {
        let main = document.getElementById('main');
        main.innerHTML = '';
    }
}
export default GetPastQuestionAndComment