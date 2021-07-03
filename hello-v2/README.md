# hello fyne v2

`Hello world` app with fyne v2.

## Desktop development

### run

```sh
go run main.go
```

## Mobile

### simulation

```sh
go run -tags mobile main.go
```

### packaging

Run the following to install the `fyne` command in advance.

```sh
go get fyne.io/fyne/v2/cmd/fyne
```

Run the following to create the `apk` file, and don't forget to create Icon.png beforehand.

```sh
fyne package -os android -appID t2wonderland.blogspot.com.hello_v2
```

### deploy to device

Connect the Android device to the PC with USB debugging enabled. Execute the following to install the application on the device.

```sh
fyne install -os android -appID t2wonderland.blogspot.com.hello_v2
```

