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
  pname.innerText = res.data.sender;

  let pmessage = document.createElement("p");
  divCol10.appendChild(pmessage);
  pmessage.innerText = res.data.body;

  let content = document.getElementById("content");
  content.appendChild(divRow);
}

function Leave() {}

function Next() {}

function Form() {
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
          rid: parseInt(sessionStorage.getItem("rid")),
          body: inputMessage.value
        })
      });
      res = await res.json();
      console.log(res)

      if (res.code !== 212){
        return
      }

      inputMessage.value = "";
    }
  });
}

function ReceiveMsg(callback) {
  fetch("/chat_stranger/api/chat/receive",{
    method: "POST",
    headers: {
      Authorization: "Bearer" + sessionStorage.getItem("token"),
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      id: parseInt(sessionStorage.getItem("rid")),
    })
  }).then(res => res.json())
  .then(res => {
    if (res.code !== 213){
      console.log(res)
      alert("Bad")
      return
    }
    callback(res)
    ReceiveMsg(callback)
  })
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

  Form();
  Leave();
  Next();
  ReceiveMsg(ShowMessage)
});
