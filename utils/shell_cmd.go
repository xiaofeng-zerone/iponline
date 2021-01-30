package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

const SHELL_CMD_EXEC_FAILED = "shell cmd exec failed"

// ShellCmdExec exec shell command
func ShellCmdExec(commandName string, params []string) (string, bool) {
	var ret_str string
	cmd := exec.Command(commandName, params...)
	fmt.Println("cmd:", cmd.Args)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd stdoutpipe err:", err)
		return SHELL_CMD_EXEC_FAILED, false
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("cmd start err:", err)
		return SHELL_CMD_EXEC_FAILED, false
	}
	var stderr bytes.Buffer
	cmd.Stdout = &stderr
	cmd.Stderr = &stderr
	reader := bufio.NewReader(stdout)
	for {
		ret, err := reader.ReadString('\n')
		ret_str += ret
		if err != nil && io.EOF != err {
			fmt.Println("cmd read string err:", err)
			return SHELL_CMD_EXEC_FAILED, false
		}
		if io.EOF == err {
			break
		}
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("cmd wait err:", err, stderr.String())
		return ret_str, false
	}
	return ret_str, true
}


// ShellCmdExecBySh exec shell command
func ShellCmdExecBySh(cmd string) (string, bool) {
	param := []string{"-c", cmd}
	cmd_ret, ok := ShellCmdExec("sh", param)
	return strings.Trim(cmd_ret, "\n"), ok
}

