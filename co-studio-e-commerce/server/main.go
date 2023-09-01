package main

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/route"
)

func main() {
	// // 5 requests per 5 minutes
	// // rate limit
	// limit := limiter.Rate{
	// 	Period: 5 * time.Minute,
	// 	Limit:  5,
	// }

	conf.SetEnv()
	app := route.NewService()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
