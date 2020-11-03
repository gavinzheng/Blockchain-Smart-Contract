# hello-wasm
Source code based on the Medium article "WebAssembly -Part II.a | Rust with WASM"

## Pre-Requisites
1. Install NodeJs. https://nodejs.org/en/download/
2. Install the Rust toolchain (for compiling programs to WebAssembly). https://www.rust-lang.org/tools/install
3. Install Wasm-Pack. https://rustwasm.github.io/wasm-pack/installer/

## Getting Started
1. In the Workspace directory run the command npm install
2. In the Workspace directory run the command npm start or npm run start-in-prod-mode*.

*Npm Start compiles the Rust code in debug mode and starts webpack-dev-server, whereas npm run start-in-prod-mode does the same but compiles the Rust code in release code. The Wasm generated in release mode is more performant of course.

## Features
* When running webpack-dev-server all Rust (.rs) files are watched. That means you can make changes to your rust code and see them immediately on the browswer!

* Rust can talk to the browser window / DOM! Inside the project, refer to the src/hello_world folder where you will find an example of Rust being integrated with the window and DOM.

## Further Reading
For more information on building web applications with Rust and WebAssembly, refer to the following sources:

* https://rustwasm.github.io/docs/book/
* https://developer.mozilla.org/en-US/docs/WebAssembly/Rust_to_wasm

## Authors
Francisco Vilches - https://github.com/fvilches17
