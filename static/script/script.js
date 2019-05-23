window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host
    let token = ''

    if (token === '') {
        if (location.href !== baseurl + '/index') {
            return
        }
    }

    let formSignIn = document.getElementById('formSignIn')
    formSignIn.addEventListener('submit', function (event) {
        event.preventDefault()
        console.log('Hello')

        let username = document.getElementById('inputUsernameSignIn').value
        let password = document.getElementById('inputPasswordSignIn').value


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
                console.log(response)
                console.log(window.location.href)
                if (response.code === 206) {
                    console.log('Log in OK')
                    token = response.token
                    // location.href = baseurl + '/welcome_user'
                } else {
                    console.log('Log in failed')
                }
            })
    })
})

