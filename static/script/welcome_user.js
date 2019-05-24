window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = localStorage.getItem('token');

    // If not log in -> force homepage
    // If log in -> welcome
    if (token === null) {
        if (location.href !== baseurl + '/index') {
            location.href = baseurl + '/index'
        }
    } else {
        if (location.href !== baseurl + '/welcome_user') {
            location.href = baseurl + '/welcome_user'
        }
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

                location.href = baseurl + '/index'
            }
        });


    let btnLogOut = document.getElementById('btnLogOut');
    btnLogOut.addEventListener('click', function () {
        localStorage.removeItem('token');

        location.href = baseurl + '/index'
    })
});
