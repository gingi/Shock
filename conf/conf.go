package conf

import (
	"flag"
	"github.com/kless/goconfig/config"
	"strings"
)

type idxOpts struct {
	unique   bool
	dropDups bool
	sparse   bool
}

// Setup conf variables
var (
	// Config File
	CONFIGFILE = ""

	// Shock 
	SITEPORT = 0
	APIPORT  = 0

	// Admin
	ADMINEMAIL = ""
	SECRETKEY  = ""

	// Directories
	DATAPATH = ""
	SITEPATH = ""
	LOGSPATH = ""

	// Mongodb 
	MONGODB = ""

	// Node Indices
	NODEIDXS map[string]idxOpts = nil
)

func init() {
	flag.StringVar(&CONFIGFILE, "conf", "/usr/local/shock/conf/shock.cfg", "path to config file")
	flag.Parse()
	c, _ := config.ReadDefault(CONFIGFILE)
	// Shock
	SITEPORT, _ = c.Int("Shock", "site-port")
	APIPORT, _ = c.Int("Shock", "api-port")

	// Admin
	ADMINEMAIL, _ = c.String("Admin", "email")
	SECRETKEY, _ = c.String("Admin", "secretkey")

	// Directories
	SITEPATH, _ = c.String("Directories", "site")
	DATAPATH, _ = c.String("Directories", "data")
	LOGSPATH, _ = c.String("Directories", "logs")

	// Mongodb
	MONGODB, _ = c.String("Mongodb", "hosts")

	// parse Node-Indices
	NODEIDXS = map[string]idxOpts{}
	nodeIdx, _ := c.Options("Node-Indices")
	for _, opt := range nodeIdx {
		val, _ := c.String("Node-Indices", opt)
		opts := idxOpts{}
		for _, parts := range strings.Split(val, ",") {
			p := strings.Split(parts, ":")
			if p[0] == "unique" {
				if p[1] == "true" {
					opts.unique = true
				} else {
					opts.unique = false
				}
			} else if p[0] == "dropDups" {
				if p[1] == "true" {
					opts.dropDups = true
				} else {
					opts.dropDups = false
				}
			} else if p[0] == "sparse" {
				if p[1] == "true" {
					opts.sparse = true
				} else {
					opts.sparse = false
				}
			}
		}
		NODEIDXS[opt] = opts
	}
}
