/**
 * Created with IntelliJ IDEA.
 * User: syncim
 * Date: 11/13/12
 * Time: 10:54 PM
 * To change this template use File | Settings | File Templates.
 */
package controllers

import ("github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/thesyncim/playground/website/template"
	"github.com/gorilla/sessions"


)

var Store = sessions.NewCookieStore([]byte("something-very-secret"))

type Entry struct {
	Title	  string
	Content	string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	entry := &Entry{}
	entry.Title = "New Go Templates (mid-2011)"
	entry.Content = "A trivial example of how to use them.  see documentation and Go source code/tests for more examples"

	template.RenderTemplate(w, "/home/syncim/IdeaProjects/playground/src/template/projectos/index.html", entry)
	session, _ := Store.Get(r, "session-name")

	session.Values["teste"] = 1

	tes := fmt.Sprint(session.Values["teste"])


	fmt.Fprintf(w, tes)
	vars := mux.Vars(r)
	category := vars["teste"]
	fmt.Fprintf(w, category)





}


