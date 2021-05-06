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
                setMainContent();
            }
        }

    }

    function setMainContent() {
        removeMainContent();
        let ul = document.createElement('ul');
        classInfo.body.map(function(content) {
            let li = document.createElement('li');
            li.innerHTML = `授業ID：${content.classId}　　授業名：${content.className}`;
            li.setAttribute('data-classid', content.classId);
            li.addEventListener('click', GetPastQuestionAndComment)
            ul.appendChild(li);
        });
        document.getElementById('leftWrapper').appendChild(ul);
    }
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
                console.log('ok');
                setMainContent();
            }
        }

    }

    function setMainContent() {
        removeMainContent();
        let ul = document.createElement('ul');
        classInfo.body.map(function(content) {
            let li = document.createElement('li');
            li.innerHTML = `授業ID：${content.classId}　　授業名：${content.className}`;
            li.setAttribute('data-classid', content.classId);
            li.addEventListener('click', GetPastQuestionAndComment)
            ul.appendChild(li);
        });
        document.getElementById('leftWrapper').appendChild(ul);
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
