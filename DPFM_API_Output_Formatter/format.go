package dpfm_api_output_formatter

import (
	"data-platform-api-purchase-requisition-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.PurchaseRequisition,
			&pm.PurchaseRequisitionDate,
			&pm.Buyer,
			&pm.PurchaseRequisitionType,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.PrecedingOrderID,
			&pm.PrecedingOrderItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.HeaderOrderStatus,
			&pm.HeaderCompleteOrderIsDefined,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			PurchaseRequisition:			data.PurchaseRequisition,
			PurchaseRequisitionDate:		data.PurchaseRequisitionDate,
			Buyer:							data.Buyer,
			PurchaseRequisitionType:		data.PurchaseRequisitionType,
			PlannedOrder:					data.PlannedOrder,
			PlannedOrderItem:				data.PlannedOrderItem,
			ProductionOrder:				data.ProductionOrder,
			ProductionOrderItem:			data.ProductionOrderItem,
			PrecedingOrderID:				data.PrecedingOrderID,
			PrecedingOrderItem:				data.PrecedingOrderItem,
			Project:						data.Project,
			WBSElement:						data.WBSElement,
			HeaderOrderStatus:				data.HeaderOrderStatus,
			HeaderCompleteOrderIsDefined:	data.HeaderCompleteOrderIsDefined,
			CreationDate:					data.CreationDate,
			CreationTime:					data.CreationTime,
			LastChangeDate:					data.LastChangeDate,
			LastChangeTime:					data.LastChangeTime,
			IsReleased:						data.IsReleased,
			IsCancelled:					data.IsCancelled,
			IsMarkedForDeletion:			data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return nil, nil
	}

	return &header, nil
}

func ConvertToItem(rows *sql.Rows) (*[]Item, error) {
	defer rows.Close()
	item := make([]Item, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Item{}

		err := rows.Scan(
			&pm.PurchaseRequisition,
			&pm.PurchaseRequisitionItem,
			&pm.PurchaseRequisitionItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.SupplyChainRelationshipProductionPlantID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromPlant,
			&pm.DeliverFromPlantStorageLocation,
			&pm.Product,
			&pm.ProductGroup,
			&pm.RequestedQuantityInBaseUnit,
			&pm.RequestedQuantityInDeliveryUnit,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.ProductionOrder,
			&pm.ProductionOrderItem,
			&pm.PrecedingOrderID,
			&pm.PrecedingOrderItem,
			&pm.FollowingOrderID,
			&pm.FollowingOrderItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.PurchaseRequisitionItemText,
			&pm.PurchaseRequisitionItemTextByBuyer,
			&pm.PurchaseRequisitionItemTextBySeller,
			&pm.PurchaseRequisitionItemPrice,
			&pm.PurchaseRequisitionItemPriceQuantity,
			&pm.ProductPlannedDeliveryDuration,
			&pm.ProductPlannedDeliveryDurationUnit,
			&pm.OrderedQuantityInBaseUnit,
			&pm.OrderedQuantityInDeliveryUnit,
			&pm.DeliveryDate,
			&pm.ItemCompleteOrderIsDefined,
			&pm.TransactionCurrency,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.GLAccount,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &item, err
		}

		data := pm
		item = append(item, Item{
			PurchaseRequisition:						data.PurchaseRequisition,
			PurchaseRequisitionItem:					data.PurchaseRequisitionItem,
			PurchaseRequisitionItemCategory:			data.PurchaseRequisitionItemCategory,
			SupplyChainRelationshipID:					data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:			data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:		data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipStockConfPlantID:	data.SupplyChainRelationshipStockConfPlantID,
			SupplyChainRelationshipProductionPlantID:	data.SupplyChainRelationshipProductionPlantID,
			Buyer:										data.Buyer,
			Seller:										data.Seller,
			DeliverToParty:								data.DeliverToParty,
			DeliverFromParty:							data.DeliverFromParty,
			DeliverToPlant:								data.DeliverToPlant,
			DeliverToPlantStorageLocation:				data.DeliverToPlantStorageLocation,
			DeliverFromPlant:							data.DeliverFromPlant,
			DeliverFromPlantStorageLocation:			data.DeliverFromPlantStorageLocation,
			Product:									data.Product,
			ProductGroup:								data.ProductGroup,
			RequestedQuantity:							data.RequestedQuantity,
			BaseUnit:									data.BaseUnit,
			DeliveryUnit:								data.DeliveryUnit,
			PlannedOrder:								data.PlannedOrder,
			PlannedOrderItem:							data.PlannedOrderItem,
			ProductionOrder:							data.ProductionOrder,
			ProductionOrderItem:						data.ProductionOrderItem,
			PrecedingOrderID:							data.PrecedingOrderID,
			PrecedingOrderItem:							data.PrecedingOrderItem,
			FollowingOrderID:							data.FollowingOrderID,
			FollowingOrderItem:							data.FollowingOrderItem,
			Project:									data.Project,
			WBSElement:									data.WBSElement,
			PurchaseRequisitionItemText:				data.PurchaseRequisitionItemText,
			PurchaseRequisitionItemTextByBuyer:			data.PurchaseRequisitionItemTextByBuyer,
			PurchaseRequisitionItemTextBySeller:		data.PurchaseRequisitionItemTextBySeller,
			PurchaseRequisitionItemPrice:				data.PurchaseRequisitionItemPrice,
			PurchaseRequisitionItemPriceQuantity:		data.PurchaseRequisitionItemPriceQuantity,
			ProductPlannedDeliveryDuration:				data.ProductPlannedDeliveryDuration,
			ProductPlannedDeliveryDurationUnit:			data.ProductPlannedDeliveryDurationUnit,
			OrderedQuantityInBaseUnit:					data.OrderedQuantityInBaseUnit,
			OrderedQuantityInDeliveryUnit:				data.OrderedQuantityInDeliveryUnit,
			DeliveryDate:								data.DeliveryDate,
			ItemCompleteOrderIsDefined:					data.ItemCompleteOrderIsDefined,
			TransactionCurrency:						data.TransactionCurrency,
			MRPArea:									data.MRPArea,
			MRPController:								data.MRPController,
			StockConfirmationBusinessPartner:			data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:						data.StockConfirmationPlant,
			ProductionPlantBusinessPartner:				data.ProductionPlantBusinessPartner,
			ProductionPlant:							data.ProductionPlant,
			GLAccount:									data.GLAccount,
			CreationDate:								data.CreationDate,
			CreationTime:								data.CreationTime,
			LastChangeDate:								data.LastChangeDate,
			LastChangeTime:								data.LastChangeTime,
			IsReleased:									data.IsReleased,
			IsCancelled:								data.IsCancelled,
			IsMarkedForDeletion:						data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
