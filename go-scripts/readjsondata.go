package go_scripts

import (
	"encoding/json"
	"fmt"
)

type Ccids []string

func ReadJsonArrayAndPrintCommand() {

	jsonStr := `["5833","5834","5835","5836","5837","5838","5839","5840","5841","5842","5843","1679","5845","5846","1683","5847","5848","2283","2285","5849","2282","5850","5851","2284","5852","5853","5855","5856","5857","5858","5859","5860","5861","5862","5863","5864","5865","1682","2277","2279","5866","2238","2278","5867","5868","2280","2290","2287","5869","2288","2289","5870","5871","5872","2291","5873","2293","337","2297","2296","5874","2294","2298","2295","2299","5875"]`
	var ccids Ccids
	json.Unmarshal([]byte(jsonStr), &ccids)
	// log.Info(context.Background(), "ccids", ccids)
	cmdTemplate := "./sfclient jc update enable-bp-ldp %v --value=true --journey=details --hostport 127.0.0.1:9281 --yes-do-it"
	for _, ccid := range ccids {
		fmt.Printf(cmdTemplate, ccid)
		fmt.Println()
	}
}
