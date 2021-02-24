pragma solidity ^0.4.5;
contract FuntionTest{
    function internalFunc() internal{}
    function externalFunc() external{}
    function callFunc(){
        //直接使用内部的方式调用
        internalFunc();
        //不能在内部调用一个外部函数，会报编译错误。
        //错误: 没有声明的标识符
        //externalFunc();
        //不能通过`external`的方式调用一个`internal`
        // 错误信息： 找不到成员函数"internalFunc" 或者不可见
        //this.internalFunc();
        //使用`this`以`external`的方式调用一个外部函数
        this.externalFunc();
    }
}
contract FunctionTest1{
    function externalCall(FuntionTest ft){
        //调用另一个合约的外部函数
        ft.externalFunc();
        //不能调用另一个合约的内部函数
        //错误信息：在合约FuntionTest里找不到成员函数"internalFunc"或者不可见
        //ft.internalFunc();
    }
}
