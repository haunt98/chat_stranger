function prepare() {
  // time
  let fromTime = new Date(0).toISOString();
  sessionStorage.setItem("fromTime", fromTime);

  // count member
  sessionStorage.setItem("countMember", "0");

  // color drop down
  $(".dropdown i").css("color", window.elementaryColor["Blueberry"][500]);
}

function showMessageLeft(name, body, color, backgroundColor) {
  let row = document.createElement("div");
  row.className = "row my-1";

  let div_name = document.createElement("div");
  div_name.className = "col-sm-2 d-flex justify-content-start";
  row.appendChild(div_name);

  let span_name = document.createElement("span");
  span_name.className = "font-weight-bold p-2";
  span_name.innerText = name;
  div_name.appendChild(span_name);

  let div_body = document.createElement("div");
  div_body.className = "col-sm-8 d-flex justify-content-start";
  row.appendChild(div_body);

  let span_body = document.createElement("span");
  span_body.className = "rounded p-2 text-break";
  span_body.style.color = color;
  span_body.style.backgroundColor = backgroundColor;
  span_body.innerText = body;
  div_body.appendChild(span_body);

  $("#chat").append(row);
}

function showMessageRight(body, color, backgroundColor) {
  let row = document.createElement("div");
  row.className = "row my-1";

  let div_empty = document.createElement("div");
  div_empty.className = "col-sm-4";
  row.appendChild(div_empty);

  let div_body = document.createElement("div");
  div_body.className = "col-sm-8 d-flex justify-content-end";
  row.appendChild(div_body);

  let span_body = document.createElement("span");
  span_body.className = "rounded p-2 text-break";
  span_body.style.color = color;
  span_body.style.backgroundColor = backgroundColor;
  span_body.innerText = body;
  div_body.appendChild(span_body);

  $("#chat").append(row);
}

function scrollTop() {
  let chat = $("#chat");
  chat.scrollTop(chat.prop("scrollHeight"));
}

async function receive() {
  let res_receive = await ChatReceiveAPI(
    sessionStorage.getItem("token"),
    sessionStorage.getItem("fromTime")
  );
  if (res_receive.code !== 800) {
    console.log(res_receive);
    return;
  }
  if (!res_receive.data || !res_receive.data.length) {
    return;
  }

  for (let i = 0; i < res_receive.data.length; i += 1) {
    if (
      res_receive.data[i].user_id === parseInt(sessionStorage.getItem("userID"))
    ) {
      showMessageRight(
        res_receive.data[i].body,
        window.elementaryColor["Silver"][100],
        window.elementaryColor["Blueberry"][500]
      );
    } else {
      showMessageLeft(
        res_receive.data[i].user_full_name,
        res_receive.data[i].body,
        window.elementaryColor["Black"][700],
        window.elementaryColor["Silver"][300]
      );
    }
  }

  let fromTime = new Date(
    res_receive.data[res_receive.data.length - 1].created_at
  ).toISOString();
  sessionStorage.setItem("fromTime", fromTime);
  scrollTop();
}

async function hello() {
  if (!sessionStorage.getItem("token")) {
    location.href = endpointWEB;
    return;
  }

  let res = await InfoAPI(sessionStorage.getItem("token"));
  if (res.code !== 300) {
    location.href = endpointWEB;
    return;
  }

  sessionStorage.setItem("userID", res.data.id);
  $(".hello span").text(res.data.full_name);
}

async function helloRoom() {
  let res_count = await ChatCountMember(sessionStorage.getItem("token"));
  if (res_count.code !== 110) {
    console.log(res_count);
    return;
  }
  if (!sessionStorage.getItem("countMember")) {
    sessionStorage.setItem("countMember", "0");
  }

  let count = parseInt(res_count.data);

  if (sessionStorage.getItem("countMember") === "1" && count === 2) {
    showMessageLeft(
      "Hệ thống",
      "Có ai đó vừa vào phòng",
      window.elementaryColor["Silver"][100],
      window.elementaryColor["Orange"][700]
    );
    scrollTop();
  } else if (sessionStorage.getItem("countMember") === "2" && count === 1) {
    showMessageLeft(
      "Hệ thống",
      "Người nói chuyện với bạn vừa rời khỏi phòng",
      window.elementaryColor["Silver"][100],
      window.elementaryColor["Orange"][700]
    );
    scrollTop();
  } else if (sessionStorage.getItem("countMember") === "0" && count === 1) {
    showMessageLeft(
      "Hệ thống",
      "Phòng đang trống, chờ ai đó vào phòng",
      window.elementaryColor["Silver"][100],
      window.elementaryColor["Orange"][700]
    );
    scrollTop();
  } else if (sessionStorage.getItem("countMember") === "0" && count === 2) {
    showMessageLeft(
      "Hệ thống",
      "Phòng đang có ai đó, hãy nhắn tin để chào",
      window.elementaryColor["Silver"][100],
      window.elementaryColor["Orange"][700]
    );
    scrollTop();
  }

  sessionStorage.setItem("countMember", count.toString());
}

