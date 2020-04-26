package app

import (
	"fmt"
	"regexp"
)

const (
	CLASS = "'>(.*?)<\\/td>\\s*<td rowspan='1'>(.*?)-(.*?)<\\/td>\\s*<td rowspan='1'>(.*?)<\\/td>\\s*<td rowspan='1'>(.*?)<\\/td><td rowspan='1' align='center'>(.*?)<br> <\\/td><td>(.*?)<\\/td>\\s*<td>(星期\\d)(第.*?节) (.*?周)<\\/td><td>(.*?)<\\/td>\\s*<td rowspan='1' align='center'><a\\s*href='kb_stuList.php\\?jxb=(.*?)' target=_blank>名单<\\/a>"
)

func Start() {
	url := "http://jwc.cqupt.edu.cn/kebiao/kb_rw.php?NhGeWgGN=4_Qmrt.KklV.ZDgsbYCwypPJ0cVF3Cf0RgbAnLS19bY0reIJdbwB3xNqq3iE8oH0qGhTiy5Ilm1J45GIj_TV1DwJByQA8hk9YIsva6aGH2SbpnRPSWAE4bq4vPVQ6q00G8JRfJxXaNLJbLNW2L2GQyOez1Td7Z5nKioZAEtbr7PqvYX0HFCD04oUmEYRmq.oVsqRbo_B0YUkefnuu5TY3KdQ8l0qG64JwlxJZUVpArgQpdAKG_bTyyo6Ce6C4pxhmjYJVL048cnhW3PWYMUT3oNgaOgiosirXO9kLBFO.vZlKHbYdtKbU7vRha3bIK2BI7m.ZvkJxn6mBvSJrRhlZQ5Y4ODu2MYRFTdEWUYKGevsZcuo2ZT0yWuXbboRO_DuRTbWOVdfb3aW1MBXpKBwtIqt2nw4AU1Vmyt96OXtzMRyb1m4YpxWvM.FSddZvdrAXMqe"
	body := GetBody(url)
	re := regexp.MustCompile(CLASS)
	classs := re.FindAllStringSubmatch(body, -1)

	fmt.Println(len(classs))
}
