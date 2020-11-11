module.exports = {
  ignorePatterns: ["/*.js", "/dist", "/coverage"],
  env: {
    node: true,
    es6: true
  },
  parserOptions: {
    project: `${__dirname}/tsconfig.json`,
    tsconfigRootDir: __dirname
  },
  rules: {
  }
};
