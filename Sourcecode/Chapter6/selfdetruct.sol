contract SelfDesctructionContract {
   public address owner;
      public string someValue;

   modifier ownerRestricted {
          require(owner == msg.sender);
        _;
      } 
// ���캯��
   function SelfDesctructionContract() {
      owner = msg.sender;
   }
// �򵥵�setter����
   function setSomeValue(string value){
      someValue = value;
   } 
// ʹ��ownerRestricted���η����޶�ֻ�к�Լ�������߲��ܵ��øú���
   function destroyContract() ownerRestricted {
     suicide(owner);
   }
}
