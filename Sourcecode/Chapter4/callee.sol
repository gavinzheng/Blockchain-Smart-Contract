pragma solidity ^0.4.24;

contract Callee {
    uint[] public values;

    function getValue(uint initial) returns(uint) {
        return initial + 150;
    }
    function storeValue(uint value) {
        values.push(value);
    }
    function getValues() returns(uint) {
        return values.length;
    }
}
