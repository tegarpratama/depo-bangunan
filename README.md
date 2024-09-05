# Project Installation

1. Clone the project
   `git clone https://github.com/tegarpratama/depo-bangunan.git`

2. Running docker compose
   `docker-compose up --build -d`

3. Migrate & seed data
   `docker exec -it depo-bangunan-app-1 sh`
   `go run migrate/migrate.go`

4. Test API
   - Hit API login:
     `http://localhost:8080/api/auth/login`
   - payload:
   ```
       {
           "email": "admin@gmail.com",
           "password": "password"
       }
   ```

# API Docs

Open url: http://127.0.0.1:8080/swagger/index.html#/
