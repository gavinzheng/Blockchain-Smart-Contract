
pragma solidity >= 0.7.0 < 0.9.0;

/*
1. Create two interfaces which each hold separate function signatures which you must look up in theSolidity docs V2

2. One interface will be called UniswapV2Factory which will contain the function signature that gets pairs

3. Another interface will called UniswapV2Pair which will contain the function signature to get reserve values

4. Create a contract which contains addresses of the paired tokens you choose as well as the factory address

5. withing the contract make a function which successfully get the pair of your tokens and set it go an address

6. Use that address to get the reserve values and finally return the reserve values

7. Switch to the Miannet and connect MetaMask but DO NOT deploy the contract (or you will loose money!!!)

8. From the Mainnet manually go through the steps (5 and 6) of retrieving the values without deploying or spending money.

*/


interface UniswapV2Factory {
    
    function getPair(address tokenA, address tokenB) external view returns (address pair);
}

interface UniswapV2Pair {
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    
}

contract MyContract {
    
    // factory address
    address private factory = 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f;
    
    // Token addresses
    address private dai = 0x6B175474E89094C44Da98b954EedeAC495271d0F;
    address private weth = 0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2;
    
    
    function getReservToken() external view returns(uint, uint){
        address pair = UniswapV2Factory(factory).getPair(dai, weth);
        (uint reserve0, uint reserve1,) = UniswapV2Pair(pair).getReserves();
        return(reserve0, reserve1);
    }
}
