package authentication

import (
    "bufio"
    "errors"
    "fmt"
    "github.com/mailhedgehog/MailHedgehog/logger"
    "golang.org/x/crypto/bcrypt"
    "os"
    "strings"
)

type userInfo struct {
    username string
    httpPass string
    smtpPass string
}

// FileAuth represents the authentication handler using file
type FileAuth struct {
    filePath string
    users    map[string]userInfo
}

func CreateFileAuthentication(authFilePath string) *FileAuth {
    authFile := &FileAuth{
        filePath: authFilePath,
    }
    authFile.authFile()

    return authFile
}

func (fileAuth *FileAuth) RequiresAuthentication() bool {
    return fileAuth.users != nil
}

func (fileAuth *FileAuth) Authenticate(authType AuthenticationType, username string, password string) bool {
    if !fileAuth.RequiresAuthentication() {
        return true
    }

    user, ok := fileAuth.users[username]
    if !ok {
        return false
    }

    passwordHashToCheck := user.httpPass
    if authType == SMTP {
        passwordHashToCheck = user.smtpPass
    }

    if err := bcrypt.CompareHashAndPassword([]byte(passwordHashToCheck), []byte(password)); err != nil {
        return false
    }

    return true
}

// authFile scan file and add users to memory
func (fileAuth *FileAuth) authFile() int {
    fileAuth.users = nil

    if len(fileAuth.filePath) <= 0 {
        logManager().Debug("File auth empty.")
        return 0
    }

    file, err := os.Open(fileAuth.filePath)
    logger.PanicIfError(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        err := fileAuth.addUserFromFileLine(scanner.Text())
        if err != nil {
            logManager().Notice(err.Error())
        } else {
        }
    }

    if fileAuth.users == nil {
        return 0
    }
    return len(fileAuth.users)
}

func (fileAuth *FileAuth) initUsers() {
    if fileAuth.users == nil {
        fileAuth.users = make(map[string]userInfo)
    }
}

func (fileAuth *FileAuth) addUserFromFileLine(line string) error {
    line = strings.TrimSpace(line)
    infoSlice := strings.Split(line, ":")
    for i := range infoSlice {
        infoSlice[i] = strings.TrimSpace(infoSlice[i])
    }

    if len(infoSlice) < 2 {
        return errors.New("at least should be present username and password")
    }

    if len(infoSlice[0]) <= 0 || len(infoSlice[1]) <= 0 {
        return errors.New("username and password can't be empty")
    }

    smtpPass := infoSlice[1]
    if len(infoSlice) > 2 && len(infoSlice[2]) > 0 {
        smtpPass = infoSlice[2]
    }

    fileAuth.initUsers()

    fileAuth.users[infoSlice[0]] = userInfo{
        username: infoSlice[0],
        httpPass: infoSlice[1],
        smtpPass: smtpPass,
    }

    logManager().Debug(fmt.Sprintf("Processes users: '%s'", infoSlice[0]))

    return nil
}

func (fileAuth *FileAuth) AddUser(username string, httpPassHash string, smtpPassHash string) error {
    if len(username) <= 0 || len(httpPassHash) <= 0 {
        return errors.New("username and httpPassHash required")
    }

    fileAuth.initUsers()

    fileAuth.users[username] = userInfo{
        username: username,
        httpPass: httpPassHash,
        smtpPass: smtpPassHash,
    }

    return fileAuth.writeToFile()
}

func (fileAuth *FileAuth) UpdateUser(username string, httpPassHash string, smtpPassHash string) error {
    if len(username) <= 0 || len(httpPassHash) <= 0 {
        return errors.New("username and httpPassHash required")
    }

    fileAuth.initUsers()

    newHttpPassHash := fileAuth.users[username].httpPass
    if len(httpPassHash) > 0 {
        newHttpPassHash = httpPassHash
    }

    newSmtpPassHash := fileAuth.users[username].smtpPass
    if len(smtpPassHash) > 0 {
        newSmtpPassHash = smtpPassHash
    }

    fileAuth.users[username] = userInfo{
        username: username,
        httpPass: newHttpPassHash,
        smtpPass: newSmtpPassHash,
    }

    return fileAuth.writeToFile()
}

func (fileAuth *FileAuth) DeleteUser(username string) error {
    delete(fileAuth.users, username)

    return fileAuth.writeToFile()
}

func (fileAuth *FileAuth) ListUsers(searchQuery string, offset, limit int) ([]UserResource, int, error) {
    keys := make([]string, 0, len(fileAuth.users))
    for k, _ := range fileAuth.users {
        if len(searchQuery) > 0 {
            if strings.Contains(k, searchQuery) {
                keys = append(keys, k)
            }
        } else {
            keys = append(keys, k)
        }
    }

    endIndex := len(keys)
    if offset+limit < len(keys) {
        endIndex = offset + limit
    }
    if offset < 0 || offset > endIndex {
        offset = 0
    }
    slice := keys[offset:endIndex]
    var resources []UserResource
    for _, username := range slice {
        resources = append(resources, UserResource{Username: username})
    }

    return resources, len(keys), nil
}

func (fileAuth *FileAuth) writeToFile() error {
    file, err := os.OpenFile(fileAuth.filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        return err
    }
    defer file.Close()

    file.Truncate(0)
    for _, userInfo := range fileAuth.users {
        smtpPass := ""
        if userInfo.httpPass != userInfo.smtpPass {
            smtpPass = userInfo.smtpPass
        }
        _, err := file.WriteString(fmt.Sprintf(
            "%s:%s:%s\n",
            userInfo.username,
            userInfo.httpPass,
            smtpPass,
        ))
        if err != nil {
            return err
        }
    }

    return nil
}

func (fileAuth *FileAuth) UsernamePresent(username string) bool {
    _, ok := fileAuth.users[username]

    return ok
}
