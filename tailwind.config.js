module.exports = {
  content: [
    "./javascript/**/*.{js,svelte}",
    "./**/*.templ",
    "./node_modules/flowbite/**/*.js",
  ],
  // https://themes.ionevolve.com/
  daisyui: {
    themes: [
      {
        lightmode: {
          // Change to any existing daisyui theme or make your own
          ...require("daisyui/src/theming/themes")["cmyk"],
          // Edit styles if required
          primary: "white",
          secondary: "#DEFBFB",
          accent: "#FF5D73",
          neutral: "#919191",
          // "base-100": "",
          info: "#623CEA",
          success: "#87FF65",
          warning: "#FFC759",
          error: "#A30000",
        },
      },
      {
        darkmode: {
          // Change to any existing daisyui theme or make your own
          ...require("daisyui/src/theming/themes")["business"],
          // Edit styles if required
          primary: "black",
          secondary: "#1D1D1D",
          accent: "blue",
          neutral: "#494949",
          "base-100": "#031622",
          info: "#623CEA",
          success: "#80D569",
          warning: "#FFC759",
          error: "#A30000",
        },
      },
    ],
  },
  plugins: [require("daisyui"), require("flowbite/plugin")],
};
