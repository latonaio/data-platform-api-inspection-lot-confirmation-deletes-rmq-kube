package requests

type Header struct {
	InspectionLot          int   `json:"InspectionLot"`
	Operations             int   `json:"Operations"`
	OperationsItem         int   `json:"OperationsItem"`
	OperationID            int   `json:"OperationID"`
	ConfirmationCountingID int   `json:"ConfirmationCountingID"`
	IsMarkedForDeletion    *bool `json:"IsMarkedForDeletion"`
}
