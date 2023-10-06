package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/kopoze/kpz/pkg/app"
)

type Response struct {
	Data []app.App
}

func List() {
	req, err := http.NewRequest("GET", BuildURL(), nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	// TODO: Add API key auth
	// req.Header.Add("x-rapidapi-key", "YOU_API_KEY")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}

	var response Response
	json.Unmarshal(body, &response)
	for _, app := range response.Data {
		fmt.Printf("%d: %s [%s -> %s]\n", app.ID, app.Name, app.Subdomain, app.Port)
	}
}

func Create(name string, subdomain string, port string) {

	appPayload := map[string]string{"name": name, "subdomain": subdomain, "port": port}
	json_data, err := json.Marshal(appPayload)

	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest("POST", BuildURL(), bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res["error"])
	}

	fmt.Println("App added successfully")
}

func Delete(id string) {
	vId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Print(err.Error())
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%d/", BuildURL(), vId), nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}

	var response Response
	json.Unmarshal(body, &response)

	if res.StatusCode != 204 {
		fmt.Println(response)
	}
}

func Update(id string, field string, value string) {

	appPayload := map[string]string{field: value}
	json_data, err := json.Marshal(appPayload)
	if err != nil {
		fmt.Print(err.Error())
	}

	vId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Print(err.Error())
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s%d/", BuildURL(), vId), bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)
		fmt.Println(res["error"])
	}

	fmt.Println("App updated successfully")
}
