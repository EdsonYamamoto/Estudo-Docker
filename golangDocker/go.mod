module main

        go 1.14

        require (
        controllers v0.0.0
        github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
        github.com/labstack/echo v3.3.10+incompatible
        github.com/labstack/gommon v0.3.0 // indirect
        golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9 // indirect
        models v0.0.0 // indirect
        )

        replace (
        controllers => ./controllers
        models => ./models
        )
