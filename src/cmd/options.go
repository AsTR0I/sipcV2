package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/AsTR0I/sipcV2/internal/app"
	"github.com/AsTR0I/sipcV2/internal/infra/config"
	"github.com/AsTR0I/sipcV2/internal/logger"
	"github.com/AsTR0I/sipcV2/internal/transport/sip"
)

func init() {
	optionsCmd.Flags().SortFlags = false

	optionsCmd.Flags().String(
		"server-host",
		"",
		"SIP server host:port",
	)

	optionsCmd.Flags().String(
		"proxy",
		"",
		"SIP proxy",
	)

	optionsCmd.Flags().String(
		"from",
		"",
		"SIP From",
	)

	optionsCmd.Flags().String(
		"to",
		"",
		"SIP To",
	)

	optionsCmd.Flags().String(
		"username",
		"",
		"SIP username",
	)

	optionsCmd.Flags().String(
		"password",
		"",
		"SIP password",
	)

	optionsCmd.Flags().String(
		"realm",
		"",
		"SIP realm",
	)

	optionsCmd.Flags().Int(
		"expire",
		300,
		"Expire seconds",
	)

	optionsCmd.Flags().String(
		"user-agent",
		fmt.Sprintf(
			"%s/%s(%s/%s)",
			rootCmd.Use,
			rootCmd.Version,
			runtime.GOOS,
			runtime.GOARCH,
		),
		"User-Agent",
	)

	if err := optionsCmd.MarkFlagRequired("server-host"); err != nil {
		panic(err)
	}

	if err := optionsCmd.MarkFlagRequired("username"); err != nil {
		panic(err)
	}

	if err := optionsCmd.MarkFlagRequired("realm"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(optionsCmd)
}

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "SIP OPTIONS",
	Long:  "Performs SIP OPTIONS transaction",

	RunE: func(cmd *cobra.Command, args []string) error {

		serverHost, err := cmd.Flags().GetString("server-host")
		if err != nil {
			return err
		}

		proxy, err := cmd.Flags().GetString("proxy")
		if err != nil {
			return err
		}

		from, err := cmd.Flags().GetString("from")
		if err != nil {
			return err
		}

		to, err := cmd.Flags().GetString("to")
		if err != nil {
			return err
		}

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			return err
		}

		realm, err := cmd.Flags().GetString("realm")
		if err != nil {
			return err
		}

		expire, err := cmd.Flags().GetInt("expire")
		if err != nil {
			return err
		}

		userAgent, err := cmd.Flags().GetString("user-agent")
		if err != nil {
			return err
		}

		cfg := config.SIPConfig{
			ServerHost: serverHost,
			Proxy:      proxy,
			From:       from,
			To:         to,
			Username:   username,
			Password:   password,
			Realm:      realm,
			Expire:     expire,
			UserAgent:  userAgent,
		}

		ctx, stop := signal.NotifyContext(
			context.Background(),
			syscall.SIGINT,
			syscall.SIGTERM,
		)
		defer stop()

		log := logger.New(logLevel)

		sipClient, err := sip.NewClient(
			log,
			cfg.UserAgent,
		)
		if err != nil {
			return err
		}

		application := app.NewApp(
			log,
			sipClient,
		)

		return application.Options(ctx, cfg)
	},
}
