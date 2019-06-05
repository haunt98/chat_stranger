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
        .then((res) => {
            console.log(res);
            if (res.code !== 201) {
                sessionStorage.removeItem('token');
                location.href = baseurl;
            }

            let user = res.user;
            let rid = location.href.split('/')[4];
            let uid = user.id;

            let wsurl = 'ws:' + '//' + location.host + '/api/public/ws' + '?rid=' + rid + '&uid=' + uid;
            let conn = new WebSocket(wsurl);

            conn.onmessage = (event) => {
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
            formChat.addEventListener('submit', (event) => {
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
            btnNext.addEventListener('click', () => {
                fetch(baseurl + '/api/me/room' + '?rid=' + rid, {
                    headers: {
                        'Authorization': 'Bearer' + token
                    }
                })
                    .then(res => res.json())
                    .then((res) => {
                        console.log(res);
                        location.href = baseurl + '/chat' + '/' + res.room
                    })
                    .catch((err) => {
                        console.log(err);
                        sessionStorage.removeItem('token');
                        location.href = baseurl;
                    })
            })

        })
        .catch((err) => {
            console.log(err);
            sessionStorage.removeItem('token');
            location.href = baseurl;
        })
});