package handler

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/model"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type User struct {
	service service.IUser
	user    model.User
}

func NewUser(service service.IUser) *User {
	return &User{service: service}
}

// GetAllUser godoc
// @Summary get all user
// @Description get all user item
// @Tags users
// @Accept json
// @Produce json
// @Router /client/get-all-user [get]
func (u *User) GetAllUser(ctx *gin.Context) {
	users, err := u.service.GetAllUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   gin.H{"users": users},
	})
}

// GetUserByRole godoc
// @Summary get user role
// @Description get user by role
// @Tags users
// @Accept json
// @Produce json
// @Router /client/get-user/role [get]
func (u *User) GetUserByRole(ctx *gin.Context) {
	var role string
	data, err := u.service.GetUserByRole(role)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
			"status":  "fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})

}

// FindUserByRole godoc
// @Summary get user role
// @Description get user by role
// @Tags users
// @Accept json
// @Produce json
// @Router /client/get-user-role [get]
func (u *User) FindUserByRole(ctx *gin.Context) {
	var role string
	data, err := u.service.FindUserByRole(role)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "message",
		"data":   data,
	})
}

// GetUserByEmail godoc
// @Summary get user email
// @Description get user by email
// @Tags users
// @Accept json
// @Produce json
// @Router /client/get-user/email [get]
func (u *User) GetUserByEmail(ctx *gin.Context) {
	var email string
	data, err := u.service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// GetUserByUsername godoc
// @Summary get user
// @Description get a new user item by username
// @Tags users
// @Accept json
// @Produce json
// @Router /client/get-user/username [get]
func (u *User) GetUserByUsername(ctx *gin.Context) {
	var username string
	data, err := u.service.GetUserByUsername(username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// GetMeV2 godoc
// @Summary get user
// @Description get a new user item
// @Tags users
// @Accept json
// @Produce json
// @Router /client/get-user [get]
func (u *User) GetMeV2(ctx *gin.Context) {
	cookie, err := ctx.Cookie("access_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "You are not login!",
		})
		return
	}

	cfg, _ := conf.LoadConfig(".")

	sub, err := util.ValidateToken(cookie, cfg.AccessTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	result, err := u.service.FindUserByID(fmt.Sprint(sub))
	resultString, err := json.Marshal(result)
	userResponse := &model.UserResponse{
		ID:        result.ID,
		Name:      result.FullName,
		Email:     result.Email,
		Photo:     result.Photo,
		Role:      result.Role,
		Enable:    result.Enable,
		Provider:  result.Provider,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": string(resultString) + "the user belonging to this token no logger exists",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"user":   userResponse,
	})
}

// UpdateUser godoc
// @Summary login user
// @Description update a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "login user"
// @Router /client/update [put]
func (u *User) UpdateUser(ctx *gin.Context) {
	cookie, err := ctx.Cookie("access_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "You are not login!",
		})
		return
	}

	cfg, _ := conf.LoadConfig(".")

	sub, err := util.ValidateToken(cookie, cfg.AccessTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result, err := u.service.FindUserByID(fmt.Sprint(sub))
	resultString, err := json.Marshal(result)
	userResponse := model.User{
		FullName:  result.FullName,
		Username:  result.Username,
		Email:     result.Email,
		Photo:     result.Photo,
		Phone:     result.Phone,
		Password:  result.Password,
		Role:      result.Role,
		Enable:    result.Enable,
		Provider:  result.Provider,
		UpdatedAt: result.UpdatedAt,
	}

	userResponse.Password, err = util.HashPassword(userResponse.Password)
	result2, err := u.service.UpdateUser(userResponse)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": string(resultString) + "the user belonging to this token no logger exists",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Updated",
		"user":    result2,
	})
}

