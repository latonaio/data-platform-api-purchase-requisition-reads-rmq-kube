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
			&pm.Buyer,
			&pm.PurchaseRequisitionType,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.PrecedingOrderID,
			&pm.PrecedingOrderItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		header = append(header, Header{
			PurchaseRequisition:		data.PurchaseRequisition,
			Buyer:						data.Buyer,
			PurchaseRequisitionType:	data.PurchaseRequisitionType,
			PlannedOrder:				data.PlannedOrder,
			PlannedOrderItem:			data.PlannedOrderItem,
			PrecedingOrderID:			data.PrecedingOrderID,
			PrecedingOrderItem:			data.PrecedingOrderItem,
			Project:					data.Project,
			WBSElement:					data.WBSElement,
			CreationDate:				data.CreationDate,
			LastChangeDate:				data.LastChangeDate,
			IsCancelled:				data.IsCancelled,
			IsMarkedForDeletion:		data.IsMarkedForDeletion,
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
			&pm.RequestedQuantity,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.PrecedingOrderID,
			&pm.PrecedingOrderItem,
			&pm.FollowingOrderID,
			&pm.FollowingOrderItem,
			&pm.Project,
			&pm.WBSElement,
			&pm.PurchaseRequisitionReleaseStatus,
			&pm.PurchaseRequisitionReleaseDate,
			&pm.PurchaseRequisitionItemText,
			&pm.PurchaseRequisitionItemTextByBuyer,
			&pm.PurchaseRequisitionItemTextBySeller,
			&pm.PurchaseRequisitionPrice,
			&pm.PurchaseRequisitionPriceQuantity,
			&pm.ProductPlannedDeliveryDuration,
			&pm.ProductPlannedDeliveryDurationUnit,
			&pm.OrderedQuantity,
			&pm.DeliveryDate,
			&pm.PurchaseRequisitionProcessingStatus,
			&pm.PurchaseRequisitionItemCurrency,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.GLAccount,
			&pm.ItemBlockStatus,
			&pm.CreationDate,
			&pm.LastChangeDate,
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
			PrecedingOrderID:							data.PrecedingOrderID,
			PrecedingOrderItem:							data.PrecedingOrderItem,
			FollowingOrderID:							data.FollowingOrderID,
			FollowingOrderItem:							data.FollowingOrderItem,
			Project:									data.Project,
			WBSElement:									data.WBSElement,
			PurchaseRequisitionReleaseStatus:			data.PurchaseRequisitionReleaseStatus,
			PurchaseRequisitionReleaseDate:				data.PurchaseRequisitionReleaseDate,
			PurchaseRequisitionItemText:				data.PurchaseRequisitionItemText,
			PurchaseRequisitionItemTextByBuyer:			data.PurchaseRequisitionItemTextByBuyer,
			PurchaseRequisitionItemTextBySeller:		data.PurchaseRequisitionItemTextBySeller,
			PurchaseRequisitionPrice:					data.PurchaseRequisitionPrice,
			PurchaseRequisitionPriceQuantity:			data.PurchaseRequisitionPriceQuantity,
			ProductPlannedDeliveryDuration:				data.ProductPlannedDeliveryDuration,
			ProductPlannedDeliveryDurationUnit:			data.ProductPlannedDeliveryDurationUnit,
			OrderedQuantity:							data.OrderedQuantity,
			DeliveryDate:								data.DeliveryDate,
			PurchaseRequisitionProcessingStatus:		data.PurchaseRequisitionProcessingStatus,
			PurchaseRequisitionItemCurrency:			data.PurchaseRequisitionItemCurrency,
			MRPArea:									data.MRPArea,
			MRPController:								data.MRPController,
			StockConfirmationBusinessPartner:			data.StockConfirmationBusinessPartner,
			StockConfirmationPlant:						data.StockConfirmationPlant,
			ProductionPlantBusinessPartner:				data.ProductionPlantBusinessPartner,
			ProductionPlant:							data.ProductionPlant,
			GLAccount:									data.GLAccount,
			ItemBlockStatus:							data.ItemBlockStatus,
			CreationDate:								data.CreationDate,
			LastChangeDate:								data.LastChangeDate,
			IsMarkedForDeletion:						data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &item, nil
	}

	return &item, nil
}
