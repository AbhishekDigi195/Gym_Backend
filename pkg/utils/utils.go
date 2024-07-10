package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"gym/pkg/chain"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/twilio/twilio-go"
	"golang.org/x/crypto/bcrypt"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

func GenerateID(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
func ParseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// // GetStartTime returns the start of the day in UTC for a given date string
// func GetStartTime(dateStr string) (time.Time, error) {
// 	t, err := ParseDate(dateStr)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	return t.In(time.UTC), nil
// }

// // GetEndTime returns the end of the day in UTC for a given date string
// func GetEndTime(dateStr string) (time.Time, error) {
// 	t, err := ParseDate(dateStr)
// 	if err != nil {
// 		return time.Time{}, err
// 	}
// 	endTime := t.Add(time.Hour*24 - time.Nanosecond).In(time.UTC)
// 	return endTime, nil
// }

func GetStartTime(dateStr string) (int64, error) {
	t, err := ParseDate(dateStr)
	if err != nil {
		return 0, err
	}
	return t.In(time.UTC).Unix(), nil
}

// GetEndTime returns the Unix timestamp for the end of the day in UTC for a given date string.
func GetEndTime(dateStr string) (int64, error) {
	t, err := ParseDate(dateStr)
	if err != nil {
		return 0, err
	}
	endTime := t.Add(time.Hour*24 - time.Nanosecond).In(time.UTC)
	return endTime.Unix(), nil
}

func RetrieveEvents(txHash common.Hash, eventName string, eventStructType interface{}) ([]interface{}, error) {

	var receipt *types.Receipt
	maxRetries := 10
	retryDelay := time.Second * 2
	var err error
	for i := 0; i < maxRetries; i++ {
		receipt, err = chain.Client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			break
		}
		log.Printf("Failed to retrieve transaction receipt: %v. Retrying in %v...", err, retryDelay)
		time.Sleep(retryDelay)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve transaction receipt after %d retries: %v", maxRetries, err)
	}
	contractAddress := os.Getenv("TOKEN_ADDRESS")
	Address := common.HexToAddress(contractAddress)

	abiString := ABI
	contractABI, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	_, exists := contractABI.Events[eventName]
	if !exists {
		return nil, fmt.Errorf("event %s not found in contract ABI", eventName)
	}

	// Initialize a slice to store the events
	var events []interface{}

	// Process logs to extract event data
	for _, vLog := range receipt.Logs {
		if vLog.Address == Address {
			// Create a new instance of the event struct
			eventInstance := reflect.New(reflect.TypeOf(eventStructType)).Interface()

			// Unpack log data into event structure
			err := contractABI.UnpackIntoInterface(eventInstance, eventName, vLog.Data)
			if err != nil {
				return nil, fmt.Errorf("failed to unpack event log data: %v", err)
			}

			// Unpack indexed topics into event structure
			// eventStructValue := reflect.ValueOf(eventInstance).Elem()
			// for i, input := range eventABI.Inputs {
			// 	if input.Indexed {
			// 		// Ensure the topics slice has enough elements
			// 		if len(vLog.Topics) <= i {
			// 			return nil, fmt.Errorf("not enough topics in event log for indexed parameter: %s", input.Name)
			// 		}
			// 		// Unpack the indexed field
			// 		unpacked, err := abi.Arguments{input}.Unpack(vLog.Topics[i+1].Bytes())
			// 		if err != nil {
			// 			return nil, fmt.Errorf("failed to unpack event log topic: %v", err)
			// 		}
			// 		// Set the field value in the struct
			// 		eventStructValue.FieldByName(strings.Title(input.Name)).Set(reflect.ValueOf(unpacked[0]))
			// 	}
			// }

			// Append the event instance to the events slice
			events = append(events, eventInstance)
		}
	}

	return events, nil
}

func SendOtp(phone string) (string, error) {

	TWILIO_ACCOUNT_SID := os.Getenv("ACC_SID")
	TWILIO_AUTH_TOKEN := os.Getenv("ACC_TOKEN")
	VERIFY_SERVICE_SID := os.Getenv("ACC_SERVICE")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})

	to := "+91" + phone
	params := &openapi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(VERIFY_SERVICE_SID, params)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("failed to generate otp")
	} else {
		fmt.Printf("verification code send '%s \n'", *resp.Sid)
		return *resp.Sid, nil
	}
}

func CheckOtp(phone, code string) error {

	TWILIO_ACCOUNT_SID := os.Getenv("ACC_SID")
	TWILIO_AUTH_TOKEN := os.Getenv("ACC_TOKEN")
	VERIFY_SERVICE_SID := os.Getenv("ACC_SERVICE")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: TWILIO_ACCOUNT_SID,
		Password: TWILIO_AUTH_TOKEN,
	})
	if code == "" {
		return errors.New("OTP code is empty")
	}
	to := "+91" + phone

	params := &openapi.CreateVerificationCheckParams{}

	params.SetTo(to)
	params.SetCode(code)
	fmt.Print(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(VERIFY_SERVICE_SID, params)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("invalid otp")
	} else if *resp.Status == "approved" {
		return nil
	} else {
		return errors.New("invalid otp")
	}
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword)
}

const ABI = `[
	{
		"inputs": [],
		"name": "InvalidInitialization",
		"type": "error"
	},
	{
		"inputs": [],
		"name": "NotInitializing",
		"type": "error"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "version",
				"type": "uint64"
			}
		],
		"name": "Initialized",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "membershipId",
				"type": "uint256"
			}
		],
		"name": "MembershipActivated",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "membershipId",
				"type": "uint256"
			}
		],
		"name": "MembershipCancelled",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "membershipId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "membershipType",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "startDate",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "endDate",
				"type": "uint256"
			}
		],
		"name": "MembershipPurchased",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "membershipId",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "newEndDate",
				"type": "uint256"
			}
		],
		"name": "MembershipRenewed",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			}
		],
		"name": "UserDeleted",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "name",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "contact",
				"type": "string"
			}
		],
		"name": "UserRegistered",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"internalType": "string",
				"name": "email",
				"type": "string"
			}
		],
		"name": "UserVerified",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_membershipId",
				"type": "uint256"
			}
		],
		"name": "activateMembership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_membershipId",
				"type": "uint256"
			}
		],
		"name": "cancelMembership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_email",
				"type": "string"
			}
		],
		"name": "checkMembershipStatus",
		"outputs": [
			{
				"internalType": "enum GymMembership.MembershipStatus",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_email",
				"type": "string"
			}
		],
		"name": "deleteUser",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_membershipId",
				"type": "uint256"
			}
		],
		"name": "getMembershipDetails",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			},
			{
				"internalType": "enum GymMembership.MembershipStatus",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "initialize",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "membershipCounter",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "memberships",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "membershipType",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "startDate",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "endDate",
				"type": "uint256"
			},
			{
				"internalType": "enum GymMembership.MembershipStatus",
				"name": "status",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "owner",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_email",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_membershipType",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "_startDate",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_endDate",
				"type": "uint256"
			}
		],
		"name": "purchaseMembership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_name",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_email",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "_contact",
				"type": "string"
			}
		],
		"name": "registerUser",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "_membershipId",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "_newEndDate",
				"type": "uint256"
			}
		],
		"name": "renewMembership",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"name": "users",
		"outputs": [
			{
				"internalType": "string",
				"name": "name",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "email",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "contact",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "membershipId",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "_email",
				"type": "string"
			}
		],
		"name": "verifyUser",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
