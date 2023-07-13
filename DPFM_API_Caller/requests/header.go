package requests

type Header struct {
	PurchaseRequisition                int    `json:"PurchaseRequisition"`
	Buyer                              int    `json:"Buyer"`
	PurchaseRequisitionType            string `json:"PurchaseRequisitionType"`
	PlannedOrder                       *int   `json:"PlannedOrder"`
	PlannedOrderItem                   *int   `json:"PlannedOrderItem"`
	PrecedingOrderID                   *int   `json:"PrecedingOrderID"`
	PrecedingOrderItem                 *int   `json:"PrecedingOrderItem"`
	Project                            *int   `json:"Project"`
	WBSElement                         *int   `json:"WBSElement"`
	CreationDate                       string `json:"CreationDate"`
	LastChangeDate                     string `json:"LastChangeDate"`
	IsCancelled                        *bool  `json:"IsCancelled"`
	IsMarkedForDeletion                *bool  `json:"IsMarkedForDeletion"`
}
