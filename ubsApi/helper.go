package ubsApi

import (
	"log"
	"net/http"
	"time"
)

func NowAsUnixMilli() int64 {
	return time.Now().UTC().UnixNano() / 1e6
}

func DbgPrintCurCookies() {
	var cookieNum int = len(UBSCookies)
	log.Printf("cookieNum=%d", cookieNum)
	for i := 0; i < cookieNum; i++ {
		var curCk *http.Cookie = UBSCookies[i]
		//log.Printf("curCk.Raw=%s", curCk.Raw)
		log.Printf("Cookie [%d]", i)
		log.Printf("Name\t=%s", curCk.Name)
		log.Printf("Value\t=%s", curCk.Value)
		log.Printf("Path\t=%s", curCk.Path)
		log.Printf("Domain\t=%s", curCk.Domain)
		log.Printf("Expires\t=%s", curCk.Expires)
		log.Printf("RawExpires=%s", curCk.RawExpires)
		log.Printf("MaxAge\t=%d", curCk.MaxAge)
		log.Printf("Secure\t=%t", curCk.Secure)
		log.Printf("HttpOnly=%t", curCk.HttpOnly)
		log.Printf("Raw\t=%s", curCk.Raw)
		log.Printf("Unparsed=%s", curCk.Unparsed)
	}
}
