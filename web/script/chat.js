function ShowMessage(res) {
  let divRow = document.createElement("div");
  divRow.className = "row";

  let divCol2 = document.createElement("div");
  divCol2.className = "col-md-2";

  let divCol10 = document.createElement("div");
  divCol10.className = "col-md-10";

  divRow.appendChild(divCol2);
  divRow.appendChild(divCol10);

  let pname = document.createElement("p");
  pname.className = "font-weight-bold";
  divCol2.appendChild(pname);
  pname.innerText = res.data.fromuser;

  let pmessage = document.createElement("p");
  divCol10.appendChild(pmessage);
  pmessage.innerText = res.data.body;

  let content = document.getElementById("content");
  content.appendChild(divRow);
}

function Leave() {
  let btnLeave = document.getElementById("btnLeave");
  btnLeave.addEventListener("click", async () => {
    let res = await fetch("/chat_stranger/api/chat/leave", {
      method: "POST",
      headers: {
        Authorization: "Bearer" + sessionStorage.getItem("token"),
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        id: parseInt(sessionStorage.getItem("roomid"))
      })
    });
    res = await res.json();
    if (res.code !== 211) {
      console.log(res);
      return;
    }

    sessionStorage.removeItem("roomid");
    location.href = "/chat_stranger/web/welcome";
  });
}

function Next() {
  let btnNext = document.getElementById("btnNext");
  btnNext.addEventListener("click", async () => {
    let res_next = await fetch("/chat_stranger/api/chat/next", {
      method: "POST",
      headers: {
        Authorization: "Bearer" + sessionStorage.getItem("token"),
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        id: parseInt(sessionStorage.getItem("roomid"))
      })
    });
    res_next = await res_next.json();
    if (res_next.code !== 214) {
      console.log(res_next);
      return;
    }

    let res_join = await fetch("/chat_stranger/api/chat/join", {
      method: "POST",
      headers: {
        Authorization: "Bearer" + sessionStorage.getItem("token"),
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        id: res_next.data
      })
    });
    res_join = await res_join.json();
    if (res_join.code !== 210) {
      console.log(res_join);
      sessionStorage.removeItem("roomid");
      return;
    }

    sessionStorage.setItem("roomid", res_next.data);
    location.reload();
  });
}

function SendMsg() {
  let formChat = document.getElementById("formChat");
  formChat.addEventListener("submit", async event => {
    event.preventDefault();

    let inputMessage = document.getElementById("inputMessage");
    if (inputMessage.value !== "") {
      let res = await fetch("/chat_stranger/api/chat/send", {
        method: "POST",
        headers: {
          Authorization: "Bearer" + sessionStorage.getItem("token"),
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          roomid: parseInt(sessionStorage.getItem("roomid")),
          body: inputMessage.value
        })
      });
      res = await res.json();
      if (res.code !== 212) {
        console.log(res);
        return;
      }

      inputMessage.value = "";
    }
  });
}

async function ReceiveMsg() {
  let res = await fetch("/chat_stranger/api/chat/receive", {
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      roomid: parseInt(sessionStorage.getItem("roomid")),
      latest: parseInt(sessionStorage.getItem("latest"))
    })
  });
  res = await res.json();
  if (res.code !== 213) {
    return;
  }

  if (res.latest !== -1) {
    ShowMessage(res);
  }
  sessionStorage.setItem("latest", res.latest);
}

function Polling() {
  setInterval(ReceiveMsg, 1000);
}

window.addEventListener("load", async () => {
  let res = await fetch("/chat_stranger/api/me", {
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token")
    }
  });
  res = await res.json();
  if (res.code !== 201) {
    console.log(res);
    sessionStorage.removeItem("token");
    location.href = "/chat_stranger/web";
    return;
  }

  sessionStorage.setItem("latest", "-1");

  SendMsg();
  Polling();
  Leave();
  Next();
});
