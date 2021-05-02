function checkInput(str, condition) {
    if (str === '' || condition === '') {
        return false
    }
    let re = new RegExp(condition);
    return re.test(str);
}

export default checkInput