package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bxcodec/faker/v4"
)

type UserObject struct {
	Type      string `json:"type"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserData struct {
	Type      string `json:"type"`
	ID        string `json:"id"`
	Href      string `json:"href"`
	ViewHref  string `json:"view_href"`
	Login     string `json:"login"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type userResStruct struct {
	Status   string   `json:"status"`
	HTTPCode int      `json:"http_code"`
	Data     UserData `json:"data"`
	Password string   `json:"password"`
}

func main() {

	// Open the CSV file
	file, err := os.Create("android_comm_3k.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Type", "ID", "Href", "ViewHref", "Login", "Email", "FirstName", "LastName", "Password"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	for i := 0; i < 3000; i++ {

		rand := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomNumber := rand.Intn(9000) + 1000
		email := fmt.Sprintf("test%d@test.com", randomNumber)

		newUser := UserObject{
			Type:      "user",
			FirstName: faker.FirstName(),
			LastName:  faker.LastName(),
			Login:     faker.Username(),
			Email:     email,
			Password:  faker.IPv4(),
		}

		jsonUser, err := json.Marshal(newUser)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		// fmt.Println("Request Data:", string(jsonUser))

		client := &http.Client{}
		req, err := http.NewRequest("POST", "https://cezur57386.stage.lithium.com/api/2.0/users", bytes.NewBuffer(jsonUser))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		var userRes userResStruct
		err = json.Unmarshal(body, &userRes)
		if err != nil {
			fmt.Println("Error unmarshaling response:", err)
			return
		}
		userRes.Password = newUser.Password

		row := []string{
			userRes.Data.Type,
			userRes.Data.ID,
			userRes.Data.Href,
			userRes.Data.ViewHref,
			userRes.Data.Login,
			userRes.Data.Email,
			userRes.Data.FirstName,
			userRes.Data.LastName,
			userRes.Password,
		}

		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing row:", err)
			return
		}
		// fmt.Printf("User %d created:\n", i)
	}

}
