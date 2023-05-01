import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import typescript from '@rollup/plugin-typescript';
import { terser } from 'rollup-plugin-terser';

//const isProduction = !process.env.ROLLUP_WATCH;

export default {
  input: 'src/main.ts',
  output: [
    {
      file: 'dist/emo.min.cjs',
      format: 'cjs',
      plugins: [terser(),]
    },
    {
      file: 'dist/emo.mjs',
      format: 'esm'
    },
    {
      file: 'dist/emo.min.js',
      format: 'iife',
      name: '$emo',
      plugins: [terser()]
    }],
  plugins: [
    typescript(),
    resolve({
      jsnext: true,
      main: true,
      browser: true,
    }),
    commonjs(),
  ],
};