function polling() {
  setInterval(async () => {
    await receive();
    await helloRoom();
  }, 250);
}

$(async () => {
  await hello();
  prepare();

  // check user is free
  let res_is_free = await ChatIsFreeAPI(sessionStorage.getItem("token"));
  if (res_is_free.code === 900) {
    // find a new room for user to join
    let res_find = await ChatFindAPI(sessionStorage.getItem("token"), "empty");
    if (res_find.code !== 400) {
      console.log(res_find);
      return;
    }

    let res_join = await ChatJoinAPI(
      sessionStorage.getItem("token"),
      res_find.data.id
    );
    if (res_join.code !== 500) {
      console.log(res_join);
      return;
    }
  }

  polling();

  $(".leave").on("click", async () => {
    let res_leave = await ChatLeaveAPI(sessionStorage.getItem("token"));
    if (res_leave.code !== 600) {
      console.log(res_leave);
      return;
    }
    sessionStorage.clear();
    location.href = endpointWEB;
  });

  $("#btnNext").on("click", async () => {
    // next
    let status = "next";
    if (sessionStorage.getItem("search") === "1") {
      status = "gender";
    } else if (sessionStorage.getItem("search") === "2") {
      status = "birth";
    }
    let res_next = await ChatFindAPI(sessionStorage.getItem("token"), status);
    if (res_next.code !== 400) {
      console.log(res_next);
      return;
    }

    // leave
    let res_leave = await ChatLeaveAPI(sessionStorage.getItem("token"));
    if (res_leave.code !== 600) {
      console.log(res_leave);
      return;
    }

    // join
    let res_join = await ChatJoinAPI(
      sessionStorage.getItem("token"),
      res_next.data.id
    );
    if (res_join.code !== 500) {
      console.log(res_join);
      return;
    }

    sessionStorage.setItem("countMember", "0");
    $("#chat").html("");
  });

  $("#formChat").on("submit", async event => {
    event.preventDefault();

    let inputMessage = $("#inputMessage");
    let msg = inputMessage.val();
    if (msg === "") {
      return;
    }
    inputMessage.val("");

    let res_send = await ChatSendAPI(sessionStorage.getItem("token"), msg);
    if (res_send.code !== 700) {
      console.log(res_send);
    }
  });

  $(".editInfo").on("click", async () => {
    let res = await InfoAPI(sessionStorage.getItem("token"));
    if (res.code !== 300) {
      location.href = endpointWEB;
      return;
    }

    $("#inputFullNameEditInfo").val(res.data.full_name);
    $("#inputGenderEditInfo").val(res.data.gender);
    $("#inputBirthYearEditInfo").val(res.data.birth_year);

    $("#modalEditInfo").modal("show");
  });

  $(".editSearch").on("click", () => {
    $("#modalEditSearch").modal("show");
  });

  $("#formEditInfo").on("submit", async event => {
    event.preventDefault();

    let res = await EditInfoAPI(
      sessionStorage.getItem("token"),
      $("#inputFullNameEditInfo").val(),
      $("#inputGenderEditInfo").val(),
      parseInt($("#inputBirthYearEditInfo").val())
    );

    if (res.code !== 120) {
      $("#errEditInfo").text(res.message);
      return;
    }
    location.reload();
  });

  $("#formEditSearch").on("submit", event => {
    event.preventDefault();

    let search = $("input[name=editSearchRadio]:checked").val();
    sessionStorage.setItem("search", search);
    location.reload();
  });

  $("#resetEditSearch").on("click", () => {
    $("#checkGenderEditSearch").prop("checked", false);
    $("#checkBirthYearEditSearch").prop("checked", false);
    sessionStorage.removeItem("search");
  });
});

$(window).on("beforeunload", () => {
  ChatLeaveAPI(sessionStorage.getItem("token"));
});
