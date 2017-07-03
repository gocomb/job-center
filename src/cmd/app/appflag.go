package app

import (
	"flag"
	"os"
	"common"
)
type ServerRunOptions struct {
	ServerDbType     string
	ServerDbIp       string
	ServerDbPort     string
	ServerDbPassword string
	ServerKubeIp     string
	ServerKubePort   string
}
type env struct {
	envDbType     string
	envDbIp       string
	envDbPort     string
	envDbPassword string
	envKubeIp     string
	envKubePort   string
}

func getOsEnv() (NewEnv env) {
	NewEnv = env{
		envDbType:     os.Getenv("DBTYPE"),
		envDbIp:       os.Getenv("DBTYPE"),
		envDbPort:     os.Getenv("DBPORT"),
		envDbPassword: os.Getenv("DBPASSWORD"),
		envKubeIp:     os.Getenv("KUBEIP"),
		envKubePort:   os.Getenv("KUBEPORT"),
	}
	return
}

func preCmdFlag(vName string, vValue string, usage string) *string {
	s := flag.String(vName, vValue, usage)
	flag.Parse()
	return s
}

func getRunFlag() map[string]*string {
	var runFlag map[string]*string
	runFlag=make(map[string]*string)
	runFlag["DbType"] = preCmdFlag("dbtype", "non", "input the Database type name")
	runFlag["DbIp"] = preCmdFlag("dbip", "non", "input the Database ip address")
	runFlag["DbPort"] = preCmdFlag("dbport", "non", "input the database port")
	runFlag["KubeIp"] = preCmdFlag("kubeip", "non", "input the KubeAPIserver ip address")
	runFlag["KubePort"] = preCmdFlag("kubeport", "non", "input the kubeAPIserver port")
	return runFlag
}

var RunFlag ServerRunOptions

func preFlag() {
	runFlag := getRunFlag()
	osEnv := getOsEnv()
	for k, v := range runFlag {
		flag.Parse()
		switch k {
		case "DbType":
			common.DebugPrint(k, *v)
			if *v == "non" {
				RunFlag.ServerDbType = osEnv.envDbType
			}
		case "DbIp":
			common.DebugPrint(k, *v)
			if *v == "non" {
				RunFlag.ServerDbIp = osEnv.envDbIp
			}
		case "DbPort":
			common.DebugPrint(k, *v)
			if *v == "non" {
				RunFlag.ServerDbPort = osEnv.envDbPort
			}
		case "KubeIp":
			common.DebugPrint(k, *v)
			if *v == "non" {
				RunFlag.ServerKubeIp = osEnv.envKubeIp
			}
		case "KubePort":
			common.DebugPrint(k, *v)
			if *v == "non" {
				RunFlag.ServerKubePort = osEnv.envKubePort
			}
		}
	}
}
func init() {
	preFlag()
}
