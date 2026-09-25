# Email Campaign

A Go command-line tool for sending email campaigns concurrently.

## Features

- Send emails concurrently
- Load campaign configuration from the project configuration package
- Render reusable email templates
- Fast and easy to use

## Requirements

- Go 1.20+

## Installation

```bash
git clone <repository-url>
cd Email_Campaign
go run .
```

## Usage

Configure the email details in the project configuration, prepare the email
template, then run:

```bash
go run .
```

The application reads the configured recipients and message details, renders
the selected template, and sends the campaign concurrently. Review the output
for send status and errors.

To run a local Mailpit SMTP server and web interface:

```bash
docker run -d \
	--restart unless-stopped \
	--name=mailpit \
	-p 8025:8025 \
	-p 1025:1025 \
	axllent/mailpit
```

## Features

### Concurrent delivery

Emails are processed concurrently so a campaign does not have to wait for
each recipient sequentially. Keep the configured concurrency appropriate for
your email provider's sending limits.

### Configuration

Campaign-specific settings, such as sender details, recipients, subject, and
SMTP or delivery options, belong in the `config/` package. Do not commit
passwords, API keys, or other secrets; use environment variables or a local
configuration file excluded from version control.

### Templates

Store reusable email bodies in `templates/`. Templates keep presentation
separate from delivery logic and can be updated without changing the campaign
runner.

## Code map



### Request flow

```text
main.go
	├── loads configuration from config/
	├── reads and renders a template from templates/
	├── creates one delivery task per recipient
	└── runs tasks concurrently and reports results
```

## Development

Format and validate the project before running it:

```bash
gofmt -w .
go test ./...
go run .
```

## License

Add your license information here.
