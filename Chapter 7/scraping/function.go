package scraping

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"

	// be_trensentimen "github.com/trensentimen/be_trensen"
	module "github.com/trensentimen/be_trensen/module"
)

func init() {
	functions.HTTP("scraping", TrensentimenGetAllTopic)
}

func TrensentimenGetAllTopic(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "https://trensentimen.my.id")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization,Token")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "https://trensentimen.my.id")
	fmt.Fprintf(w, module.GCFHandlerScraping("PASETOPUBLICKEY", "MONGOSTRING", "trensentimen", r))
}
