package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	olURL = "https://api.us.onelogin.com"
)

type Sample struct {
	GrantType string `json:"grant_type"`
}

type Tokens struct {
	AccessToken  string    `json:"access_token"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	TokenType    string    `json:"token_type"`
	AccountID    int       `json:"account_id"`
}

type OLUser struct {
	ActivatedAt          time.Time   `json:"activated_at"`
	DistinguishedName    interface{} `json:"distinguished_name"`
	ExternalID           interface{} `json:"external_id"`
	Firstname            string      `json:"firstname"`
	LastLogin            time.Time   `json:"last_login"`
	Lastname             string      `json:"lastname"`
	Company              string      `json:"company"`
	DirectoryID          interface{} `json:"directory_id"`
	InvitationSentAt     interface{} `json:"invitation_sent_at"`
	MemberOf             interface{} `json:"member_of"`
	UpdatedAt            time.Time   `json:"updated_at"`
	PreferredLocaleCode  interface{} `json:"preferred_locale_code"`
	CreatedAt            time.Time   `json:"created_at"`
	Userprincipalname    interface{} `json:"userprincipalname"`
	TrustedIdpID         interface{} `json:"trusted_idp_id"`
	Comment              string      `json:"comment"`
	Title                string      `json:"title"`
	RoleIds              []int       `json:"role_ids"`
	Department           string      `json:"department"`
	ID                   int         `json:"id"`
	CustomAttributes     OLUserCustomAttributes `json:"custom_attributes"`
	InvalidLoginAttempts int         `json:"invalid_login_attempts"`
	ManagerUserID        interface{} `json:"manager_user_id"`
	LockedUntil          time.Time   `json:"locked_until"`
	Username             string      `json:"username"`
	ManagerAdID          interface{} `json:"manager_ad_id"`
	//Email                string      `json:"email"`
	Phone                string      `json:"phone"`
	State                int         `json:"state"`
	GroupID              interface{} `json:"group_id"`
	PasswordChangedAt    time.Time   `json:"password_changed_at"`
	Status               int         `json:"status"`
	Samaccountname       interface{} `json:"samaccountname"`
}


type OLUserCustomAttributes struct {
	Employeenumber interface{} `json:"employeenumber"`
	Food           string      `json:"food"`
}

func main(){
	// set access information for the API
	userName := os.Getenv("CLIENT_ID")
	password := os.Getenv("CLIENT_SECRET")

	// get token
	token, err := generateToken(userName, password)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("[log] token is %s\n", token)

	// get user info
	userID := 143348903
	olUser, err := getUser(userID,token)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("[log] User:\n%+v\n", olUser)
	//fmt.Printf("[log] Email: %s\n", olUser.Email)
}

func generateToken(userName string, password string) (string, error){
	data := new(Sample)
	data.GrantType = "client_credentials"

	dataJSON, _ := json.Marshal(data)
	//fmt.Printf("[log] bodyJSON: %+v\n", dataJSON)

	req, err := http.NewRequest(http.MethodPost, olURL + "/auth/oauth2/v2/token", bytes.NewBuffer(dataJSON))
	if err != nil{
		fmt.Println("[Err] ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Basic " + encodedCredentials)
	req.Header.Set("Authorization", "client_id:" + userName + " ,client_secret:" + password)


	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Err] ", err.Error())
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return "", errors.New("[Err] Request OneLogin API token: HTTP Status is " + resp.Status)
	}

	fmt.Printf("status: %s\n", resp.Status)
	var r io.Reader = resp.Body
	//r = io.TeeReader(r, os.Stderr)

	var tokens Tokens
	json.NewDecoder(r).Decode(&tokens)

	return tokens.AccessToken, nil
}

func getUser(userID int, token string) (*OLUser, error) {
	req, err := http.NewRequest(http.MethodGet, olURL + "/api/2/users/" + fmt.Sprint(userID), nil)
	if err != nil{
		fmt.Println("[Err] ", err.Error())
	}
	req.Header.Set("Authorization", "beare " + token)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[Err] ", err.Error())
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return nil, errors.New("[Err] Get User: HTTP Status is " + resp.Status)
	}

	fmt.Printf("status: %s\n", resp.Status)
	var r io.Reader = resp.Body
	//r = io.TeeReader(r, os.Stderr)

	var olUser OLUser
	json.NewDecoder(r).Decode(&olUser)

	return &olUser, nil
}
