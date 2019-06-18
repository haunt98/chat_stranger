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

function Chat() {
  let btnStartChat = document.getElementById("btnStartChat");
  btnStartChat.addEventListener("click", async () => {
    let res_empty = await fetch("/chat_stranger/api/chat/empty", {
      headers: {
        Authorization: "Bearer" + sessionStorage.getItem("token")
      }
    });
    res_empty = await res_empty.json();
    sessionStorage.setItem("rid", res_empty.data);

    let res_join = await fetch("/chat_stranger/api/chat/join", {
      method: "POST",
      headers: {
        Authorization: "Bearer" + sessionStorage.getItem("token"),
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        id: res_empty.data
      })
    });
    res_join = await res_join.json();

    location.href = "/chat_stranger/web/chat";
  });
}

window.addEventListener("load", async () => {
  let res = await fetch("/chat_stranger/api/me", {
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token")
    }
  });
  res = await res.json();
  if (res.code !== 201) {
    sessionStorage.removeItem("token");
    location.href = "/chat_stranger/web";
    return;
  }

  Show(res);
  LogOut();
  Chat();
});
