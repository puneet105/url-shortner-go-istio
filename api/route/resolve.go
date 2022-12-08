package route

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/puneet105/url-shortner-go/api/database"
)

func ResolveUrl(c *fiber.Ctx) error{

	url := c.Params("url")
	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error":"short url not found in db"})
	} else if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error":"Internal server error"})
	}

	rIncr := database.CreateClient(1)
	defer rIncr.Close()

	_ = rIncr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}
