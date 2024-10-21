package main

import (
	"flag"
)

var (
	flagAgentBinaryPath   string
	flagServerBinaryPath  string
	flagTargetSourcePath  string
	flagServerHost        string
	flagServerPort        string
	flagServerBaseURL     string
	flagFileStoragePath   string
	flagDatabaseDSN       string
	flagSHA256Key         string
	flagBaseProfilePath   string
	flagResultProfilePath string
	flagPackageName       string
)

func init() {
	flag.StringVar(&flagAgentBinaryPath, "agent-binary-path", "", "path to target agent binary")
	flag.StringVar(&flagServerBinaryPath, "binary-path", "", "path to target server binary")
	flag.StringVar(&flagTargetSourcePath, "source-path", "", "path to target server source")
	flag.StringVar(&flagServerHost, "server-host", "", "host of target address")
	flag.StringVar(&flagServerPort, "server-port", "", "port of target address")
	flag.StringVar(&flagServerBaseURL, "server-base-url", "", "base URL of target address")
	flag.StringVar(&flagFileStoragePath, "file-storage-path", "", "path to persistent file storage")
	flag.StringVar(&flagDatabaseDSN, "database-dsn", "", "connection string to database")
	flag.StringVar(&flagSHA256Key, "key", "", "sha256 key for hashing")
	flag.StringVar(&flagBaseProfilePath, "base-profile-path", "", "path to base pprof profile")
	flag.StringVar(&flagResultProfilePath, "result-profile-path", "", "path to result pprof profile")
	flag.StringVar(&flagPackageName, "package-name", "", "name of package to be tested")
}
