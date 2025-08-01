package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
)

type Owner struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	Level         int    `json:"level"`
}

type Resource struct {
	ID     string  `json:"id"`
	Owners []Owner `json:"owners"`
}

type Service struct {
	ID        string     `json:"id"`
	Resources []Resource `json:"resources"`
}

func main() {
	url := flag.String("url", "http://localhost:8080/api/v1/services", "API URL")
	parallel := flag.Int("parallel", 10, "Parallel requests per step")
	steps := flag.Int("steps", 5, "Number of steps (batches)")
	flag.Parse()

	for step := 1; step <= *steps; step++ {
		var wg sync.WaitGroup
		for i := 1; i <= *parallel; i++ {
			wg.Add(1)
			go func(i, step int) {
				defer wg.Done()
				index := (step-1)*(*parallel) + i
				service := Service{
					ID: fmt.Sprintf("service_id_%d", index),
					Resources: []Resource{
						{
							ID: fmt.Sprintf("resource_id_%d", index),
							Owners: []Owner{
								{
									ID:            fmt.Sprintf("owner_id_%d", index),
									Name:          fmt.Sprintf("name%d", index),
									AccountNumber: fmt.Sprintf("Account%d", index),
									Level:         index,
								},
							},
						},
					},
				}
				payload, _ := json.Marshal(service)
				resp, err := http.Post(*url, "application/json", bytes.NewBuffer(payload))
				if err != nil {
					fmt.Printf("Failed to send request %d: %v\n", index, err)
					return
				}
				defer resp.Body.Close()
				fmt.Printf("Sent request %d → Status: %s\n", index, resp.Status)
			}(i, step)
		}
		wg.Wait()
	}
}
