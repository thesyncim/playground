/**
 * Created with IntelliJ IDEA.
 * User: syncim
 * Date: 11/13/12
 * Time: 3:50 PM
 * To change this template use File | Settings | File Templates.
 */
package models

import (
	"github.com/ungerik/go-start/media"
	"github.com/ungerik/go-start/model"
	"github.com/ungerik/go-start/mongo"
	"github.com/ungerik/go-start/user"


)

func init() {


}

var Users = mongo.NewCollection("projectos")

func NewUser() *User {
	var doc User
	Users.InitDocument(&doc)
	return &doc
}

type User struct {
	user.User `bson:",inline"`

	Image  media.ImageRef
	Gender model.Choice `model:"options=,Male,Female"`
}
