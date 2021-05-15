package main

import (
	"blog/internal/user/v1/conf"
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"gopkg.in/yaml.v2"
	"os"
)

var (
	Name     string
	Version  string
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path,eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(hs, gs),
	)
}

func main() {
	flag.Parse()
	logger := log.NewStdLogger(os.Stdout)
	config := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := config.Load(); err != nil {
		panic(err)
	}
	var bc conf.Bootstrap
	if err := config.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanUp, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanUp()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
