pragma solidity >= 0.7.0 < 0.9.0;

contract Fallback {
    
    event log(uint gas);
    
    fallback() external payable{
        
        emit log(gasleft());
    }
    
    function getBalance() public view returns(uint){
        return address(this).balance;
    }
}


contract SendToFallback {
    
    function transferToFallback(address payable _to) public payable{
        
        _to.transfer(msg.value);
    }
    
    function callFallBack(address payable _to) public payable {
        (bool sent,) = _to.call{value:msg.value}('');
        require(sent, 'Failed to send! ');
    }
}