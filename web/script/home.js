$(() => {
  sessionStorage.clear();

  // only show log in
  $("#cardSignUp").addClass("d-none");
  $("#cardLogIn").removeClass("d-none");

  $("#cardSignUp a").on("click", () => {
    $("#cardSignUp").addClass("d-none");
    $("#cardLogIn").removeClass("d-none");
  });

  $("#cardLogIn a").on("click", () => {
    $("#cardSignUp").removeClass("d-none");
    $("#cardLogIn").addClass("d-none");
  });

  $("#cardSignUp form").on("submit", async event => {
    event.preventDefault();

    let res = await SingUpAPI(
      $("#inputRegisterNameSignUp").val(),
      $("#inputPasswordSignUp").val(),
      $("#inputFullNameSignUp").val()
    );

    if (res.code !== 100) {
      $("#errSignUp").text(res.message);
      return;
    }

    // token
    sessionStorage.setItem("token", res.data);
    location.href = endpointWEB + "/chat";
  });

  $("#cardLogIn form").on("submit", async event => {
    event.preventDefault();

    let res = await LogInAPI(
      $("#inputRegisterNameLogIn").val(),
      $("#inputPasswordLogIn").val()
    );

    if (res.code !== 200) {
      $("#errLogIn").text(res.message);
      return;
    }

    // token
    sessionStorage.setItem("token", res.data);
    location.href = endpointWEB + "/chat";
  });
});
