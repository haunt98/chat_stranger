window.addEventListener('load', function () {
    let baseurl = location.protocol + '//' + location.host;
    let token = sessionStorage.getItem('token');

    if (token === null) {
        location.href = baseurl
    }

    // Get user info
    let user;
    fetch(baseurl + '/api/me', {
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
                sessionStorage.removeItem('token');
                location.href = baseurl;
            }
        });

    let roomid = location.href.split('/')[4];

    let wsurl = 'ws:' + '//' + location.host + '/api/public/ws' + '?id=' + roomid;
    let conn = new WebSocket(wsurl);

    let msgOpen = 'has joined';
    let msgClose = 'has closed';

    conn.onopen = function (event) {
        // conn.send(JSON.stringify({
        //     fullname: user.fullname,
        //     body: msgOpen
        // }))
        // conn.send('Hi')
    };

    conn.onclose = function (event) {
        // conn.send(JSON.stringify({
        //     fullname: user.fullname,
        //     body: msgClose
        // }))
        // conn.send('Bye')
    };

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

        if (inputMessage.value !== '') {
            conn.send(JSON.stringify({
                fullname: user.fullname,
                body: inputMessage.value
            }));

            inputMessage.value = ''
        }
    });

    let btnLeave = document.getElementById('btnLeave');
    btnLeave.addEventListener('click', function () {
        conn.close();
        location.href = baseurl + '/welcome_user'
    });

    let btnNext = document.getElementById('btnNext');
    btnNext.addEventListener('click', function () {
        fetch(baseurl + '/api/me/room' + '?id=' + roomid, {
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