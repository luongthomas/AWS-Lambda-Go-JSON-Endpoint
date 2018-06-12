package main

import (
	"fmt"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

type PersonJson struct {
	Name           string `json:"name"`
	FavoriteColor  string `json:"favoriteColor"`
	FavoriteAnimal string `json:"favoriteAnimal"`
}

type DataJson struct {
	EachPerson []PersonJson `json:"data"`
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	bytes := []byte(request.Body)
	d := DataJson{}
	err := json.Unmarshal(bytes, &d)
	if err != nil {
		panic(err)
	}
	
	arr := []string{}
	for _, item := range d.EachPerson {
		arr = append(arr, item.Name)
	}
	
	var formattedString = ""
	for _, name := range arr {
		formattedString += name
		formattedString += "\n"
	}


	// Create new JSON to send back
	// https://stackoverflow.com/questions/40429296/converting-string-to-json-or-struct-in-golang
	jsonStr := `{"names": ` + formattedString + `}`

	fmt.Println(jsonStr)

	// jsonBytes := []byte(jsonStr)
	// var raw map[string]interface{}
	// json.Unmarshal(jsonBytes, &raw)
	// out, _ := json.Marshal(raw)
	
	return events.APIGatewayProxyResponse{Body: formattedString, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
