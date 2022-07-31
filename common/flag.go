package common

import (
	"flag"
	"runtime"
	"runtime/debug"
	"time"
)

func init() {
	go func() {
		for {
			runtime.GC()
			debug.FreeOSMemory()
			time.Sleep(10 * time.Second)
		}
	}()
}

func Banner() {
	banner := `
   ___                              _    
  / _ \     ___  ___ _ __ __ _  ___| | __ 
 / /_\/____/ __|/ __| '__/ _` + "`" + ` |/ __| |/ /
/ /_\\_____\__ \ (__| | | (_| | (__|   <    
\____/     |___/\___|_|  \__,_|\___|_|\_\   
                     fscan version: ` + version + `
`
	print(banner)
	print(`
<-------------- 加入了很多地方的POC,目前POC数436 -------------->
`)
}

func Flag(Info *HostInfo) {
	Banner()
	flag.StringVar(&Info.Host, "h", "", "IP address of the host you want to scan,for example: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12")
	flag.StringVar(&NoHosts, "hn", "", "the hosts no scan,as: -hn 192.168.1.1/24")
	flag.StringVar(&Info.Ports, "p", DefaultPorts, "Select a port,for example: 22 | 1-65535 | 22,80,3306")
	flag.StringVar(&PortAdd, "pa", "", "add port base DefaultPorts,-pa 3389")
	flag.StringVar(&UserAdd, "usera", "", "add a user base DefaultUsers,-usera user")
	flag.StringVar(&PassAdd, "pwda", "", "add a password base DefaultPasses,-pwda password")
	flag.StringVar(&NoPorts, "pn", "", "the ports no scan,as: -pn 445")
	flag.StringVar(&Command, "c", "", "exec command (ssh)")
	flag.StringVar(&SshKey, "sshkey", "", "sshkey file (id_rsa)")
	flag.StringVar(&Domain, "domain", "", "smb domain")
	flag.StringVar(&Username, "user", "", "username")
	flag.StringVar(&Password, "pwd", "", "password")
	flag.Int64Var(&Timeout, "time", 3, "Set timeout")
	flag.StringVar(&Scantype, "m", "all", "Select scan type ,as: -m ssh")
	flag.StringVar(&Path, "path", "", "fcgi、smb romote file path")
	flag.IntVar(&Threads, "t", 600, "Thread nums")
	flag.IntVar(&LiveTop, "top", 10, "show live len top")
	flag.StringVar(&HostFile, "hf", "", "host file, -hf ip.txt")
	flag.StringVar(&Userfile, "userf", "", "username file")
	flag.StringVar(&Passfile, "pwdf", "", "password file")
	flag.StringVar(&PortFile, "portf", "", "Port File")
	flag.StringVar(&PocPath, "pocpath", "", "poc file path")
	flag.StringVar(&RedisFile, "rf", "", "redis file to write sshkey file (as: -rf id_rsa.pub) ")
	flag.StringVar(&RedisShell, "rs", "", "redis shell to write cron file (as: -rs 192.168.1.1:6666) ")
	flag.BoolVar(&IsWebCan, "nopoc", false, "not to scan web vul")
	flag.BoolVar(&IsBrute, "nobr", false, "not to Brute password")
	flag.IntVar(&BruteThread, "br", 1, "Brute threads")
	flag.BoolVar(&IsPing, "np", false, "not to ping")
	flag.BoolVar(&Ping, "ping", false, "using ping replace icmp")
	flag.StringVar(&TmpOutputfile, "o", "result.txt", "Outputfile")
	flag.BoolVar(&TmpSave, "no", false, "not to save output log")
	flag.Int64Var(&WaitTime, "debug", 60, "every time to LogErr")
	flag.BoolVar(&Silent, "silent", false, "silent scan")
	flag.BoolVar(&PocFull, "full", false, "poc full scan,as: shiro 100 key")
	flag.StringVar(&URL, "u", "", "url")
	flag.StringVar(&UrlFile, "uf", "", "urlfile")
	flag.StringVar(&Pocinfo.PocName, "pocname", "", "use the pocs these contain pocname, -pocname weblogic")
	flag.StringVar(&Proxy, "proxy", "", "set poc proxy, -proxy http://127.0.0.1:8080")
	flag.StringVar(&Socks5Proxy, "socks5", "", "set socks5 proxy, will be used in tcp connection, timeout setting will not work")
	flag.StringVar(&Cookie, "cookie", "", "set poc cookie,-cookie rememberMe=login")
	flag.Int64Var(&WebTimeout, "wt", 5, "Set web timeout")
	flag.IntVar(&PocNum, "num", 20, "poc rate")
	flag.StringVar(&SC, "sc", "", "ms17 shellcode,as -sc add")
	flag.Parse()
}
