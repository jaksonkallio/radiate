package main

import (
	"bufio"
	"fmt"
	"github.com/jaksonkallio/radiate/internal/service"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Command struct {
	Verb      string
	Noun      string
	Arguments []string
	Fn        func(args map[string]string) error
}

var Commands = []Command{
	{
		Verb: "list",
		Noun: "library",
		Fn: func(args map[string]string) error {
			fmt.Printf("listing libraries now :)\n")
			return nil
		},
	},
	{
		Verb: "add",
		Noun: "library",
		Arguments: []string{
			"identifier",
		},
		Fn: func(args map[string]string) error {
			fmt.Printf("listing libraries now :) provided args: %#v\n", args)
			return nil
		},
	},
}

var ExitKeywords = []string{
	"exit",
	"quit",
	"escape",
	"leave",
	"done",
	"goodbye",
	"bye",
	"close",
}

func StartInteractiveCLI(service *service.Service) {
	stdInReader := bufio.NewReader(os.Stdin)
	var shouldExit bool

	for !shouldExit {
		inputStr, _ := stdInReader.ReadString('\n')
		inputStr = strings.TrimSpace(inputStr)

		if len(inputStr) == 0 {
			// Empty command, do nothing.
			continue
		}

		for _, exitKeyword := range ExitKeywords {
			if strings.ToLower(inputStr) == exitKeyword {
				shouldExit = true
				break
			}
		}

		if shouldExit {
			break
		}

		inputStrParts := strings.Split(inputStr, " ")

		if len(inputStrParts) < 2 {
			log.Warn().
				Str("input", inputStr).
				Msg("format \"noun verb [arg1 [arg2 [arg3 ...]]]\" expected")
			continue
		}

		var commandExists bool

		for _, command := range Commands {
			if strings.ToLower(inputStrParts[0]) == command.Noun && strings.ToLower(inputStrParts[1]) == command.Verb {
				commandExists = true
				args := inputStrParts[2:]

				if len(command.Arguments) != len(args) {
					log.Warn().
						Int("expectedArgumentCount", len(command.Arguments)).
						Int("providedArgumentCount", len(args)).
						Msg("did not provide the correct number of arguments")
					break
				}

				argMap := make(map[string]string)
				for i := range args {
					argMap[command.Arguments[i]] = args[i]
				}

				go func() {
					err := command.Fn(argMap)
					if err != nil {
						log.Error().
							Err(err).
							Msg("executing command failed")
					}
				}()

				break
			}
		}

		if !commandExists {
			log.Warn().
				Str("input", inputStr).
				Msg("command not found")
		}
	}
}
