// Copyright © 2018 Anduin Transactions Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/anduintransaction/raven/api/raven/admin"
	"github.com/anduintransaction/raven/api/raven/config"
	"github.com/anduintransaction/raven/api/raven/database"
	"github.com/anduintransaction/raven/api/raven/mailgun"
	"github.com/anduintransaction/raven/api/raven/servers"
	"github.com/anduintransaction/raven/api/raven/smtpserver"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var uiData string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [config file]",
	Short: "Run raven servers",
	Long:  "Run raven servers",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.ParseConfig(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		setupLogger(config.Logging)
		connectDatabase(config.Database)
		startServers(config)
		database.Close()
	},
}

func setupLogger(loggingConfig *config.LoggingConfig) {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	switch strings.ToUpper(loggingConfig.Level) {
	case "DEBUG":
		logrus.SetLevel(logrus.DebugLevel)
	case "INFO":
		logrus.SetLevel(logrus.InfoLevel)
	case "WARN":
		logrus.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logrus.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logrus.SetLevel(logrus.FatalLevel)
	case "PANIC":
		logrus.SetLevel(logrus.PanicLevel)
	}
	switch loggingConfig.Output {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "stderr":
		logrus.SetOutput(os.Stderr)
	default:
		output, err := os.Create(loggingConfig.Output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot create log file %s\n, error: %s", loggingConfig.Output, err)
			os.Exit(1)
		}
		logrus.SetOutput(output)
	}
}

func connectDatabase(databaseConfig *config.DatabaseConfig) {
	err := database.Connect(databaseConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startServers(config *config.Config) {
	servers := servers.NewServers()
	servers.AddServer("admin", admin.NewAPIServer(config.Admin, uiData))
	servers.AddServer("mailgun", mailgun.NewAPIServer(config.Mailgun))
	servers.AddServer("smtp", smtpserver.NewSMTPServer(config.SMTPServer))
	servers.ListenAndServe()
}

func init() {
	RootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVar(&uiData, "ui-data", "", "frontend folder")
}
