package route

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/tanakrid/accounting/transaction"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func transactionHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
		case "POST":
			body, err := io.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error: %v", err)
				return
			}

			record := transaction.Record{}
			err = json.Unmarshal(body, &record)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error: %v", err)
				return
			}

			transaction.Add(uuid.New().String(), record)
			w.WriteHeader(http.StatusOK)
			return
		case "GET":
			b, err := json.Marshal(transaction.Show())
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "error: %v", err)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Write(b)

		case "PUT":
			body, err := io.ReadAll(req.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error: %v", err)
				return
			}

			id := req.URL.Query().Get("id")
			fmt.Println("id =>", id)

			record := transaction.Record{}
			err = json.Unmarshal(body, &record)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error: %v", err)
				return
			}

			transaction.Edit(id, record)
			w.WriteHeader(http.StatusOK)
			return
		case "DELETE":
			id := req.URL.Query().Get("id")
			fmt.Println("id =>", id)
			transaction.Del(id)
			w.WriteHeader(http.StatusOK)
			return
	}
}

func reportHandler(w http.ResponseWriter, req *http.Request) {

}

func filterHandler(w http.ResponseWriter, req *http.Request) {

}
func InitRoute() {
	fmt.Println("init route and start service")
	http.HandleFunc("/transaction", transactionHandler)
	http.HandleFunc("/report", reportHandler)
	http.HandleFunc("/filter", filterHandler)

	err := http.ListenAndServe("localhost:4444", nil)
	log.Fatal(err)
}