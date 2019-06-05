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
            .then(res => res.json())
            .then((res) => {
                console.log(res);
                if (res.code !== 206) {
                    return
                }
                sessionStorage.setItem('token', res.token);
                location.href = baseurl + '/welcome_user'
            })
            .catch((err) => {
                console.log(err)
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
            .then(res => res.json())
            .then((res) => {
                console.log(res);
                if (res.code !== 205) {
                    return
                }
                location.reload()
            })
    })
});

