package twilio

import (
	"backend-starter-project/domain/interfaces"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioService struct {
	client *twilio.RestClient
	from   string
}

func NewTwilioService(accountSID, authToken, from string) interfaces.TwilioService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSID,
		Password: authToken,
	})
	return &twilioService{client: client, from: from}
}

func (service *twilioService) SendOTP(phoneNumber, otp string) error {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo("+251"+phoneNumber)
	params.SetFrom("+16502295720")
	params.SetBody("Your OTP is: " + otp)
	_, err := service.client.Api.CreateMessage(params)

	return err
}
