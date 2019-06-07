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
        .then(res => res.json())
        .then(res => {
            console.log(res);
            if (res.code !== 201) {
                sessionStorage.removeItem('token');
                location.href = baseurl
            }

            let welcome = document.getElementById('welcome');
            welcome.innerText = res.data.fullname;

            let btnLogOut = document.getElementById('btnLogOut');
            btnLogOut.addEventListener('click', () => {
                sessionStorage.removeItem('token');
                location.href = baseurl
            });

            let btnStartChat = document.getElementById('btnStartChat');
            btnStartChat.addEventListener('click', () => {
                fetch(baseurl + '/api/me/room', {
                    headers: {
                        'Authorization': 'Bearer' + token
                    }
                })
                    .then(res => res.json())
                    .then(res => {
                        console.log(res);
                        location.href = baseurl + '/chat' + '/' + res.room
                    })
                    .catch(err => {
                        console.log(err);
                        sessionStorage.removeItem('token');
                        location.href = baseurl
                    })
            })
        })
        .catch(err => {
            console.log(err);
            sessionStorage.removeItem('token');
            location.href = baseurl
        })
});
