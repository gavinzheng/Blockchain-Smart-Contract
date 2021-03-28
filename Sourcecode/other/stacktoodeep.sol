// SPDX-License-Identifier: MIT
pragma solidity 0.7.1;

contract StackTooDeepTest1 {
    function addUints(
        uint256 a,uint256 b,uint256 c,uint256 d,uint256 e,uint256 f,uint256 g,uint256 h,uint256 i
    ) external pure returns(uint256) {

        //return a+b+c+d+e+f+g+h;
        return a+2+3+4+5+6+7+8;
    }
}