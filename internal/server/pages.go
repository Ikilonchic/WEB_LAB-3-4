package server

import (
	"os"
	"io/ioutil"
	"net/http"
)

func (serve *Server) getPage(page string) http.HandlerFunc {
	file, err := os.Open(serve.config.Templates + page + ".html")
	if err != nil {
		serve.logger.Error("Cannot find file ./" + serve.config.Templates + page + ".html")
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
    if err != nil {
        serve.logger.Error("Cannot read file ./" + serve.config.Templates + page + ".html")
	}

	return func(res http.ResponseWriter, req *http.Request) {
		if _, err := res.Write(data); err == nil {
			res.Header().Set("Content-Type", "text/html")
		}
	}
}