package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/AsTR0I/sipcV2/internal/logger"
	"github.com/AsTR0I/sipcV2/internal/sip"
)

var (
	serverHost    string
	serverPort    int
	from          string
	to            string
	username      string
	password      string
	realm         string
	expire        int
	userAgentName string
)

func init() {
	registerCmd.Flags().SortFlags = false
	registerCmd.Flags().StringVar(
		&serverHost,
		"server-host",
		"",
		"SIP server host",
	)

	registerCmd.Flags().IntVar(
		&serverPort,
		"server-port",
		5060,
		"SIP server port",
	)

	registerCmd.Flags().StringVar(
		&from,
		"from",
		"",
		"SIP From URI",
	)

	registerCmd.Flags().StringVar(
		&to,
		"to",
		"",
		"SIP To URI",
	)

	registerCmd.Flags().StringVar(
		&username,
		"username",
		"",
		"SIP auth username",
	)

	registerCmd.Flags().StringVar(
		&password,
		"password",
		"",
		"SIP password",
	)

	registerCmd.Flags().StringVar(
		&realm,
		"realm",
		"",
		"SIP realm",
	)

	registerCmd.Flags().IntVar(
		&expire,
		"expire",
		300,
		"Register expire",
	)

	registerCmd.Flags().StringVar(
		&userAgentName,
		"UAName",
		fmt.Sprintf("%s/%s", rootCmd.Use, rootCmd.Version),
		"User Agent name",
	)

	if err := registerCmd.MarkFlagRequired("server-host"); err != nil {
		panic(err)
	}
	// if err := registerCmd.MarkFlagRequired("server-port"); err != nil {
	// 	panic(err)
	// }
	if err := registerCmd.MarkFlagRequired("from"); err != nil {
		panic(err)
	}
	if err := registerCmd.MarkFlagRequired("to"); err != nil {
		panic(err)
	}
	if err := registerCmd.MarkFlagRequired("username"); err != nil {
		panic(err)
	}
	if err := registerCmd.MarkFlagRequired("realm"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "SIP REGISTER",
	Long:  "Performs SIP REGISTER transaction with authentication",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, stop := signal.NotifyContext(
			context.Background(),
			syscall.SIGINT,
			syscall.SIGTERM,
		)
		defer stop()

		log := logger.New(logLevel)

		client, err := sip.NewClient(
			log,
			sip.Options{
				ServerHost:    serverHost,
				ServerPort:    serverPort,
				From:          from,
				To:            to,
				Username:      username,
				Password:      password,
				Realm:         realm,
				Expire:        expire,
				UserAgentName: userAgentName,
			},
		)
		if err != nil {
			return err
		}

		return client.Register(ctx)
	},
}
