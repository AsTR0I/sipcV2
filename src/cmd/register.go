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
	registerCmd.Flags().SortFlags = false

	registerCmd.Flags().String(
		"user-port",
		"",
		"SIP client port",
	)

	registerCmd.Flags().String(
		"server-host",
		"",
		"SIP server host:port",
	)

	registerCmd.Flags().String(
		"proxy",
		"",
		"SIP proxy",
	)

	registerCmd.Flags().String(
		"from",
		"",
		"SIP From",
	)

	registerCmd.Flags().String(
		"to",
		"",
		"SIP To",
	)

	registerCmd.Flags().String(
		"realm",
		"",
		"realm",
	)

	registerCmd.Flags().String(
		"username",
		"",
		"username",
	)

	registerCmd.Flags().String(
		"password",
		"",
		"password",
	)

	registerCmd.Flags().Int(
		"expires",
		600,
		"expires",
	)

	registerCmd.Flags().String(
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

	if err := registerCmd.MarkFlagRequired("server-host"); err != nil {
		panic(err)
	}

	if err := registerCmd.MarkFlagRequired("username"); err != nil {
		panic(err)
	}
	if err := registerCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "SIP REGISTER",
	Long:  "Make SIP REGISTER",

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

		realm, err := cmd.Flags().GetString("realm")
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

		expires, err := cmd.Flags().GetInt("expires")
		if err != nil {
			return err
		}

		userAgent, err := cmd.Flags().GetString("user-agent")
		if err != nil {
			return err
		}

		req := domain.RegisterRequest{
			UserPort:   userPort,
			ServerHost: serverHost,
			Proxy:      proxy,
			From:       from,
			To:         to,
			Realm:      realm,
			Username:   username,
			Password:   password,
			Expires:    expires,
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
			userAgent,
		)
		if err != nil {
			return err
		}

		application := app.NewApp(
			sipClient,
		)

		return application.Register(ctx, req)
	},
}
