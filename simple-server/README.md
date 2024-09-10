# SignIn

```bash
 curl --location 'localhost:8080/signIn' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'username=test1' --data-urlencode 'password=testtest'
```

# Sign Up

```bash
 curl --location 'localhost:8080/signUp' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'username=test1' --data-urlencode 'password=testtest'
```

# UpdateInfo

```bash
curl --location 'localhost:8080/updateInfo' --form 'username=test' --form 'password=testtest' --form 'info=some_info' --form 'file=@/Users/macos/Desktop/screenshots/dummy.png'
```
