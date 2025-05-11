# RUNNING THE APP
### 1. Copy the .env.example file to .env and fill in the required values.
```bash
cp .env.example .env
```
### 2. Build the Docker image and run the app.
```bash
docker compose up --build
```
wait for all the containers to start up, then run the following command to create the database and run migrations.
### 3. Running Migrations
```bash
docker compose exec app /bin/migrate
```

### 4. Running Mock Data
```bash
docker compose exec app /bin/mock
```