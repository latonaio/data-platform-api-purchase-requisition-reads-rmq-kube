package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header *[]Header `json:"Header"`
	Item   *[]Item   `json:"Item"`
}

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

type Item struct {
	PurchaseRequisition                           int      `json:"PurchaseRequisition"`
	PurchaseRequisitionItem                       int      `json:"PurchaseRequisitionItem"`
	PurchaseRequisitionItemCategory               string   `json:"PurchaseRequisitionItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID       *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	SupplyChainRelationshipProductionPlantID      *int     `json:"SupplyChainRelationshipProductionPlantID"`
	Buyer                                         int      `json:"Buyer"`
	Seller                                        *int     `json:"Seller"`
	DeliverToParty                                int      `json:"DeliverToParty"`
	DeliverFromParty                              *int     `json:"DeliverFromParty"`
	DeliverToPlant                                string   `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation                 *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                              *string  `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation               *string  `json:"DeliverFromPlantStorageLocation"`
	Product                                       string   `json:"Product"`
	ProductGroup                                  *string  `json:"ProductGroup"`
	RequestedQuantity                             float32  `json:"RequestedQuantity"`
	BaseUnit                                      string   `json:"BaseUnit"`
	DeliveryUnit                                  *string  `json:"DeliveryUnit"`
	PlannedOrder                                  *int     `json:"PlannedOrder"`
	PlannedOrderItem                              *int     `json:"PlannedOrderItem"`
	PrecedingOrderID                              *int     `json:"PrecedingOrderID"`
	PrecedingOrderItem                            *int     `json:"PrecedingOrderItem"`
	FollowingOrderID                              *int     `json:"FollowingOrderID"`
	FollowingOrderItem                            *int     `json:"FollowingOrderItem"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	PurchaseRequisitionReleaseStatus              *string  `json:"PurchaseRequisitionReleaseStatus"`
	PurchaseRequisitionReleaseDate                *string  `json:"PurchaseRequisitionReleaseDate"`
	PurchaseRequisitionItemText                   *string  `json:"PurchaseRequisitionItemText"`
	PurchaseRequisitionItemTextByBuyer            *string  `json:"PurchaseRequisitionItemTextByBuyer"`
	PurchaseRequisitionItemTextBySeller           *string  `json:"PurchaseRequisitionItemTextBySeller"`
	PurchaseRequisitionPrice                      *float32 `json:"PurchaseRequisitionPrice"`
	PurchaseRequisitionPriceQuantity              *int     `json:"PurchaseRequisitionPriceQuantity"`
	ProductPlannedDeliveryDuration                *float32 `json:"ProductPlannedDeliveryDuration"`
	ProductPlannedDeliveryDurationUnit            *int     `json:"ProductPlannedDeliveryDurationUnit"`
	OrderedQuantity                               *float32 `json:"OrderedQuantity"`
	DeliveryDate                                  *string  `json:"DeliveryDate"`
	PurchaseRequisitionItemStatus                 *string  `json:"PurchaseRequisitionItemStatus"`
	PurchaseRequisitionItemCurrency               *string  `json:"PurchaseRequisitionItemCurrency"`
	MRPArea                                       *string  `json:"MRPArea"`
	MRPController                                 *string  `json:"MRPController"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	GLAccount                                     *string  `json:"GLAccount"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	CreationDate                                  string   `json:"CreationDate"`
	LastChangeTime                                string   `json:"LastChangeTime"`
	IsCancelled                                   *int     `json:"IsCancelled"`
	IsMarkedForDeletion                           *int     `json:"IsMarkedForDeletion"`
}
