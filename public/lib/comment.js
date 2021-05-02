let commentsWrapper = document.createElement('div');
commentsWrapper.setAttribute('class', 'commentsWrapper');
let main = document.getElementById('main');


function GetComments(classId) {
    let httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                let commentInfo = JSON.parse(httpRequest.responseText);
                console.log(commentInfo);
                SetComments(commentInfo, classId);
            }
        }
    }
    httpRequest.open('GET', '/comments/main/?classId='+classId+'&commentId=0', true);
    httpRequest.send();
    removeMainContent();
}


function SetComments(commentInfo, classId) {
    let comments = commentInfo.body;

    comments.map((comment)=>{SetComment(comment)});
    let text = document.createElement('input');
    let button = document.createElement('button');
    text.setAttribute('type', 'text');
    text.addEventListener('change', function() {
        button.setAttribute('data-text', text.value)
    });
    button.innerText = 'send';
    button.setAttribute('flg', 'main');
    button.setAttribute('data-classid', classId);
    button.addEventListener('click', PostComment);
    main.appendChild(commentsWrapper);
    main.appendChild(text);
    main.appendChild(button);
}

function SetComment(comment) {
        let details = document.createElement('details');
        let summary = document.createElement('summary');
        let divup = document.createElement('div');
        let divdown = document.createElement('div');
        let text = document.createElement('input');
        let button = document.createElement('button');
        let ul = document.createElement('ul');
        summary.setAttribute('data-classid', comment.classId);
        summary.setAttribute('data-commentid', comment.commentId);
        ul.setAttribute('id', 'commentId'+comment.commentId);
        text.setAttribute('type', 'text');
        text.addEventListener('change', function() {
            button.setAttribute('data-text', text.value)
        });
        button.innerText = 'send';
        button.setAttribute('flg', 'reply');
        button.setAttribute('data-classid', comment.classId);
        button.setAttribute('data-commentid', comment.commentId);
        button.addEventListener('click', PostComment)

        divup.innerHTML = `${comment.studentId} ${comment.comment}`;
        // good bad　未作成

        divdown.innerHTML = `<i class="fas fa-thumbs-up"></i>${comment.good} <i class="fas fa-thumbs-down"></i>${comment.bad}`
        summary.appendChild(divup);
        summary.appendChild(divdown);

        // 改良する必要があり
        summary.addEventListener('click', function() {
            // let classId = event.target.getAttribute('id-classid');
            // let commentId = event.target.getAttribute('id-commentId');
            // console.log(event.target);
        
            let httpRequest = new XMLHttpRequest();
            httpRequest.onreadystatechange = function() {
                if (httpRequest.readyState === XMLHttpRequest.DONE) {
                    if (httpRequest.status === 200) {
                        let replyComments = JSON.parse(httpRequest.responseText);
                        replyComments.body.map((comment)=>{SetReplyComments(comment)});
                    }
                }
            }
            httpRequest.open('GET', '/comments/reply/?classId='+comment.classId+'&commentId='+comment.commentId, true);
            httpRequest.send();
        });
        
        details.appendChild(summary);
        details.appendChild(ul);
        details.appendChild(text);
        details.appendChild(button);
        commentsWrapper.appendChild(details);
}

function PostComment(event) {
    let flg = event.target.getAttribute('flg');
    let comment = event.target.getAttribute('data-text');
    let httpRequest = new XMLHttpRequest();
    const studentId = document.getElementById('studentId').textContent;
    const classId = event.target.getAttribute('data-classId');

    if (flg === 'main') {
        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let comment = JSON.parse(httpRequest.responseText);
                    SetComment(comment.body);
                }
            }
        }
        httpRequest.open('POST', '/comments/main/', true);
        httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        httpRequest.send('classId='+classId+'&studentId='+studentId+'&comment='+comment);
    } else if (flg === 'reply') {
        const commentId = event.target.getAttribute('data-commentid');

        httpRequest.onreadystatechange = function() {
            if (httpRequest.readyState === XMLHttpRequest.DONE) {
                if (httpRequest.status === 200) {
                    let comment = JSON.parse(httpRequest.responseText);
                    SetReplyComments(comment.body);
                }
            }
        }
        httpRequest.open('POST', '/comments/reply/', true);
        httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        httpRequest.send('classId='+classId+'&studentId='+studentId+'&comment='+comment+'&commentId='+commentId);
    }
}


// function GetReplyComment(classId, commentId) {
//     // let classId = event.target.getAttribute('id-classid');
//     // let commentId = event.target.getAttribute('id-commentId');
//     // console.log(event.target);

//     let httpRequest = new XMLHttpRequest();
//     httpRequest.onreadystatechange = function() {
//         if (httpRequest.readyState === XMLHttpRequest.DONE) {
//             if (httpRequest.status === 200) {
//                 let replyComments = JSON.parse(httpRequest.responseText);
//                 replyComments.body.map(SetReplyComments(comment));
//             }
//         }
//     }
//     httpRequest.open('GET', '/comments/reply/?classId='+classId&'commentId='+commentId, true);
//     httpRequest.send();
// }

function SetReplyComments(comment) {
    console.log(comment.commentId);
    let ul = document.getElementById('commentId'+comment.commentId);
    let li = document.createElement('li');
    console.log(ul);
    li.innerHTML = `${comment.comment}`;
    ul.appendChild(li);
}

function removeMainContent() {
    let main = document.getElementById('main');
    main.innerHTML = '';
}

export default GetComments;