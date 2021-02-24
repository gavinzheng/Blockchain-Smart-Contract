package main

import (
    "syscall/js"
    "strconv"
)

// func add(i []js.Value) {
//     js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
//     println(js.ValueOf(i[0].Int() + i[1].Int()).String())
// }
func add(this js.Value, i []js.Value) interface{}{
    value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
    value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

    js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
    return nil
}
// func subtract(i []js.Value) {
//     js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
//     println(js.ValueOf(i[0].Int() - i[1].Int()).String())
// }
func subtract(this js.Value, i []js.Value) interface{} {
//func subtract(i []js.Value) {
    value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
    value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

    js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
    return nil
}
func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))
    js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}