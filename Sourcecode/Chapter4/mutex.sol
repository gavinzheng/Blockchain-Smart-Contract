pragma solidity >0.4.24;
contract owned {
    constructor() public { owner = msg.sender; }
    address payable owner;
    // �����Լֻ����һ�����η����ǲ�û���õ�������������������Լ��ʹ��
    // ������ᱻ���뵽���η���������������`_;`֮��
    // ����ζ�������Լ�����ߵ���ֵ��������ǣ�����������ִ�У�����ͻ������쳣
    modifier onlyOwner {
        require(
            msg.sender == owner,
            "Only owner can call this function."
        );
        _;
    }
}
contract mortal is owned {
    // �����Լ��Owned��Լ�̳�`onlyOwner`���η�������`close`����
    // �𵽵�������ֻ��������/owner���ܵ���`close`
    function close() public onlyOwner {
        selfdestruct(owner);
    }
}
contract priced {
    // ���η��ɽ��ܲ���
    modifier costs(uint price) {
        if (msg.value >= price) {
            _;   // _��ʾ������
        }
    }
}
contract Register is priced, owned {
    mapping (address => bool) registeredAddresses;
    uint price;
    constructor(uint initialPrice) public { price = initialPrice; }
    // ����ע����ǣ�`payable` �ؼ����Ǳ���ģ����������Զ��ܾ��͸�����Ether
    function register() public payable costs(price) {
        registeredAddresses[msg.sender] = true;
    }
    function changePrice(uint _price) public onlyOwner {
        price = _price;
    }
}
// ����֪�����Ժ�Լ�Ĺ��������빥����Mutex��Լ�൱��������һ����Լ������
// Mutex��Լͨ�����η���������Լ����ֹ���빥����
contract Mutex {
    bool locked;
    modifier noReentrancy() {
        require(
            !locked,
            "Reentrant call."
        );
        locked = true;
        _;  // �����εĺ�����
        locked = false;
    }
    /// ��������Ǳ�Mutex�����������ġ�����ζ��`msg.sender.call`�ﲻ���ٵ���f 
/// �Ӷ���ֹ�����빥��
    /// `return 7`��佫7��Ϊ����ֵ��Ȼ����ִ����������η����`locked = false` 
    function f() public noReentrancy returns (uint) {
        (bool success,) = msg.sender.call("");
        require(success);
        return 7;
    }
}
