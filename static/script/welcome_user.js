window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = localStorage.getItem('token');

    if (token === null) {
            location.href = baseurl + '/index'
    }

    // Get user info
    fetch(baseurl + '/api/privateForUser', {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer' + token
        }
    })
        .then(response => response.json())
        .then(function (response) {
            console.log(response);
            if (response.code === 201) {
                console.log(response.message);

                document.getElementById('welcomeUsername').innerText = response.user.fullname
            } else {
                console.log(response.message);

                localStorage.removeItem('token');
                location.href = baseurl + '/index'
            }
        });


    let btnLogOut = document.getElementById('btnLogOut');
    btnLogOut.addEventListener('click', function () {
        localStorage.removeItem('token');

        location.href = baseurl + '/index'
    });

    let btnStartChat = document.getElementById('btnStartChat');
    btnStartChat.addEventListener('click', function () {
        location.href = baseurl + '/chat'
    })
});
