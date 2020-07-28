package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strings"
)

type IExecutor interface {
	name() string
	input(in string)
	subExecutors() []IExecutor
	setUpper(IExecutor)
	getUpper() IExecutor
	getHistories() []string
	completer(doc prompt.Document) []prompt.Suggest
}

type BaseExecutor struct {
	upper     IExecutor
	histories []string
}

func (e *BaseExecutor) name() string {
	return "base"
}

func (e *BaseExecutor) setUpper(executor IExecutor) {
	e.upper = executor
}

func (e *BaseExecutor) getUpper() IExecutor {
	return e.upper
}

func (e *BaseExecutor) getHistories() []string {
	return e.histories
}

var LivePrefixState struct {
	LivePrefix string
	IsEnable   bool
}

func changeLivePrefix() (string, bool) {
	return LivePrefixState.LivePrefix, LivePrefixState.IsEnable
}

type mainExecutor struct {
	BaseExecutor
	Subs []IExecutor
}

func (c *mainExecutor) name() string {
	return "main"
}

func (c *mainExecutor) subExecutors() []IExecutor {
	return c.Subs
}

func (c *mainExecutor) input(in string) {
	// fmt.Println("Your input: " + in)
	c.histories = append(c.histories, in)
	if in == "" {
		LivePrefixState.IsEnable = false
		LivePrefixState.LivePrefix = in
		return
	}
	se, ok := parseInputSubExecutor(c, in)
	if ok {
		se.input(in)
		return
	} else {
		LivePrefixState.LivePrefix = in + "> "
		LivePrefixState.IsEnable = false
	}

}

func (c *mainExecutor) completer(doc prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ssh", Description: "make a ssh connection to some server"},
		{Text: "history", Description: "show history in current session"},
	}

	se, ok := parseInputSubExecutor(c, doc.Text)
	if !ok {
		return prompt.FilterHasPrefix(s, doc.GetWordBeforeCursor(), true)
	} else {
		// sub commands
		// fmt.Println(doc.GetWordBeforeCursor())
		return se.completer(doc)
	}

}

func parseInputSubExecutor(executor IExecutor, in string) (se IExecutor, ok bool) {
	toWords := strings.Split(in, " ")

	for _, se := range executor.subExecutors() {
		if se.name() == toWords[0] {
			return se, true
		}
	}

	return nil, false
}

type sshExecutor struct {
	BaseExecutor
}

func (c *sshExecutor) name() string {
	return "ssh"
}

func (e *sshExecutor) subExecutors() []IExecutor {
	return []IExecutor{}
}

func (c *sshExecutor) input(in string) {

}
func (e *sshExecutor) completer(doc prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "root@192.168.1.109"},
		{Text: "dongchaofeng@192.168.1.249"},
	}
	return prompt.FilterHasPrefix(s, doc.GetWordBeforeCursor(), true)
}

type historyExecutor struct {
	BaseExecutor
}

func (c *historyExecutor) name() string {
	return "history"
}

func (c *historyExecutor) subExecutors() []IExecutor {
	return []IExecutor{}
}

func (c *historyExecutor) input(in string) {
	var histories = c.getUpper().getHistories()
	for _, cmdStr := range histories {
		fmt.Println(cmdStr)
	}
}

func (c *historyExecutor) completer(doc prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
	}
	return prompt.FilterHasPrefix(s, doc.GetWordBeforeCursor(), true)
}

func main() {
	var cmd = new(mainExecutor)
	cmd.Subs = []IExecutor{
		new(sshExecutor),
		new(historyExecutor),
	}
	for _, se := range cmd.Subs {
		se.setUpper(cmd)
	}
	pt := prompt.New(
		cmd.input,
		cmd.completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionLivePrefix(changeLivePrefix),
		prompt.OptionTitle("my-cli"),
		prompt.OptionHistory(cmd.histories),
	)
	pt.Run()
}
