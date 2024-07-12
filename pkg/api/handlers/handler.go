package handlers

import (
	"encoding/csv"
	"fmt"
	"gym/pkg/contract/store"
	middleware "gym/pkg/middlewares"
	"gym/pkg/models"
	"gym/pkg/usecase"
	"gym/pkg/utils"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var (
	users       = make(map[string]*models.User)
	memberships = make(map[string]*models.Membership)
	mutex       sync.Mutex
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

type Handler struct {
	Usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{usecase}
}

// func JSONResponse(c *gin.Context, statusCode int, status string, data interface{}, message string) {
// 	c.JSON(statusCode, Response{
// 		Status:  status,
// 		Data:    data,
// 		Message: message,
// 	})
// }

func JSONResponse(c *gin.Context, statusCode int, status string, data interface{}, err error, message string) {
	response := Response{
		Status:  status,
		Data:    data,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	c.JSON(statusCode, response)
}

func (uu *Handler) SendOtp(c *gin.Context) {
	var input struct {
		Contact string `json:"contact"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Invalid input")
		return
	}

	const maxAttempts = 3
	const retryDelay = 2 * time.Second // 3 seconds delay between retries

	var res string
	var err error
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		res, err = utils.SendOtp(input.Contact)
		if err == nil {
			break
		}
		log.Printf("Attempt %d: Failed to send OTP to %s, error: %v", attempt, input.Contact, err)
		time.Sleep(retryDelay)
	}
	JSONResponse(c, 200, "success", res, nil, "Otp Sent succesfully")

}

func (uu *Handler) CheckOTP(c *gin.Context) {
	var input struct {
		Contact string `json:"contact"`
		Otp     string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Invalid input")
		return
	}
	err := utils.CheckOtp(input.Contact, input.Otp)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "OTP Check failed")
		return
	}
	JSONResponse(c, 200, "success", nil, nil, "Otp Verified succesfully")
}

func (uu *Handler) RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Invalid input")
		return
	}
	user.ID = utils.GenerateID(user.Email)
	user.Verified = false
	users[user.ID] = &user

	cont := store.Session
	res, err := cont.RegisterUser(user.Name, user.Email, user.Contact)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}
	var event models.UserRegistered
	tx := res.Hash()
	fmt.Println("tx", tx)
	events, err := utils.RetrieveEvents(tx, "UserRegistered", event)
	if err != nil {
		fmt.Println("err", err)
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Failed to get events")
		return
	}

	err = uu.Usecase.RegisterUser(user)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Invalid input")
		return
	}
	JSONResponse(c, 200, "success", gin.H{"user": user, "events": events}, nil, "User registered successfully")
}

func (uu *Handler) VerifyUser(c *gin.Context) {
	var input struct {
		EmailId string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, 400, "error", nil, err, "Invalid input")
		return
	}

	// user,err:=uu.Usecase.GetUserDetailsById(input.UserID)
	// if err != nil{
	// 	if err != nil {
	// 		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "DB error")
	// 		return
	// 	}
	// }
	cont := store.Session
	res, err := cont.VerifyUser(input.EmailId)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}
	err = uu.Usecase.SetUserVerified(input.EmailId)
	if err != nil {
		if err != nil {
			JSONResponse(c, http.StatusBadRequest, "error", nil, err, "DB error")
			return
		}
	}

	JSONResponse(c, 200, "success", res.Hash(), nil, "User Verified successfully")
}

func (uu *Handler) PurchaseMembership(c *gin.Context) {
	var input struct {
		Email          string `json:"email"`
		MembershipType string `json:"membership_type"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, 400, "error", nil, err, "Invalid input")
		return
	}
	startdate, err := utils.GetStartTime(input.StartDate)
	if err != nil {
		JSONResponse(c, 400, "error", nil, err, "Time convertion error")
		return
	}

	enddate, err := utils.GetEndTime(input.EndDate)
	if err != nil {
		JSONResponse(c, 400, "error", nil, err, "Time convertion error")
		return
	}

	// user, err := uu.Usecase.GetUserDetailsById(input.UserID)
	// if err != nil {
	// 	JSONResponse(c, 400, "error", nil, err, "User Not Found")
	// 	return
	// }
	// if !user.Verified {
	// 	JSONResponse(c, 400, "error", nil, nil, "User Not Found or Not Verified")

	// 	return
	// }

	cont := store.Session

	verified, err := cont.VerifyUser(input.Email)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}
	fmt.Println("res", verified)

	res, err := cont.PurchaseMembership(input.Email, input.MembershipType, big.NewInt(startdate), big.NewInt(enddate))
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}
	tx := res.Hash()
	var event models.MembershipPurchasedEvent
	events, err := utils.RetrieveEvents(tx, "MembershipPurchased", event)
	if err != nil {
		fmt.Println("err", err)
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Failed to get events")
		return
	}
	fmt.Println("event", events)

	membershipEvent, ok := events[0].(*models.MembershipPurchasedEvent)
	if !ok {
		JSONResponse(c, http.StatusBadRequest, "error", nil, nil, "Event type mismatch")
		return
	}

	membership := &models.Membership{
		ID:        membershipEvent.MembershipId.String(),
		Email:     input.Email,
		Type:      input.MembershipType,
		StartDate: time.Unix(startdate, 0),
		EndDate:   time.Unix(enddate, 0),
		Status:    "inactive",
	}
	fmt.Println("meme", membership)

	err = uu.Usecase.PurchaseMembership(*membership)
	if err != nil {
		JSONResponse(c, 400, "error", nil, err, " failed to save details")
		return
	}

	JSONResponse(c, 200, "success", gin.H{"user": membership, "events": events}, nil, "Membership Purchased")
}

