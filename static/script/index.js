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

    let formSignIn = document.getElementById('formSignIn');
    formSignIn.addEventListener('submit', function (event) {
        event.preventDefault();

        let username = document.getElementById('inputUsernameSignIn').value;
        let password = document.getElementById('inputPasswordSignIn').value;

        fetch(baseurl + '/api/public/users/authenticate', {
            method: 'POST',
            body: JSON.stringify(
                {
                    name: username,
                    password: password
                }
            )
        })
            .then(response => response.json())
            .then(function (response) {
                console.log(response);
                if (response.code === 206) {
                    console.log(response.message);

                    localStorage.setItem('token', response.token);
                    location.href = baseurl + '/welcome_user'
                } else {
                    console.log(response.message)
                }
            })
    });

    let formSignUp = document.getElementById('formSignUp');
    formSignUp.addEventListener('submit', function (event) {
        event.preventDefault();

        let username = document.getElementById('inputUsernameSignUp').value;
        let password = document.getElementById('inputPasswordSignUp').value;
        let fullname = document.getElementById('inputFullNameSignUp').value;

        fetch(baseurl + '/api/public/users/register', {
            method: 'POST',
            body: JSON.stringify(
                {
                    name: username,
                    password: password,
                    fullname: fullname
                }
            )
        })
            .then(response => response.json())
            .then(function (response) {
                if (response.code === 205) {
                    console.log(response.message);
                    location.reload()
                } else {
                    console.log(response.message);
                }
            })
    })
});

