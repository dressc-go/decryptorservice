package server

import (
	"decryptorservice/pkg/config"
	"decryptorservice/pkg/cryptkey"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dressc-go/decryptors/base64OeapSha1"
	"github.com/dressc-go/zlogger"
)

type JsonRequestData struct {
	C    string
	Ctxt string
}

type JsonResponse struct {
	Txt string
}

func decryptDataHandler(privateKey *cryptkey.CryptKey) func(http.ResponseWriter, *http.Request) {
	realFun := func(w http.ResponseWriter, req *http.Request) {
		logger := zlogger.GetLogger("server.decryptBpk")
		logger.Info().
			Str("RemoteAddr", req.RemoteAddr).
			Str("User-Agent", req.Header.Get("User-Agent")).
			Msg("received_request")
		reqBody, _ := ioutil.ReadAll(req.Body)
		logger.Debug().Str("RemoteAddr", req.RemoteAddr).
			Str("data", string(reqBody)).
			Msg("received_data")

		var jsonRequestData JsonRequestData
		errJson := json.Unmarshal(reqBody, &jsonRequestData)
		if errJson != nil {
			logger.Error().
				Err(errJson).
				Msg("Error while parsing the payload")
			return
		}
		resp := JsonResponse{Txt: ""}
		if jsonRequestData.C == "base64OeapSha1" {
			privKey, _ := privateKey.GetPrivateKey()
			clearText, _ := base64OeapSha1.Decrypt(jsonRequestData.Ctxt, privKey)
			resp = JsonResponse{Txt: clearText}
		}
		str_resp, _ := json.Marshal(resp)
		if len(resp.Txt) > 0 {
			_, err := fmt.Fprintf(w, string(str_resp))
			if err != nil {
				logger.Error().
					Err(err).
					Msg("Error while handling the request")
				return
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	return realFun
}

func Start(cnf *config.Config) {
	listenOn := cnf.IpAddress + ":" + strconv.Itoa(int(cnf.Port))
	logger := zlogger.GetLogger("server")
	logger.Info().Msg("Listening on " + listenOn)

	decryptionKey := new(cryptkey.CryptKey)
	err := decryptionKey.New(cnf.DecPrivateKeyFile)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}

	http.HandleFunc("/decrypt", decryptDataHandler(decryptionKey))

	if cnf.TLSKeyFile != "" && cnf.TLSCertFile != "" {
		err = http.ListenAndServeTLS(listenOn, cnf.TLSCertFile, cnf.TLSKeyFile, nil)
	} else {
		err = http.ListenAndServe(listenOn, nil)
	}
	if err != nil {
		logger.Fatal().Err(err).Msg("Can not start listener")
	}
}
