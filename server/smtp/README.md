## SMTP server

Has very simple

```go
// Start server
go smtp.Listen(context, exitChannel)
// And close on demand (by default it will ve closed on close application)
smtp.Close()
```

All required data an configuration goes through `context` variable
