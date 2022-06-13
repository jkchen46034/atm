## ATM 
````
        ATM is an ATM machine implmented in Go.
````
## Installation
````
        Download from https://github.com/jkchen46034/atm
````
## Build
````
        $go mod tidy
        $go build -o atm 
````

## Execution

````
        jk@jk-HP:~/dev/golang/atm$ ./atm

        > balance
        Authorization required

        > history
        Authorization required

        > withdraw 10000
        Authorization required

        > deposit 2000
        Authorization required

        > logout
        No account is currently authorized

        > authorize 2859459814 738
        Authorization failed

        > authorize 2859459814 7386
        Account 2859459814 successfully authorized

        > withdraw 8900
        Amount dispensed: $80.00. Current balance: $20.24.

        > withdraw 4000
        Amount dispensed: $40.00. You have been charged overdraft fee of $5.00. Current balance: $-24.-76

        > withdraw 8000
        Your account is overdrawn! You may not make withdrawals at this time.

        > deposit 20000
        Current balance: $175.24.

        > deposit 3500
        Current balance: $210.24.

        > history
        2022-06-12 18:34:49 35.00 210.24
        2022-06-12 18:34:36 200.00 175.24
        2022-06-12 18:34:22 -40.00 -24.-76
        2022-06-12 18:33:58 -80.00 20.24

        > balance
        Current balance $210.24

        > logout
        Account 2859459814 logged out.

        > authorize 1434597300 4557
        Account 1434597300 successfully authorized

        > balance
        Current balance $90000.55

        > withdraw 2000000
        Amount dispensed: $10115.00. Current balance: $79885.55.

        > withdraw 2000000
        Unable to process your withdrawal at this time.

        > balance
        Current balance $79885.55

        > history
        2022-06-12 18:41:39 -10115.00 79885.55

        > logout
        Account 1434597300 logged out.

        > authorize 2859459814 7386
        Account 2859459814 successfully authorized

        > balance
        Current balance $210.24

        > withdraw 10000
        Unable to process your withdrawal at this time.

        > deposit 10000
        Current balance: $310.24.

        > withdraw 12000
        Amount dispensed: $100.00. Current balance: $210.24.

        > history
        2022-06-12 18:44:39 -100.00 210.24
        2022-06-12 18:44:20 100.00 310.24
        2022-06-12 18:34:49 35.00 210.24
        2022-06-12 18:34:36 200.00 175.24
        2022-06-12 18:34:22 -40.00 -24.-76
        2022-06-12 18:33:58 -80.00 20.24

        > logout
        Account 2859459814 logged out.

        > exit
        jk@jk-HP:~/dev/golang/atm$ 

