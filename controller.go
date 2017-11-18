// require_relative './models/user'
// require_relative './view'
package main
import "fmt"
import "io/ioutil"
import "encoding/json"
import "os"
import "os/exec"

func clearScreen(){
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}

func registrationCtrl()  {
  // f := make(map[string]string)
  // registrationView(f)

  // u := User {
  //   Name: f["Name"],
  //   Email: f["Email"],
  //   Phone: f["Phone"],
  //   Password: f["Password"],
  // }
  //registrationView()
  // u := User{}
  // u.Save("user.json")
  
}


// func editProfileCtrl(){
  
//   editProfile(f)

//    u := User {
//     Name: f["Name"],
//     Email: f["Email"],
//     Phone: f["Phone"],
//     Password: f["Password"],
    
//   }
  
//   u.Save("user.json")
// }
// func loginCntrl(){


//     func order_goride(){
//       clear_screen(opts)
//       output = order_goride(opts)

//       file = load_locations_file
//       if !file{
//         output[:flash_msg] = "Route not covered."
//         main_menu(output)
//       }

//       locations = JSON.parse(file)
//       output[:origin] = locations.
//         select {|location| location['name'] == output[:origin]}.first
//       output[:destination] = locations.
//         select {|location| location['name'] == output[:destination]}.first
//       if output[:origin].nil? || output[:destination].nil?{
//         output[:flash_msg] = "Route not covered."
//         main_menu(output)
//       end

//       output[:est_price] = calculate_est_price(output[:origin], output[:destination])
//       output[:origin] = output[:origin]['name']
//       output[:destination] = output[:destination]['name']

//       order_goride_confirm(output)
//     }

//     func order_goride_confirm(){
//       clear_screen(opts)
//       output = order_goride_confirm(opts)

//       case output[:steps].last[:option].to_i

//       when 1
//         save_order(output)
//         main_menu(output)
//       when 2
//         order_goride(output)
//       else
//         main_menu(output)
//     }

//     protected
//       func clear_screen(){
//         Gem.win_platform? ? (system "cls") : (system "clear")
//         if opts[:flash_msg] {
//           puts opts[:flash_msg]
//           puts ''
//           opts[:flash_msg] = nil
//         }
//       }



//       func load_locations_file(){
//         file = nil

//         if File.file?("#{File.expand_path(File.dirname(__FILE__))}/../data/locations.json"){
//           file = File.read("#{File.expand_path(File.dirname(__FILE__))}/../data/locations.json")
//         }
//         return file
//       }

//       func calculate_est_price(origin, destination){
//         0.0
//       }
      

//       func save_order(){
//         if !File.file?("#{File.expand_path(File.dirname(__FILE__))}/../data/orders.json")
//           orders = []
//         } else {
//           file = File.read("#{File.expand_path(File.dirname(__FILE__))}/../data/orders.json")
//           orders = JSON.parse(file)
//         }

//         orders << {
//           timestamp: Time.now,
//           origin: output[:origin],
//           destination: output[:destination],
//           est_price: output[:est_price]
//         }

//         File.open("#{File.expand_path(File.dirname(__FILE__))}/../data/orders.json", "w") do |f|
//           f.write JSON.generate(orders)
//         end
//       }
//   end
// end

// Read from file
    func userFromFile(f string) User {
      bs, err := ioutil.ReadFile(f)

      if err != nil {
        return User{}
      }
        return userFromJson(bs)
      }



    func locationFromFile(f string) Location {
      bs, err := ioutil.ReadFile(f)

      if err != nil {
        return Location{}
      }
        return locationFromJson(bs)
      }

    func orderFromFile(f string) Order {
      bs, err := ioutil.ReadFile(f)

      if err != nil {
        return Order{}
      }
        return orderFromJson(bs)
      }


// Save to File
    func (u User) Save(f string) error {
      return ioutil.WriteFile(f, u.toJson(), 0644)
    }


    func (u Location) SaveLoc(f string) error {
      return ioutil.WriteFile(f, u.toJson(), 0644)
    }

  
  // Convert to json
    func (u User) toJson() []byte {
      bs, err := json.Marshal(u)

      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      return bs
      }

    // Read from json
    func userFromJson(bs []byte) User {
      var u User
      err := json.Unmarshal(bs, &u)
      
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      return u
      }

    func locationFromJson(bs []byte) Location {
      var u Location
      err := json.Unmarshal(bs, &u)
      
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      return u
      }

    func orderFromJson(bs []byte) Order {
      var u Order
      err := json.Unmarshal(bs, &u)
      
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      return u
      }
