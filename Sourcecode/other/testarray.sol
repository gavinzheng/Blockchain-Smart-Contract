pragma solidity ^0.4.17;

contract Test {
    
    struct structUser {
        uint value;
        uint index;
        bool exists;
    }

    mapping(address => structUser) public arrayStructs;
    
    address[] public addressIndexes;
    
    function addAddress(uint _val) public returns (bool){
        // if user exists, add _val
        if (arrayStructs[msg.sender].exists > true) {
            arrayStructs[msg.sender].value += _val;
        }
        else {
            // else its new user
            addressIndexes.push(msg.sender);
            arrayStructs[msg.sender].value = _val;
            arrayStructs[msg.sender].index = addressIndexes.length-1;
            arrayStructs[msg.sender].exists = true;
        }
        return true;
    }
    
    function deleteAddress() public returns (bool) {
        // 如果地址已存在
        if (arrayStructs[msg.sender].exists) {
            structUser memory deletedUser = arrayStructs[msg.sender];
            // 如果index不是最后一个
            if (deletedUser.index != addressIndexes.length-1) {
                // delete addressIndexes[deletedUser.index];
                // last strucUser
                address lastAddress = addressIndexes[addressIndexes.length-1];
                addressIndexes[deletedUser.index] = lastAddress;
                arrayStructs[lastAddress].index = deletedUser.index; 
            }
            delete arrayStructs[msg.sender];
            addressIndexes.length--;
            return true;
        }
    }
    
    function getAddresses() public view returns (address[]){
        return addressIndexes;    
    }
    
    function getTotalValue() public view returns (uint) {
        uint arrayLength = addressIndexes.length;
        uint total = 0;
        for (uint i=0; i<arrayLength; i++) {
            total += arrayStructs[addressIndexes[i]].value;
        }
        return total;
    }
    
    function getTotalUsers() public view returns (uint) {
        return addressIndexes.length;
    }
}
