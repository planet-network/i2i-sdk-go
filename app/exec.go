package app

import (
	"fmt"
	"os/exec"
)

type Executor struct {
	graphQlPort    int
	workdir        string
	keychainSource string
	i2iPath        string
	cmd            *exec.Cmd
}

type ExecutorOpts struct {
	GraphQlPort    int
	I2IPath        string
	Workdir        string
	KeychainSource string
}

func NewExecutor(opt *ExecutorOpts) *Executor {
	return &Executor{
		graphQlPort:    opt.GraphQlPort,
		workdir:        opt.Workdir,
		keychainSource: opt.KeychainSource,
		i2iPath:        opt.I2IPath,
		cmd:            nil,
	}
}

type ConsoleOutput struct {
}

func (c *ConsoleOutput) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

func (e *Executor) Run() error {
	e.cmd = exec.Command(e.i2iPath, "listen", "--debug",
		"--keychain-source", e.keychainSource,
		"--graphql-addr", fmt.Sprintf(":%d", e.graphQlPort),
		"--workdir", e.workdir)

	e.cmd.Stdout = &ConsoleOutput{}
	e.cmd.Stderr = &ConsoleOutput{}

	fmt.Println(e.cmd.String())

	return e.cmd.Run()
}

func (e *Executor) Stop() error {
	return e.cmd.Process.Kill()
}

func (a *App) Executor(name string, port int) (*Executor, error) {
	node, err := a.NodeByName(name)
	if err != nil {
		return nil, err
	}

	if node.Meta.LocalExec.I2IPath == "" {
		return nil, fmt.Errorf("missing i2i path")
	}

	executor := NewExecutor(&ExecutorOpts{
		GraphQlPort:    port,
		Workdir:        a.workdirPath(name),
		KeychainSource: a.keychainSourceFile(name),
		I2IPath:        node.Meta.LocalExec.I2IPath,
	})

	return executor, nil
}
