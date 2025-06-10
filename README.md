# How to run? Get in cmd dir and 
```bash 
go run main.go
```

# Comands (JQ - formatted JSON output. Advice: download or suffer)

# Add task (Duration, seconds - int)
with JQ
```bash
curl -X POST http://localhost:8080/task -H "Content-Type: application/json" -d '{"duration":15}' | jq
```
without JQ
```bash
curl -X POST http://localhost:8080/task -H "Content-Type: application/json" -d '{"duration":15}'
```
# Get ALL tasks 
with JQ
```bash
curl http://localhost:8080/task | jq
```
without JQ
```bash
curl http://localhost:8080/task
```
# Delete task by ID
with JQ
```bash 
curl -X DELETE http://localhost:8080/task/1 | jq
```
without JQ
```bash 
curl -X DELETE http://localhost:8080/task/1
```
