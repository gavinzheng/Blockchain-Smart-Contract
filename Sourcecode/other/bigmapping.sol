pragma solidity ^0.4.24;
contract bigmapping {
    mapping(uint256 => Tuple) tuplemap;
    struct Tuple {
      uint256 a;
      uint256 b;
      uint256 c;
    }
    constructor () public {
      tuplemap [0x1].a = 0x7A;
      tuplemap [0x1].b = 0x7B;
      tuplemap [0x1].c = 0x7C;
    }
}
