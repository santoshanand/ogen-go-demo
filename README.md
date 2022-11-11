### Pet Store API

# To Generate code 
  ```
  go generate ./...   
  
  ```

# Run main server

  ```
  go run main.go
  ```

# Check out put

  ```
  curl -X "POST" -H "Content-Type: application/json" --data "{\"name\":\"Cat\"}" http://localhost:5000/pet
  ```