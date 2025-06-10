# How to run? Get in cmd dir and 
```bash 
go run main.go
```

# Comands (JQ - formatted JSON output. Advice: download or suffer)

# Add task (Duration, seconds - int)
```bash
curl -X POST http://localhost:8080/task -H "Content-Type: application/json" -d '{"duration":15}'
```
# Get ALL tasks 
```bash
curl http://localhost:8080/task | jq
```
# Delete task by ID
```bash 
curl -X DELETE http://localhost:8080/task/1 | jq
```
