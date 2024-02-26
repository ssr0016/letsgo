# migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
# migrate create -seq -ext=.sql -dir=./migrations create_users_table

# Rate and Limiting && IP-base Rate Limiting && Deleting old limiters
# go run ./cmd/api/ -limiter-burst=2
# go run ./cmd/api/ -limiter-enabled=false
# for i in {1..6}; do curl http://localhost:4000/v1/healthcheck; done

# Sending Shutdown Signals
# pgrep -l api
# pkill -SIGKILL api
# pkill -SIGTERM api
# ctrl + \

# migrate create -seq -ext .sql -dir ./migrations create_tokens_table