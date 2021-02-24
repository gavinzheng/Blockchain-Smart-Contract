pragma solidity >0.4.24;
contract owned {
    constructor() public { owner = msg.sender; }
    address payable owner;
    // 这个合约只定义一个修饰符但是并没有用到它。它可以在派生合约里使用
    // 函数体会被加入到修饰符定义里的特殊符号`_;`之后
    // 这意味着如果合约所有者调用值这个函数是，函数在这里执行，否则就会引发异常
    modifier onlyOwner {
        require(
            msg.sender == owner,
            "Only owner can call this function."
        );
        _;
    }
}
contract mortal is owned {
    // 这个合约从Owned合约继承`onlyOwner`修饰符并用在`close`函数
    // 起到的作用是只有所有者/owner才能调用`close`
    function close() public onlyOwner {
        selfdestruct(owner);
    }
}
contract priced {
    // 修饰符可接受参数
    modifier costs(uint price) {
        if (msg.value >= price) {
            _;   // _表示函数体
        }
    }
}
contract Register is priced, owned {
    mapping (address => bool) registeredAddresses;
    uint price;
    constructor(uint initialPrice) public { price = initialPrice; }
    // 必须注意的是：`payable` 关键字是必须的，否则函数会自动拒绝送给它的Ether
    function register() public payable costs(price) {
        registeredAddresses[msg.sender] = true;
    }
    function changePrice(uint _price) public onlyOwner {
        price = _price;
    }
}
// 我们知道，对合约的攻击有重入攻击。Mutex合约相当于设置了一个合约互斥锁
// Mutex合约通过修饰符来保护合约，防止重入攻击。
contract Mutex {
    bool locked;
    modifier noReentrancy() {
        require(
            !locked,
            "Reentrant call."
        );
        locked = true;
        _;  // 被修饰的函数体
        locked = false;
    }
    /// 这个函数是被Mutex互斥锁保护的。这意味着`msg.sender.call`里不能再调用f 
/// 从而防止了重入攻击
    /// `return 7`语句将7设为返回值，然后再执行语句在修饰符里的`locked = false` 
    function f() public noReentrancy returns (uint) {
        (bool success,) = msg.sender.call("");
        require(success);
        return 7;
    }
}
