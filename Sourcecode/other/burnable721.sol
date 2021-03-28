pragma solidity ^0.4.24;
import "/openzeppelin-solidity/contracts/token/ERC721/ERC721Full.sol";
import "/openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "/openzeppelin-solidity/contracts/token/ERC721/ERC721Burnable.sol";
contract BoticsToken is ERC721Full, Ownable, ERC721Burnable{
  
  constructor() ERC721Full("Botics", "BTS")  public {}

function mint(address to, uint256 tokenId) public onlyOwner {
   _mint(to, tokenId);
  }

function _mint(address to) public onlyOwner{
   mint(to, totalSupply().add(1));
  }

function _burn(uint256 tokenId) public{
   burn(tokenId);
  }
}