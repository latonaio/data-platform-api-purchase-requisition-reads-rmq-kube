package requests

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
