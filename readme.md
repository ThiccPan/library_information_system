# Library Information System

A library information system application developed using golang with echo framework and postgresql

start running the application in your local machine using this following steps:
1. clone the repository
```bash
git clone git@github.com:ThiccPan/warehouse-system.git
```
2. step into the project directory
```bash
cd library_information_system
```
3. configure your .env file
```bash
# db setup example
DB_CONNECTION=pgsql
DB_HOST=dbhost
DB_PORT=5433
DB_DATABASE=dbname
DB_USERNAME=yourusername
DB_PASSWORD=yourpassword
JWT_SECRET=secret
```
4. run your container script
```bash
docker-compose up -d --build
```
5. configure your application key
```bash
docker exec warehouse_system php artisan key:generate
```
6. done! you can start hitting the application api

Documentation:
- API documentation: [postman link](https://documenter.getpostman.com/view/23637484/2sAXxJgu76)