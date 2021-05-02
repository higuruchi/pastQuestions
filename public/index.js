import {Home, searchClass} from './lib/home.js'
// import checkInput from './lib/common.js'

window.onload = function() {
    document.querySelector('.home').addEventListener('click', Home);
    document.querySelector('header div.search button').addEventListener('click', searchClass);
    // document.querySelector('.questionBoard').addEventListener('click', )
    Home();
    // home.setMainContent();

// const homeNav = document.querySelector(".home");

// homeNav.addEventListener("click",)

}