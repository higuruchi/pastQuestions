// let pastQuestionWrapper = document.createElement('div');
// pastQuestionWrapper.setAttribute('class', 'pastQuestionWrapper');
let pastQuestionWrapper = document.getElementById('leftWrapper');
let pastQuestionWrapperOver = document.createElement('div');

function GetPastQuestion(classId) {
    let httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                let pastQuestionInfo = JSON.parse(httpRequest.responseText);
                SetPastQuestions(pastQuestionInfo, classId);
            }
        }
    }
    httpRequest.open('GET', '/pastQuestion/'+classId+'/0/', true);
    httpRequest.send();
}

function SetPastQuestions(pastQuestionInfo, classId) {
    let pastQuestions = pastQuestionInfo.body;
    let form = document.createElement('form');
    let inputFile = document.createElement('input');
    let inputSubmit = document.createElement('input');
    let inputYear = document.createElement('input');
    let selectSemester = document.createElement('select');
    let selectorOption = {
        1: '第1セメスター',
        2: '第2セメスター',
        3: '第3セメスター',
        4: '第4セメスター'
    };
    let now = new Date();
    let nowSemester;
    if (4 <= now.getMonth()+1 <= 5) {
        nowSemester = 1;
    } else if (6 <= now.getMonth()+1 <= 9) {
        nowSemester = 2;
    } else if (10 <= now.getMonth()+1 <= 11) {
        nowSemester = 3;
    } else if (12 === now.getMonth()+1 || 1 <= now.getMonth()+1 <= 3) {
        nowSemester = 4;
    }

    inputYear.setAttribute('type', 'tel');
    inputYear.setAttribute('maxlentgh', '4');
    form.setAttribute('id', 'postPastQuestionForm');
    inputFile.setAttribute('type', 'file');
    inputFile.setAttribute('name', 'pastQuestion');
    inputFile.setAttribute('accept', 'application/pdf');
    inputSubmit.setAttribute('type', 'button');
    inputSubmit.setAttribute('value', 'send');
    inputSubmit.setAttribute('data-year', now.getFullYear());
    inputSubmit.setAttribute('data-semester', nowSemester);
    inputSubmit.setAttribute('data-classid', classId);
    inputSubmit.addEventListener('click', PostPastQuestion);
    selectSemester.addEventListener('change', function() {
        inputSubmit.setAttribute('data-semester', this.value);
    });
    inputYear.addEventListener('change', function() {
        inputSubmit.setAttribute('data-year', this.value);
    })

    for (let i=1; i <= 4; i++) {
        let selectorOptionTmp = document.createElement('option');
        selectorOptionTmp.setAttribute('value', i);
        selectorOptionTmp.innerText = selectorOption[i];
        selectSemester.appendChild(selectorOptionTmp);
    }

    if (pastQuestions !== null) {
        let ul = document.createElement('ul');
        pastQuestions.map(function(pastQuestion){
            let li = document.createElement('li');
            let link = document.createElement('a');
            link.setAttribute('href', '/pastQuestions/'+pastQuestion.fileName);
            link.setAttribute('data-fileid', pastQuestion.fileId)
            link.innerText = pastQuestion.year+'年：第'+pastQuestion.semester+'学期';
            li.appendChild(link);
            ul.appendChild(li);
        });
        pastQuestionWrapperOver.appendChild(ul);
    }
    pastQuestionWrapperOver.appendChild(inputYear);
    pastQuestionWrapperOver.append(selectSemester);
    form.appendChild(inputFile);
    form.appendChild(inputSubmit);
    pastQuestionWrapperOver.appendChild(form);
    pastQuestionWrapper.appendChild(pastQuestionWrapperOver);
    
}

function PostPastQuestion(event) {


    let pastQuestion = document.getElementById('postPastQuestionForm');
    let httpRequest = new XMLHttpRequest();
    let data = new FormData(pastQuestion);
    let year = event.target.getAttribute('data-year');
    let semester = event.target.getAttribute('data-semester');
    let classId = event.target.getAttribute('data-classid');

    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                // let comment = JSON.parse(httpRequest.responseText);
                
            } else {
                console.log('fail');
            }
        }
    }
    httpRequest.open('POST', `/pastQuestion/${classId}/${year}/${semester}/`, true);
    httpRequest.send(data);
}
export default GetPastQuestion;