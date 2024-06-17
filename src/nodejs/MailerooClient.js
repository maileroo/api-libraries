const fs = require('fs');
const crypto = require('crypto');
const { Blob } = require('buffer');

class MailerooClient {

    static API_ENDPOINT = 'https://smtp.maileroo.com/';

    constructor(api_key) {
        this.api_key = api_key;
        this.reset();
    }

    reset() {
        this.data = {
            from: '',
            to: '',
            cc: '',
            bcc: '',
            reply_to: '',
            subject: '',
            html: '',
            plain: '',
            tracking: 'yes',
            reference_id: '',
            tags: [],
            template_id: '',
            template_data: ''
        };
        this.attachments = [];
        this.inline_attachments = [];
    }

    setFrom(name, address) {
        this.data.from += `${name} <${address}>,`;
        return this;
    }

    setTo(name, address) {
        this.data.to += `${name} <${address}>,`;
        return this;
    }

    setCc(name, address) {
        this.data.cc += `${name} <${address}>,`;
        return this;
    }

    setBcc(name, address) {
        this.data.bcc += `${name} <${address}>,`;
        return this;
    }

    setReplyTo(name, address) {
        this.data.reply_to += `${name} <${address}>,`;
        return this;
    }

    setSubject(subject) {
        this.data.subject = subject;
        return this;
    }

    setHtml(html) {
        this.data.html = html;
        return this;
    }

    setPlain(plain) {
        this.data.plain = plain;
        return this;
    }

    addAttachment(file_path, file_name, file_type) {

        if (fs.existsSync(file_path)) {
            const file = fs.readFileSync(file_path);
            this.attachments.push({ file, file_name, file_type });
        }

        return this;

    }

    addInlineAttachment(file_path, file_name, file_type) {

        if (fs.existsSync(file_path)) {
            const file = fs.readFileSync(file_path);
            this.inline_attachments.push({ file, file_name, file_type });
        }

        return this;

    }

    setReferenceId(reference_id) {
        this.data.reference_id = reference_id;
        return this;
    }

    setTags(tags) {
        this.data.tags = JSON.stringify(tags);
        return this;
    }

    setTracking(tracking) {
        this.data.tracking = tracking ? 'yes' : 'no';
        return this;
    }

    setTemplateId(template_id) {
        this.data.template_id = template_id;
        return this;
    }

    setTemplateData(template_data) {
        this.data.template_data = JSON.stringify(template_data);
        return this;
    }

    removeTrailingCommas() {

        const keys = ['from', 'to', 'cc', 'bcc', 'reply_to'];

        keys.forEach(key => {
            if (this.data[key]) {
                this.data[key] = this.data[key].replace(/,\s*$/, '');
            }
        });

    }

    async sendRequest(endpoint, method) {

        const url = `${MailerooClient.API_ENDPOINT}${endpoint}`;

        const headers = {
            'X-API-Key': this.api_key
        };

        const form = new FormData();

        Object.keys(this.data).forEach(key => {
            form.append(key, this.data[key]);
        });

        this.attachments.forEach((attachment, index) => {
            form.append(`attachments[${index}]`, new Blob([attachment.file], { type: attachment.file_type }), attachment.file_name);
        });

        this.inline_attachments.forEach((inline_attachment, index) => {
            form.append(`inline_attachments[${index}]`, new Blob([inline_attachment.file], { type: inline_attachment.file_type }), inline_attachment.file_name);
        });

        try {

            const response = await fetch(url, {
                method,
                headers,
                body: form
            });

            return await response.json();

        } catch (error) {

            return {
                success: false,
                message: error.message
            };

        }

    }

    async sendBasicEmail() {

        this.removeTrailingCommas();

        return this.sendRequest('/send', 'POST');

    }

    async sendTemplateEmail() {

        this.removeTrailingCommas();

        return this.sendRequest('/send-template', 'POST');

    }

    generateReferenceId() {
         return crypto.randomBytes(12).toString('hex');
    }

}

module.exports = MailerooClient;