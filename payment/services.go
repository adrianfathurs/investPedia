package payment

import (
	"investPedia/user"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	midtrans "github.com/veritrans/go-midtrans"
)

type services struct {
}

type Services interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

func NewService() *services {
	return &services{}
}

func (s *services) GetPaymentURL(transaction Transaction, user user.User) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// getting env variables SITE_TITLE and DB_HOST
	midtransClientKey := os.Getenv("MIDTRANS_CLIENT_KEY")
	midtransServerKey := os.Getenv("MIDTRANS_SERVER_KEY")

	midclient := midtrans.NewClient()
	midclient.ServerKey = midtransServerKey
	midclient.ClientKey = midtransClientKey
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.FullName,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", nil
	}
	return snapTokenResp.RedirectURL, nil
}
