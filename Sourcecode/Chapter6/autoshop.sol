contract AutoShop {
   address[] autoAssets;
   function createChildContract(string brand, string model) public     payable {
      // insert check if the sent ether is enough to cover the car asset ...
      address newAutoAsset = new AutoAsset(brand, model, msg.sender);            
      autoAssets.push(newAutoAsset);   
   }
function getDeployedChildContracts() public view returns (address[]) {
      return autoAssets;
   }
}
contract AutoAsset {
string public brand;
    string public model;
    address public owner;
function AutoAsset(string _brand, string _model, address _owner) public 
{
      brand = _brand;
      model = _model;
      owner = _owner;
   }
}