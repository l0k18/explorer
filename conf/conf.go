package conf

import (
	"git.parallelcoin.io/marcetin/explorer/jdb"
)

func Conf() (c *mod.Config) {
	jdb.JDB.Read("config", "config", &c)
	return
}
