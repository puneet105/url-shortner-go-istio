package route

import (
	"bufio"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/puneet105/url-shortner-go/api/database"
	"log"
	"os"
	"strings"
)

func ResolveUrl(c *fiber.Ctx) error{

	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()

	//fetch from redis db
	value, err := r.Get(database.Ctx, url).Result()
	fmt.Println(value)
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"short url not found in db"})
	} else if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"})
	}

	rIncr := database.CreateClient(1)
	defer rIncr.Close()

	_ = rIncr.Incr(database.Ctx, "counter")

	//fetch from file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make(map[string]string, 1024*4)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineRead := scanner.Text()
		stringsSlice := strings.Split(lineRead, " ")
		data[stringsSlice[0]] = stringsSlice[1]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Content retrieved from file",data[value])
	fmt.Println("Your Request has been redirected to", value)
	return c.Redirect(value, 301)
}
