window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = sessionStorage.getItem('token');

    if (token === null) {
        location.href = baseurl
    }

    fetch(baseurl + '/api/me', {
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
                sessionStorage.removeItem('token');
                location.href = baseurl
            }
        });


    let btnLogOut = document.getElementById('btnLogOut');
    btnLogOut.addEventListener('click', function () {
        sessionStorage.removeItem('token');
        location.href = baseurl
    });

    let btnStartChat = document.getElementById('btnStartChat');
    btnStartChat.addEventListener('click', function () {
        fetch(baseurl + '/api/me/room', {
            headers: {
                'Authorization': 'Bearer' + token
            }
        })
            .then(response => response.json())
            .then(function (response) {
                console.log(response);
                location.href = baseurl + '/chat' + '/' + response.room
            })
    })
});
