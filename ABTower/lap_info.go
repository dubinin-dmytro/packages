package ABTower

type LapInfoResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    LapInfo `json:"data"`
}

type LapInfo struct {
	EventList struct {
		Original struct {
			Total struct {
				Number int `json:"number"`
			} `json:"total"`
			Registered struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"registered"`
			ChoiceShow struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"choice_show"`
			GoToEditor struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"go_to_editor"`
			Payment struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"payment"`
			AmountToTotal struct {
				Number int `json:"number"`
			} `json:"amount_to_total"`
			SupportTicket struct {
				Number int `json:"number"`
			} `json:"support_ticket"`
			SubscriptionCancel struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"subscription_cancel"`
			Refund struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"refund"`
			WillPay struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"will_pay"`
			WillPayAmount struct {
				Number int `json:"number"`
			} `json:"will_pay_amount"`
			WillPayAmountToTotal struct {
				Number int `json:"number"`
			} `json:"will_pay_amount_to_total"`
			CanceledBeforePay struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"canceled_before_pay"`
			CanceledBeforePayAmount struct {
				Number int `json:"number"`
			} `json:"canceled_before_pay_amount"`
			PartialRefund struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"partial_refund"`
			SupportChat struct {
				Number int `json:"number"`
			} `json:"support_chat"`
			SupportCall struct {
				Number int `json:"number"`
			} `json:"support_call"`
			PaymentAmount struct {
				Number int `json:"number"`
			} `json:"payment_amount"`
			CcGiven struct {
				Number int `json:"number"`
			} `json:"cc_given"`
			CcGivenAmount struct {
				Number float64 `json:"number"`
			} `json:"cc_given_amount"`
		} `json:"original"`
		MultiCurl struct {
			Total struct {
				Number int `json:"number"`
			} `json:"total"`
			Registered struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"registered"`
			ChoiceShow struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"choice_show"`
			GoToEditor struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"go_to_editor"`
			Payment struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"payment"`
			AmountToTotal struct {
				Number float64 `json:"number"`
			} `json:"amount_to_total"`
			SupportTicket struct {
				Number int `json:"number"`
			} `json:"support_ticket"`
			SubscriptionCancel struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"subscription_cancel"`
			Refund struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"refund"`
			WillPay struct {
				Number     int `json:"number"`
				Conversion int `json:"conversion"`
			} `json:"will_pay"`
			WillPayAmount struct {
				Number int `json:"number"`
			} `json:"will_pay_amount"`
			WillPayAmountToTotal struct {
				Number int `json:"number"`
			} `json:"will_pay_amount_to_total"`
			CanceledBeforePay struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"canceled_before_pay"`
			CanceledBeforePayAmount struct {
				Number float64 `json:"number"`
			} `json:"canceled_before_pay_amount"`
			PartialRefund struct {
				Number     int     `json:"number"`
				Conversion float64 `json:"conversion"`
			} `json:"partial_refund"`
			SupportChat struct {
				Number int `json:"number"`
			} `json:"support_chat"`
			SupportCall struct {
				Number int `json:"number"`
			} `json:"support_call"`
			PaymentAmount struct {
				Number float64 `json:"number"`
			} `json:"payment_amount"`
			CcGiven struct {
				Number int `json:"number"`
			} `json:"cc_given"`
			CcGivenAmount struct {
				Number float64 `json:"number"`
			} `json:"cc_given_amount"`
		} `json:"multi-curl"`
	} `json:"eventList"`
	PValues struct {
		Original struct {
			Registered struct {
				WinRate float64 `json:"win-rate"`
				PValue  float64 `json:"p-value"`
			} `json:"registered"`
			GoToEditor struct {
				WinRate float64 `json:"win-rate"`
				PValue  float64 `json:"p-value"`
			} `json:"go_to_editor"`
			Payment struct {
				WinRate float64 `json:"win-rate"`
				PValue  float64 `json:"p-value"`
			} `json:"payment"`
		} `json:"original"`
		MultiCurl struct {
			Registered struct {
				WinRate float64 `json:"win-rate"`
				PValue  float64 `json:"p-value"`
			} `json:"registered"`
			GoToEditor struct {
				WinRate float64 `json:"win-rate"`
				PValue  float64 `json:"p-value"`
			} `json:"go_to_editor"`
			Payment struct {
				WinRate float64 `json:"win-rate"`
				PValue  float64 `json:"p-value"`
			} `json:"payment"`
		} `json:"multi-curl"`
	} `json:"pValues"`
}

func (l LapInfoResponse) isSuccess() bool {
	return l.Success
}

func (l LapInfoResponse) errorMessage() string {
	if l.Success {
		return ""
	}
	return l.Message
}

func (l LapInfoResponse) getData() LapInfo {
	if !l.Success {
		return LapInfo{}
	}
	return l.Data
}
