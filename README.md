# A simple CLI password manager

---

This CLI application can be used to manage passwords. To begin using the application clone the app and
an "app.env" file at the root folder. Inside the **_app.env_** file add an environment variable **SECRET_KEY**
of length **32bit**. To start using the program run: "go build". You can start adding your passwords after registering
using the **safe register** cmd

**Example**: SECRET_KEY=blabasdfblsbalsadfrtasdflnfasdbf

## Commands

- **add**: Add's a password
- **completion**: Generate the autocompletion script for the specified shell
- **delete**: Delete password
- **help**: Help about any command
- **list**: List passwords after logging in
- **register**: Register as a user

## Examples

**Example**: safe register -u=username -p=password

**Example**: safe add -u=username -p=password -e=1@2.com -n=random text

**Example**: safe list -l=username

**Example**: safe delete 1
