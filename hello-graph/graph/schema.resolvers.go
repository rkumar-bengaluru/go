package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/rkumar-bengaluru/hello-graph/graph/generated"
	"github.com/rkumar-bengaluru/hello-graph/graph/model"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link model.Link
	var user model.User
	link.ID = fmt.Sprintf("%d", rand.Int63())
	link.Address = input.Address
	link.Title = input.Title
	user.Name = "test"
	link.User = &user
	return &link, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

type FCMNotification struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
}

type FCMMessage struct {
	Notification FCMNotification `json:"notification"`
	Token        string          `json:"token"`
}

type FCMPushNotification struct {
	Message FCMMessage `json:"message"`
}

// CreatePushNotification is the resolver for the createPushNotification field.
func (r *mutationResolver) CreatePushNotification(ctx context.Context, input model.PushNotification) (string, error) {
	fmt.Printf("title : %s , body : %s, image %s, token %s\n", input.Title, input.Body, input.Image, input.DeviceToken)
	fcm := FCMPushNotification{
		Message: FCMMessage{
			Notification: FCMNotification{
				Title: input.Title,
				Body:  input.Body,
				Image: input.Image,
			},
			Token: input.DeviceToken,
		},
	}
	data, _ := json.Marshal(fcm)
	reqBody := ioutil.NopCloser(strings.NewReader(string(data)))
	fmt.Println(string(data))
	// req url
	reqUrl, _ := url.Parse("https://fcm.googleapis.com/v1/projects/zuellingpharma/messages:send")
	// create request object & replace oath2 token from https://developers.google.com/oauthplayground/
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {" Bearer ya29.a0AVA9y1sN8u6XdoTzyiCJAO9VMxNecms9lurpxox4MOT_kp0c-KonFhIsfoYS3vymEQfPt_6kIB62T1LbbF2-z2ZAMQcmwu2dlgbPo70An_GeWbylF3wRUm3Xnn9C7jZT4kmQ_vxgzg0AaueVGsaQ7iZOwVcAaCgYKATASARESFQE65dr8Wt5njYt-V95_kESVHJgMfQ0163"},
		},
		Body: reqBody,
	}
	// send the request
	res, err := http.DefaultClient.Do(req)
	// check for error
	if err != nil {
		log.Fatal("Error:", err)
	}
	// read response body
	dataR, _ := ioutil.ReadAll(res.Body)
	// close response body
	res.Body.Close()
	// print response status
	fmt.Printf("status: %d\n", res.StatusCode)
	fmt.Printf("body : %s\n", dataR)
	fmt.Println("initiating campaign...")
	return string(data), nil
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var links []*model.Link

	dummyLink := model.Link{
		Title:   "our dummy link",
		Address: "https://address.org",
		User:    &model.User{Name: "admin"},
	}
	links = append(links, &dummyLink)
	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var GlobalLinks = make(map[string]model.Link)
