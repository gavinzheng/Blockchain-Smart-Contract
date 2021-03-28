pragma solidity >=0.5.8;

contract Proxy {
  /**
   * The logic contract address.
   */
  address public implementation;

  /**
  * @dev This event will be emitted every time the implementation gets upgraded
  * @param implementation representing the address of the upgraded implementation
  */
  event Upgraded(address indexed implementation, string version);

  /**
   * Constructor.
   */
  constructor (address _implementation) public {
    require(_implementation != address(0), 'implementation must be valid');
    implementation = _implementation;
  }

  /**
   * @dev Set new logic contract address.
   */
  function setImplementation(address _implementation) public {
    require(_implementation != address(0), 'implementation must be valid');
    require(_implementation != implementation, 'already this implementation');

    implementation = _implementation;

    emit Upgraded(_implementation, version);
  }

  /**
  * @dev Fallback function allowing to perform a delegatecall to the given implementation.
  * This function will return whatever the implementation call returns.
  * (Credits: OpenZepellin)
  */
  function () payable external {
    address _impl = getImplementation();
    require(_impl != address(0), 'implementation not set');

    assembly {
      let ptr := mload(0x40)
      calldatacopy(ptr, 0, calldatasize)
      let result := delegatecall(gas, _impl, ptr, calldatasize, 0, 0)
      let size := returndatasize
      returndatacopy(ptr, 0, size)
      switch result
      case 0 { revert(ptr, size) }
      default { return(ptr, size) }
    }
  }
}

contract ExistingWithoutABI001  {
    
    address dc;
    
    function ExistingWithoutABI001(address _t) public {
        dc = _t;
    }
    
    function setA_ASM(uint _val) public returns (uint answer) {
        
        bytes4 sig = bytes4(keccak256("setA(uint256)"));
        assembly {
            // move pointer to free memory spot
            let ptr := mload(0x40)
            // put function sig at memory spot
            mstore(ptr,sig)
            // append argument after function sig
            mstore(add(ptr,0x04), _val)

            let result := call(
              15000, // gas limit
              sload(dc_slot),  // to addr. append var to _slot to access storage variable
              0, // not transfer any ether
              ptr, // Inputs are stored at location ptr
              0x24, // Inputs are 36 bytes long
              ptr,  //Store output over input
              0x20) //Outputs are 32 bytes long
            
            if eq(result, 0) {
                revert(0, 0)
            }
            
            answer := mload(ptr) // Assign output to answer var
            mstore(0x40,add(ptr,0x24)) // Set storage pointer to new space
        }
    }
}