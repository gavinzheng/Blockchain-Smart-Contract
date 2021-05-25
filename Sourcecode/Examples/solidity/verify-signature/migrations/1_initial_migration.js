const Migrations = artifacts.require("Migrations");
const VerifyPoC = artifacts.require("VerifyPoC");

module.exports = function (deployer) {
  deployer.deploy(Migrations);
  deployer.deploy(VerifyPoC);
};
