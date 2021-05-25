const assert = require("assert").strict;
const VerifyPoC = artifacts.require("VerifyPoC");

contract("VerifyPoC", (accounts) => {
  it(`uintToStr() must working correct`, async () => {
    let pocContract = await VerifyPoC.deployed();
    let randomValue = (Math.random()*0xffffff) >>> 0;
    let result = await pocContract.uintToStr(randomValue);
    assert.equal(Buffer.from(result.substr(2), 'hex').toString(), randomValue.toString(), `Result wasn't ${randomValue}`);
  });

  it(`Message should be signed and serialized properly with ${accounts[0]}`, async () => {
    let pocContract = await VerifyPoC.deployed();
    // Message is bytes32
    let message = Buffer.from("Hello, I'm Chiro");
    let messageHexString = `0x${message.toString("hex")}`;

    let signature = await web3.eth.sign(messageHexString, accounts[0]);
    let sig = {
      r: `0x${signature.substr(2, 64)}`,
      s: `0x${signature.substr(66, 64)}`,
      v: `0x${signature.substr(-2)}`,
    };
    console.log(`\tMessage:\t${messageHexString} (${message.toString()})`);
    console.log(`\tSigner:\t\t${accounts[0]}`);
    console.log(`\tSignature:\t${signature}`);
    console.log(`\tR:\t\t${sig.r}`);
    console.log(`\tS:\t\t${sig.s}`);
    console.log(`\tV:\t\t${sig.v}`);
    // Call a method on smart contract to verify signature
    let result = await pocContract.verifySerialized( messageHexString, signature);
    assert.equal(result, accounts[0], `Signer wasn't ${accounts[0]}`);
  });


  it(`Message should be signed without serialized properly with ${accounts[0]}`, async () => {
    let pocContract = await VerifyPoC.deployed();
    // Message is bytes32
    let message = Buffer.from("Hello, I'm Chiro");
    let messageHexString = `0x${message.toString("hex")}`;

    let signature = await web3.eth.sign(messageHexString, accounts[0]);
    let sig = {
      r: `0x${signature.substr(2, 64)}`,
      s: `0x${signature.substr(66, 64)}`,
      v: `0x${signature.substr(-2)}`,
    };
    console.log(`\tMessage:\t${messageHexString} (${message.toString()})`);
    console.log(`\tSigner:\t\t${accounts[0]}`);
    console.log(`\tSignature:\t${signature}`);
    console.log(`\tR:\t\t${sig.r}`);
    console.log(`\tS:\t\t${sig.s}`);
    console.log(`\tV:\t\t${sig.v}`);
    // Call a method on smart contract to verify signature
    let result = await pocContract.verify( messageHexString, sig.r, sig.s, sig.v);
    assert.equal(result, accounts[0], `Signer wasn't ${accounts[0]}`);
  });
});
