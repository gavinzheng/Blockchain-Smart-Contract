pragma solidity >= 0.7.0 < 0.9.0;


contract AssemblyCodes {
    
    function addToEVM() public view returns(bool success){
        
        uint size;
        address addr;
        
        assembly {
            size := extcodesize(addr)
        }
        
        if(size > 0){
            return true;
        }else{
            return false;
        }
    }
    
    function addToEVM2() external view{
        
        bytes memory data = new bytes(10);
        
        //bytes32 dataB32 = new bytes32(data);
    }
}
