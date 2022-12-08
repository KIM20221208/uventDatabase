package util

import (
	"context"
	"encoding/json"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
)

// read json file, assignment to object
func JsonUnmarshal(data []byte, object interface{}) {
	err := json.Unmarshal(data, object)
	if err != nil {
		log.Fatalf("error unmarshal %v", err)
	}
	log.Print("Successfully unmarshal user json file.")
}

// account part
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	User
	EmailVerified bool   `json:"emailVerified"`
	PhoneNumber   string `json:"phoneNumber"`
	DisplayName   string `json:"displayName"`
	PhotoURL      string `json:"photoURL"`
	Disabled      bool   `json:"disabled"`
}

// create user， if uid, email, phonenumber had been used，report err
func (N *NewUser) CreateUser(ctx context.Context, client *auth.Client) *auth.UserRecord {
	// [START create_user_golang]
	params := (&auth.UserToCreate{}).
		Email(N.Email).
		EmailVerified(N.EmailVerified).
		PhoneNumber(N.PhoneNumber).
		Password(N.Password).
		DisplayName(N.DisplayName).
		PhotoURL(N.PhotoURL).
		Disabled(N.Disabled)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v, %v\n", N.Email, err)
	}
	log.Printf("Successfully created user: %v\n", u)
	// [END create_user_golang]
	return u
}

// get user data py email
func GetUserByEmail(ctx context.Context, app *firebase.App, email string) *auth.UserRecord {
	// [START get_user_golang]
	// Get an auth client from the firebase.App
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUserByEmail(ctx, email)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", email, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]
	return u
}

// update user(incomplete, unusable)
func UpdateUser(ctx context.Context, client *auth.Client) {
	// user@example.com
	uid := "FNK82cJjsmXfJyiJtESn0zVowQ13"
	// [START update_user_golang]
	params := (&auth.UserToUpdate{}).
		Email("user@example.com").
		EmailVerified(true).
		PhoneNumber("+15555550100").
		Password("newPassword").
		DisplayName("JIN CHENGYU").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(true)
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully updated user: %v\n", u)
	// [END update_user_golang]
}

// delete user
func DeleteUser(ctx context.Context, client *auth.Client, uid string) {
	// [START delete_user_golang]
	err := client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	// [END delete_user_golang]
}

// event part
type EventDates struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Places struct {
	City     string `json:"city"`
	District string `json:"district"`
	Details  string `json:"details"`
}

type Event struct {
	Category        string     `json:"category"`
	Title           string     `json:"title"`
	Text            string     `json:"text"`
	EventDate       EventDates `json:"eventDate"`
	Place           Places     `json:"place"`
	RecruitDeadline string     `json:"recruitDeadline"`
	Images          string     `json:"images"`
}

// add event to database
func (E *Event) UploadEvent(ctx context.Context, client *firestore.Client, email string) {
	// Collection represent a user，Doc represent a event，Set contains the detail of event
	_, err := client.Collection(email).Doc(E.Title).Set(ctx, map[string]interface{}{
		"category":        E.Category,
		"text":            E.Text,
		"eventDate":       E.EventDate,
		"place":           E.Place,
		"recruitDeadline": E.RecruitDeadline,
		"image":           E.Images,
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Fatalf("An error has occurred: %s", err)
	}
	log.Printf("Successfully upload event: %v", E.Title)
	operate := "upload"
	UpdateLog(ctx, client, email, E.Title, operate)
}

// update log by user operator
func UpdateLog(ctx context.Context, client *firestore.Client, email, title, operate string) {
	_, err := client.Collection("log").Doc(email).Set(ctx, map[string]interface{}{
		"title":     title,
		"timestamp": firestore.ServerTimestamp,
		"operate":   operate,
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	log.Printf("Successfully update log. user: %v, event: %v, operate: %v", email, title, operate)
}

// update event by path(event details)
func UpdateEventPath(ctx context.Context, client *firestore.Client, email, title, field, value string) {
	_, err := client.Collection(email).Doc(title).Update(ctx, []firestore.Update{
		{
			Path:  field,
			Value: value,
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Fatalf("An error has occurred: %s", err)
	}
	log.Printf("Successfully update event path: %v, value: %v", title, value)
}

// get event by title（获取数据）
func GetEventByTitle(ctx context.Context, client *firestore.Client, email, title string, event Event) {
	dsnap, err := client.Collection(email).Doc(title).Get(ctx)
	if err != nil {
		log.Fatalf("error getting Doc: %v. \n%v", title, err)
	}
	dsnap.DataTo(&event)
	log.Printf("event: %v, the detail is:\n%v", title, event)
}

// print event by path
func GetEventByPathValue(ctx context.Context, client *firestore.Client, email, path, value string) {
	iter := client.Collection(email).Where(path, "==", value).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error getting event by path: %v, value: %v.\n%v", path, value, err)
		}
		log.Println(doc.Data())
	}
	log.Printf("Successfully find all event by email: %v, path: %v, value: %v.", email, path, value)
}

// get total event by user email
func GetEventByEmail(ctx context.Context, client *firestore.Client, email string) {
	iter := client.Collection(email).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error getting event by email: %v", email)
		}
		log.Println(doc.Data())
	}
	log.Printf("Successfully find all event by email: %v.", email)
}

// delete event（doc）
func DeleteEvent(ctx context.Context, client *firestore.Client, email, title string) {
	_, err := client.Collection(email).Doc(title).Delete(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	log.Printf("Successfully delete event: %v", title)
	operate := "delete"
	UpdateLog(ctx, client, email, title, operate)
}

// delete event details（path）
func DeleteEventPath(ctx context.Context, client *firestore.Client, email, title, field string) {
	_, err := client.Collection(email).Doc(title).Update(ctx, []firestore.Update{
		{
			Path:  field,
			Value: firestore.Delete,
		},
	})
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Fatalf("An error has occurred: %s", err)
	}
	log.Printf("Successfully delete event.")
}
