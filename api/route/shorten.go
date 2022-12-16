package route

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"url-shortner-go/api/database"
	"url-shortner-go/api/handler"
	"os"
	"time"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry     	    time.Duration `json:"expiry"`
}

var filePath = "../api/data.txt"
func ShortenUrl(c *fiber.Ctx) error {

	body := new(request)
	err := c.BodyParser(&body)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"cannot parse JSON"})
	}
	fmt.Println(err)
	//check if URL is valid
	if !govalidator.IsURL(body.URL){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error":"Invalid URL"})
	}

	//check any domain error
	if !handler.RemoveDomainError(body.URL){
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"Cannot perform operation...!!"})
	}

	//enforce http
	body.URL = handler.EnforceHTTP(body.URL)

	//generate shot hash value of string
	var id string
	if body.CustomShort == ""{
		id = uuid.New().String()[:6]
	}else{
		id = body.CustomShort
	}

	r := database.CreateClient(0)
	defer r.Close()

	val, _ := r.Get(database.Ctx, id).Result()

	if val != ""{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "URL Custom short is already in use",
		})
	}

	if body.Expiry == 0{
		body.Expiry = 24
	}
	err = r.Set(database.Ctx, id, body.URL, body.Expiry*3600*time.Second).Err()

	//writing into a file
	fmt.Printf("Request received by StoreInFile: shortUrl: %v originalUrl:%v\n", id, body.URL)
	modifiedData := id + " " + body.URL

	fileHandler, errorResponse := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if errorResponse != nil {
		fmt.Println("StoreInFile | Error while opening file", errorResponse)
	}
	noOfWrittenBytes, errorResponse := fmt.Fprintln(fileHandler, modifiedData)
	fmt.Println("No. of bytes written", noOfWrittenBytes)
	if errorResponse != nil {
		fmt.Println("StoreInFile | Error while writing to file: %v", errorResponse)
		fileHandler.Close()
	}
	errorResponse = fileHandler.Close()
	if errorResponse != nil {
		fmt.Println("StoreInFile | Error while closing file connection", errorResponse)
	}
	fmt.Println("Request successfully processed by StoreInFile")

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to connect to server",
		})
	}

	resp := response{
		URL: 			body.URL,
		CustomShort: 	"",
		Expiry: 		body.Expiry,
	}

	resp.CustomShort = os.Getenv("DOMAIN") + "/" + id

	return c.Status(fiber.StatusOK).JSON(resp)

}
