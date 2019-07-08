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

function SingUpAPI(register_name, password, full_name) {
  return fetch(BuildUrl(endpointAPI + "/auth/signup", {}), {
    method: "POST",
    body: JSON.stringify({
      register_name: register_name,
      password: password,
      full_name: full_name
    })
  }).then(res => res.json());
}

function LogInAPI(register_name, password) {
  return fetch(BuildUrl(endpointAPI + "/auth/login", {}), {
    method: "POST",
    body: JSON.stringify({
      register_name: register_name,
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

function ChatReceiveAPI(token, fromTime) {
  return fetch(
    BuildUrl(endpointAPI + "/chat/receive", {
      fromTime: fromTime
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
