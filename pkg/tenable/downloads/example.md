# Tenable Downloads API Client Usage

This document provides an example of how to use the Tenable Downloads API client in a Go application.

## Setup

First, ensure that you have the Tenable Downloads API client implemented in your project, as detailed in `pkg/tenable/downloads/client.go`.

## Example Usage

Below is a basic example of how to use the client to interact with the Tenable Downloads API.

```go
package main

import (
    "context"
    "fmt"
    "tenable-tf/pkg/tenable/downloads"
)

func main() {
    client := downloads.NewClient("YourAuthToken")

    // List all product pages
    pages, err := client.ListProductPages(context.Background())
    if err != nil {
        fmt.Println("Error listing product pages:", err)
        return
    }
    fmt.Println("Product Pages:", string(pages))

    // List downloads for a specific product (example: 'nessus')
    downloads, err := client.ListDownloadsForProduct(context.Background(), "nessus")
    if err != nil {
        fmt.Println("Error listing downloads for product:", err)
        return
    }
    fmt.Println("Downloads for Product:", string(downloads))

    // Download a specific file (commented out for example purposes)
    // fileContent, err := client.DownloadFile(context.Background(), "nessus", "Nessus-latest-x64.msi")
    // if err != nil {
    //     fmt.Println("Error downloading file:", err)
    //     return
    // }
    // fmt.Println("Downloaded File Content:", string(fileContent))
}
```

## Notes

- Replace `"YourAuthToken"` with an actual bearer token provided by Tenable.
- The `ListProductPages` and `ListDownloadsForProduct` methods are used in this example.
- The `DownloadFile` method is commented out but can be used similarly to download specific files.

Remember to handle the API responses appropriately in your application, potentially parsing the JSON response into Go structs for easier data manipulation.
