package main

import (
	"context"
	_ "fmt"
	"log"
	_ "os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	// connection with firebase
	// [my admin sdk file]
	adminSdkFile := "E:/GO/src/go_code/projectWeb/fresh-catwalk-351117-firebase-adminsdk-vrzdd-c07d74f71f.json"
	opt := option.WithCredentialsFile(adminSdkFile)
	log.Println("Successfully Set opt: ", opt)

	// [START initialize_ap_golang]
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error making NewApp: %v", err)
	}
	// [END initialize_app_golang]

	// // USER PART:
	// // read json file
	// path := "E:/GO/src/go_code/projectWeb/log/accountRegister.json"
	// jsonFile, err := os.ReadFile(path)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// var newUser util.NewUser
	// newUser.JsonUnmarshal(jsonFile)
	// // if email is new: maek a new Auth.client，create a user in firebase
	// // [START inicialize_auth.client_golang]
	// client, err := app.Auth(ctx)
	// if err != nil {
	// 	log.Fatalf("error making Auth: %v", err)
	// }
	// // create new user, check duplicate registration
	// newUser.CreateUser(ctx, client)

	// // user sign in
	// path := "E:/GO/src/go_code/projectWeb/log/accountRegister.json"
	// jsonFile, err := os.ReadFile(path)
	// var user util.User
	// user.JsonUnmarshal(jsonFile)
	// // find user by email, return UID
	// wantedUser := util.GetUserByEmail(ctx, app, user.Email)
	// fmt.Println(wantedUser.DisplayName)

	// EVENT PART:
	//email := "monday500jcy@gamail.com"
	// title := "Let's Roll!"
	// field := "text"
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	// // json unmarshal
	// path := "E:/GO/src/go_code/projectWeb/log/uploadEvent.json"
	// jsonFile, err := os.ReadFile(path)
	// if err != nil {
	// 	log.Fatalf("error read json file %v", err)
	// }
	// var event util.Event
	// util.JsonUnmarshal(jsonFile, &event)
	// fmt.Println(event)
	// event.UploadEvent(ctx, client, email)

	// // CRUD
	// util.DeleteEvent(ctx, client, email, title)
	// var event util.Event
	// util.GetEventByTitle(ctx, client, email, title, event)
	// util.GetEventByEmail(ctx, client, email)
}
