pragma solidity >= 0.7.0 < 0.9.0;

contract ledgerBalance {
    
    mapping(address => uint) balance;
    
    function updateBalance(uint newBalance) public {
        balance[msg.sender] = newBalance;
    }
} 


contract Updated {
    
    function updatesBalance () public {
        ledgerBalance ledgerBalance = new ledgerBalance();
        
        ledgerBalance.updateBalance(30);
    }
}

contract simpleStorage {
    
    uint storedData;
    
    function set(uint x) public {
        //storedData = x;
        //storedData = block.difficulty;
        //storedData = block.timestamp;
         storedData = block.number;
    }
    
    function get() public view returns(uint){
        return storedData;
    }
}
