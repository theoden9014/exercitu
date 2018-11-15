package cmd

import (
	"github.com/spf13/cobra"

	eclient "github.com/theoden9014/exercitu/client"
	"github.com/theoden9014/exercitu/worker"
)

const (
	defaultHttpPort = 8080
)

var (
	httpPort int
	file     string
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "http mode",
	Long:  `HTTP mode`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := http(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")
	httpCmd.PersistentFlags().IntVarP(&httpPort, "port", "p", defaultHttpPort, "port number")
	httpCmd.PersistentFlags().StringVarP(&file, "file", "F", "", "send file")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func http() error {
	switch role {
	// case serverRole:
	// 	s := eserver.NewHTTP()

	// 	if err := s.Run(); err != nil {
	// 		return err
	// 	}
	case clientRole:
		w := worker.NewHTTP()
		c := eclient.New()
		c.SetWorker(w)

		if err := c.Run(); err != nil {
			return err
		}
	}

	return nil
}
