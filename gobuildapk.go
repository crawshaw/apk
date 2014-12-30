// A version of the Go tool's build command that produces an .apk if
// there is an AndroidManifest.xml file in the package directory.
//
// This is a temporary hack. The idea, if this works, is to merge this
// into the Go tool so `go build` just works.

//+build ignore

package main

import (
	"crypto/x509"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/crawshaw/apk"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	flag.Parse()
	pkg, err := ctx.Import(flag.Args()[0], cwd, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pkg)

	manifestPath := filepath.Join(pkg.Dir, "AndroidManifest.xml")
	if _, err := os.Stat(manifestPath); err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(2)
		}
		// Just do an ordinary build.
		cmd := exec.Command("go", "build", flag.Args()[0])
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			os.Exit(2)
		}
		os.Exit(0)
	}

	workPath, err := ioutil.TempDir("", "gobuildapk-work-")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(2)
	}
	defer os.RemoveAll(workPath)
	libPath := filepath.Join(workPath, "lib"+pkg.Name+".so")

	cmd := exec.Command(`go`, `build`, `-ldflags="-shared"`, `-o`, libPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{
		"CGO_ENABLED=1",
		"GOOS=android",
		"GOARCH=arm",
		"GOARM=7",
		"GOPATH=" + os.Getenv("GOPATH"),
	}
	if err := cmd.Run(); err != nil {
		os.Exit(2)
	}

	block, _ := pem.Decode([]byte(debugCert))
	if block == nil {
		log.Fatal("no cert")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(privKey)

	out, err := os.Create(filepath.Base(pkg.Dir) + ".apk")
	if err != nil {
		log.Fatal(err) // TODO: overwrite, and -o.
	}

	apkw := apk.NewWriter(out, privKey)

	manifestData, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		log.Fatal(err)
	}
	libName, err := manifestLibName(manifestData)
	if err != nil {
		log.Fatal(err)
	}
	w, err := apkw.Create("AndroidManifest.xml")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := w.Write(manifestData); err != nil {
		log.Fatal(err)
	}

	r, err := os.Open(libPath)
	if err != nil {
		log.Fatal(err)
	}
	w, err = apkw.Create("lib/armeabi/lib" + libName + ".so")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	// TODO: icons and such, maybe gdbserver.

	if err := apkw.Close(); err != nil {
		log.Fatal(err)
	}
}

// manifestLibName parses the AndroidManifest.xml and finds the library
// name of the NativeActivity.
func manifestLibName(data []byte) (string, error) {
	manifest := new(manifestXML)
	if err := xml.Unmarshal(data, manifest); err != nil {
		return "", err
	}
	if manifest.Activity.Name != "android.app.NativeActivity" {
		return "", fmt.Errorf("can only build an .apk for NativeActivity, not %q", manifest.Activity.Name)
	}
	libName := ""
	for _, md := range manifest.Activity.MetaData {
		if md.Name == "android.app.lib_name" {
			libName = md.Value
			break
		}
	}
	if libName == "" {
		return "", errors.New("AndroidManifest.xml missing meta-data android.app.lib_name")
	}
	return libName, nil
}

type manifestXML struct {
	Activity activityXML `xml:"application>activity"`
}

type activityXML struct {
	Name     string        `xml:"name,attr"`
	MetaData []metaDataXML `xml:"meta-data"`
}

type metaDataXML struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

var ctx = build.Default

// A random uninteresting private key.
const debugCert = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAy6ItnWZJ8DpX9R5FdWbS9Kr1U8Z7mKgqNByGU7No99JUnmyu
NQ6Uy6Nj0Gz3o3c0BXESECblOC13WdzjsH1Pi7/L9QV8jXOXX8cvkG5SJAyj6hcO
LOapjDiN89NXjXtyv206JWYvRtpexyVrmHJgRAw3fiFI+m4g4Qop1CxcIF/EgYh7
rYrqh4wbCM1OGaCleQWaOCXxZGm+J5YNKQcWpjZRrDrb35IZmlT0bK46CXUKvCqK
x7YXHgfhC8ZsXCtsScKJVHs7gEsNxz7A0XoibFw6DoxtjKzUCktnT0w3wxdY7OTj
9AR8mobFlM9W3yirX8TtwekWhDNTYEu8dwwykwIDAQABAoIBAA2hjpIhvcNR9H9Z
BmdEecydAQ0ZlT5zy1dvrWI++UDVmIp+Ve8BSd6T0mOqV61elmHi3sWsBN4M1Rdz
3N38lW2SajG9q0fAvBpSOBHgAKmfGv3Ziz5gNmtHgeEXfZ3f7J95zVGhlHqWtY95
JsmuplkHxFMyITN6WcMWrhQg4A3enKLhJLlaGLJf9PeBrvVxHR1/txrfENd2iJBH
FmxVGILL09fIIktJvoScbzVOneeWXj5vJGzWVhB17DHBbANGvVPdD5f+k/s5aooh
hWAy/yLKocr294C4J+gkO5h2zjjjSGcmVHfrhlXQoEPX+iW1TGoF8BMtl4Llc+jw
lKWKfpECgYEA9C428Z6CvAn+KJ2yhbAtuRo41kkOVoiQPtlPeRYs91Pq4+NBlfKO
2nWLkyavVrLx4YQeCeaEU2Xoieo9msfLZGTVxgRlztylOUR+zz2FzDBYGicuUD3s
EqC0Wv7tiX6dumpWyOcVVLmR9aKlOUzA9xemzIsWUwL3PpyONhKSq7kCgYEA1X2F
f2jKjoOVzglhtuX4/SP9GxS4gRf9rOQ1Q8DzZhyH2LZ6Dnb1uEQvGhiqJTU8CXxb
7odI0fgyNXq425Nlxc1Tu0G38TtJhwrx7HWHuFcbI/QpRtDYLWil8Zr7Q3BT9rdh
moo4m937hLMvqOG9pyIbyjOEPK2WBCtKW5yabqsCgYEAu9DkUBr1Qf+Jr+IEU9I8
iRkDSMeusJ6gHMd32pJVCfRRQvIlG1oTyTMKpafmzBAd/rFpjYHynFdRcutqcShm
aJUq3QG68U9EAvWNeIhA5tr0mUEz3WKTt4xGzYsyWES8u4tZr3QXMzD9dOuinJ1N
+4EEumXtSPKKDG3M8Qh+KnkCgYBUEVSTYmF5EynXc2xOCGsuy5AsrNEmzJqxDUBI
SN/P0uZPmTOhJIkIIZlmrlW5xye4GIde+1jajeC/nG7U0EsgRAV31J4pWQ5QJigz
0+g419wxIUFryGuIHhBSfpP472+w1G+T2mAGSLh1fdYDq7jx6oWE7xpghn5vb9id
EKLjdwKBgBtz9mzbzutIfAW0Y8F23T60nKvQ0gibE92rnUbjPnw8HjL3AZLU05N+
cSL5bhq0N5XHK77sscxW9vXjG0LJMXmFZPp9F6aV6ejkMIXyJ/Yz/EqeaJFwilTq
Mc6xR47qkdzu0dQ1aPm4XD7AWDtIvPo/GG2DKOucLBbQc2cOWtKS
-----END RSA PRIVATE KEY-----
`
