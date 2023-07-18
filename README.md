# go-ecom-yt
Go eCom YT is a powerful eCommerce platform developed using the GoLang programming language and the Gin framework. This repository serves as a comprehensive starter kit, offering essential features such as user management (signup, login), product browsing with search functionality, shopping cart management (add, remove, list), secure checkout process, and address management (add, edit, delete). With its clean codebase and modular structure, Go eCom YT provides an ideal foundation for building your own online store. Jumpstart your eCommerce journey with confidence using this flexible and scalable solution.

## Project Diagram
The project has the following diagram:
<div align="center">
    <img src="./readme_img/GoEcom-YT.drawio.svg"  height="400" alt="">
</div>

## File Structure

The project has the following file structure:

```
- main.go
- go.mod
- go.sum
- README.md
- LICENSE
- docker-compose.yaml
- controllers/
  - address.go
  - cart.go
  - controllers.go
- database/
  - cart.go
  - databasetup.go
- middleware/
  - middleware.go
- models/
  - models.go
- routes/
  - routes.go
- tokens/
  - tokensgen.go
```

## File Relationship

The project has the following file relationship:
<div align="center">
    <img src="./readme_img/go-ecom-yt-fr.drawio.svg"  height="400" alt="">
</div>

## Database Diagram
The project has the following database diagram:
<div align="center">
    <img src="./readme_img/go-ecom-yt.svg"height="500" alt="">
</div>

## Prerequisites
Before running the application, make sure you have the following prerequisites installed on your system:

- Go (version X.X.X)
- Docker (version X.X.X)

