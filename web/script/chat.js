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
    let res = await LeaveAPI(parseInt(sessionStorage.getItem("roomid")));
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
    let res_next = await NextAPI(parseInt(sessionStorage.getItem("roomid")));
    if (res_next.code !== 214) {
      console.log(res_next);
      return;
    }

    let res_join = await JoinAPI(res_next.data);
    if (res_join.code !== 210) {
      console.log(res_join);
      sessionStorage.removeItem("roomid");
      return;
    }

    sessionStorage.setItem("roomid", res_next.data);
    location.reload();
  });
}

function FormSend() {
  let formChat = document.getElementById("formChat");
  formChat.addEventListener("submit", async event => {
    event.preventDefault();

    let inputMessage = document.getElementById("inputMessage");
    if (inputMessage.value !== "") {
      let res = await SendAPI(
        parseInt(sessionStorage.getItem("roomid")),
        inputMessage.value
      );
      if (res.code !== 212) {
        console.log(res);
        return;
      }

      inputMessage.value = "";
    }
  });
}

function Polling() {
  if (!sessionStorage.getItem("latest")) {
    sessionStorage.setItem("latest", "-1");
  }

  setInterval(async () => {
    let res = await ReceiveAPI(
      parseInt(sessionStorage.getItem("roomid")),
      parseInt(sessionStorage.getItem("latest"))
    );
    if (res.code !== 213) {
      console.log(res);
    }
    if (res.data) {
      ShowMessage(res);
    }

    sessionStorage.setItem("latest", res.latest);
  }, 1000);
}

async function LongPolling() {
  if (!sessionStorage.getItem("latest")) {
    sessionStorage.setItem("latest", "-1");
  }

  let res = await ReceiveAPI(
    parseInt(sessionStorage.getItem("roomid")),
    parseInt(sessionStorage.getItem("latest"))
  );

  if (res.code !== 213) {
    console.log(res);
  }
  if (res.data) {
    ShowMessage(res);
  }

  sessionStorage.setItem("latest", res.latest);
  LongPolling();
}

window.addEventListener("load", async () => {
  FormSend();
  Polling();
  // LongPolling();
  Leave();
  Next();
});
