/// <reference types="react-scripts" />

declare module '*.yaml' {
  const value: any;
  export default value;
}

declare module 'url:*' {
  const value: string;
  export default value;
}
