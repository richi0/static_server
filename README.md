# Static file server

A simple static file server with basic auth and automatic shutdown after the given time.

## Usage

```
  -folder string
        Folder to serve from (default ".")
  -password string
        BasicAuth password (default "password")
  -port string
        Port number (default "4000")
  -runtime int
        Time the serve will be available in minutes (0 == always)
  -username string
        BasicAuth username (default "username")
```
