package main
// import "fmt"

func main(){ 
  user := User{}
  //user := make(map[string]string)
  user = userFromFile("user.json")

  if user.Name == "" {
    clearScreen()
    user.registrationView()
  }

  if user.Name != "" {
    clearScreen()
    user.login()
  }
}
