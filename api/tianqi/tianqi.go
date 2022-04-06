package tianqi

import (
	"fmt"
	"regexp"

	"github.com/imroc/req/v3"
)

const KEY = "4r9bergjetiv1tsd"

func getTianqiTuplesByCity(city string) []string {

	res := req.MustGet(fmt.Sprintf("https://api.seniverse.com/v3/weather/now.json?key=%s&location=%s&language=zh-Hans&unit=c", KEY, city))
	content := res.String()

	r := regexp.MustCompile(`name"\s*:\s*"(\S+?)".*?text"\s*:\s*"(\S+?)".*?temperature"\s*:\s*"(\S+?)"`)
	arrs := r.FindStringSubmatch(content)
	fmt.Println(arrs)

	return arrs[1:]
}
