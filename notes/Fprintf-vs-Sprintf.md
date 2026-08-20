# Fprintf vs Sprintf

## fmt.Sprintf

- **What it does:** String print
- **Where the output goes:** Returns a brand new string variable back to the code
- **Example use case:** Creating a dynamic error message or building a URL string

## fmt.Fprintf

- **What it does:** File/Stream print
- **Where the output goes:** Sends the text straight to a stream (an `io.Writer`)
- **Example use case:** Logging text to a file or writing HTML data to a web browser
