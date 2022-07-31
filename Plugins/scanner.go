package Plugins

import (
	"errors"
	"fmt"
	"github.com/shadow1ng/fscan/WebScan/lib"
	"github.com/shadow1ng/fscan/common"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

func Scan(info common.HostInfo) {
	fmt.Println("start infoscan")
	Hosts, err := common.ParseIP(info.Host, common.HostFile, common.NoHosts)
	if err != nil {
		fmt.Println("len(hosts)==0", err)
		return
	}
	lib.Inithttp(common.Pocinfo)
	var ch = make(chan struct{}, common.Threads)
	var wg = sync.WaitGroup{}
	if len(Hosts) > 0 || len(common.HostPort) > 0 {
		if common.IsPing == false && len(Hosts) > 0 {
			Hosts = CheckLive(Hosts, common.Ping)
			fmt.Println("[*] Icmp alive hosts len is:", len(Hosts))
		}
		if common.Scantype == "icmp" {
			common.LogWG.Wait()
			return
		}
		var AlivePorts []string
		if common.Scantype == "webonly" {
			AlivePorts = NoPortScan(Hosts, info.Ports)
		} else if len(Hosts) > 0 {
			AlivePorts = PortScan(Hosts, info.Ports, common.Timeout)
			fmt.Println("[*] alive ports len is:", len(AlivePorts))
			if common.Scantype == "portscan" {
				common.LogWG.Wait()
				return
			}
		}
		if len(common.HostPort) > 0 {
			AlivePorts = append(AlivePorts, common.HostPort...)
			AlivePorts = common.RemoveDuplicate(AlivePorts)
			common.HostPort = nil
			fmt.Println("[*] AlivePorts len is:", len(AlivePorts))
		}
		var severports []string //severports := []string{"21","22","135"."445","1433","3306","5432","6379","9200","11211","27017"...}
		for _, port := range common.PORTList {
			severports = append(severports, strconv.Itoa(port))
		}
		fmt.Println("start vulscan")
		for _, targetIP := range AlivePorts {
			info.Host, info.Ports = strings.Split(targetIP, ":")[0], strings.Split(targetIP, ":")[1]
			if common.Scantype == "all" || common.Scantype == "main" {
				switch {
				case info.Ports == "135":
					AddScan(info.Ports, info, ch, &wg) //findnet
				case info.Ports == "445":
					//AddScan(info.Ports, info, ch, &wg)  //smb
					AddScan("1000001", info, ch, &wg) //ms17010
					//AddScan("1000002", info, ch, &wg) //smbghost
				case info.Ports == "9000":
					AddScan(info.Ports, info, ch, &wg) //fcgiscan
					AddScan("1000003", info, ch, &wg)  //http
				case IsContain(severports, info.Ports):
					AddScan(info.Ports, info, ch, &wg) //plugins scan
				default:
					AddScan("1000003", info, ch, &wg) //webtitle
				}
			} else {
				port, _ := common.PORTList[common.Scantype]
				scantype := strconv.Itoa(port)
				AddScan(scantype, info, ch, &wg)
			}
		}
	}
	for _, url := range common.Urls {
		info.Url = url
		AddScan("1000003", info, ch, &wg)
	}
	wg.Wait()
	common.LogWG.Wait()
	close(common.Results)
	fmt.Println(fmt.Sprintf("已完成 %v/%v", common.End, common.Num))
}

var Mutex = &sync.Mutex{}

func AddScan(scantype string, info common.HostInfo, ch chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		Mutex.Lock()
		common.Num += 1
		Mutex.Unlock()
		ScanFunc(PluginList, scantype, &info)
		Mutex.Lock()
		common.End += 1
		Mutex.Unlock()
		<-ch
		wg.Done()
	}()
	ch <- struct{}{}
}

func ScanFunc(m map[string]interface{}, name string, infos ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(infos) != f.Type().NumIn() {
		err = errors.New("The number of infos is not adapted ")
		fmt.Println(err.Error())
		return result, nil
	}
	in := make([]reflect.Value, len(infos))
	for k, info := range infos {
		in[k] = reflect.ValueOf(info)
	}
	result = f.Call(in)
	return result, nil
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
