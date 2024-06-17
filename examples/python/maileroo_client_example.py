from maileroo_client import MailerooClient

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
print(response)