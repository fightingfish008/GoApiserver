package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/util"
	"github.com/lexkong/log/lager"

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
	//
	//if errno.IsErrUserNotFound(err) {
	//	log.Debug("err type is ErrUserNotFound")
	//}


	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	if err := r.checkParam(); err != nil {
		SendResponse(c, err, nil)
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

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}