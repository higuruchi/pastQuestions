import React, {Component} from 'react';

class SelectClass extends Component {
    constructor(props) {
        super(props);
        this.state = {
            class: this.getClass(),
            searchText: props.searchText
        }
        this.setClass = this.setClass.bind(this);
        this.getClass = this.getClass.bind(this);
    }

    getClass() {
        // 授業の取得
    }
    setClass() {
        this.setState({
            class: this.getClass()
        })
    }

    render() {
        let classLi = this.state.class.map((data) => {
            return <li onClick={this.selectClass} data-classId={this.data.classId}>{data.className}</li>
        })
        return(
            <ul>
                {classLi}
            </ul>
        )
    };

}

export default SelectClass;