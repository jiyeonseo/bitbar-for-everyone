# bitbar-for-everyone

This is a skeleton repository that I used when I introduced a [bitbar](https://github.com/matryer/bitbar) to non-tech colleagues in the office. They don't need to know how to use [brew](https://brew.sh/) or move python files to an unknown path for the convenience of engineers. You can folk this and add more plugins under [`/plugins/`](./plugins) directory to publish your users. 

## Prerequisites
- macOS

## In the script [`./install.go`](./install.go)

- Put your own git repo address for `pluginURL` which contains plugins in `/plugins` directory.
  - In case you are using GitGub Enterprise, you need to use token and edit code like below. 
    ```go
    func downloadPlugins() []string {
	    result := downloadZip(pluginURL, "main.zip", true) // make here true
        return result
    }
    ```

## deploy the installation 

In [`install.go`](./install.go), it downloads Bitbar.app and this repo for move `plugins/` files into 

- build `install.go`

```sh
$ go build -o install-bitbar
```

- move it into version folder(ex. `0.0.1`) and zip the folder.

```sh
$ mkdir 0.0.1 && mv install-bitbar 0.0.1 & zip -r 0.0.1.zip 0.0.1
``` 

- release your latest version in the repository and attach the zipfile in the release decription.

(screenshot)

## notes 

Mac probably warns the user that the app can't be opened because it is from an unidentified developer. Inform the users that they can allow the app via `System Preferences` > `Security & Privacy` like below: 

![](https://user-images.githubusercontent.com/2231510/109412315-8b38dd80-79ea-11eb-9d91-8d30733d2e2d.png)
