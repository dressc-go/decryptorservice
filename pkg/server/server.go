package server

import (
	"crypto/rsa"
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
	Txt    string
	Error  int16
	Errmsg string
}

func parsePayload(reqBody []byte) (JsonRequestData, error) {
	var jsonRequestData JsonRequestData

	logger := zlogger.GetLogger("server.parsePayload")
	errJson := json.Unmarshal(reqBody, &jsonRequestData)
	if errJson != nil {
		logger.Error().
			Err(errJson).
			Msg("Error while parsing the payload")
		return jsonRequestData, errJson
	}
	return jsonRequestData, nil
}

func handleRequest(reqBody []byte, privateKey *cryptkey.CryptKey) JsonResponse {
	var jsonRequestData JsonRequestData
	var err error

	logger := zlogger.GetLogger("server.handleRequest")
	jsonRequestData, err = parsePayload(reqBody)
	if err != nil {
		return JsonResponse{Error: 1, Errmsg: err.Error(), Txt: ""}
	}
	if jsonRequestData.C == "base64OeapSha1" {
		var privKey *rsa.PrivateKey
		privKey, err = privateKey.GetPrivateKey()
		if err != nil {
			logger.Error().Err(err).Msg("Private Key Error")
			return JsonResponse{Error: 1, Errmsg: err.Error(), Txt: ""}
		}
		clearText, err := base64OeapSha1.Decrypt(jsonRequestData.Ctxt, privKey)
		if err != nil {
			logger.Error().Err(err).Msg("Decryption Error")
			return JsonResponse{Error: 1, Errmsg: err.Error(), Txt: ""}
		}
		return JsonResponse{Error: 0, Txt: clearText}
	}
	logger.Error().Msg("Unknown crypto function requested in request")
	return JsonResponse{Error: 1, Errmsg: "Unknown crypto function", Txt: ""}
}

func decryptDataHandler(privateKey *cryptkey.CryptKey) func(http.ResponseWriter, *http.Request) {
	realFun := func(w http.ResponseWriter, req *http.Request) {
		var jsonResponse JsonResponse
		var jsonString []byte
		var err error
		logger := zlogger.GetLogger("server.decryptDataHandler")
		logger.Info().
			Str("RemoteAddr", req.RemoteAddr).
			Str("User-Agent", req.Header.Get("User-Agent")).
			Msg("received_request")
		reqBody, _ := ioutil.ReadAll(req.Body)
		logger.Debug().Str("RemoteAddr", req.RemoteAddr).
			Str("data", string(reqBody)).
			Msg("received_data")

		jsonResponse = handleRequest(reqBody, privateKey)

		jsonString, _ = json.Marshal(jsonResponse)
		_, err = fmt.Fprint(w, string(jsonString))
		if err != nil {
			logger.Error().
				Err(err).
				Msg("Error while handling the request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	return realFun
}

func Start(cnf *config.Config) {
	listenOn := cnf.IpAddress + ":" + strconv.Itoa(int(cnf.Port))
	logger := zlogger.GetLogger("server.Start")
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