## Installation
To install and set up the project, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/your-project.git
   ```
2. Change into the project directory:
   ```bash
   cd your-project
   ```
3. Install the project dependencies:
   ```bash
   go mod download
   ```
4. Set up the database:
    ```bash
    docker-compose up -d
    ```
5. Build and run the application:
    ```bash
    go build main.go
    go run main.go
    ```

## Application Routers
- **SIGNUP FUNCTION API CALL (POST REQUEST)**

  ```
  http://localhost:8000/users/signup
  ```

  ```json
  {
    "first_name": "Md",
    "last_name": "Romi",
    "email": "mdromi14@gmail.com",
    "password": "pas123",
    "phone": "+8801793202"
  }
  ```

Response :"Successfully Signed Up!!"

- **LOGIN FUNCTION API CALL (POST REQUEST)**

  ```
  http://localhost:8000/users/login
  ```

  ```json
  {
    "email": "mdromi14@gmail.com",
    "password": "pas123"
  }
  ```

response will be like this
  ```json
  {
    "_id": "64b46006c6857a8bc5449ff1",
    "first_name": "Md",
    "last_name": "Romi",
    "password": "$2a$14$xokVkVJXwR3xq8uG42jA/uXoMS/hr5DPzVSEtDdnGadz/wu61aUdG",
    "email": "mdromi14@gmail.com",
    "phone": "+8801793202",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im1kcm9taTE0QGdtYWlsLmNvbSIsImZpcnN0X25hbWUiOiJNZCIsImxhc3RfbmFtZSI6IlJvbWkiLCJ1aWQiOiI2NGI0NjAwNmM2ODU3YThiYzU0NDlmZjEiLCJleHAiOjE2ODk2MjkwNjJ9.6tek3_MXd8yLKZED6YpZb3EQGFKzwXqTFXJVFnA7Hc8",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IiIsImZpcnN0X25hbWUiOiIiLCJsYXN0X25hbWUiOiIiLCJ1aWQiOiIiLCJleHAiOjE2OTAxNDc0NjJ9.TBk7aukEwp9w0U_Ume1br8rRM537E4q3oy_VpcrcD1w",
    "created_at": "2023-07-16T21:24:22Z",
    "updated_at": "2023-07-16T21:24:22Z",
    "user_id": "64b46006c6857a8bc5449ff1",
    "usercart": [],
    "address": [],
    "orders": []
  }
  ```

- **Admin add Product Function (POST REQUEST)**

  ```
  http://localhost:8000/admin/addproduct
  ```

  ```json
  {
    "product_name": "Alienware x15",
    "price": 2500,
    "rating": 10,
    "image": "alienware.jpg"
  }
  ```

Response : "Successfully added our Product Admin!!"

- **View all the Products in db GET REQUEST**

  ```
  http://localhost:8000/users/productview
  ```

Response

  ```json
  [
    {
        "Product_ID": "64b6a6724219114bbdbec147",
        "product_name": "Alienware x15",
        "price": 2500,
        "rating": 10,
        "image": "alienware.jpg"
    },
    {
        "Product_ID": "64b6a6914219114bbdbec149",
        "product_name": "Mac Pro",
        "price": 2500,
        "rating": 5,
        "image": "mac_pro.jpg"
    },
    {
        "Product_ID": "64b6a6c04219114bbdbec14b",
        "product_name": "giner ale",
        "price": 900,
        "rating": 5,
        "image": "gin.jpg"
    },
    {
        "Product_ID": "64b6a6cd4219114bbdbec14d",
        "product_name": "iphone 13",
        "price": 1700,
        "rating": 4,
        "image": "ipho.jpg"
    },
    {
        "Product_ID": "64b6a6db4219114bbdbec14f",
        "product_name": "whiskey",
        "price": 100,
        "rating": 7,
        "image": "whis.jpg"
    },
    {
        "Product_ID": "64b6a6e84219114bbdbec151",
        "product_name": "acer predator",
        "price": 3000,
        "rating": 10,
        "image": "acer.jpg"
    }
  ]
  ```

- **Search Product by regex function (GET REQUEST)**

defines the word search sorting
  ```
  http://localhost:8000/users/search?name=al
  ```

response:

  ```json
  [
    {
      "Product_ID": "616152fa9f29be942bd9df91",
      "product_name": "Alienware x15",
      "price": 1500,
      "rating": 10,
      "image": "1.jpg"
    },
    {
      "Product_ID": "616153039f29be942bd9df92",
      "product_name": "ginger Ale",
      "price": 300,
      "rating": 10,
      "image": "1.jpg"
    }
  ]
  ```

- **Adding the Products to the Cart (GET REQUEST)**

  ```
  http://localhost:8000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx
  ```

  Corresponding mongodb query

- **Removing Item From the Cart (GET REQUEST)**

  ```
  http://localhost:8000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx
  ```

- **Listing the item in the users cart (GET REQUEST) and total price**

  ```
  http://localhost:8000/listcart?id=xxxxxxuser_idxxxxxxxxxx
  ```

- **Addding the Address (POST REQUEST)**

  ```
  http://localhost:8000/addadress?id=user_id**\*\***\***\*\***
  ```

  The Address array is limited to two values home and work address more than two address is not acceptable

  ```json
  {
    "house_name": "white house",
    "street_name": "white street",
    "city_name": "washington",
    "pin_code": "332423432"
  }
  ```

- **Editing the Home Address(PUT REQUEST)**

  ```
  http://localhost:8000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx
  ```

- **Editing the Work Address(PUT REQUEST)**

  ```
  http://localhost:8000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx
  ```

- **Delete Addresses(GET REQUEST)**

  ```
  http://localhost:8000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx
  ```

  delete both addresses

- **Cart Checkout Function and placing the order(GET REQUEST)**

  After placing the order the items have to be deleted from cart functonality added

  ```
  http://localhost:8000?id=xxuser_idxxx
  ```

- **Instantly Buying the Products(GET REQUEST)**
  ```
  http://localhost:8000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx
  ```



## Contributing
Contributions to the project are welcome! If you have any suggestions, improvements, or bug fixes, please submit a pull request. For major changes, please open an issue first to discuss the changes.

When contributing to this repository, please follow the existing code style and conventions, and make sure to update the README with any necessary changes.

## License
The project is licensed under the MIT License. Feel free to use, modify, and distribute the code as per the terms of the license.