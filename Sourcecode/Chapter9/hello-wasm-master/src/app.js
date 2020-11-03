async function handleGreetButtonClickEvent(event) {
    event.preventDefault();
    const { greet } = await import('../src/hello_world/pkg');
    greet(); //When greet() is called, .wasm will update the <h1> title in index.html
}

document.querySelector('#btn-greet').onclick = handleGreetButtonClickEvent;