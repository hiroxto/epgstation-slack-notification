package command

import (
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
	"github.com/hiroxto/epgstation-slack-notification/pkg/service"
	"github.com/k0kubun/pp/v3"
	"github.com/urfave/cli/v2"
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

func dumpConfigCommandAction(context *cli.Context) error {
	fileService := service.NewFileService()
	configData, err := fileService.ReadConfigFile(context.String("config"))
	if err != nil {
		return err
	}

	conf, err := config.LoadConfigFromYaml(configData)
	if err != nil {
		return err
	}

	pp.Default.SetColoringEnabled(context.Bool("color"))
	pp.Println(conf)

	return nil
}
