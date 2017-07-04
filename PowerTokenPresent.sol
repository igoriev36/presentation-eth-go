//start pragma OMIT
pragma solidity ^0.4.10;
//end pragma OMIT

contract Power {
//start basic OMIT

event ErrorEvent(string message, address meter, address user, uint256 last, uint256 current);

function newReading(uint256 newBalance, address user) MustBeMyMeter(user) {
    if (newBalance == balances[user].last_reading) return;
    if (newBalance > balances[user].last_reading) {
        balances[user].balance += newBalance - balances[user].last_reading;
        balances[user].last_reading = newBalance;
    } else {
        ErrorEvent("negative usage", msg.sender, user, balances[user].last_reading, newBalance);
    }
}

//end basic OMIT


//start fa OMIT
    struct reading  {
        address     meter_id;
        uint256     balance;
        uint256     last_reading;
    }

    /* This creates an array with all balances */
    mapping (address => reading) public balances;

    /* and cost of power */
    uint256   cost;
//end fa OMIT

//start mod OMIT
    /* This is an easy way to run a check */

modifier MustBeTeller() {
    if (msg.sender != teller_) throw;

    _;
}

modifier MustBeMyMeter(address user) {
     if (balances[user].meter_id != msg.sender) throw;

     _;
}
//end mod OMIT

    modifier MustBeTech() {
        if (msg.sender != tech_) throw;

        _;
    }


    modifier MustBeOwner() {
        if (msg.sender != owner) throw;

        _;
    }


//start constructor OMIT
address teller_;
address tech_;
address owner;

    function Power(
        address teller,
        address tech
        ) {
        owner = msg.sender;
        teller_ = teller;
        tech_ = tech;
    }
//end constructor OMIT

//start balance OMIT
    function balanceOf(address user) constant returns (uint256) {
        return balances[user].balance * cost;
    }
//end balance OMIT

    /* Send coins */
//start xfer OMIT
    function payBill(address user, uint256 value) MustBeTeller {
        balances[user].balance -= value / cost;
    }

//end xfer OMIT

//start cs OMIT
    /* Approve the account for operation */
    function enableAccount(address user, address meter,uint256  start) MustBeTech {
        balances[user].balance = 0;
        balances[user].meter_id = meter;
        balances[user].last_reading = start;
    }
//end cs OMIT

    function newOwner(address newOwner) MustBeOwner {
        owner = newOwner;
    }


}
