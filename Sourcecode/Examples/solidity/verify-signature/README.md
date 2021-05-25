# A better code to verify Ethereum signature

I read this article, [Signatures: Verifying Messages in Solidity](https://blog.ricmoo.com/verifying-messages-in-solidity-50a94f82b2ca). It's such a good article but this code made me uncomfortable, It's too wasting and dirty then I implement mine so you could do compare.

## Message signing in Ethereum

Almost messages on Ethereum will be sign with a given format:

```
"\x19Ethereum Signed Message:\n" + message.length + message
```

E.g: "\x19Ethereum Signed Message:\n16Hello, I'm Chiro"

This prefix `"\x19Ethereum Signed Message:\n"` was used to make sure signed message won't be **"reused"** somewhere else.

## Improvement

As you see, the only thing that prevent us to archive verification onchain is `message.length`. You need to encode a number to string, here are code of @RicMoo:

```js
contract Verifier {
  // Returns the address that signed a given string message
  function verifyString(
    string memory message,
    uint8 v,
    bytes32 r,
    bytes32 s
  ) public pure returns (address signer) {
    // The message header; we will fill in the length next
    string memory header = '\x19Ethereum Signed Message:\n000000';
    uint256 lengthOffset;
    uint256 length;
    assembly {
      // The first word of a string is its length
      length := mload(message) // The beginning of the base-10 message length in the prefix
      lengthOffset := add(header, 57)
    } // Maximum length we support
    require(length <= 999999); // The length of the message's length in base-10
    uint256 lengthLength = 0; // The divisor to get the next left-most message length digit
    uint256 divisor = 100000; // Move one digit of the message length to the right at a time
    while (divisor != 0) {
      // The place value at the divisor
      uint256 digit = length / divisor;
      if (digit == 0) {
        // Skip leading zeros
        if (lengthLength == 0) {
          divisor /= 10;
          continue;
        }
      } // Found a non-zero digit or non-leading zero digit
      lengthLength++; // Remove this digit from the message length's current value
      length -= digit * divisor; // Shift our base-10 divisor over
      divisor /= 10;

      // Convert the digit to its ASCII representation (man ascii)
      digit += 0x30; // Move to the next character and write the digit
      lengthOffset++;
      assembly {
        mstore8(lengthOffset, digit)
      }
    } // The null string requires exactly 1 zero (unskip 1 leading 0)
    if (lengthLength == 0) {
      lengthLength = 1 + 0x19 + 1;
    } else {
      lengthLength += 1 + 0x19;
    } // Truncate the tailing zeros from the header
    assembly {
      mstore(header, lengthLength)
    } // Perform the elliptic curve recover operation
    bytes32 check = keccak256(abi.encodePacked(header, message));
    return ecrecover(check, v, r, s);
  }
}
```

The basic concept could be simplify:

```js
function uintToStr(value) {
  let result = [];
  for (let i = value; i > 0; i = (i / 10) >>> 0) {
    result.push(String.fromCharCode((i % 10) + 0x30));
  }
  return result.reverse().join('');
}
```

Radix is `10`, so we need modulo for `10` to get the digit and add the result with `0x30` to get its ASCII code.

E.g:
123 % 10 = 3
3 + 0x30 = 0x33 (ASCII of "3")

Here are my implementation without place holder: 

```js
//SPDX-License-Identifier: MIT
pragma solidity ^0.6.4;

contract VerifyPoC {
    function verifySerialized(bytes memory message, bytes memory signature) public pure returns (address) {
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            // Singature need to be 65 in length
            // if (signature.length !== 65) revert();
            if iszero(eq(mload(signature), 65)) {
                revert(0, 0)
            }
            // r = signature[:32]
            // s = signature[32:64]
            // v = signature[64]
            r := mload(add(signature, 0x20))
            s := mload(add(signature, 0x40))
            v := byte(0, mload(add(signature, 0x60)))
            // Invalid v value, for Ethereum it's only possible to be 27, 28 and 0, 1 in legacy code
            if lt(v, 27) {
                v := add(v, 27)
            }
            if iszero(or(eq(v, 27), eq(v, 28))) {
                revert(0, 0)
            }
        }

        // Get hashes of message with Ethereum proof prefix
        bytes32 hashes = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n", uintToStr(message.length), message));

        return ecrecover(hashes, v, r, s);
    }

    function verify(bytes memory message, bytes32 r, bytes32 s, uint8 v) public pure returns (address) {
        if(v < 27) {
            v += 27;
        }
        // V must be 27 or 28
        require(v == 27 || v == 28, "Invalid v value");
        // Get hashes of message with Ethereum proof prefix
        bytes32 hashes = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n", uintToStr(message.length), message));

        return ecrecover(hashes, v, r, s);
    }

    function uintToStr(uint256 value) public pure returns (bytes memory result) {
        assembly {
            switch value
                case 0 {
                    // In case of 0, we just return "0"
                    result := mload(0x20)
                    // result.length = 1
                    mstore(result, 0x01)
                    // result = "0"
                    mstore(add(result, 0x20), 0x30)
                }
                default {
                    let length := 0x0
                    // let result = new bytes(32)
                    result := mload(0x20)

                    // Get length of render number
                    // for (let v := value; v > 0; v = v / 10)
                    for { let v := value } gt(v, 0x00) { v := div(v, 0x0a) } {
                        length := add(length, 0x01)
                    }

                    // We're only support number with 32 digits
                    // if (length > 32) revert();
                    if gt(length, 0x20) {
                        revert(0, 0)
                    }

                    // Set length of result
                    mstore(result, length)

                    // Start render result
                    // for (let v := value; length > 0; v = v / 10)
                    for { let v := value } gt(length, 0x00) { v := div(v, 0x0a) } {
                        // result[--length] = 48 + (v % 10)
                        length := sub(length, 0x01)
                        mstore8(add(add(result, 0x20), length), add(0x30, mod(v, 0x0a)))
                    }
                }
        }
    }
}
```

Result of `truffle test`:

```
  Contract: VerifyPoC
    ✓ uintToStr() must working correct
        Message:        0x48656c6c6f2c2049276d20436869726f (Hello, I'm Chiro)
        Signer:         0x3624dDc1940f71f8a2b33df0f85fF19f5114310d
        Signature:      0xd92905b550ef4a044d890f15e57fd58e12738b4233921fb6de54ce3e4df46c800f568e29c6b521e19c80711a563696b9478ea05d65c216692df94aea7006bdfc01
        R:              0xd92905b550ef4a044d890f15e57fd58e12738b4233921fb6de54ce3e4df46c80
        S:              0x0f568e29c6b521e19c80711a563696b9478ea05d65c216692df94aea7006bdfc
        V:              0x01
    ✓ Message should be sign properly with 0x3624dDc1940f71f8a2b33df0f85fF19f5114310d
```

The result and whole code could be found here: [Improving Ethereum signature verification](https://github.com/chiro-hiro/examples/tree/master/solidity/verify-signature)

## Conclusion

- My implementation cost `5846 Gas`, RicMoo's implementation cost `6440 Gas`
- New feature to verify serialized signature cost `6160 Gas`
- In theory, It could verify message with length equal to 10^32-1