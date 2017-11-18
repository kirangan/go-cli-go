 // module GoCLI
 //   class View
package main
import "fmt"

    type User struct {
      Name string
      Email string
      Phone string
      Password string
    }

    type Order struct{
      timeStamp string
      origin string
      destination string
      estPrice int
    }

    type Location struct{
      name string
      coord []int
    }

    func (u *User) registrationView() {
       // u := make(map[string]string)
       fmt.Println("Registration")
       fmt.Println("")

       fmt.Println("Please tell me your name: ")
       fmt.Scanln(&u.Name)

       fmt.Println("Your e-mail: ")
       fmt.Scanln(&u.Email)

       fmt.Println("Your phone: ")
       fmt.Scanln(&u.Phone)

       fmt.Println("Your password: ")
       fmt.Scanln(&u.Password)
       u.Save("user.json")

       //registrationCtrl(u)
    }


    func (u *User) login(){
     
      var email string
      var password string

      fmt.Println("Login")

      fmt.Println("Enter your login: ")
      fmt.Scanln(&email)

      fmt.Println("Enter your password: ")
      fmt.Scanln(&password)
      
      if u.Email == email && (u.Password == password) {
        mainMenu()
      } else {
        fmt.Println("Login Failed")
      }

    }

    func mainMenu(){
      clearScreen()
      var option int

      location := Location{}
      location = locationFromFile("location.json")

      fmt.Println("Welcome to Go-CLI!")
      fmt.Println("")

      fmt.Println("Main Menu")
      fmt.Println("1. View Profile")
      fmt.Println("2. Order Go-Ride")
      fmt.Println("3. View Order History")
      fmt.Println("4. Exit")
      fmt.Scanln(&option)

      switch int(option){
      case 1:
        viewProfile()
      case 2:
        location.orderGoride()
      case 3:
        fmt.Println("History")
      case 4:
        fmt.Println("Exit")
      default:
        fmt.Println("Undefined Input")
      }
    }

    func viewProfile(){
      clearScreen()
      var option int

      user := User{}
      user = userFromFile("user.json")

      fmt.Println("View Profile")
      fmt.Println("")

      fmt.Println(user.Name)
      fmt.Println(user.Email)
      fmt.Println(user.Phone)

      fmt.Println("1. Edit Profile")
      fmt.Println("2. Back")

      fmt.Scanln(&option)
      switch int(option){
      case 1:
        user.editProfile()
      case 2:
        fmt.Println("Back")
      default:
        fmt.Println("Undefined Input")
      }
    }

    func (u *User) editProfile(){
     // u := make(map[string]string)
      clearScreen()

      fmt.Println("Edit Profile")
      fmt.Println("")

      fmt.Println("Please tell me your new name: ")
      fmt.Scanln(&u.Name)

      fmt.Println("Your new e-mail: ")
      fmt.Scanln(&u.Email)

      fmt.Println("Your new phone: ")
      fmt.Scanln(&u.Phone)

      fmt.Println("Your new password: ")
      fmt.Scanln(&u.Password)

      u.Save("user.json")

      // editProfileCtrl(u)
    }

    func (l *Location) orderGoride(){
      clearScreen()
      var origin string
      var destination string

      fmt.Println("Order Go-Ride")
      fmt.Println("")

      fmt.Println("Where are you know ? ")
      fmt.Scanln(&origin)

      fmt.Println("Where do you want to go ? ")
      fmt.Scanln(&destination)


      // u := Location{}
      // u = locationFromFile("locations.json")
      // fmt.Println(u.name)
    }

//     func order_goride_confirm(){
//       output = opts

//       puts 'Your Order Confirmation'
//       puts ''

//       puts "Origin: #{output[:origin]}"
//       puts "Destination: #{output[:destination]}"
//       puts "Est. Price: #{output[:est_price]}"
//       puts ''

//       puts 'Is that correct?'
//       puts '1. Order'
//       puts '2. Restart from beginning'
//       puts '3. Back'

//       print 'Enter your option: '
//       output[:steps] << {id: __method__, option: gets.chomp}

//       output
//     }

//     # def self.view_order_history(opts = {})
//     # end
//   }
// }


  