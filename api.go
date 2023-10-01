package main

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	JariJariLingkaran int `json:"jariJariLingkaran"`
	SisiPersegi       int `json:"sisiPersegi"`
	AlasSegitiga      int `json:"alasSegitiga"`
	TinggiSegitiga    int `json:"tinggiSegitiga"`
}

type ResponseData struct {
	LuasLingkaran     float64 `json:"luasLingkaran"`
	KelilingLingkaran float64 `json:"kelilingLingkaran"`
	LuasPersegi       float64 `json:"luasPersegi"`
	KelilingPersegi   float64 `json:"kelilingPersegi"`
	LuasSegitiga      float64 `json:"luasSegitiga"`
	KelilingSegitiga  float64 `json:"kelilingSegitiga"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		request := new(RequestData)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		response := new(ResponseData)

		response.LuasLingkaran = math.Pi * float64(request.JariJariLingkaran*request.JariJariLingkaran)
		response.KelilingLingkaran = 2 * math.Pi * float64(request.JariJariLingkaran)
		response.LuasPersegi = math.Pow(float64(request.SisiPersegi), 2)
		response.KelilingPersegi = 4 * float64(request.SisiPersegi)
		response.LuasSegitiga = 0.5 * float64(request.AlasSegitiga) * float64(request.TinggiSegitiga)
		response.KelilingSegitiga = float64(request.AlasSegitiga) + 2*math.Sqrt(math.Pow(float64(request.TinggiSegitiga), 2)+math.Pow(float64(request.AlasSegitiga)/2, 2))

		return c.JSON(response)
	})

	app.Listen(":3000")
}
