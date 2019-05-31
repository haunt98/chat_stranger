window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = sessionStorage.getItem('token');

    if (token !== null) {
        location.href = baseurl + '/welcome_user'
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
                    sessionStorage.setItem('token', response.token);
                    location.href = baseurl + '/welcome_user'
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
                console.log(response);
                if (response.code === 205) {
                    location.reload()
                }
            })
    })
});