func (uu *Handler) ActivateMembership(c *gin.Context) {
	var input struct {
		MembershipID string `json:"membership_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, 400, "error", nil, err, "Invalid input")
		return
	}

	membershipID := new(big.Int)
	_, ok := membershipID.SetString(input.MembershipID, 10)
	if !ok {
		JSONResponse(c, 400, "error", nil, nil, "Invalid MembershipID format")
		return
	}

	cont := store.Session
	res, err := cont.ActivateMembership(membershipID)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}

	err = uu.Usecase.UpdateMembershipStatus(input.MembershipID, "active")
	if err != nil {
		JSONResponse(c, 400, "error", nil, err, " failed to save details")
		return
	}

	JSONResponse(c, 200, "success", res.Hash(), nil, "Membership Activated")
}

func (uu *Handler) AccessControl(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusForbidden, "error", nil, err, "Invalid input")
		return
	}

	cont := store.Session
	ress, err := cont.CheckMembershipStatus(input.Email)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}
	if ress != 1 {
		JSONResponse(c, http.StatusBadRequest, "error", nil, nil, "Access denied")
		return
	}

	// res, err := uu.Usecase.GetMembershipStatus(input.MembershipID)
	// if err != nil {
	// 	JSONResponse(c, http.StatusForbidden, "error", nil, err, "DB Error")
	// 	return
	// }
	// if res == "inactive" {
	// 	JSONResponse(c, http.StatusForbidden, "error", res, nil, "Access Denied")
	// 	return
	// }

	JSONResponse(c, 200, "success", nil, nil, "Access granted")
}

func (uu *Handler) UpdateMembershipStatus(c *gin.Context) {
	var input struct {
		MembershipID string `json:"membership_id"`
		Status       string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusForbidden, "error", nil, err, "Invalid input")
		return
	}
	membershipID := new(big.Int)
	_, ok := membershipID.SetString(input.MembershipID, 10)
	if !ok {
		JSONResponse(c, 400, "error", nil, nil, "Invalid MembershipID format")
		return
	}

	cont := store.Session
	var Res *types.Transaction
	if input.Status == "active" {
		res, err := cont.ActivateMembership(membershipID)
		if err != nil {
			JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
			return
		}
		Res = res
	} else if input.Status == "inactive" {
		res, err := cont.CancelMembership(membershipID)
		if err != nil {
			JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
			return
		}
		Res = res
	}
	err := uu.Usecase.UpdateMembershipStatus(input.MembershipID, input.Status)
	if err != nil {
		JSONResponse(c, 400, "error", nil, err, "Db Updationfailed")
		return
	}
	JSONResponse(c, 200, "success", Res.Hash(), nil, "Membership Updated")
}

func (uu *Handler) CheckMembershipStatus(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusForbidden, "error", nil, err, "Invalid input")
		return
	}
	// membershipID := new(big.Int)
	// _, ok := membershipID.SetString(input.MembershipID, 10)
	// if !ok {
	// 	JSONResponse(c, 400, "error", nil, nil, "Invalid MembershipID format")
	// 	return
	// }

	cont := store.Session
	ress, err := cont.CheckMembershipStatus(input.Email)
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}
	var status string

	if ress == 0 {
		status = "none"
	} else if ress == 1 {
		status = "active"
	} else if ress == 2 {
		status = "inactive"
	}
	// res, err := uu.Usecase.GetMembershipStatus(input.MembershipID)
	// if err != nil {
	// 	JSONResponse(c, http.StatusForbidden, "error", err, " DB Error")
	// 	return
	// }

	JSONResponse(c, 200, "success", status, nil, "Membership details found")
}

func (uu *Handler) RenewMembership(c *gin.Context) {
	var input struct {
		MembershipID string `json:"membership_id"`
		EndDate      string `json:"end_date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusForbidden, "error", nil, err, "Invalid input")
		return
	}
	enddate, err := utils.GetEndTime(input.EndDate)
	if err != nil {
		JSONResponse(c, 400, "error", nil, err, "Time Convertion error")
		return
	}
	membershipID := new(big.Int)
	_, ok := membershipID.SetString(input.MembershipID, 10)
	if !ok {
		JSONResponse(c, 400, "error", nil, nil, "Invalid MembershipID format")
		return
	}

	cont := store.Session
	res, err := cont.RenewMembership(membershipID, big.NewInt(enddate))
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Contract error")
		return
	}

	var event models.MembershipRenewedEvent
	tx := res.Hash()
	fmt.Println("tx", tx)
	events, err := utils.RetrieveEvents(tx, "MembershipRenewed", event)
	if err != nil {
		fmt.Println("err", err)
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Failed to get events")
		return
	}

	err1 := uu.Usecase.RenewMembership(input.MembershipID, time.Unix(enddate, 0))
	if err1 != nil {
		JSONResponse(c, 400, "error", nil, err1, "Db updation error")
		return
	}
	JSONResponse(c, 200, "success", gin.H{"tx": tx, "Renew Details": events}, nil, "Membership Renewed")
}

