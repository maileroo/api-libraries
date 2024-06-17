package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	APIEndpoint = "https://smtp.maileroo.com"
)

type MailerooClient struct {
	APIKey            string
	Data              map[string]interface{}
	Attachments       []string
	InlineAttachments []string
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewMailerooClient(apiKey string) *MailerooClient {

	client := &MailerooClient{
		APIKey: apiKey,
	}

	client.Reset()

	return client

}

func (c *MailerooClient) Reset() {

	c.Data = map[string]interface{}{
		"from":          "",
		"to":            "",
		"cc":            "",
		"bcc":           "",
		"reply_to":      "",
		"subject":       "",
		"html":          "",
		"plain":         "",
		"tracking":      "yes",
		"reference_id":  "",
		"tags":          []string{},
		"template_id":   "",
		"template_data": "",
	}

	c.Attachments = []string{}
	c.InlineAttachments = []string{}

}

func (c *MailerooClient) SetFrom(name, address string) *MailerooClient {
	c.Data["from"] = fmt.Sprintf("%s <%s>", name, address)
	return c
}

func (c *MailerooClient) SetTo(name, address string) *MailerooClient {
	c.Data["to"] = fmt.Sprintf("%s <%s>", name, address)
	return c
}

func (c *MailerooClient) SetCc(name, address string) *MailerooClient {
	c.Data["cc"] = fmt.Sprintf("%s <%s>", name, address)
	return c
}

func (c *MailerooClient) SetBcc(name, address string) *MailerooClient {
	c.Data["bcc"] = fmt.Sprintf("%s <%s>", name, address)
	return c
}

func (c *MailerooClient) SetReplyTo(name, address string) *MailerooClient {
	c.Data["reply_to"] = fmt.Sprintf("%s <%s>", name, address)
	return c
}

func (c *MailerooClient) SetSubject(subject string) *MailerooClient {
	c.Data["subject"] = subject
	return c
}

func (c *MailerooClient) SetHtml(html string) *MailerooClient {
	c.Data["html"] = html
	return c
}

func (c *MailerooClient) SetPlain(plain string) *MailerooClient {
	c.Data["plain"] = plain
	return c
}

func (c *MailerooClient) AddAttachment(filePath string) *MailerooClient {

	if _, err := os.Stat(filePath); err == nil {
		c.Attachments = append(c.Attachments, filePath)
	}

	return c

}

func (c *MailerooClient) AddInlineAttachment(filePath string) *MailerooClient {

	if _, err := os.Stat(filePath); err == nil {
		c.InlineAttachments = append(c.InlineAttachments, filePath)
	}

	return c

}

func (c *MailerooClient) SetReferenceId(referenceId string) *MailerooClient {

	c.Data["reference_id"] = referenceId

	return c

}

func (c *MailerooClient) SetTags(tags map[string]string) *MailerooClient {

	jsonData, _ := json.Marshal(tags)

	c.Data["tags"] = string(jsonData)

	return c

}

func (c *MailerooClient) SetTracking(tracking bool) *MailerooClient {

	if tracking {
		c.Data["tracking"] = "yes"
	} else {
		c.Data["tracking"] = "no"
	}

	return c

}

func (c *MailerooClient) SetTemplateId(templateId string) *MailerooClient {

	c.Data["template_id"] = templateId

	return c

}

func (c *MailerooClient) SetTemplateData(templateData map[string]string) *MailerooClient {

	jsonData, _ := json.Marshal(templateData)

	c.Data["template_data"] = string(jsonData)

	return c

}

func (c *MailerooClient) sendRequest(endpoint, method string) (APIResponse, error) {

	var apiResponse APIResponse

	url := APIEndpoint + endpoint

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range c.Data {

		switch v := val.(type) {

		case string:
			_ = writer.WriteField(key, v)

		case []string:
			for _, item := range v {
				_ = writer.WriteField(key, item)
			}
		}

	}

	for _, filePath := range c.Attachments {

		file, err := os.Open(filePath)

		if err != nil {
			return apiResponse, err
		}

		defer file.Close()

		part, err := writer.CreateFormFile("attachments", file.Name())

		if err != nil {
			return apiResponse, err
		}

		_, err = io.Copy(part, file)

		if err != nil {
			return apiResponse, err
		}

	}

	for _, filePath := range c.InlineAttachments {

		file, err := os.Open(filePath)

		if err != nil {
			return apiResponse, err
		}

		defer file.Close()

		part, err := writer.CreateFormFile("inline_attachments", file.Name())

		if err != nil {
			return apiResponse, err
		}

		_, err = io.Copy(part, file)

		if err != nil {
			return apiResponse, err
		}

	}

	err := writer.Close()

	if err != nil {
		return apiResponse, err
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return apiResponse, err
	}

	req.Header.Set("X-API-Key", c.APIKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return apiResponse, err
	}

	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)

	if err != nil {
		return apiResponse, err
	}

	err = json.Unmarshal(responseData, &apiResponse)

	if err != nil {
		return apiResponse, err
	}

	return apiResponse, nil

}

func (c *MailerooClient) RemoveTrailingCommas() {

	keys := []string{"from", "to", "cc", "bcc", "reply_to"}

	for _, key := range keys {

		if val, ok := c.Data[key].(string); ok {

			if len(val) > 0 && val[len(val)-1] == ',' {
				c.Data[key] = val[:len(val)-1]
			}

		}

	}

}

func (c *MailerooClient) SendBasicEmail() (APIResponse, error) {
	c.RemoveTrailingCommas()
	return c.sendRequest("/send", "POST")
}

func (c *MailerooClient) SendTemplateEmail() (APIResponse, error) {
	c.RemoveTrailingCommas()
	return c.sendRequest("/send-template", "POST")
}

func (c *MailerooClient) GenerateReferenceId() string {
	return fmt.Sprintf("%x", randomBytes(12))
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}