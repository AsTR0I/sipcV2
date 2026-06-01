package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/AsTR0I/sipcV2/internal/app"
	"github.com/AsTR0I/sipcV2/internal/domain"
	"github.com/AsTR0I/sipcV2/internal/logger"
	"github.com/AsTR0I/sipcV2/internal/transport/sip"
	"github.com/spf13/cobra"
)

func init() {
	inviteCmd.Flags().SortFlags = false

	inviteCmd.Flags().String(
		"user-port",
		"",
		"SIP client port",
	)

	inviteCmd.Flags().String(
		"server-host",
		"",
		"SIP server host[:port]",
	)

	inviteCmd.Flags().String(
		"proxy",
		"",
		"SIP proxy",
	)

	inviteCmd.Flags().String(
		"from",
		"",
		"SIP From header",
	)

	inviteCmd.Flags().String(
		"to",
		"",
		"SIP To header",
	)

	inviteCmd.Flags().String(
		"realm",
		"",
		"Auth realm",
	)

	inviteCmd.Flags().String(
		"username",
		"",
		"Auth username",
	)

	inviteCmd.Flags().String(
		"password",
		"",
		"Auth password",
	)

	inviteCmd.Flags().Int(
		"expires",
		600,
		"SIP Expires header",
	)

	inviteCmd.Flags().String(
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

	if err := inviteCmd.MarkFlagRequired("server-host"); err != nil {
		panic(err)
	}

	if err := inviteCmd.MarkFlagRequired("username"); err != nil {
		panic(err)
	}
	if err := inviteCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(inviteCmd)
}

var inviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "SIP INVITE",
	Long:  "Make SIP INVITE",

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

		req := domain.InviteRequest{
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
			userPort,
		)
		if err != nil {
			return err
		}

		application := app.NewApp(
			sipClient,
		)

		err = application.Invite(ctx, req)

		fmt.Println("invite returned", err)

		return err
	},
}
