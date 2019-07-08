$(() => {
  sessionStorage.clear();

  $("#formSignUp").on("submit", async event => {
    event.preventDefault();

    let res = await SingUpAPI(
      $("#inputRegisterNameSignUp").val(),
      $("#inputPasswordSignUp").val(),
      $("#inputFullNameSignUp").val()
    );

    if (res.code !== 1) {
      $("#errSignUp").text(res.message);
      return;
    }

    // token
    sessionStorage.setItem("token", res.data);
    location.href = endpointWEB + "/chat";
  });

  $("#formLogIn").on("submit", async event => {
    event.preventDefault();

    let res = await LogInAPI(
      $("#inputRegisterNameLogIn").val(),
      $("#inputPasswordLogIn").val()
    );

    if (res.code !== 2) {
      $("#errLogIn").text(res.message);
      return;
    }

    // token
    sessionStorage.setItem("token", res.data);
    location.href = endpointWEB + "/chat";
  });
});
