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
	PurchaseRequisitionDate            string `json:"PurchaseRequisitionDate"`
	Buyer                              int    `json:"Buyer"`
	PurchaseRequisitionType            string `json:"PurchaseRequisitionType"`
	PlannedOrder                       *int   `json:"PlannedOrder"`
	PlannedOrderItem                   *int   `json:"PlannedOrderItem"`
	ProductionOrder                    *int   `json:"ProductionOrder"`
	ProductionOrderItem                *int   `json:"ProductionOrderItem"`
	PrecedingOrderID                   *int   `json:"PrecedingOrderID"`
	PrecedingOrderItem                 *int   `json:"PrecedingOrderItem"`
	Project                            *int   `json:"Project"`
	WBSElement                         *int   `json:"WBSElement"`
	HeaderOrderStatus				   string `json:"HeaderOrderStatus"`
	HeaderCompleteOrderIsDefined	   *bool  `json:"HeaderCompleteOrderIsDefined"`
	CreationDate                       string `json:"CreationDate"`
	CreationTime                       string `json:"CreationTime"`
	LastChangeDate                     string `json:"LastChangeDate"`
	LastChangeTime                     string `json:"LastChangeTime"`
	IsReleased                         *bool  `json:"IsReleased"`
	IsCancelled                        *bool  `json:"IsCancelled"`
	IsMarkedForDeletion                *bool  `json:"IsMarkedForDeletion"`
}

type Item struct {
	PurchaseRequisition                           int      `json:"PurchaseRequisition"`
	PurchaseRequisitionItem                       int      `json:"PurchaseRequisitionItem"`
	PurchaseRequisitionItemCategory               string   `json:"PurchaseRequisitionItemCategory"`
	SupplyChainRelationshipID                     *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID             *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID        *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
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
	RequestedQuantityInBaseUnit                   float32  `json:"RequestedQuantityInBaseUnit"`
	RequestedQuantityInDeliveryUnit               float32  `json:"RequestedQuantityInDeliveryUnit"`
	BaseUnit                                      string   `json:"BaseUnit"`
	DeliveryUnit                                  string   `json:"DeliveryUnit"`
	RequestedDeliveryDate                         string   `json:"RequestedDeliveryDate"`
	PlannedOrder                                  *int     `json:"PlannedOrder"`
	PlannedOrderItem                              *int     `json:"PlannedOrderItem"`
	ProductionOrder                    			  *int     `json:"ProductionOrder"`
	ProductionOrderItem                			  *int     `json:"ProductionOrderItem"`
	PrecedingOrderID                              *int     `json:"PrecedingOrderID"`
	PrecedingOrderItem                            *int     `json:"PrecedingOrderItem"`
	FollowingOrderID                              *int     `json:"FollowingOrderID"`
	FollowingOrderItem                            *int     `json:"FollowingOrderItem"`
	Project                                       *int     `json:"Project"`
	WBSElement                                    *int     `json:"WBSElement"`
	PurchaseRequisitionItemText                   *string  `json:"PurchaseRequisitionItemText"`
	PurchaseRequisitionItemTextByBuyer            *string  `json:"PurchaseRequisitionItemTextByBuyer"`
	PurchaseRequisitionItemTextBySeller           *string  `json:"PurchaseRequisitionItemTextBySeller"`
	PurchaseRequisitionItemPrice                  *float32 `json:"PurchaseRequisitionItemPrice"`
	PurchaseRequisitionItemPriceQuantity          *int     `json:"PurchaseRequisitionItemPriceQuantity"`
	ProductPlannedDeliveryDuration                *float32 `json:"ProductPlannedDeliveryDuration"`
	ProductPlannedDeliveryDurationUnit            *int     `json:"ProductPlannedDeliveryDurationUnit"`
	OrderedQuantityInBaseUnit                     *float32 `json:"OrderedQuantityInBaseUnit"`
	OrderedQuantityInDeliveryUnit                 *float32 `json:"OrderedQuantityInDeliveryUnit"`
	ItemCompleteOrderIsDefined             		  *bool    `json:"ItemCompleteOrderIsDefined"`
	TransactionCurrency               			  *string  `json:"TransactionCurrency"`
	MRPArea                                       *string  `json:"MRPArea"`
	MRPController                                 *string  `json:"MRPController"`
	StockConfirmationBusinessPartner              *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                        *string  `json:"StockConfirmationPlant"`
	ProductionPlantBusinessPartner                *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                               *string  `json:"ProductionPlant"`
	GLAccount                                     *string  `json:"GLAccount"`
	ItemBlockStatus                               *bool    `json:"ItemBlockStatus"`
	CreationDate                       			  string   `json:"CreationDate"`
	CreationTime                       			  string   `json:"CreationTime"`
	LastChangeDate                     			  string   `json:"LastChangeDate"`
	LastChangeTime                     			  string   `json:"LastChangeTime"`
	IsReleased                         			  *bool    `json:"IsReleased"`
	IsCancelled                                   *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                           *bool    `json:"IsMarkedForDeletion"`
}
