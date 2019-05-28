window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = localStorage.getItem('token');

    if (token === null) {
        location.href = baseurl
    }

    // Get user info
    fetch(baseurl + '/api/me', {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer' + token
        }
    })
        .then(response => response.json())
        .then(function (response) {
            console.log(response);
            if (response.code === 201) {
                document.getElementById('welcomeUsername').innerText = response.user.fullname
            } else {
                localStorage.removeItem('token');
                location.href = baseurl
            }
        });


    let btnLogOut = document.getElementById('btnLogOut');
    btnLogOut.addEventListener('click', function () {
        localStorage.removeItem('token');

        location.href = baseurl
    });

    let btnStartChat = document.getElementById('btnStartChat');
    btnStartChat.addEventListener('click', function () {
        fetch(baseurl + '/api/public/users/roomid', {
            method: 'GET'
        })
            .then(response => response.json())
            .then(function (response) {
                console.log(response);
                location.href = baseurl + '/chat' + '/' + response.roomid
            })
    })
});
