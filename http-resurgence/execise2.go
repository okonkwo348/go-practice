package main

import (
	"fmt"
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// text := r.FormValue("text")
	// w.Header().Set("Content-Type", "text/plain")
	// w.Write(([]byte((text))))

	//putting w.Write() before w.Header()Set() rearrenge the ecntents in the header.
	// changes (content-Type, Date and Content-Length) to (Date, content-Length and Content-Type)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	if len(body) == 0 {

		http.Error(w, "Bad Request", http.StatusBadRequest)
		return

	}

	fmt.Fprint(w, string(body))
}

/*Think about —
1. What does io.ReadAll return if the request has no body at all?
	It returns an empty []byte and nil err. that is an empty body, len(body)--0
2. Is len(body) == 0 the same as body == nil? Try both and see.
	No, they are not the same.
	body==nil means the slice was never initialized at all
	len(body)==0 means the slice exists but has no elements.
	io.ReadAll always returns an initialized slice, never nil.
	So body == nil would never be true after io.ReadAll.
3. What happens to r.Body if you read it twice without closing it?
	Second read gets nothing — empty bytes. The stream cursor is at
	the end after the first rea

*/
