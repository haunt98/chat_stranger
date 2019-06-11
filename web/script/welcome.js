function Show(res) {
  let welcome = document.getElementById("welcome");
  welcome.innerText = res.data.fullname;
}

function LogOut() {
  let btnLogOut = document.getElementById("btnLogOut");
  btnLogOut.addEventListener("click", () => {
    sessionStorage.removeItem("token");
    location.href = "/chat_stranger/web";
  });
}

function Chat(token) {
  let btnStartChat = document.getElementById("btnStartChat");
  btnStartChat.addEventListener("click", async () => {
    let res = await fetch("/chat_stranger/api/me/room", {
      headers: {
        Authorization: "Bearer" + token
      }
    });
    res = await res.json();
    location.href = "/chat_stranger/web/chat" + "?rid=" + res.room;
  });
}

window.addEventListener("load", async () => {
  let token = sessionStorage.getItem("token");
  if (!token) {
    location.href = "/chat_stranger/web";
    return
  }

  let res = await fetch("/chat_stranger/api/me", {
    headers: {
      Authorization: "Bearer" + token
    }
  });
  res = await res.json();
  if (res.code !== 201) {
    sessionStorage.removeItem("token");
    location.href = "/chat_stranger/web";
    return
  }

  Show(res);
  LogOut();
  Chat(token);
});
