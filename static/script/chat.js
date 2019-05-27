window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = localStorage.getItem('token');

    if (token === null) {
        location.href = baseurl
    }

    let wsurl = 'ws:' + '//' + location.host + '/ws'
    console.log(wsurl)
    conn = new WebSocket(wsurl)

    conn.onopen = function (event) {
        console.log(event)
    }

    conn.onclose = function (event) {
        console.log(event)
    }

    conn.onerror = function (event) {
        console.log(event)
    }

    conn.onmessage = function (event) {
        console.log(event.data)
    }

    let formChat = document.getElementById('formChat')
    formChat.addEventListener('submit', function (event) {
        event.preventDefault()

        let inputMessage = document.getElementById('inputMessage').value

        conn.send(inputMessage)
    })
});