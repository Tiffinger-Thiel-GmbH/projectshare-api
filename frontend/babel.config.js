// this file is only required for Jest
module.exports = function(api) {
  api.cache(true);

  const presets = [
    "@babel/env",
    "@babel/react"
   ];

  const plugins = [
    ["@babel/plugin-transform-typescript", { isTSX: true }]
  ];

  return {
    presets,
    plugins
  };
};