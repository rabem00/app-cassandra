package Messages

import (
	"net/http"

	"github.com/gocql/gocql"
	"github.com/rabem00/app-cassandra/Cassandra"
)

// Post -- handles POST request to /messages/new to create a new message
// params:
// w - response writer for building JSON payload response
// r - request reader to fetch form data or url params
func Post(w http.ResponseWriter, r *http.Request) {
	var errs []string
	var errStr, userIDStr, message string

	if userIDStr, errStr = processFormField(r, "userID"); len(errStr) != 0 {
		errs = append(errs, errStr)
	}
	userID, err := gocql.ParseUUID(userIDStr)
	if err != nil {
		errs = append(errs, "Parameter 'userID' not a UUID")
	}

	if message, errStr = processFormField(r, "message"); len(errStr) != 0 {
		errs = append(errs, errStr)
	}

	gocqlUUID := gocql.TimeUUID()

	if len(errs) == 0 {
		if err := Cassandra.Session.Query(`
		INSERT INTO messages (id, userID, message) VALUES (?, ?, ?)`,
			gocqlUUID, userID, message).Exec(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	/*
		if created {
			// send message to Stream
			globalMessages, err := Stream.Client.FlatFeed("messages", "global")
			if err == nil {
				globalMessages.AddActivity(&getstream.Activity{
					Actor:  getstream.FeedID(userID.String()),
					Verb:   "post",
					Object: getstream.FeedID(gocqlUUID.String()),
					MetaData: map[string]string{
						// add as many custom keys/values here as you like
						"message": message,
					},
				})
			}

			json.NewEncoder(w).Encode(NewMessageResponse{ID: gocqlUUID})
		} else {
			json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
		}
	*/
}
