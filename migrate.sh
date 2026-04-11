export $(grep -v '^#' .env | xargs)

migrate -path migrations -database "$DATABASE_URL" up