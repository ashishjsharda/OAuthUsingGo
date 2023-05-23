# OAuth Authentication Example using Golang

This repository contains an example implementation of OAuth authentication using Golang. It demonstrates how to perform the OAuth flow, exchange an authorization code for an access token, and make authenticated API requests using the obtained access token.

## Prerequisites

Before running the code, make sure you have the following prerequisites:

- Golang installed on your system. You can download it from the official website: https://golang.org/dl/
- OAuth client credentials (client ID and client secret) from the OAuth provider you are integrating with.

## Configuration

1. Create a JSON file named `config.json` with the following structure:

   ```json
   {
     "client_id": "your_client_id",
     "client_secret": "your_client_secret",
     "redirect_url": "http://localhost:8080/callback"
   }
Replace "your_client_id", "your_client_secret", and "http://localhost:8080/callback" with your actual OAuth client credentials and redirect URL.
