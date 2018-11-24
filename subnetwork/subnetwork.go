package subnetwork

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"os/user"
	"path"
	"time"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
)

// InitCustomGenesis Initialises a new genesis block with data directory for a new network
func InitCustomGenesis() {
	// get os home directory location
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting os user occurred: ", err.Error())
		return
	}

	// create path of directory to init node data within
	dirName := fmt.Sprint("geth-data-", time.Now().Unix())
	dirPath := path.Join(usr.HomeDir, dirName)

	// create node directory
	err = os.Mkdir(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
		return
	}

	// add new genesis file to newly created directory
	newGenesisPath := path.Join(dirPath, "genesis.json")

	genesis := core.Genesis{
		Config: &params.ChainConfig{
			ChainID:        big.NewInt(15),
			HomesteadBlock: big.NewInt(0),
			EIP155Block:    big.NewInt(0),
			EIP158Block:    big.NewInt(0),
		},
		Difficulty: big.NewInt(1),
		GasLimit:   uint64(2100000),
		Alloc:      core.GenesisAlloc{},
	}

	genesisJSON, _ := json.Marshal(genesis)
	err = ioutil.WriteFile(newGenesisPath, genesisJSON, 0644)

	// run init genesis block command
	initGenesisCmd := exec.Command("/home/mr/go/src/github.com/ethereum/go-ethereum/build/bin/geth",
		"init", newGenesisPath, "--datadir", dirPath)

	output, err := initGenesisCmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println("New network data is added at path: '", dirPath, "'")
}
