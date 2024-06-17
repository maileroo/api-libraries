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
const MailerooClient = require('MailerooClient');

const apiKey = 'YOUR_API_KEY';
const client = new MailerooClient(apiKey);

client.setFrom('Sender Name', 'sender@example.com')
    .setTo('Recipient Name', 'receiver@example.com')
    .setSubject('Test Email')
    .setHtml('<p>This is a test email</p>')
    .setPlain('This is a test email')
    .setReferenceId(client.generateReferenceId())
    .setTags({
        'tag1': 'value1',
        'tag2': 'value2'
    })
    .setTracking(true);

client.addAttachment('path/to/file', 'file_name', 'file_type');

client.sendBasicEmail().then(response => {

    if (response.success) {
        console.log('Email sent successfully:', response.data);
    } else {
        console.error('Error sending email:', response.error);
    }

}).catch(error => {

    console.error('Error sending email:', error);

});
```

### Python

```
client = MailerooClient('API_KEY')

response = client.set_from('Sender Name', 'sender@example.com') \
    .set_to('Recipient Name', 'recipient@example.com') \
    .set_subject('Test Email') \
    .set_html('<h1>Hello</h1>') \
    .set_plain('Hello') \
    .set_tracking(True) \
    .set_reference_id(client.generate_reference_id()) \
    .set_tags({'tag1': 'value1', 'tag2': 'value2'}) \
    .add_attachment('test.txt', 'test.txt', 'text/plain') \
    .send_basic_email()
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