package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedExtension = "vlc"

var ErrEmptyPath = errors.New("path to file is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	data, err := ioutil.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	// packed := Encode(data)
	packed := ""
	fmt.Println(string(data))
	err = ioutil.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}
}

func packedFileName(path string) string {

	filename := filepath.Base(path)

	return strings.TrimSuffix(filename, filepath.Ext(filename)) + "." + packedExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
