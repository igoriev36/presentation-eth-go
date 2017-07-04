pragma solidity ^0.4.10;


contract Power {

struct reading  {
    address     meter_id;
    uint256     balance;
    uint256     last_reading;
}

mapping (address => reading) public balances;
uint256   cost;  // cost of the power unit


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

modifier MustBeMyMeter(address user) {
     if (balances[user].meter_id != msg.sender) throw;

     _;
}

modifier MustBeTeller() {
    if (msg.sender != teller_) throw;

    _;
}


modifier MustBeTech() {
    if (msg.sender != tech_) throw;

    _;
}


modifier MustBeOwner() {
    if (msg.sender != owner) throw;

    _;
}


    function balanceOf(address user) constant returns (uint256) {
        return balances[user].balance * cost;
    }

    function payBill(address user, uint256 value) MustBeTeller {
        balances[user].balance -= value / cost;
    }


    /* Approve the account for operation */
    function enableAccount(address user, address meter,uint256  start) MustBeTech {
        balances[user].balance = 0;
        balances[user].meter_id = meter;
        balances[user].last_reading = start;
    }


    function newOwner(address newOwner) MustBeOwner {
        owner = newOwner;
    }


}
