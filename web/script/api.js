const endpointAPI = "/chat_stranger/api";
const endpointWEB = "/chat_stranger/web";

/**
 * @return {string}
 */
function BuildUrl(url, params) {
  if (jQuery.isEmptyObject(params)) {
    return url;
  }
  let searchParams = new URLSearchParams(params);
  return url + "?" + searchParams;
}

function SingUpAPI(registername, password, showname) {
  return fetch(BuildUrl(endpointAPI + "/auth/signup", {}), {
    method: "POST",
    body: JSON.stringify({
      registername: registername,
      password: password,
      showname: showname
    })
  }).then(res => res.json());
}

function LogInAPI(registername, password) {
  return fetch(BuildUrl(endpointAPI + "/auth/login", {}), {
    method: "POST",
    body: JSON.stringify({
      registername: registername,
      password: password
    })
  }).then(res => res.json());
}

function InfoAPI(token) {
  return fetch(BuildUrl(endpointAPI + "/me", {}), {
    headers: {
      Authorization: `Bearer ${token}`
    }
  }).then(res => res.json());
}

function EditInfoAPI(token, showname, gender, birthyear) {
  return fetch(BuildUrl(endpointAPI + "/me", {}), {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${token}`
    },
    body: JSON.stringify({
      showname: showname,
      gender: gender,
      birthyear: birthyear
    })
  }).then(res => res.json());
}

function ChatFindAPI(token, status) {
  return fetch(
    BuildUrl(endpointAPI + "/chat/find", {
      status: status
    }),
    {
      headers: {
        Authorization: `Bearer ${token}`
      }
    }
  ).then(res => res.json());
}

function ChatJoinAPI(token, roomID) {
  return fetch(
    BuildUrl(endpointAPI + "/chat/join", {
      roomID: roomID
    }),
    {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`
      }
    }
  ).then(res => res.json());
}

function ChatLeaveAPI(token) {
  return fetch(BuildUrl(endpointAPI + "/chat/leave", {}), {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`
    }
  }).then(res => res.json());
}

function ChatSendAPI(token, body) {
  return fetch(BuildUrl(endpointAPI + "/chat/send", {}), {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      body: body
    })
  }).then(res => res.json());
}

function ChatReceiveAPI(token, from) {
  return fetch(
    BuildUrl(endpointAPI + "/chat/receive", {
      from: from
    }),
    {
      headers: {
        Authorization: `Bearer ${token}`
      }
    }
  ).then(res => res.json());
}

function ChatIsFreeAPI(token) {
  return fetch(BuildUrl(endpointAPI + "/chat/is_free", {}), {
    headers: {
      Authorization: `Bearer ${token}`
    }
  }).then(res => res.json());
}

function ChatCountMember(token) {
  return fetch(BuildUrl(endpointAPI + "/chat/count_member", {}), {
    headers: {
      Authorization: `Bearer ${token}`
    }
  }).then(res => res.json());
}
