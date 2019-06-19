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

function Leave() {}

function Next() {}

function SendMsg(userid) {
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
          fromuserid: userid,
          body: inputMessage.value
        })
      });
      res = await res.json();
      console.log(res);

      if (res.code !== 212) {
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

  let latest = parseInt(sessionStorage.getItem("latest"));
  latest += 1;
  sessionStorage.setItem("latest", String(latest));

  ShowMessage(res);
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

  SendMsg(res.data.id);
  Polling();
  Leave();
  Next();
});
