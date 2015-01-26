package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var file = filepath.Join(os.TempDir(), "./gvpf.html")

func main() {
	if err := search(); err != nil {
		panic(err)
	}

	cmd := exec.Command("open", file)

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func search() error {
	doc, err := goquery.NewDocument("https://www.gv.com/portfolio/")
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}

	defer f.Close()

	f.WriteString("<html><body><table border='1'>")

	doc.Find("div.portfolio-all div.investments ul.more-investments li").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(strings.Split(s.Text(), "\n")[2])
		desc := strings.TrimSpace(strings.Split(s.Text(), "\n")[4])
		url, _ := s.Find("a").Attr("href")

		f.WriteString("<tr><td>" + strconv.Itoa(i+1) + "</td><td><a href='" + url + "' target='_blank'>" + name + "</a></td><td>" + desc + "</td></tr>")
	})

	f.WriteString("</table></body></html>")

	return nil
}
