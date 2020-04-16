package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		//c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		SendResponse(c, errno.ErrBind, nil)

		return
	}
	desc :=c.Query("desc")
	log.Infof("desc:%s ",desc)

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)


	lastname := c.DefaultQuery("lastname", "none")
	log.Infof("lastname: %s",lastname)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		//err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		//log.Errorf(err, "Get an error")
		return
	}
	//
	//if errno.IsErrUserNotFound(err) {
	//	log.Debug("err type is ErrUserNotFound")
	//}

	if r.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
		//err = fmt.Errorf("password is empty")
		return
	}

	//code, message := errno.DecodeErr(err)
	//c.JSON(http.StatusOK, gin.H{"code": code, "message": message})
	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
