# Go Link Shortener

A simple link shortener written in GO


## Features

- Makes links shorter


## Prerequisites

- [GoLang](https://go.dev/doc/install)
- [Sqlite3](https://www.sqlite.com/)

## Setup

Follow these steps to set up the project locally.

1. **Clone the repository:**

   ```bash
   git clone https://github.com/tsugent/link_shortener_go.git
   ```

2. **Navigate to directory:**

   ```bash
   cd ./link_shortener_go/
   ```

3. **Build/Run**
    Production:
    ```bash
    go build -o go-url-short .
    ./go-url-short
    ```

    Development:
    ```bash
    go run .
    ```

## Usage

1. **Create a Shortened Link:**
    
    #POST:


    ```bash
    curl -X POST localhost:8080/shorten -d '{"originalURL": "https://facebook.com", "created": "10/19/2021 10:19:01"}' 
    ```

    #STDOUT:


    ```bash
    "{\"shortened_link\": \"http://localhost:8080/short/?r=e70f7e67\"}"
    ```

2. **Retrieve URL to be directed to:**
    
    #GET:


    ```bash
        curl -X GET "http://localhost:8080/short/?r=e70f7e67" 
    ```
    
    #STDOUT:


    ```bash
        "{\"original_url\": \"https://facebook.com\"}"
    ```

## Contributing

I don't know why you'd do something like that.