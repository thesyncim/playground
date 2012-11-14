/**
 * Created with IntelliJ IDEA.
 * User: syncim
 * Date: 11/13/12
 * Time: 2:54 PM
 * To change this template use File | Settings | File Templates.
 */

package website

import ("github.com/gorilla/mux"
	"net/http"


)

func main() {
	r := mux.NewRouter()
//	r.HandleFunc("/{teste}/{id}", controllers.HomeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8090", r)

}
