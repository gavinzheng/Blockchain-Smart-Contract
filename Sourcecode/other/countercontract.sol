contract counterContract {
  uint32 public counter;
  bool private stopped = false; //断路器的标志
  address private owner;  

  /**
    返回合约状态
  */
  modifier isNotStopped {
    require(!stopped, 'Contract is stopped.');
    _;
  }  

  /**
    检查调用者是否是合约所有者
  */
  modifier isOwner {
    require(msg.sender == owner, 'Sender is not owner.');
    _;
  }  
  constructor() public {
    counter = 0;
    owner = msg.sender;
  }  
  
  /**
   如果合约为活跃状态，则将计数器加2
   如果合约为终止状态，则什么都不做。因为这里有isNotStopped修饰符
   */
  function incrementCounter() isNotStopped public {
    counter += 2; 
  }  
  
  /**
    终止或者重开
  */
  function toggleContractStopped() isOwner public {
      stopped = !stopped;
  }
}