// LoginWithUserName godoc
// @Summary login user
// @Description login user with username
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "login user"
// @Router /client/login/username [post]
func (u *User) LoginWithUserName(ctx *gin.Context) {
	// Lấy thông tin từ request
	userRequest := model.SignInWithUsername{}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	data, err := u.service.LoginUserByUsername(userRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	cfg, _ := conf.LoadConfig(".")

	// Generate token
	accessToken, err := util.CreateToken(cfg.AccessTokenExpiresIn, data.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	refreshToken, err := util.CreateToken(cfg.RefreshTokenExpiresIn, data.ID, cfg.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.SetCookie("access_token", accessToken, cfg.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, cfg.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", cfg.AccessTokenMaxAge*60, "/", "localhost", false, false)

	// Trả về thông báo login thành công
	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login successful",
		"user":    data,
	})
}

// LoginWithEmail godoc
// @Summary login user
// @Description login user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.SignInInput true "login user"
// @Router /client/login/email [post]
func (u *User) LoginWithEmail(ctx *gin.Context) {
	//  Lấy thông tin từ request
	userRequest := model.SignInWithEmail{}

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	data, err := u.service.LoginUserByEmail(userRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	cfg, _ := conf.LoadConfig(".")

	// Generate token
	accessToken, err := util.CreateToken(cfg.AccessTokenExpiresIn, data.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error()},
		)
		return
	}

	refreshToken, err := util.CreateToken(cfg.RefreshTokenExpiresIn, data.ID, cfg.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error()},
		)
		return
	}

	ctx.SetCookie("access_token", accessToken, cfg.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, cfg.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", cfg.AccessTokenMaxAge*60, "/", "localhost", false, false)

	// Trả về thông báo login thành công
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

// Register godoc
// @Summary register user
// @Description Create a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "register user"
// @Router /client/register [post]
func (u *User) Register(ctx *gin.Context) {
	var user model.SignUpInput
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	if !util.EmailValid(user.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email không hợp lệ !",
		})
		return
	}

	if !util.PhoneValid(user.Phone) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Số điện thoại không đúng chuẩn",
		})
		return
	}

	// Bên phía client sẽ phải so sánh password thêm một lần nữa đã đúng chưa
	if !util.PasswordStrong(user.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Mật khẩu phải có ít nhất 8 ký tự, bao gồm chữ hoa, chữ thường và số !",
		})
		return
	}

	// Băm mật khẩu
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	user.Password = hashedPassword
	user.Password = util.Santize(user.Password)
	user.Username = util.Santize(user.Username)
	user.Email = util.Santize(user.Email)
	user.Email = strings.ToLower(user.Email)

	// Generate Verification Code
	//code := randstr.String(8)
	//verificationCode := util.Encode(code)

	newUser := model.User{
		FullName: user.FullName,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
		Role:     "user",
		//VerificationCode: verificationCode,
		Verified:   false,
		Provider:   "Cỏ Studio",
		WhoUpdates: user.FullName,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Enable:     1,
	}

	// thực hiện đăng ký người dùng
	userResponse, err := u.service.RegisterUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error()},
		)
		return
	}

	//var firstName = newUser.FullName
	//if strings.Contains(firstName, "") {
	//	firstName = strings.Split(firstName, " ")[0]
	//}
	//
	//config, _ := conf.LoadConfig(".")
	//
	//emailData := util.EmailData{
	//	URL:       config.ClientOrigin + "/verify_email/" + code,
	//	FirstName: firstName,
	//	Subject:   "Your account verification code",
	//}
	//
	//util.SendEmail(&newUser, &emailData, "*/verificationCode.html")
	//message := "We sent an email with a verification code to " + newUser.Email
	//// Trả về phản hồi thành công
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		//"message": message,
		"data": userResponse,
	})
}

// RefreshAccessToken godoc
// @Summary refresh token user
// @Description refresh token item
// @Tags users
// @Accept json
// @Produce json
// @Router /client/refresh [get]
func (u *User) RefreshAccessToken(ctx *gin.Context) {
	message := "could not fresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": message},
		)
		return
	}

	cfg, _ := conf.LoadConfig(".")

	sub, err := util.ValidateToken(cookie, cfg.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": err.Error()},
		)
		return
	}

	result, err := u.service.FindUserByID(fmt.Sprint(sub))
	resultString, err := json.Marshal(result)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": string(resultString) + "the user belonging to this token no logger exists"},
		)
		return
	}

	accessToken, err := util.CreateToken(cfg.AccessTokenExpiresIn, result.ID, cfg.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"status":  "fail",
			"message": err.Error()},
		)
		return
	}

	ctx.SetCookie("access_token", accessToken, cfg.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", cfg.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{
		"status":       "success",
		"access_token": accessToken,
	})
}

// LogoutUser godoc
// @Summary logout user
// @Description logout item
// @Tags users
// @Accept json
// @Produce json
// @Router /client/logout [get]
func (u *User) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

// ResolvePassword godoc
// @Summary resolve password of user
// @Description resolve password of user
// @Tags users
// @Accept json
// @Produce json
// @Router /client/logout [get]
func (u *User) ResolvePassword(ctx *gin.Context) {

}
