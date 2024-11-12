package github

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var OAuth2Config = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
}

type User struct {
	ID       string
	Username string
}

// 認証URLを生成
func GenerateAuthURL(state string) (string, error) {
	if state == "" {
		return "", fmt.Errorf("ステートが空です")
	}
	fmt.Println("aaaa", OAuth2Config.ClientID)
	return OAuth2Config.AuthCodeURL(state, oauth2.AccessTypeOnline), nil
}

// 認証コードをトークンに交換
func ExchangeCodeForToken(ctx context.Context, code string) (*oauth2.Token, error) {
	if code == "" {
		return nil, fmt.Errorf("認証コードが空です")
	}
	return OAuth2Config.Exchange(ctx, code)
}

func GetUser(ctx context.Context, accessToken string) (*User, error) {
	client := OAuth2Config.Client(ctx, &oauth2.Token{AccessToken: accessToken})
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil

}
