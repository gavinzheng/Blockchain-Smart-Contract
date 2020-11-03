
======= testmapping.sol:testmapping =======
EVM assembly:
    /* "testmapping.sol":26:165  contract testmapping {... */
  mstore(0x40, 0x80)
    /* "testmapping.sol":96:162  constructor()  public {... */
  callvalue
    /* "--CODEGEN--":8:17   */
  dup1
    /* "--CODEGEN--":5:7   */
  iszero
  tag_1
  jumpi
    /* "--CODEGEN--":30:31   */
  0x0
    /* "--CODEGEN--":27:28   */
  dup1
    /* "--CODEGEN--":20:32   */
  revert
    /* "--CODEGEN--":5:7   */
tag_1:
    /* "testmapping.sol":96:162  constructor()  public {... */
  pop
    /* "testmapping.sol":150:154  0x5d */
  0x5d
    /* "testmapping.sol":127:134  itemmap */
  0x0
    /* "testmapping.sol":127:147  itemmap [0x6f0f9b4e] */
  dup1
    /* "testmapping.sol":136:146  0x6f0f9b4e */
  0x6f0f9b4e
    /* "testmapping.sol":127:147  itemmap [0x6f0f9b4e] */
  dup2
  mstore
  0x20
  add
  swap1
  dup2
  mstore
  0x20
  add
  0x0
  keccak256
    /* "testmapping.sol":127:154  itemmap [0x6f0f9b4e] = 0x5d */
  dup2
  swap1
  sstore
  pop
    /* "testmapping.sol":26:165  contract testmapping {... */
  dataSize(sub_0)
  dup1
  dataOffset(sub_0)
  0x0
  codecopy
  0x0
  return
stop

sub_0: assembly {
        /* "testmapping.sol":26:165  contract testmapping {... */
      mstore(0x40, 0x80)
      0x0
      dup1
      revert

    auxdata: 0xa165627a7a72305820f971ebf4a6772dcb0291e4e5f8f115c5e1b9974869f6311cc11938e3f2cddbd40029
}

Binary: 
6080604052348015600f57600080fd5b50605d600080636f0f9b4e81526020019081526020016000208190555060358060396000396000f3006080604052600080fd00a165627a7a72305820f971ebf4a6772dcb0291e4e5f8f115c5e1b9974869f6311cc11938e3f2cddbd40029
