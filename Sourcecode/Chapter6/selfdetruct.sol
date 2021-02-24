contract SelfDesctructionContract {
   public address owner;
      public string someValue;

   modifier ownerRestricted {
          require(owner == msg.sender);
        _;
      } 
// 构造函数
   function SelfDesctructionContract() {
      owner = msg.sender;
   }
// 简单的setter函数
   function setSomeValue(string value){
      someValue = value;
   } 
// 使用ownerRestricted修饰符来限定只有合约的所有者才能调用该函数
   function destroyContract() ownerRestricted {
     suicide(owner);
   }
}
