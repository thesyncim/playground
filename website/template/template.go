/**
 * Created with IntelliJ IDEA.
 * User: syncim
 * Date: 11/13/12
 * Time: 11:16 PM
 * To change this template use File | Settings | File Templates.
 */
package template

import (
	"net/http"
	"html/template"
	"github.com/thesyncim/playground/website/controllers"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *controllers.Entry) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

