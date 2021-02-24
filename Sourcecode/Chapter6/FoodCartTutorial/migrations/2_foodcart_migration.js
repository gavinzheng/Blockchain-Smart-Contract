var FoodCart = artifacts.require("./FoodCart.sol");

module.exports = function(deployer) {
  deployer.deploy(FoodCart);
};
