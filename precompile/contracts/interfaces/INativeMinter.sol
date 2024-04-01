// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

interface INativeMinter {

    /////////////////////////////////////// READ METHODS //////////////////////////////////////////

    // Returns the current owner of the contract
    function owner() view external returns (address);

    // Returns the current minter of the contract
    function minter() view external returns (address);

    ////////////////////////////////////// WRITE METHODS //////////////////////////////////////////

    // Mints the specified amount of tokens to the specified address
    function mint(address addr, uint256 amount) external returns (bool);

    // Burns the specified amount of tokens from the specified address
    function burn(address addr, uint256 amount) external returns (bool);

    // Sets a new minter for the contract
    function setMinter(address newMinter) external returns (bool);

    // Transfers the ownership of the contract to a new owner
    function transferOwnership(address newOwner) external returns (bool);

    // Renounces ownership of the contract
    function renounceOwnership() external returns (bool);

}
