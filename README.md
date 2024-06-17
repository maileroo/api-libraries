# Maileroo API Libraries

Welcome to the official repository for Maileroo's API libraries. This repository contains SDKs to help you integrate with Maileroo's email sending API efficiently and effectively.

Sending emails with our easy-to-use SDKs, providing seamless integration and powerful features. Whether you're a seasoned developer or just getting started, Maileroo makes it simple to send emails and manage your email delivery effortlessly.

## Usage

### PHP

```
require_once("MailerooClient.php");

$api_key = 'YOUR_API_KEY';

$client = new MailerooClient($api_key);

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
TODO
```

### C#

```
TODO
```

