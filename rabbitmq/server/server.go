package server

import (
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"

    amqp "github.com/rabbitmq/amqp091-go"
)

func addBasicHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
    w.Header().Set("Content-Type", "application/json")
}

type responseData struct {
    Note    string  `json:"note"`
}

type answerData struct {
    Response    string  `json:"response"`
}

func readJson(w http.ResponseWriter, req *http.Request, respdata *responseData) (error) {
    resp, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal([]byte(resp), respdata)
    if err != nil {
        return err
    }

    return err
}


func SendNote(w http.ResponseWriter, req *http.Request) {
    var respdata responseData

    log.Println("POST /sendNote")

    addBasicHeaders(w);

    err := readJson(w, req, &respdata)
    if err != nil {
        log.Print("Error unmarshalling JSON")
        w.WriteHeader(500)

        answer := answerData {
                            Response: "Wrong input params",
                         }
        json.NewEncoder(w).Encode(answer)

        return
    } else {
        w.WriteHeader(200)
        answer := answerData {
                            Response: "OK",
                         }
        json.NewEncoder(w).Encode(answer)
    }
}
