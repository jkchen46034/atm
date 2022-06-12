# BST 
atm is an ATM machine implmented in Go.

## Installation
```
    Download from git@github.com:jkchen46034/atm.git

    $go mod tidy
    $go build ./...
    $./atm

## Test 
```
        jk@jk-HP:~/dev/golang/atm$ go test -v ./...

        === RUN   Test_Integration_Test1
        --- PASS: Test_Integration_Test1 (0.00s)
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

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
