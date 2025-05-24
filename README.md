# How to run?
## Get in cmd dir and 
`go run main.go`

# Comands
# JQ - format output for JSON. Advice: download or suffer
# Fast add quotes
```bash
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Shaulin", "quote":"Life is good."}'
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Shuka", "quote":"Life is bad."}'
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Confucius", "quote":"Life is gorgeus and simple."}'
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Shaulin", "quote":"Life is good and beauty."}'
curl -X POST http://localhost:8080/quotes -H "Content-Type: application/json" -d '{"author":"Shuka", "quote":"Life is rock and roll."}'
```

# Get ALL quotes (JQ - format outputing. Advice: download)
```bash
curl http://localhost:8080/quotes | jq
```

# Get random quote
```bash
curl http://localhost:8080/quotes/random | jq
```

# Output quotes from a specific author
```bash
curl http://localhost:8080/quotes?author=Confucius | jq
```

# Delete quote by ID
```bash 
curl -X DELETE http://localhost:8080/quotes/1 | jq
```