````
## Test 
```
        jk@jk-HP:~/dev/golang/atm$ go test -v ./...

        === RUN   Test_Main_Demo
        --- PASS: Test_Main_Demo (0.00s)
        PASS
        ok  	takeoff.com/atm	0.007s
        === RUN   Test_Deposit_Success
        --- PASS: Test_Deposit_Success (0.00s)
        === RUN   Test_Deposit_NegativeAccount
        --- PASS: Test_Deposit_NegativeAccount (0.00s)
        === RUN   Test_GetHistory_0
        --- PASS: Test_GetHistory_0 (0.00s)
        === RUN   Test_GetHistory_1D
        --- PASS: Test_GetHistory_1D (0.00s)
        === RUN   TestLogin_Success
        --- PASS: TestLogin_Success (0.00s)
        === RUN   TestLogin_BadPIN
        --- PASS: TestLogin_BadPIN (0.00s)
        === RUN   TestLogin_BadAccountID
        --- PASS: TestLogin_BadAccountID (0.00s)
        === RUN   TestLogout_0
        --- PASS: TestLogout_0 (0.00s)
        === RUN   TestLogout_No_one_login
        --- PASS: TestLogout_No_one_login (0.00s)
        === RUN   Test_Withdraw
        --- PASS: Test_Withdraw (0.00s)
        === RUN   Test_Withdraw_OnlyIn20
        --- PASS: Test_Withdraw_OnlyIn20 (0.00s)
        === RUN   Test_Withdraw_ATM_Inventory_Not_Enough
        --- PASS: Test_Withdraw_ATM_Inventory_Not_Enough (0.00s)
        === RUN   Test_Withdraw_ATM_Inventory_is_0
        --- PASS: Test_Withdraw_ATM_Inventory_is_0 (0.00s)
        === RUN   Test_Withdraw_Account_Overdraft
        --- PASS: Test_Withdraw_Account_Overdraft (0.00s)
        === RUN   Test_Withdraw_Account_No_Fund
        --- PASS: Test_Withdraw_Account_No_Fund (0.00s)
        PASS
        ok  	takeoff.com/atm/handler	0.008s
        === RUN   TestAccount_Create
        --- PASS: TestAccount_Create (0.00s)
        === RUN   TestAccount_GetAccountID
        --- PASS: TestAccount_GetAccountID (0.00s)
        === RUN   TestAccount_GetAccountPIN
        --- PASS: TestAccount_GetAccountPIN (0.00s)
        === RUN   TestAccount_GetAccountBalance
        --- PASS: TestAccount_GetAccountBalance (0.00s)
        === RUN   TestAccount_Withdraw
        --- PASS: TestAccount_Withdraw (0.00s)
        === RUN   TestAccount_Deposit
        --- PASS: TestAccount_Deposit (0.00s)
        === RUN   TestAccount_GetHistory_0_Transaction
        --- PASS: TestAccount_GetHistory_0_Transaction (0.00s)
        === RUN   TestAccount_GetHistory_1_Transaction
        --- PASS: TestAccount_GetHistory_1_Transaction (0.00s)
        === RUN   TestAccount_GetHistory_Another_1_Transaction
        --- PASS: TestAccount_GetHistory_Another_1_Transaction (0.00s)
        === RUN   TestAccount_GetHistory_2_Transactions
        --- PASS: TestAccount_GetHistory_2_Transactions (0.00s)
        PASS
        ok  	takeoff.com/atm/model	(cached)
        jk@jk-HP:~/dev/golang/atm$ 

```
## Demo()
````
        jk@jk-HP:~/dev/golang/atm$ ./atm
        > demo

        Account 2859459814 successfully authorized
        Current balance $100.24
        Withdrawing $90.00
        Amount dispensed: $80.00. Current balance: $20.24.
        Depositing $30.00
        Current balance: $50.24.
        Withdrawing $40.00
        Amount dispensed: $40.00. Current balance: $10.24.
        Depositing $40.00
        Current balance: $50.24.
        Withdrawing $30.00
        Amount dispensed: $20.00. Current balance: $30.24.
        Withdrawing $200.00
        Amount dispensed: $200.00. You have been charged overdraft fee of $5.00. Current balance: $-174.-76
        Withdrawing $400.00
        Your account is overdrawn! You may not make withdrawals at this time.
        Depositing $400.00
        Current balance: $225.24.
        Withdrawing $90.00
        Amount dispensed: $80.00. Current balance: $145.24.
        Getting transaction history
        2022-06-12 16:22:47 -80.00 145.24
        2022-06-12 16:22:47 400.00 225.24
        2022-06-12 16:22:47 -200.00 -174.-76
        2022-06-12 16:22:47 -20.00 30.24
        2022-06-12 16:22:47 40.00 50.24
        2022-06-12 16:22:47 -40.00 10.24
        2022-06-12 16:22:47 30.00 50.24
        2022-06-12 16:22:47 -80.00 20.24

        Account 2859459814 logged out.
        Withdrawing $20.00
        Authorization required
        Getting balance
        Authorization required
        Account 1434597300 successfully authorized
        Current balance $90000.55
        Withdrawing $20000.00
        Amount dispensed: $10050.00. Current balance: $79950.55.
        Withdrawing $10000.00
        Unable to process your withdrawal at this time. 
        Getting transaction history
        2022-06-12 16:22:47 -10050.00 79950.55

        Account 1434597300 logged out.
        Account 2859459814 successfully authorized
        Withdrawing $20.00
        Unable to process your withdrawal at this time. 

        > exit
        jk@jk-HP:~/dev/golang/atm$ 
````

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
