package make

import (
	cicd "github.com/NoUseFreak/cicd/pkg/helper"
)

func stepMake() *cicd.Step {
	return &cicd.Step{
		Schema: map[string]*cicd.Schema{
			"target": {
				Type:     cicd.TypeString,
				Required: true,
			},
			"arguments": {
				Type:     cicd.TypeString,
				Required: false,
			},
		},
		Run: func(data *cicd.RunArguments, ctx *cicd.Context) (*cicd.StepResponse, error) {
			cmdParts := []string{}
			if target := data.GetString("target"); target != "" {
				cmdParts = append(cmdParts, target)
			}
			if args := data.GetString("arguments"); args != "" {
				cmdParts = append(cmdParts, args)
			}

			return &cicd.StepResponse{}, ctx.HandleExitMessage(ctx.RunCommand("make", cmdParts...), makeExitMessages)
		},
	}
}
