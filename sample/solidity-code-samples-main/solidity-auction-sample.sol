pragma solidity >= 0.7.0 < 0.9.0;

// An Auction contract that transfers tokens to the beneficiary
contract Auction{
    
    address payable public beneficiary;
    uint public auctionEndTime;
    
    // current state of auctionEndTime
    address public highestBidder;
    uint public highestBid;
    bool ended;
    
    mapping(address => uint) pendingReturns;
    
    event highestBidIncreased(address bidder, uint amount);
    event auctionEnded (address winner, uint amount);
    
    constructor (uint _biddingtime, address payable _beneficiary){
        beneficiary = _beneficiary;
        auctionEndTime = block.timestamp + _biddingtime;
    }
    
    //
    function bid() public payable{
        if(block.timestamp > auctionEndTime) revert('The aucion has ended');
        
        if(msg.value <= highestBid) revert('Sorry, the bid not high enough');
        
        if(highestBid != 0){
            pendingReturns[highestBidder] += highestBid;
        }
        
        highestBidder = msg.sender;
        highestBid = msg.value;
        
        emit highestBidIncreased(msg.sender, msg.value);
    }
    
    function withdraw() public payable returns(bool){
        uint amount = pendingReturns[msg.sender];
        
        if(amount > 0){
            pendingReturns[msg.sender] = 0;
        }
        
        if(!payable(msg.sender).send(amount)){
            pendingReturns[msg.sender] = amount;
        }
            return true;
    }
    
    function auctionEnd() public{
        if(block.timestamp < auctionEndTime) revert('The auction has not ended yet');
        if(ended) revert('The acution already ended');
        ended = true;
        emit auctionEnded (highestBidder, highestBid);
        beneficiary.transfer(highestBid);
    }
}