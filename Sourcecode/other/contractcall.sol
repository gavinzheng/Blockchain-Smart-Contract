pragma solidity ^0.4.18;
contract Deployed {
    
    function setA(uint) public returns (uint) {}
    
    function a() public pure returns (uint) {}
    
}
contract Existing  {
    
    Deployed dc;
    
    function Existing(address _t) public {
        dc = Deployed(_t);
    }
 
    function getA() public view returns (uint result) {
        return dc.a();
    }
    
    function setA(uint _val) public returns (uint result) {
        dc.setA(_val);
        return _val;
    }
    
}

contract ExistingWithoutABI  {
    
    address dc;
    
    function ExistingWithoutABI(address _t) public {
        dc = _t;
    }
    
    function setA_Signature(uint _val) public returns(bool success){
        require(dc.call(bytes4(keccak256("setA(uint256)")),_val));
        return true;
    }
}
