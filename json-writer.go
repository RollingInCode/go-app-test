import (
    "encoding/json"
    "os"
)

func writeJSON(filename string, users <-chan User) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    writer := json.NewEncoder(file)
    
    for user := range users {
        if err := writer.Encode(user); err != nil {
            return err
        }
    }
    
    return nil
}
