package main

import (
	"log"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/server"
	servercmd "github.com/cosmos/cosmos-sdk/server/cmd"

	hubtypes "github.com/sentinel-official/hub/types"
	_ "github.com/sentinel-official/hub/x/node/legacy/v1/types"
	_ "github.com/sentinel-official/hub/x/plan/legacy/v1/types"
	_ "github.com/sentinel-official/hub/x/provider/legacy/v1/types"
	_ "github.com/sentinel-official/hub/x/session/legacy/v1/types"
	_ "github.com/sentinel-official/hub/x/subscription/legacy/v1/types"
)

func HomeDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	dir = path.Join(dir, ".sentinelhub")
	return dir, nil
}

func main() {
	cfg := hubtypes.GetConfig()
	cfg.Seal()

	homeDir, err := HomeDir()
	if err != nil {
		log.Fatal(err)
	}

	cmd := NewRootCmd(homeDir)
	if err = servercmd.Execute(cmd, homeDir); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
