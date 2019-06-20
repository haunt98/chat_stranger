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
    let res_empty = await EmptyAPI();
    if (res_empty.code !== 209) {
      console.log(res_empty);
      return;
    }
    sessionStorage.setItem("roomid", res_empty.data);

    let res_join = await JoinAPI(res_empty.data);
    if (res_join.code !== 210) {
      console.log(res_join);
      sessionStorage.removeItem("roomid");
      return;
    }

    location.href = "/chat_stranger/web/chat";
  });
}

window.addEventListener("load", async () => {
  let res = await MeGETAPI();
  if (res.code !== 201) {
    sessionStorage.removeItem("token");
    location.href = "/chat_stranger/web";
    return;
  }

  Show(res);
  LogOut();
  Chat();
});
