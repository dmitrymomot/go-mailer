package main

import (
	"fmt"
	"time"

	"github.com/dmitrymomot/go-mailer/builder"
)

func main() {
	b := builder.New(builder.EmailBuilderOptions{
		ProductName: "Test App",
		ProductLink: "http://localhost:8080",
		Logo:        "http://localhost:8080/logo.png",
	})

	_, p, err := b.Build(
		builder.H1("Test Email Builder"),
		builder.Text("Commodo culpa consectetur aute commodo ea nostrud consequat enim aute tempor irure. Cupidatat sit nostrud incididunt deserunt amet id ea ad adipisicing dolor. Ullamco ullamco excepteur consequat voluptate voluptate laboris enim veniam. Enim non incididunt nisi incididunt non sit."),
		builder.DefaultButton("Test Button", "http://localhost"),
		builder.Text("Laboris laborum tempor laboris deserunt eiusmod consectetur laborum ullamco laboris reprehenderit. Elit amet pariatur dolore eiusmod commodo. Irure reprehenderit id esse dolor culpa aliquip nostrud officia sint quis quis eu duis. Pariatur esse sit quis aliquip."),
		builder.Dictionary(map[string]interface{}{
			"Some string": "value",
			"Price":       12.30,
		}),
		// builder.Space(),
		builder.Discount("10% off for your next purchase", "Redeem this discount until 31.12.2019 23:00", "http://localhost/discount", "Redeem now"),
		// builder.Space(),
		builder.Ul([]string{
			"first item",
			"second item",
			"third item in list",
		}),

		builder.Ol([]string{
			"first item",
			"second item",
			"third item in list",
		}),

		builder.Options([]builder.OptionsItem{
			builder.OptionsItem{
				Title: "First item",
				Link:  "http://localhost",
				Text:  "Laboris laborum tempor laboris deserunt eiusmod consectetur laborum ullamco laboris reprehenderit. Elit amet pariatur dolore eiusmod commodo.",
			},

			builder.OptionsItem{
				Title: "Second item",
				Link:  "http://localhost",
				Text:  "Laboris laborum tempor laboris deserunt eiusmod consectetur laborum ullamco laboris reprehenderit. ",
			},

			builder.OptionsItem{
				Title: "The Last Item",
				Link:  "http://localhost",
				Text:  "Laboris laborum tempor laboris deserunt eiusmod consectetur laborum ullamco laboris reprehenderit. Elit amet pariatur dolore eiusmod commodo.",
			},
		}),

		builder.Ol([]string{
			"first item",
			"second item",
			"third item in list",
		}),

		builder.PurchaseTable(builder.PurchaseTableOptions{
			ID:          "invoice #123",
			Date:        time.Now().Format(time.RFC822),
			TotalAmount: "$100.90",
			Items: map[string]string{
				"Some product title":       "55.90",
				"Yet another product item": "45.00",
			},
		}),

		builder.Separator(),
		builder.SubBody([]string{
			"Laboris laborum tempor laboris deserunt eiusmod consectetur laborum ullamco laboris reprehenderit. Elit amet pariatur dolore eiusmod commodo.",
			"Irure reprehenderit id esse dolor culpa aliquip nostrud officia sint quis quis eu duis. Pariatur esse sit quis aliquip.",
		}),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}
