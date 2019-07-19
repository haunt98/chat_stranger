// https://elementary.io/docs/human-interface-guidelines#color
window.elementaryColor = {
  Orange: {
    100: "#ffc27d",
    500: "#f37329",
    700: "#cc3b02"
  },
  Blueberry: {
    100: "#8cd5ff",
    300: "#64baff",
    500: "#3689e6",
    700: "#0d52bf",
    900: "#002e99"
  },
  Grape: {
    100: "#e4c6fa",
    900: "#452981"
  },
  Silver: {
    100: "#fafafa",
    300: "#d4d4d4"
  },
  Slate: {
    700: "#273445",
    900: "#0e141f"
  },
  Black: {
    700: "#1a1a1a"
  }
};

$(() => {
  let btn_primary = $(".btn_primary");
  btn_primary.css("color", window.elementaryColor["Silver"][100]);
  btn_primary.css("backgroundColor", window.elementaryColor["Blueberry"][500]);
  btn_primary.css("borderColor", window.elementaryColor["Blueberry"][500]);

  // link in log in
  $(".card a").css("color", window.elementaryColor["Blueberry"][700]);

  // title chat stranger
  $(".title").css("color", window.elementaryColor["Orange"][500]);

  let btn_next = $("#btnNext");
  btn_next.css("color", window.elementaryColor["Silver"][100]);
  btn_next.css("backgroundColor", window.elementaryColor["Orange"][700]);
  btn_next.css("borderColor", window.elementaryColor["Orange"][700]);

  // background
  $("body").css("backgroundColor", window.elementaryColor["Slate"][700]);
  $(".card").css("backgroundColor", window.elementaryColor["Silver"][100]);
  $("#chat").css("backgroundColor", window.elementaryColor["Silver"][100]);
});
