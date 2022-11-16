package blocks

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *Task
var contractAddr string = "0x4780Bd8937247bd72FfE5aa81Ba5aE97c82C411B"
var keystr = `{"address":"06830605aab8b1e4c3e2ad6ce7761b52b7d7086e","crypto":{"cipher":"aes-128-ctr","ciphertext":"0cab47167fbf2d27f2a15c2ff8385cd5d04e8b1353f9d0352faad88de220bb57","cipherparams":{"iv":"518ac16b5b64557fcfa1e2dee8c60c6b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"a0197aa6394fd3e770101e60f5b605d83d7e043fa2a1cf5ff84d8a247a84b2b2"},"mac":"64db80db40e22ba9d13e6b16663edeffe9b4dfb5ee1ed9b27c6eb96f7c00dbf5"},"id":"b27b4f02-2cb8-48d4-a6dd-1d133906874e","version":3}`

type TaskInfox struct {
	TaskID  string `json:"task_id"`
	Issuer  string `json:"issuer"`
	Worker  string `json:"task_user"`
	Desc    string `json:"task_name"`
	Bonus   int64  `json:"bonus"`
	Comment string `json:"comment"`
	Status  uint8  `json:"task_status"`
}

//init会自动触发运行
func init() {
	cli, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Panic("failed to Dial", err)
	}

	task, err := NewTask(common.HexToAddress(contractAddr), cli)
	if err != nil {
		log.Panic("failed to NewTask", err)
	}
	instance = task
}

func Register(username, passwd string) error {
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("failed to NewTransactor", err)
		return err
	}
	_, err = instance.Register(auth, username, passwd)
	if err != nil {
		fmt.Println("failed to Register", err)
		return err
	}
	return nil
}

func Login(username, passwd string) (bool, error) {

	return instance.Login(nil, username, passwd)
}

func Issue(username, passwd, desc string, bonus int64) error {
	//创建用户身份，需要私钥文件生成
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("failed to NewTransactor", err)
		return err
	}
	//合约调用
	_, err = instance.Issue(auth, username, passwd, desc, big.NewInt(bonus))
	if err != nil {
		fmt.Println("failed to Issue", err)
		return err
	}
	return nil
}

//status : 1- 接受 2- 提交 3- 确认 4- 退回
func Update(username, passwd, comment string, status uint8, taskID int64) error {
	//创建用户身份，需要私钥文件生成
	reader := strings.NewReader(keystr)
	auth, err := bind.NewTransactor(reader, "123")
	if err != nil {
		fmt.Println("failed to NewTransactor", err)
		return err
	}
	//合约调用
	if status == 1 {
		_, err = instance.Take(auth, username, passwd, big.NewInt(taskID))
		if err != nil {
			fmt.Println("failed to Take task", err)
			return err
		}
	} else if status == 2 {
		_, err = instance.Commit(auth, username, passwd, big.NewInt(taskID))
		if err != nil {
			fmt.Println("failed to Take Commit", err)
			return err
		}
	} else if status == 3 || status == 4 {
		if status == 4 {
			status = 1
		}
		_, err = instance.Confirm(auth, username, passwd, big.NewInt(taskID), comment, status)
		if err != nil {
			fmt.Println("failed to Take Confirm", err)
			return err
		}
	}

	return nil
}

func Tasklist() ([]TaskInfox, error) {
	var tasks []TaskInfox
	taskbs, err := instance.QryAllTasks(nil)
	if err != nil {
		fmt.Println("failed to QryAllTasks", err)
		return nil, err
	}
	var task TaskInfox

	for k, v := range taskbs {
		task.TaskID = fmt.Sprintf("%d", k)
		task.Bonus = v.Bonus.Int64()
		task.Comment = v.Comment
		task.Desc = v.Desc
		task.Issuer = v.Issuer
		task.Status = v.Status
		task.Worker = v.Worker
		tasks = append(tasks, task)
	}

	return tasks, nil
}
