pragma solidity ^0.4.25;

import "https://github.com/CaptainJavaScript/Solidity/usingCaptainJS_v2.sol";

contract IamASeaman is usingCaptainJS {
    
    constructor() public {}
    
    function UseVoucher() public    {
        // let's save some ether :-)
        ActivateVoucher("MobyDick");
    }
    
    function () public payable {
        // accept some cash later ;-)
    }
    
    // query kraken to find out the USD price of ether
    function KrakenQuery() public {
        // now call Run of #ScriptIt
        Run(
            123, 
            "json:https://api.kraken.com/0/public/Ticker?pair=ETHUSD", // here comes the url of Kraken
            "result.XETHZUSD.a[0]", // here comes the json query 
            "", // no npm modules required
            1, // 1 slice = 10 seconds is required to perform the query
            100000, // use 100k gas units max for the return of the query Result
            DEFAULT_GAS_PRICE // use the default gas price as defined in usingCaptainJS
        );
        
        // throw an event
        emit KrakenQuery_Pending();
    }
    
    // fetch a successful result of the Kraken query
    function CaptainsResult(uint JobID, string result) external onlyCaptainsOrdersAllowed {
        // only the captain will be able to invoke this function
        // process the result now...
        //...
        // emit the event
        Flag = true;
        emit KrakenQuery_Success(result);
    }
    
    // fetch a failure 
    function CaptainsError(uint JobID, string ErrorMsg) external onlyCaptainsOrdersAllowed {
        // bla bla bla
        
        // emit the event
        Flag = true;
        emit KrakenQuery_Failure(ErrorMsg);
    }
    
    // we need a flag to force events submission
    bool Flag;
    
    // define events for pending, success and failure
    event KrakenQuery_Success(string Result);
    event KrakenQuery_Failure(string ErrorMsg);
    event KrakenQuery_Pending();
}
