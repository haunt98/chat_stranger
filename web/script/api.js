const endpointAPI = "/chat_stranger/api";
const endpointWEB = "/chat_stranger/web";

/**
 * @return {string}
 */
function BuildUrl(url, params) {
  let searchParams = new URLSearchParams(params);
  return url + "?" + searchParams;
}

async function EmptyAPI() {
  let res = await fetch("/chat_stranger/api/chat/empty", {
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token")
    }
  });
  res = await res.json();
  return res;
}

async function NextAPI(oldroomid) {
  let res = await fetch("/chat_stranger/api/chat/next", {
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      id: oldroomid
    })
  });
  res = await res.json();
  return res;
}

async function JoinAPI(roomid) {
  let res = await fetch("/chat_stranger/api/chat/join", {
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      id: roomid
    })
  });
  res = await res.json();
  return res;
}

async function LeaveAPI(roomid) {
  let res = await fetch("/chat_stranger/api/chat/leave", {
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      id: roomid
    })
  });
  res = await res.json();
  return res;
}

async function SendAPI(roomid, body) {
  let res = await fetch("/chat_stranger/api/chat/send", {
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      roomid: roomid,
      body: body
    })
  });
  res = await res.json();
  return res;
}

async function ReceiveAPI(roomid, latest) {
  let res = await fetch("/chat_stranger/api/chat/receive", {
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      roomid: roomid,
      latest: latest
    })
  });
  res = await res.json();
  return res;
}

async function MeGETAPI() {
  let res = await fetch("/chat_stranger/api/me", {
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token")
    }
  });
  res = await res.json();
  return res;
}

async function infoAPI(roomid, status) {
  return await fetch(
    BuildUrl(endpointAPI + "/chat/info", {
      roomid: roomid,
      status: status
    }),
    {
      headers: {
        Authorization: "Bearer" + sessionStorage.getItem("token")
      }
    }
  ).then(res => res.json());
}
