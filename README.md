# Maileroo API Libraries

Welcome to the official repository for Maileroo's API libraries. This repository contains SDKs to help you integrate with Maileroo's email sending API efficiently and effectively.

Sending emails with our easy-to-use SDKs, providing seamless integration and powerful features. Whether you're a seasoned developer or just getting started, Maileroo makes it simple to send emails and manage your email delivery effortlessly.

## Usage

### PHP

```
$client = new MailerooClient("YOUR_API_KEY");

$client->setFrom('Maileroo', 'no.reply@mail.maileroo.com')
    ->setTo('John Doe', 'john.doe@maileroo.com')
    ->setSubject('Hello World')
    ->setHtml('<p>Hello World</p>')
    ->setPlain('Hello World')
    ->sendBasicEmail();
```

### Node.js

```
TODO
```

### Python

```
TODO
```

### Golang

```
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
```

### C#

```
TODO
```

