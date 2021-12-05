pragma solidity >= 0.7.0 < 0.9.0;

contract Oracle {
    address admin;
    uint public rand;
    
    
    constructor() public{
        admin = msg.sender;
    }
    
    function feedRand(uint _rand) external {
        require(msg.sender == admin);
        rand = _rand;
    }
}

contract GenerateRandomNumber {
    
    Oracle oracle;
    
    constructor(address oracleAddress){
        oracle = Oracle(oracleAddress);
    }
    
    function randomNumberGenerator (uint range) public view returns(uint){
        return uint(keccak256(abi.encodePacked(oracle.rand(), block.timestamp, block.difficulty, msg.sender))) % range;
    }
}