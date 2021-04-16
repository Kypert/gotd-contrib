package redis_test

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	redisclient "github.com/go-redis/redis/v8"
	"golang.org/x/xerrors"

	"github.com/gotd/contrib/auth"
	"github.com/gotd/contrib/auth/redis"
	"github.com/gotd/contrib/auth/terminal"
	"github.com/gotd/td/telegram"
)

func redisAuth(ctx context.Context) error {
	redisClient := redisclient.NewClient(&redisclient.Options{})
	cred := redis.NewCredentials(redisClient).
		WithPhoneKey("phone").
		WithPasswordKey("password")

	client, err := telegram.ClientFromEnvironment(telegram.Options{})
	if err != nil {
		return xerrors.Errorf("create client: %w", err)
	}

	return client.Run(ctx, func(ctx context.Context) error {
		return client.AuthIfNecessary(
			ctx,
			telegram.NewAuth(auth.Build(cred, terminal.OS()), telegram.SendCodeOptions{}),
		)
	})
}

func ExampleCredentials() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := redisAuth(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}