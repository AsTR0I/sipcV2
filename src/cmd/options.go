package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/AsTR0I/sipcV2/internal/app"
	"github.com/AsTR0I/sipcV2/internal/domain"
	"github.com/AsTR0I/sipcV2/internal/logger"
	"github.com/AsTR0I/sipcV2/internal/transport/sip"
)

func init() {
	optionsCmd.Flags().SortFlags = false

	optionsCmd.Flags().String(
		"user-port",
		"",
		"SIP client port",
	)
	optionsCmd.Flags().String(
		"server-host",
		"",
		"SIP server host[:port]",
	)

	optionsCmd.Flags().String(
		"proxy",
		"",
		"SIP proxy",
	)

	optionsCmd.Flags().String(
		"from",
		"",
		"SIP From header",
	)

	optionsCmd.Flags().String(
		"to",
		"",
		"SIP To header",
	)

	optionsCmd.Flags().String(
		"user-agent",
		fmt.Sprintf(
			"%s(%s/%s)/v.%s",
			rootCmd.Use,
			rootCmd.Version,
			runtime.GOOS,
			runtime.GOARCH,
		),
		"SIP User-Agent header",
	)

	if err := optionsCmd.MarkFlagRequired("server-host"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(optionsCmd)
}

var optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "SIP OPTIONS",
	Long:  "Make SIP OPTIONS",

	RunE: func(cmd *cobra.Command, args []string) error {
		userPort, err := cmd.Flags().GetString("user-port")
		if err != nil {
			return err
		}

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

		userAgent, err := cmd.Flags().GetString("user-agent")
		if err != nil {
			return err
		}

		// cfg := config.SIPConfig{
		// 	UserPort:   userPort,
		// 	ServerHost: serverHost,
		// 	Proxy:      proxy,
		// 	From:       from,
		// 	To:         to,
		// 	UserAgent:  userAgent,
		// }

		req := domain.OptionsRequest{
			UserPort:   userPort,
			ServerHost: serverHost,
			Proxy:      proxy,
			From:       from,
			To:         to,
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
			req.UserAgent,
			userPort,
		)

		if err != nil {
			return err
		}

		application := app.NewApp(
			sipClient,
		)

		return application.Options(ctx, req)
	},
}
