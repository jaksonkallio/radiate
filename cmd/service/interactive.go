package main

import (
	"bufio"
	"github.com/jaksonkallio/radiate/internal/service"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Command struct {
	Name      string
	Arguments []string
	Fn        func(service *service.Service, args map[string]string) error
}

var ExitKeywords = []string{
	"exit",
	"quit",
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

		if len(inputStrParts) < 1 {
			log.Warn().
				Str("input", inputStr).
				Msg("format \"command [arg1 arg2...]\" expected")
			continue
		}

		var commandExists bool

		for _, command := range Commands {
			if strings.ToLower(inputStrParts[0]) == command.Name {
				commandExists = true
				args := inputStrParts[1:]

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
					err := command.Fn(service, argMap)
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
