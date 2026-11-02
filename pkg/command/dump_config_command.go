package command

import (
	"context"
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
	"github.com/k0kubun/pp/v3"
	"github.com/urfave/cli/v3"
)

// DumpConfigCommand dump-config コマンド
var DumpConfigCommand = &cli.Command{
	Name:  "dump:config",
	Usage: "設定を出力するデバッグ用コマンド",
	Description: `
   設定を出力するデバッグ用コマンド
`,
	Action: dumpConfigCommandAction,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "color",
			Value: false,
		},
	},
}

func dumpConfigCommandAction(_ context.Context, cmd *cli.Command) error {
	fileService := service.NewFileService()
	configData, err := fileService.ReadConfigFile(cmd.String("config"))
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml(configData)
	if err != nil {
		return err
	}

	pp.Default.SetColoringEnabled(cmd.Bool("color"))
	pp.Println(conf)

	return nil
}
