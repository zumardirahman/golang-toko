# golang-toko
Toko Online dengan Golang | Gorm, Mux, PostgreSQL, Bootstrap

Creted by Zumardi Rahman

Basic E-commerce Web App built with GO
## Features
- Home Page
- Register, Login, Profile User
- Products Page
- Single Product Page
- Shopping Cart
- Checkout
  - Shipping Cost Calculation
  - Payment Gateway
- Order History
- Admin Panel:
  - Dashboard
  - Manage Products
  - Manage Orders
  - Manage Users
  - Reports
  - Configuration

## Technology Stack
- Golang :
  - Gorilla Mux (https://github.com/gorilla/mux)
  - Gorm (http://gorm.io)
  - Render (https://github.com/unrolled/render)
- PostgreSQL
- Bootstrap

## Step Productions
#### Install Gorilla Mux
For Routing
    go get github.com/gorilla/mux

#### Install Gorm
For ORM
    go get -u gorm.io/gorm

#### Install Driver Postgres
For PostgreSQL DB Driver

    go get -u gorm.io/driver/postgres


#### Install Go .env
For Read File `.ENV`
    go get github.com/joho/godotenv

#### Install Render
For Render view : HTML 
    go get github.com/unrolled/render

#### Setup Struktur Direktori, GO Modules, dan Route
- MVC



## References
- https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121
- https://dasarpemrogramangolang.novalagung.com/