# Tennet-Wallet-App

# Instalation 
1. Clone repository from branch main
2. Set up .env file with local environment (according to .env.example)
3. Install Package
4. Run 'go run main.go'
### Package used
1. go get -u gorm.io/gorm (to insatll gorm for orm)
2. go get -u gorm.io/driver/mysql (to use mysql from gorm)
3. go get -u github.com/gin-gonic/gin (to install gin for framework)
4. go get -u golang.org/x/crypto/bcrypt (for hasing password)
5. go get -u github.com/golang-jwt/jwt/v4 (to create token for authentication)
6. go get github.com/joho/godotenv (to process .env)

# END POINTS
# USER API

## SIGNUP

---

### Method POST (http://localhost:3000/signup)

### Request Header

> none

### Request Body

> "Email" : "<asset_email>" <br>"Password" : "<asset_password>"

### Response (200)

> "message" : "<success_msg>" <br>"user" : "<asset_email>" <br>Header : <set_cookie>

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## SIGNIN

---

### Method POST (http://localhost:3000/signin)

### Request Header

> none

### Request Body

> "Email" : "<asset_email>" <br>"Password" : "<asset_password>"

### Response (200)

> "message": "<success_msg>" <br>"user": "<asset_email>" <br>Header : <set_cookie>

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

# WALLET API

## CREATE WALLET

---

### Method POST (http://localhost:3000/v1/wallet)

### Request Header

> <get_cookie>

### Request Body

> "Name" : "<asset_name>"

### Response (200)

> "message": "<success_msg>" <br>"data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## GET ALL WALLET

---

### Method GET (http://localhost:3000/v1/wallet)

### Request Header

> <get_cookie>

### Request Body

> none

### Response (200)

> "message": "<success_msg>" <br> "data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## GET A WALLET

---

### Method GET (http://localhost:3000/v1/wallet/:id)

### Request Header

> <get_cookie>

### Request Body

> none

### Response (200)

> "message": "<success_msg>" <br> "data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## UPDATE WALLET

---

### Method PUT (http://localhost:3000/v1/wallet/:id)

### Request Header

> <get_cookie>

### Request Body

> "Name" : "<asset_name>"

### Response (200)

> "message": "<success_msg>" <br> "data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## DELETE WALLET

---

### Method DELETE (http://localhost:3000/v1/wallet/:id)

### Request Header

> <get_cookie>

### Request Body

> none

### Response (200)

> "message": "<success_msg>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

# ASSET API

## CREATE ASSET

---

### Method POST (http://localhost:3000/v1/asset)

### Request Header

> <get_cookie>

### Request Body

> "Name" : "<asset_name>" <br> "Wallet_id" : "<asset_wallet_id>" <br> "Symbol" : "<asset_symbol>" <br> "Network" : "<asset_network>" <br> "Address" : "<asset_address>" <br> "Balance" : "<asset_balance>"

### Response (200)

> "message": "<success_msg>" <br>"data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## GET ALL ASSET

---

### Method GET (http://localhost:3000/v1/asset)

### Request Header

> <get_cookie>

### Request Body

> none

### Response (200)

> "message": "<success_msg>" <br> "data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## GET AN ASSET

---

### Method GET (http://localhost:3000/v1/asset/:id)

### Request Header

> <get_cookie>

### Request Body

> none

### Response (200)

> "message": "<success_msg>" <br> "data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## UPDATE ASSET

---

### Method PUT (http://localhost:3000/v1/wallet/:id)

### Request Header

> <get_cookie>

### Request Body

> "Name" : "<asset_name>" <br> "Wallet_id" : "<asset_wallet_id>" <br> "Symbol" : "<asset_symbol>" <br> "Network" : "<asset_network>" <br> "Address" : "<asset_address>" <br> "Balance" : "<asset_balance>"

### Response (200)

> "message": "<success_msg>" <br> "data": "<data_created>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

## DELETE ASSET

---

### Method DELETE (http://localhost:3000/v1/asset/:id)

### Request Header

> <get_cookie>

### Request Body

> none

### Response (200)

> "message": "<success_msg>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>

# TRANSACTION API

## CREATE TRANSACTION

---

### Method POST (http://localhost:3000/v1/transaction)

### Request Header

> <get_cookie>

### Request Body

> "Src_wallet_id" : "<asset_src_wallet_id>" <br> "Src_asset_id" : "<asset_src_asset_id>" <br> "Dest_wallet_id" : "<asset_Dest_wallet_id>" <br> "Dest_asset_id" : "<asset_Dest_asset_id>" <br> "Amount" : "<asset_amount>" <br> "Gas_fee" : "<asset_Gas_fee>" <br> "Total" : "<asset_total>"

### Response (200)

> "message": "<success_msg>" <br> "transaction": "<data_created_transaction>" <br> "source_asset" : "<data_updated_src_asset>" <br> "dest_asset" : "<data_updated_dest_asset>"

### Response (400 - Bad Request)

> "error": "<error_msg>"

---

<br>



