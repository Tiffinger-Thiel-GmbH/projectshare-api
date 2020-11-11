module.exports = {
  ignorePatterns: ["/*.js", "/dist"],
  extends: [
    'plugin:react/recommended',
    'prettier/react'
  ],
  plugins: [
    'react-hooks'
  ],
  parserOptions: {
    project: `${__dirname}/tsconfig.json`,
    tsconfigRootDir: __dirname
  },
  rules: {
    'react/prop-types': 'off',
    'react/jsx-fragments': 'warn',
    'react/no-unescaped-entities': 'off',
    'react/self-closing-comp': 'warn',
    'react-hooks/rules-of-hooks': 'error',
    'react-hooks/exhaustive-deps': 'warn'
  },
  settings: {
    react: {
      version: 'detect' // Tells eslint-plugin-react to automatically detect the version of React to use
    }
  }
};