func (uu *Handler) BulkRegisterUser(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Failed to get uploaded file")
		return
	}

	filePath := filepath.Join("/tmp", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		JSONResponse(c, http.StatusInternalServerError, "error", nil, err, "Failed to save uploaded file")
		return
	}
	defer os.Remove(filePath)

	csvFile, err := os.Open(filePath)
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, "error", nil, err, "Failed to open CSV file")
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, "error", nil, err, "Failed to read CSV file")
		return
	}
	var membership []models.Membership
	var users []models.User
	for _, record := range records[1:] {
		if len(record) < 3 {
			continue
		}
		user := models.User{
			Name:     record[0],
			Email:    record[1],
			Contact:  record[2],
			IDProof:  record[3],
			Verified: true,
		}
		startdate, err := utils.GetStartTime(record[5])
		if err != nil {
			JSONResponse(c, 400, "error", nil, err, "Time convertion error")
			return
		}

		enddate, err := utils.GetEndTime(record[6])
		if err != nil {
			JSONResponse(c, 400, "error", nil, err, "Time convertion error")
			return
		}

		users = append(users, user)
		memberships := models.Membership{
			Email:     record[1],
			Type:      record[4],
			StartDate: time.Unix(startdate, 0),
			EndDate:   time.Unix(enddate, 0),
			Status:    record[7],
		}
		membership = append(membership, memberships)
	}
	var txHashes []common.Hash
	cont := store.Session
	for _, user := range users {

		tx, err := cont.RegisterUser(user.Name, user.Email, user.Contact)
		if err != nil {
			JSONResponse(c, http.StatusInternalServerError, "error", nil, err, fmt.Sprintf("Contract error for user %s", user.Email))
			return
		}
		txHashes = append(txHashes, tx.Hash())
		user.ID = utils.GenerateID(user.Email)
		user.Verified = false
		err = uu.Usecase.RegisterUser(user)
		if err != nil {
			JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Invalid input")
			return
		}
	}
	for _, mem := range membership {
		st := mem.StartDate.Unix()
		ed := mem.EndDate.Unix()

		res, err := cont.PurchaseMembership(mem.Email, mem.Type, big.NewInt(st), big.NewInt(ed))
		if err != nil {
			JSONResponse(c, http.StatusInternalServerError, "error", nil, err, fmt.Sprintf("Contract error for user %s", mem.Email))
			return
		}
		tx := res.Hash()
		var event models.MembershipPurchasedEvent
		events, err := utils.RetrieveEvents(tx, "MembershipPurchased", event)
		if err != nil {
			fmt.Println("err", err)
			JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Failed to get events")
			return
		}
		membershipEvent, ok := events[0].(*models.MembershipPurchasedEvent)
		if !ok {
			JSONResponse(c, http.StatusBadRequest, "error", nil, nil, "Event type mismatch")
			return
		}
		mem.ID = membershipEvent.MembershipId.String()
		fmt.Println("mem", mem)
		err = uu.Usecase.PurchaseMembership(mem)
		if err != nil {
			JSONResponse(c, http.StatusBadRequest, "error", nil, err, "Invalid input")
			return
		}

	}

	JSONResponse(c, http.StatusOK, "success", gin.H{"txhashes": txHashes}, nil, "Users registered successfully")
}

func (uu *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusBadGateway, "error", nil, err, "Invalid input")
		return
	}
	res, err := uu.Usecase.Login(input.Email)
	if err != nil {
		JSONResponse(c, http.StatusBadGateway, "error", nil, err, " Logine error")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password)); err != nil {
		JSONResponse(c, http.StatusUnauthorized, "error", nil, err, "Invalid credentials")
		return
	}
	token, err := middleware.GenerateJWT(input.Email, res.Role)
	if err != nil {
		JSONResponse(c, http.StatusInternalServerError, "error", nil, err, "Failed to generate token")
		return
	}
	JSONResponse(c, 200, "success", token, nil, "Login Success")

}

func (uu *Handler) RegisterAdmin(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		JSONResponse(c, http.StatusBadGateway, "error", nil, err, "Invalid input")
		return
	}
	admin := models.Admins{
		Email:    input.Email,
		Password: utils.HashPassword(input.Password),
		Role:     "admin",
	}
	err := uu.Usecase.RegisterAdmin(admin)
	if err != nil {
		JSONResponse(c, http.StatusBadGateway, "error", nil, err, "DB Error")
		return
	}
	JSONResponse(c, 200, "success", nil, nil, "admin created succesfully")

}
