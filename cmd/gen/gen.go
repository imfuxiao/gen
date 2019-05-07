package gen

import (
	"fmt"
	"gen/model"
	"gen/template"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

var outPath string
var yamlPath string

var rootCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen是代码生成器",
	Long:  `目前仅支持mongo代码生成`,
	Run: func(cmd *cobra.Command, args []string) {
		checkParam()
		model, err := model.Read(yamlPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = template.ExecuteModelTemplate(outPath+"/model", model, "entity")
		if err != nil {
			fmt.Println("entity generator error: ", err.Error())
			os.Exit(1)
		}
		err = template.ExecuteModelTemplate(outPath+"/repository", model, "repository")
		if err != nil {
			fmt.Println("repository generator error: ", err.Error())
			os.Exit(1)
		}
		err = template.ExecuteModelTemplate(outPath+"/services", model, "service")
		if err != nil {
			fmt.Println("service generator error: ", err.Error())
			os.Exit(1)
		}
		err = template.ExecuteModelTemplate(outPath+"/controller/"+model.Version, model, "controller")
		if err != nil {
			fmt.Println("controller generator error: ", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	pflag.StringVar(&outPath, "o", "", "代码输出路径")
	pflag.StringVar(&yamlPath, "i", "", "yaml模型文件路径")
	pflag.Parse()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkParam() {
	if outPath == "" {
		fmt.Println("outPath is null")
		os.Exit(1)
	}
	if yamlPath == "" {
		fmt.Println("yamlPath is null")
		os.Exit(1)
	}
}
