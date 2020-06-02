package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kckecheng/grafana_annotate/gapi"
	"github.com/spf13/cobra"
)

var errExit int = 1

func validateGlobalOpts(cmd *cobra.Command) (string, int64, string, string, error) {
	var (
		address  string
		port     int64
		user     string
		password string
		err      error
	)

	flags := cmd.Flags()

	address, err = flags.GetString("address")
	if err != nil {
		return address, port, user, password, err
	}
	address = strings.Trim(address, " ")
	if len(address) == 0 {
		return address, port, user, password, errors.New("Grafana server address is empty")
	}

	port, err = flags.GetInt64("port")
	if err != nil {
		return address, port, user, password, err
	}

	user, err = flags.GetString("user")
	if err != nil {
		return address, port, user, password, err
	}
	if len(user) == 0 {
		return address, port, user, password, errors.New("Grafana server user name is empty")
	}

	password, err = flags.GetString("password")
	if err != nil {
		return address, port, user, password, err
	}
	if len(password) == 0 {
		return address, port, user, password, errors.New("Grafana server password is empty")
	}

	return address, port, user, password, nil
}

func bailout(err error, msg string, code int) {
	if err != nil {
		emsg := err.Error()
		if emsg != "" {
			emsg = ": " + emsg
		}
		fmt.Println(msg + emsg)
		os.Exit(code)
	}
}

func login(cmd *cobra.Command) *gapi.Grafana {
	address, port, user, password, err := validateGlobalOpts(cmd)
	bailout(err, "Fail to parse command options", errExit)

	g, err := gapi.NewGrafana(address, port, user, password)
	if err != nil {
		bailout(errors.New(""), fmt.Sprintf("Fail to connect to Grafana server http://%s:%d with %s:%s", address, port, user, password), errExit)
	}
	return g
}
