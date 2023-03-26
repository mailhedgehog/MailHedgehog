## Parse email

Parse email string to appropriate struct for easy access data and manipulate it.
More information about fields you can find [here](rfc5322.txt)

```go
email, err := email.Parse(ioReaderWithEmailData)
// or
email, err := email.Parse(strings.NewReader(emailRawStringData))
if err != nil {
    // handle error
}

fmt.Print(email.Subject)
```
