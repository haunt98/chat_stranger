function SignIn() {
  let formSignIn = document.getElementById("formSignIn");
  formSignIn.addEventListener("submit", async event => {
    event.preventDefault();

    let regname = document.getElementById("inputRegNameSignIn").value;
    let password = document.getElementById("inputPasswordSignIn").value;

    let res = await fetch("/chat_stranger/api/auth/login", {
      method: "POST",
      body: JSON.stringify({
        regname: regname,
        password: password
      })
    });
    res = await res.json();
    if (res.code !== 206) {
      signInErr = document.getElementById("signInErr")
      signInErr.innerText = res.message
      return;
    }
    sessionStorage.setItem("token", res.token);
    location.href = "/chat_stranger/web/welcome";
  });
}

function SignUp() {
  let formSignUp = document.getElementById("formSignUp");
  formSignUp.addEventListener("submit", async event => {
    event.preventDefault();

    let regname = document.getElementById("inputRegNameSignUp").value;
    let password = document.getElementById("inputPasswordSignUp").value;
    let fullname = document.getElementById("inputFullNameSignUp").value;

    let res = await fetch("/chat_stranger/api/auth/register", {
      method: "POST",
      body: JSON.stringify({
        regname: regname,
        password: password,
        fullname: fullname
      })
    });
    res = await res.json();
    if (res.code !== 205) {
      signOutErr = document.getElementById("signOutErr")
      signOutErr.innerText = res.message
      return;
    }
    location.reload()
  });
}

window.addEventListener("load", () => {
  SignIn();
  SignUp();
});
