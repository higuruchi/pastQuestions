import React from 'react'

const List = (props) => {
    let li = props.list.map((data) =>
        <li
        onClick={props.handleClick}
        key={data.commentReplyId}
        >{data}</li>
    );
    return (
        <ul>
            {li}
        </ul>
    )
}
export default List;