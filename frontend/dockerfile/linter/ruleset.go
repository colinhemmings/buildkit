package linter

import (
	"fmt"
)

var (
	RuleStageNameCasing = LinterRule[func(string) string]{
		NumericID:   "0001",
		Name:        "StageNameCasing",
		Description: "Stage names should be lowercase",
		URL: 				 "http://docs.docker.com/go/build-rules/stage-name-casing",
		Format: func(stageName string) string {
			return fmt.Sprintf("Stage name '%s' should be lowercase", stageName)
		},
	}
	RuleFromAsCasing = LinterRule[func(string, string) string]{
		NumericID:  "0002",
		Name:        "FromAsCasing",
		Description: "The 'as' keyword should match the case of the 'from' keyword",
		URL:				 "http://docs.docker.com/go/build-rules/from-as-casing",
		Format: func(from, as string) string {
			return fmt.Sprintf("'%s' and '%s' keywords' casing do not match", as, from)
		},
	}
	RuleNoEmptyContinuations = LinterRule[func() string]{
		NumericID:   "0003",
		Name:        "NoEmptyContinuations",
		Description: "Empty continuation lines will become errors in a future release",
		URL: 	  		 "http://docs.docker.com/go/build-rules/no-empty-continuations",
		Format: func() string {
			return "Empty continuation line"
		},
	}
	RuleConsistentInstuctionCasing = LinterRule[func(string) string]{
		NumericID:   "0004",
		Name:        "ConsistentInstructionCasing",
		Description: "Instructions should be in consistent casing (all lower or all upper)",
		URL: 			   "http://docs.docker.com/go/build-rules/consistent-instruction-casing",
		Format: func(command string) string {
			return fmt.Sprintf("Command '%s' should be consistently cased", command)
		},
	}
	//
	// I'm not sure why we need two instruction casing rules. Can we not combine them into one?
	//
	// RuleFileConsistentCommandCasing = LinterRule[func(string, string) string]{
	// 	Name:        "FileConsistentCommandCasing",
	// 	Description: "All commands within the Dockerfile should use the same casing (either upper or lower)",
	// 	Format: func(violatingCommand, correctCasing string) string {
	// 		return fmt.Sprintf("Command '%s' should match the case of the command majority (%s)", violatingCommand, correctCasing)
	// 	},
	// }
	RuleDuplicateStageName = LinterRule[func(string) string]{
		NumericID:   "0005",
		Name:        "DuplicateStageName",
		Description: "Stage names should be unique",
		URL:				 "http://docs.docker.com/go/build-rules/duplicate-stage-name",
		Format: func(stageName string) string {
			return fmt.Sprintf("Duplicate stage name %q, stage names should be unique", stageName)
		},
	}
	RuleReservedStageName = LinterRule[func(string) string]{
		NumericID:   "0006",
		Name:        "ReservedStageName",
		Description: "Reserved words should not be used as stage names",
		URL: 				 "http://docs.docker.com/go/build-rules/reserved-stage-name",
		Format: func(reservedStageName string) string {
			return fmt.Sprintf("%q is reserved and should not be ", reservedStageName)
		},
	}
	RuleJSONArgsRecommended = LinterRule[func(instructionName string) string]{
		NumericID:  "0007",
		Name:        "JSONArgsRecommended",
		Description: "JSON arguments recommended for ENTRYPOINT/CMD to prevent unintended behavior related to OS signals, unless using a custom shell",
		URL: 				 "http://docs.docker.com/go/build-rules/json-args-recommended",
		Format: func(instructionName string) string {
			return fmt.Sprintf("JSON arguments recommended for %s to prevent unintended behavior related to OS signals", instructionName)
		},
	}
	RuleMaintainerDeprecated = LinterRule[func() string]{
		NumericID:   "0008",
		Name:        "MaintainerDeprecated",
		Description: "The MAINTAINER instruction is deprecated, use a label instead to define an image author",
		URL: 				 "http://docs.docker.com/go/build-rules/maintainer-deprecated",
		Format: func() string {
			return "Maintainer instruction is deprecated in favor of using label"
		},
	}
	RuleUndefinedArgInFrom = LinterRule[func(string, string) string]{
		NumericID:   "0009",
		Name:        "UndefinedArgInFrom",
		Description: "FROM command must use declared ARGs",
		URL: 				 "http://docs.docker.com/go/build-rules/undefined-arg-in-from",
		Format: func(baseArg, suggest string) string {
			out := fmt.Sprintf("FROM argument '%s' is not declared", baseArg)
			if suggest != "" {
				out += fmt.Sprintf(" (did you mean %s?)", suggest)
			}
			return out
		},
	}
	RuleUndefinedArg = LinterRule[func(string) string]{
		NumericID:   "0010",
		Name:        "UndefinedArg",
		Description: "ARGs should be defined before their use",
		URL: 				 "http://docs.docker.com/go/build-rules/undefined-arg",
		Format: func(arg string) string {
			return fmt.Sprintf("Usage of undefined variable '$%s'", arg)
		},
	}
	RuleUndefinedVar = LinterRule[func(string, string) string]{
		NumericID:   "0011",
		Name:        "UndefinedVar",
		Description: "Variables should be defined before their use",
		URL: 				 "http://docs.docker.com/go/build-rules/undefined-var",
		Format: func(arg, suggest string) string {
			out := fmt.Sprintf("Usage of undefined variable '$%s'", arg)
			if suggest != "" {
				out += fmt.Sprintf(" (did you mean $%s?)", suggest)
			}
			return out
		},
	}

	RuleWorkdirRelativePath = LinterRule[func(workdir string) string]{
		NumericID:   "0012",
		Name:        "WorkdirRelativePath",
		Description: "Relative workdir without an absolute workdir declared within the build can have unexpected results if the base image changes",
		URL: 				 "http://docs.docker.com/go/build-rules/workdir-relative-path",
		Format: func(workdir string) string {
			return fmt.Sprintf("Relative workdir %q can have unexpected results if the base image changes", workdir)
		},
	}
	
	
	RuleMultipleInstructionsDisallowed = LinterRule[func(instructionName string) string]{
		NumericID:   "0013",
		Name:        "MultipleInstructionsDisallowed",
		Description: "Multiple instructions of the same type should not be used in the same stage",
		URL: 				 "http://docs.docker.com/go/build-rules/multiple-instructions-disallowed",
		Format: func(instructionName string) string {
			return fmt.Sprintf("Multiple %s instructions should not be used in the same stage because only the last one will be used", instructionName)
		},
	}

	RuleLegacyKeyValueFormat = LinterRule[func(cmdName string) string]{
		NumericID:   "0014",
		Name:        "LegacyKeyValueFormat",
		Description: "Legacy key/value format with whitespace separator should not be used",
		URL: 				 "http://docs.docker.com/go/build-rules/legacy-key-value-format",
		Format: func(cmdName string) string {
			return fmt.Sprintf("\"%s key=value\" should be used instead of legacy \"%s key value\" format", cmdName, cmdName)
		},
	}
)
