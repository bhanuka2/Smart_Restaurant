package scheduler

import (
    "fmt"
    "time"
    "smart-restaurant/services"
)

func StartScheduler() {
    go func() {
        for {
            fmt.Println("Running offer engine...")
            offers := services.GenerateOffers()
            fmt.Println("Generated offers:", offers)

            time.Sleep(30 * time.Second) // run every 30 sec
        }
    }()
}