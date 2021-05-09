function Login(event) {
    let studentId = event.target.parentElement.children.item(0).value;
    let password = event.target.parentElement.children.item(1).value;

    let httpRequest = new XMLHttpRequest();
    httpRequest.onreadystatechange = function() {
        if (httpRequest.readyState === XMLHttpRequest.DONE) {
            if (httpRequest.status === 200) {
                let ret = JSON.parse(httpRequest.responseText);
                let larbel = document.querySelector('.search label');
                document.getElementById('login').checked = false;
                document.getElementById('studentId').setAttribute('value', ret.body.studentId);
                larbel.innerHTML = ret.body.studentName;
                larbel.removeAttribute('for');
            }
        }
    }
    httpRequest.open('POST', '/login', true);
    httpRequest.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    httpRequest.send(`studentId=${studentId}&password=${password}`);

}

export default Login;