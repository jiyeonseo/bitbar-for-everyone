package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const pluginURL = "https://github.com/jiyeonseo/bitbar-for-everyone/archive/main.zip"
const bitbarURL = "https://github.com/matryer/bitbar/releases/download/v1.9.2/BitBar-v1.9.2.zip"

const token = ""

func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		fpath := filepath.Join(dest, f.Name)
		
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func downloadZip(url string, name string, needAuth bool) []string {
    req, err := http.NewRequest("GET", url, nil)
    if needAuth {
        req.Header.Set("Authorization", "token "+token)
    }
	
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil
	}

	out, err := os.Create(name)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	fmt.Printf("err: %s", err)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %s", err)
	}

	files, err := Unzip(name, dir)

	if err != nil {
		fmt.Printf("err: %s", err)
	}

    return files
}

func downloadPlugins() []string {
	result := downloadZip(pluginURL, "main.zip", false)
    return result
}

func downloadBitbar() []string {
    result := downloadZip(bitbarURL, "bitbar.zip", false)
    return result
}

func main() {
	plugins := downloadPlugins()
	
	home, _  := os.UserHomeDir()
    pluginLocation := plugins[0] + "/plugins/"
    configLocation := home+"/Documents/bitbar/"
	os.RemoveAll(configLocation)
	rerr := os.Rename(pluginLocation, configLocation)
	if rerr != nil {
		fmt.Printf("Rename: %s\n", rerr)
	}

	cerr := os.Chdir(home+"/Downloads")
	if cerr != nil {
		fmt.Printf("Chdir err: %s\n", cerr)
	}

	bitbar := downloadBitbar()

	cmd := exec.Command("open", bitbar[0])
	stdout, oerr := cmd.Output()

	if oerr != nil {
		fmt.Printf("Open err: %s\n", oerr)
	}
	fmt.Printf(string(stdout))
}
