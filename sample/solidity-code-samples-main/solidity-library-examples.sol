pragma solidity >= 0.7.0 < 0.9.0;


library Search {
    
    // We want this loop to return the index value of some interger we are searcing for
    function indexOfArray(uint[] storage self, uint value) public view returns(uint){
        for(uint i = 0; i < self.length; i++){
            if(self[i] == value){
                return i;
            }
        }
    }
}


contract Test{
    uint[] data;
    
    // insert value into the Array in the constructor
    constructor() public {
        data.push(1);
        data.push(2);
        data.push(3);
        data.push(4);
        data.push(5);
    }
    
    // Takes a value and connect to the Search library to use it's indexOfArray method to search an array and return a match value
    function isValuePresent(uint val) external view returns(uint){
        uint value = val;
        uint index = Search.indexOfArray(data, value);
        return index;
    }
    
}


library Search2 {
    
    // We want this loop to return the index value of some interger we are searcing for
    function indexOfArray(uint[] storage self, uint value) public view returns(uint){
        for(uint i = 0; i < self.length; i++){
            if(self[i] == value){
                return i;
            }
        }
    }
}


contract Test2{
    
    using Search2 for uint[];
    
    uint[] data;
    
    // insert value into the Array in the constructor
    constructor() public {
        data.push(1);
        data.push(2);
        data.push(3);
        data.push(4);
        data.push(5);
    }
    
    // Takes a value and connect to the Search library to use it's indexOfArray method to search an array and return a match value
    function isValuePresent() external view returns(uint){
        uint value = 4;
        uint index = data.indexOfArray(value);
        return index;
    }
    
}