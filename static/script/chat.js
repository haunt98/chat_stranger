window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = localStorage.getItem('token');

    if (token === null) {
        location.href = baseurl
    }

});