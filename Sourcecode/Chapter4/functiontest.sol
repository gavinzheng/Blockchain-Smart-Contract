pragma solidity ^0.4.5;
contract FuntionTest{
    function internalFunc() internal{}
    function externalFunc() external{}
    function callFunc(){
        //ֱ��ʹ���ڲ��ķ�ʽ����
        internalFunc();
        //�������ڲ�����һ���ⲿ�������ᱨ�������
        //����: û�������ı�ʶ��
        //externalFunc();
        //����ͨ��`external`�ķ�ʽ����һ��`internal`
        // ������Ϣ�� �Ҳ�����Ա����"internalFunc" ���߲��ɼ�
        //this.internalFunc();
        //ʹ��`this`��`external`�ķ�ʽ����һ���ⲿ����
        this.externalFunc();
    }
}
contract FunctionTest1{
    function externalCall(FuntionTest ft){
        //������һ����Լ���ⲿ����
        ft.externalFunc();
        //���ܵ�����һ����Լ���ڲ�����
        //������Ϣ���ں�ԼFuntionTest���Ҳ�����Ա����"internalFunc"���߲��ɼ�
        //ft.internalFunc();
    }
}
