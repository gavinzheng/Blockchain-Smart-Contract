
======= simpleexample.sol:simplexample =======
EVM assembly:
    /* "simpleexample.sol":26:210  contract simplexample {... */
  mstore(0x40, 0x80)
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
    /* "simpleexample.sol":26:210  contract simplexample {... */
  pop
  dataSize(sub_0)
  dup1
  dataOffset(sub_0)
  0x0
  codecopy
  0x0
  return
stop

sub_0: assembly {
        /* "simpleexample.sol":26:210  contract simplexample {... */
      mstore(0x40, 0x80)
      jumpi(tag_1, lt(calldatasize, 0x4))
      and(div(calldataload(0x0), 0x100000000000000000000000000000000000000000000000000000000), 0xffffffff)
      0x20965255
      dup2
      eq
      tag_2
      jumpi
      dup1
      0x55241077
      eq
      tag_3
      jumpi
    tag_1:
      0x0
      dup1
      revert
        /* "simpleexample.sol":132:207  function getValue() public view returns(uint256)  {... */
    tag_2:
      callvalue
        /* "--CODEGEN--":8:17   */
      dup1
        /* "--CODEGEN--":5:7   */
      iszero
      tag_4
      jumpi
        /* "--CODEGEN--":30:31   */
      0x0
        /* "--CODEGEN--":27:28   */
      dup1
        /* "--CODEGEN--":20:32   */
      revert
        /* "--CODEGEN--":5:7   */
    tag_4:
        /* "simpleexample.sol":132:207  function getValue() public view returns(uint256)  {... */
      pop
      tag_5
      jump(tag_6)
    tag_5:
      0x40
      dup1
      mload
      swap2
      dup3
      mstore
      mload
      swap1
      dup2
      swap1
      sub
      0x20
      add
      swap1
      return
        /* "simpleexample.sol":69:126  function setValue(uint256 para) public {... */
    tag_3:
      callvalue
        /* "--CODEGEN--":8:17   */
      dup1
        /* "--CODEGEN--":5:7   */
      iszero
      tag_7
      jumpi
        /* "--CODEGEN--":30:31   */
      0x0
        /* "--CODEGEN--":27:28   */
      dup1
        /* "--CODEGEN--":20:32   */
      revert
        /* "--CODEGEN--":5:7   */
    tag_7:
      pop
        /* "simpleexample.sol":69:126  function setValue(uint256 para) public {... */
      tag_8
      calldataload(0x4)
      jump(tag_9)
    tag_8:
      stop
        /* "simpleexample.sol":132:207  function getValue() public view returns(uint256)  {... */
    tag_6:
        /* "simpleexample.sol":172:179  uint256 */
      0x0
        /* "simpleexample.sol":197:200  val */
      sload
        /* "simpleexample.sol":132:207  function getValue() public view returns(uint256)  {... */
      swap1
      jump	// out
        /* "simpleexample.sol":69:126  function setValue(uint256 para) public {... */
    tag_9:
        /* "simpleexample.sol":112:115  val */
      0x0
        /* "simpleexample.sol":112:121  val= para */
      sstore
        /* "simpleexample.sol":69:126  function setValue(uint256 para) public {... */
      jump	// out

    auxdata: 0xa165627a7a7230582097c3213864b2674ababc34376c8a991b87535e4bdf6b63c72e55d93ee61da4730029
}

Binary: 
608060405234801561001057600080fd5b5060bf8061001f6000396000f30060806040526004361060485763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663209652558114604d57806355241077146071575b600080fd5b348015605857600080fd5b50605f6088565b60408051918252519081900360200190f35b348015607c57600080fd5b506086600435608e565b005b60005490565b6000555600a165627a7a7230582097c3213864b2674ababc34376c8a991b87535e4bdf6b63c72e55d93ee61da4730029
