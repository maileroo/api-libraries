package main

import "fmt"

func main() {

	client := NewMailerooClient("YOUR_API_KEY")

	client.SetFrom("John Doe", "john.doe@example.com").
		SetTo("Jane Doe", "jane.doe@example.com").
		SetSubject("Hello").
		SetHtml("<h1>Hello World</h1>").
		SetPlain("Hello World").
		SetReferenceId(client.GenerateReferenceId()).
		SetTracking(true).
		SetTags(map[string]string{"tag1": "value1", "tag2": "value2"})

	response, err := client.SendBasicEmail()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", response)
	}

}
