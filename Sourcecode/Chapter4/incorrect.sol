pragma solidity ^0.4.21;
contract Incorrect{
  function Incorrect () public { // 构造函数
    this.test();
  }
  function test() public pure returns(uint256) {
    return 2;
  }
}
