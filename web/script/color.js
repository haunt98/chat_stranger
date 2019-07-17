// https://elementary.io/docs/human-interface-guidelines#color
window.elementaryColor = {
  Orange: {
    500: "#f37329"
  },
  Blueberry: {
    100: "#8cd5ff",
    300: "#64baff",
    500: "#3689e6",
    700: "#0d52bf",
    900: "#002e99"
  },
  Grape: {
    100: "#e4c6fa"
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
  btn_primary.css("background-color", window.elementaryColor["Blueberry"][500]);
  btn_primary.css("border-color", window.elementaryColor["Blueberry"][500]);

  $("form a").css("color", window.elementaryColor["Blueberry"][500]);

  $(".title").css("color", window.elementaryColor["Blueberry"][500]);
});
