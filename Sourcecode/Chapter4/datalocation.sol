pragma solidity ^0.4.0;
contract DataLocation{
  uint valueType;
  mapping(uint => uint) public refrenceType;
  function changeMemory(){
    var tmp = valueType;
    tmp = 100;
  }
  function changeStorage(){
    var tmp = refrenceType;
    tmp[1] = 100;
  }
 function getAll() returns (uint, uint){
    return (valueType, refrenceType[1]);
  }
}
