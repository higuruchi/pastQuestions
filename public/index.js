import {Home, searchClass} from './lib/home.js'
import QuestionBoard from './lib/questionBoard.js'
import Login from './lib/login.js'
// import checkInput from './lib/common.js'

window.onload = function() {
    document.querySelector('.home').addEventListener('click', Home);
    // document.querySelector('.questionBoard').addEventListener('click', QuestionBoard);
    document.querySelector('header div.search button').addEventListener('click', searchClass);
    document.querySelector('.questionBoard').addEventListener('click', QuestionBoard);
    document.getElementById('loginButton').addEventListener('click', Login)
    Home();
}