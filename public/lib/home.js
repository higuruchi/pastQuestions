import GetPastQuestionAndComment from './pastQuestionAndComment.js'

function Home() {
    let httpRequest = new XMLHttpRequest();
    let classInfo;
    
    if (!httpRequest) {
            return false;
    }
    httpRequest.onreadystatechange = getClassInfoHandle;
    httpRequest.open('GET', '/classes/', true);
    httpRequest.send();

    function getClassInfoHandle() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                classInfo = JSON.parse(httpRequest.responseText);
                setMainContent(classInfo);
            }
        }

    }

}
function setMainContent(classInfo) {
    removeMainContent();
    let ul = document.createElement('ul');
    ul.setAttribute('class', 'classes');
    classInfo.body.map(function(content) {
        let li = document.createElement('li');
        let className = document.createElement('div');
        className.innerText = `${content.className}`
        li.appendChild(className);
        li.setAttribute('data-classid', content.classId);
        className.setAttribute('data-classid', content.classId);
        li.addEventListener('click', GetPastQuestionAndComment)
        ul.appendChild(li);
    });
    document.getElementById('leftWrapper').appendChild(ul);
}

function searchClass() {
    let searchCondition = document.querySelector("header div.search input[type=text]").value
    console.log(searchCondition)
    let httpRequest = new XMLHttpRequest();
    let classInfo;
    
    if (!httpRequest) {
        return false;
    }
    httpRequest.onreadystatechange = getClassInfoHandle;
    httpRequest.open('GET', '/classes/?className='+searchCondition, true);
    httpRequest.send();

    function getClassInfoHandle() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                classInfo = JSON.parse(httpRequest.responseText);
                console.log(httpRequest.responseText);
                setMainContent(classInfo);
            }
        }

    }
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
export {Home, searchClass};
