package command

import (
	"github.com/hiroxto/epgstation-slack-notification/pkg/config"
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
	configFilePath := getConfigFilePath(context)
	configData, err := readConfigFile(configFilePath)
	if err != nil {
		return err
	}

	config, err := config.LoadConfigFromYaml([]byte(configData))
	if err != nil {
		return err
	}

	pp.Default.SetColoringEnabled(context.Bool("color"))
	pp.Println(config)

	return nil
}
