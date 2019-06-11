function ShowMessage(message) {
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
  pname.innerText = message.fullname;

  let pmessage = document.createElement("p");
  divCol10.appendChild(pmessage);
  pmessage.innerText = message.body;

  let content = document.getElementById("content");
  content.appendChild(divRow);
}

function Leave(conn) {
  let btnLeave = document.getElementById("btnLeave");
  btnLeave.addEventListener("click", () => {
    conn.close();
    location.href = "/chat_stranger/web/welcome";
  });
}

function Next(token, rid) {
  let btnNext = document.getElementById("btnNext");
  btnNext.addEventListener("click", async () => {
    let res = await fetch("/chat_stranger/api/me/room" + "?rid=" + rid, {
      headers: {
        Authorization: "Bearer" + token
      }
    });
    res = await res.json();
    location.href = "/chat_stranger/web/chat" + "?rid=" + res.room;
  });
}

function Form(res, conn) {
  let formChat = document.getElementById("formChat");
  formChat.addEventListener("submit", event => {
    event.preventDefault();

    let inputMessage = document.getElementById("inputMessage");
    if (inputMessage.value !== "") {
      conn.send(
        JSON.stringify({
          fullname: res.data.fullname,
          body: inputMessage.value
        })
      );
      inputMessage.value = "";
    }
  });
}

window.addEventListener("load", async () => {
  let token = sessionStorage.getItem("token");
  if (!token) {
    location.href = "/chat_stranger/web";
    return;
  }

  let res = await fetch("/chat_stranger/api/me", {
    headers: {
      Authorization: "Bearer" + token
    }
  });
  res = await res.json();
  if (res.code !== 201) {
    console.log(res);
    sessionStorage.removeItem("token");
    location.href = "/chat_stranger/web";
    return;
  }

  let uid = res.data.id;
  let params = await new this.URL(location.href);
  let rid = await params.searchParams.get("rid");
  if (!rid) {
    location.href = "/chat_stranger/web";
    return;
  }

  let wsurl =
    "ws:" +
    "//" +
    location.host +
    "/chat_stranger/api/ws" +
    "?rid=" +
    rid +
    "&uid=" +
    uid;
  let conn = new WebSocket(wsurl);
  conn.onmessage = event => {
    let message = JSON.parse(event.data);
    ShowMessage(message);
  };

  Form(res, conn);
  Leave(conn);
  Next(token, rid);
});
