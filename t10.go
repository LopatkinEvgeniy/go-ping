package main

import (
	"flag"
	"fmt"
	"github.com/sparrc/go-ping"
//	"go/types"
	"net"
	"strconv"
	"time"
)

var usage = `
Usage:
    pping [-s start IP] [-e end IP]

Example:
    # pping -s 172.20.13.1 -e 172.20.13.225
`

func f1(){
	jobs := make(chan string)
	results := make(chan string)


	start := flag.String("s", "172.20.13.1", "start IP")
	end := flag.String("e", "172.20.13.224", "end IP")

	hhelp := flag.String("h", "", "")
	hhelp1 := flag.String("help", "", "")


	flag.Usage = func() {
		fmt.Printf(usage)
	}

	flag.Parse()

	if *hhelp == "h" {
		flag.Usage()
		return
	}

	if *hhelp1 == "help" {
		flag.Usage()
		return
	}

	ip := net.ParseIP(*start)
	ip2 := net.ParseIP(*end)

	ip = ip.To4()
	ip2 = ip2.To4()

	ips := int(ip2[3]-ip[3])


	if ip[0] != ip2[0]{
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	}else if ip[1] != ip2[1]{
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	}else if ip[2] != ip2[2]{
		fmt.Printf("Error! Start and end IP's have different subnets\n")
		return
	}




	for w := 0; w <= ips; w++ {
		go func(id int, jobs <-chan string, results chan<- string) {
			for j := range jobs {

				pinger, err := ping.NewPinger(j)
				if err != nil {
					fmt.Printf("ERROR: %s\n", err.Error())
					return
				}


				pinger.OnFinish = func(stats *ping.Statistics) {

					var0 := fmt.Sprintf("%v", stats.PacketLoss)
					var10, err := strconv.ParseFloat(var0, 64)
					if err == nil {
					}

					var2 := float64(0)
					var3 := float64(100)


					if var10 == var2 {
						//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST ACTIVE")
						fmt.Printf("%s %s\n", pinger.Addr(), "- HOST ACTIVE")
					}else if var10 < var3 {
						fmt.Printf("%s %s, %v%% packet loss\n", pinger.Addr(), "- HOST WARNING", stats.PacketLoss)
					}else if var10 == var3 {
						//	fmt.Printf("%s, %v%% packet loss, %s\n", pinger.Addr(), stats.PacketLoss, "HOST DISABLED")
						fmt.Printf("%s %s\n", pinger.Addr(), "- HOST DISABLED")
					}



				}


				pinger.Count = 3
				pinger.Timeout = 10000000000


				pinger.Run()

				time.Sleep(time.Second * 10)
				results <- j
			}

		}(w, jobs, results)
	}




	for i := ip[3]; i <= ip2[3]; i++ {

		test := ip[0]
		test1 := ip[1]
		test2 := ip[2]

		output := fmt.Sprintf("%d.%d.%d.%d", test, test1, test2, i)
		jobs <- output
	}


}

func main(){
//	var res
//	res <- f1()
//f1()
	f1()
	time.Sleep(time.Second * 12)
	return

}
