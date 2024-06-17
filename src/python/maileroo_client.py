import requests
import json
import os
import random
import string


class MailerooClient:
    API_ENDPOINT = 'https://smtp.maileroo.com/'

    def __init__(self, api_key):
        self.api_key = api_key
        self.reset()

    def reset(self):
        self.data = {
            'from': '',
            'to': '',
            'cc': '',
            'bcc': '',
            'reply_to': '',
            'subject': '',
            'html': '',
            'plain': '',
            'tracking': 'yes',
            'reference_id': '',
            'tags': [],
            'template_id': '',
            'template_data': ''
        }
        self.attachments = []
        self.inline_attachments = []

    def set_from(self, name, address):
        self.data['from'] += f'{name} <{address}>,'
        return self

    def set_to(self, name, address):
        self.data['to'] += f'{name} <{address}>,'
        return self

    def set_cc(self, name, address):
        self.data['cc'] += f'{name} <{address}>,'
        return self

    def set_bcc(self, name, address):
        self.data['bcc'] += f'{name} <{address}>,'
        return self

    def set_reply_to(self, name, address):
        self.data['reply_to'] += f'{name} <{address}>,'
        return self

    def set_subject(self, subject):
        self.data['subject'] = subject
        return self

    def set_html(self, html):
        self.data['html'] = html
        return self

    def set_plain(self, plain):
        self.data['plain'] = plain
        return self

    def add_attachment(self, file_path, file_name, file_type):
        if os.path.exists(file_path):
            self.attachments.append(('attachments', (file_name, open(file_path, 'rb'), file_type)))
        return self

    def add_inline_attachment(self, file_path, file_name, file_type):
        if os.path.exists(file_path):
            self.inline_attachments.append(('inline_attachments', (file_name, open(file_path, 'rb'), file_type)))
        return self

    def set_reference_id(self, reference_id):
        self.data['reference_id'] = reference_id
        return self

    def set_tags(self, tags):
        self.data['tags'] = json.dumps(tags)
        return self

    def set_tracking(self, tracking):
        self.data['tracking'] = 'yes' if tracking else 'no'
        return self

    def set_template_id(self, template_id):
        self.data['template_id'] = template_id
        return self

    def set_template_data(self, template_data):
        self.data['template_data'] = json.dumps(template_data)
        return self

    def send_request(self, endpoint, method):
        url = self.API_ENDPOINT + endpoint

        headers = {
            'X-API-Key': self.api_key
        }

        post_fields = self.data.copy()
        files = {}

        if self.attachments:
            for i, (key, value) in enumerate(self.attachments.items()):
                files[f'file{i}'] = value

        if self.inline_attachments:
            for i, (key, value) in enumerate(self.inline_attachments.items()):
                files[f'inline_file{i}'] = value

        if method.upper() == 'GET':
            response = requests.get(url, headers=headers, params=post_fields)
        else:
            response = requests.request(method, url, headers=headers, data=post_fields, files=files)

        return response.json()

    def remove_trailing_commas(self):
        keys = ['from', 'to', 'cc', 'bcc', 'reply_to']

        for key in keys:
            if key in self.data:
                self.data[key] = self.data[key].rstrip(',')

    def send_basic_email(self):
        self.remove_trailing_commas()
        return self.send_request('/send', 'POST')

    def send_template_email(self):
        self.remove_trailing_commas()
        return self.send_request('/send-template', 'POST')

    def generate_reference_id(self):
        hex_chars = string.hexdigits.lower()[:16]
        return ''.join(random.choices(hex_chars, k=24))