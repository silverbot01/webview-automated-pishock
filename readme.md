Go/templ/webview/viper desktop app transformation for Lisa's automated shocking app.

Viper is utilized for reading the .yaml config file. Multiple sharecodes can be entered as shown in the sample config.yaml. Ensure its in the same directory as autoshock.exe

Dev notes:

Utilize "go build -ldflags "-linkmode 'external' -extldflags '-static'" -o bin_deploy_static/autoshock.exe" to deploy a statically-linked executable for portability. Deploy with this readme and the sample config file