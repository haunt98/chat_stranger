window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = localStorage.getItem('token');

    if (token === null) {
        location.href = baseurl
    }

    // Get user info
    let user;
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
                user = response.user
            } else {
                localStorage.removeItem('token');
                location.href = baseurl;
            }
        });

    let roomid = location.href.split('/')[4];

    let wsurl = 'ws:' + '//' + location.host + '/ws' + '?roomid=' + roomid;
    conn = new WebSocket(wsurl);

    conn.onmessage = function (event) {
        let message = JSON.parse(event.data);

        let divRow = document.createElement('div');
        divRow.className = 'row';

        let divCol2 = document.createElement('div');
        divCol2.className = 'col-md-2';

        let divCol10 = document.createElement('div');
        divCol10.className = 'col-md-10';

        divRow.appendChild(divCol2);
        divRow.appendChild(divCol10);

        let pname = document.createElement('p');
        pname.className = 'font-weight-bold';
        divCol2.appendChild(pname);
        pname.innerText = message.fullname;

        let pmessage = document.createElement('p');
        divCol10.appendChild(pmessage);
        pmessage.innerText = message.body;

        let content = document.getElementById('content');
        content.appendChild(divRow);
    };

    let formChat = document.getElementById('formChat');
    formChat.addEventListener('submit', function (event) {
        event.preventDefault();

        let inputMessage = document.getElementById('inputMessage');

        conn.send(JSON.stringify({
            fullname: user.fullname,
            body: inputMessage.value
        }))

        inputMessage.value = ''
    })
});