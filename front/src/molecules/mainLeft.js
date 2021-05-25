import React from 'react';

let MainLeft = (props) => {
    if (props.mainContentState === "home") {
        let classData = "";

        classData = props.classesData.map(function(data) {
            return (
                    <li
                        key={data.classId}
                        data-classid={data.classId}
                        onClick={props.handleGetMainComment}
                    >
                        {data.className}
                    </li>
                    )
        });
        return (
            <div id="mainLeft">
                <ul class="classes">
                    {classData}
                </ul>
            </div>
        )
    }
}

export default MainLeft;