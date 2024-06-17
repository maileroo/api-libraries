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