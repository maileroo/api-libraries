<?php

require_once("../../src/php/MailerooClient.php");

$api_key = 'YOUR_API_KEY';

$client = new MailerooClient($api_key);

// Send a basic email

$basic_email_response = $client->setFrom('Maileroo', 'no.reply@mail.maileroo.com')
    ->setTo('John Doe', 'john.doe@maileroo.com')
    ->setCc('Jane Doe', 'jane.doe@maileroo.com') // Optional
    ->setBcc('Jim Doe', 'jim.doe@maileroo.com') // Optional
    ->setReplyTo('Administrator', 'admin@maileroo.com') // Optional
    ->setReferenceId($client->generateReferenceId()) // Optional
    ->setTags(['tag1' => 'value1', 'tag2' => 'value2']) // Optional
    ->setTracking(true) // Optional
    ->setSubject('Hello World')
    ->setHtml('<p>Hello World</p>')
    ->setPlain('Hello World') // Optional
    ->addAttachment('path/to/file', 'file_name', 'file_type') // Optional
    ->addInlineAttachment('path/to/file', 'file_name', 'file_type') // Optional
    ->sendBasicEmail();

var_dump($basic_email_response);

// Reset

$client->reset();

// Send a template email

$template_email_response = $client->setFrom('Maileroo', 'no.reply@mail.maileroo.com')
    ->setTo('John Doe', 'john.doe@maileroo.com')
    ->setCc('Jane Doe', 'jane.doe@maileroo.com') // Optional
    ->setBcc('Jim Doe', 'jim.doe@maileroo.com') // Optional
    ->setReplyTo('Administrator', 'admin@maileroo.com') // Optional
    ->setReferenceId($client->generateReferenceId()) // Optional
    ->setTags(['tag1' => 'value1', 'tag2' => 'value2']) // Optional
    ->setTracking(true) // Optional
    ->setSubject('Hello World')
    ->setTemplateId(1)
    ->setTemplateData(['name' => 'John Doe'])
    ->addAttachment('path/to/file', 'file_name', 'file_type') // Optional
    ->addInlineAttachment('path/to/file', 'file_name', 'file_type') // Optional
    ->sendTemplateEmail();

var_dump($template_email_response);