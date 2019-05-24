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