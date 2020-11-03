pragma solidity ^0.4.24;
contract simplexample {
	uint256 val;

	function setValue(uint256 para) public {
	val= para;
	}
	
	function getValue() public view returns(uint256)  {
    	return val;
  	}
}
