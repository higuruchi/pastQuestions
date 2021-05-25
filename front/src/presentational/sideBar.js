import React from 'react'
import NavIcon from '../molecules/navIcon';


let SideBar = (props) => {
    
    return (
        <aside id="aside">
            <nav>
                <NavIcon/>
                <ul>
                    <li class="home" onClick={props.handleHome}>
                        <i class="fas fa-home"></i>
                        ホーム
                    </li>
                    <li class="questionBoard" onClick={props.handleQuestionBoard}>
                        <i class="fas fa-question"></i>
                        質問掲示板
                    </li>
                    <li class="student" onClick={props.handleStudent}>
                        <i class="fas fa-user"></i>
                        ユーザ情報
                    </li>
                </ul>
            </nav>
            <div class="sideFooter">
                side Footer
            </div>
        </aside>
    )
}

export default SideBar;