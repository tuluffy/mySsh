package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"errors"

	"github.com/gcmurphy/getpass"
)

var (
	f       string
	host    string
	port    int
	list    bool
	h       bool
	help    bool
	version bool
)

func init() {
	flag.StringVar(&f, "f", "", "specify config file")
	flag.StringVar(&host, "host", "", "specity remote host")
	flag.IntVar(&port, "port", 22, "specity remote port")
	flag.BoolVar(&list, "list", false, "list remote hosts")
	flag.BoolVar(&h, "h", false, "Usage")
	flag.BoolVar(&help, "help", false, "Usage")
	flag.BoolVar(&version, "version", false, "Print version")

}

func Args(app *App) error {
	last := flag.Args()
	if len(last) == 1 {
		//re, _ := regexp.Compile(`(?P<user>[0-9A-Za-z]+)@(?P<host>[0-9.]+)`)
		hostinfo := strings.Split(last[0], "@")

		fmt.Println(hostinfo)
		USER := hostinfo[0]
		HOST := hostinfo[1]
		password, err := getpass.GetPassWithOptions("\033[32mEnter Password: \033[0m", 0, getpass.DefaultMaxPass)
		if err != nil {
			//Printer.Errorln(err)
			//os.Exit(2)
			return err
		}
		server := Server{
			Name:     "",
			User:     USER,
			Password: password,
			Ip:       HOST,
			Port:     port,
			Method:   "password",
			Key:      "",
		}
		app.servers = append(app.servers, server)
		app.servers[0].Connection()

	} else if len(last) >= 2 {
		return errors.New("args error")
	}

	if h || help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		Printer.Infoln("version=", VERSION, "\nGithub: https://github.com/doovemax/mySsh")
		os.Exit(0)
	}
	if f != "" {
		f, _ = filepath.Abs(f)
		servers, err := ConfigPath(f)
		if err != nil {
			//Printer.Errorln(err)
			//os.Exit(2)
			return err
		}
		app.servers = servers
		return nil

	}
	//用户指定IP和port登录
	if host != "" {
		var user string
		var password string

		Printer.Info("Enter user: ")
		_, err := fmt.Scanln(&user)
		//Printer.Info("Enter ")
		password, err = getpass.GetPassWithOptions("\033[32mEnter Password: \033[0m", 1, getpass.DefaultMaxPass)
		if err != nil {
			//Printer.Errorln(err)
			//os.Exit(2)
			return err
		}
		server := Server{
			Name:     "",
			User:     user,
			Password: password,
			Ip:       host,
			Port:     port,
			Method:   "password",
			Key:      "",
		}
		app.servers = append(app.servers, server)

	}
	return nil
}
