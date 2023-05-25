package main

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    users, err := readCSV("users.csv")
    if err != nil {
        log.Fatalf("Failed to read CSV file: %v", err)
    }
    
    if err := writeJSON("users.json", users); err != nil {
        log.Fatalf("Failed to write JSON file: %v", err)
    }
}
