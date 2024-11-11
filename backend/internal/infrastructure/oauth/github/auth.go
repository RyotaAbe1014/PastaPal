package github

import (
	"context"
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
