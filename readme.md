Go/templ/webview desktop app transformation for Lisa's automated shocking app.

Please run in command line with the following arguments in order: "username" "shareCode" "api-key".

Example starting command: ./autoshock.exe "cutepuppy" "345ADAD" "d9DADDYa9da8-wda0dwa8da0d08-wdada80da8d0-wada"

Can add a fourth optional argument to change your nickname sent with API calls from the default of "Autoshock-Default"

Dev notes:

Utilize "go build -ldflags "-linkmode 'external' -extldflags '-static'" -o bin_deploy_static/autoshock.exe" to deploy a statically-linked executable for portability