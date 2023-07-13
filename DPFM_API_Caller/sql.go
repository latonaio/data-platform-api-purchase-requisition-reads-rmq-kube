package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-purchase-requisition-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-purchase-requisition-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,


	) interface{} {
		var header *[]dpfm_api_output_formatter.Header
		var item *[]dpfm_api_output_formatter.Item
//		var partner *[]dpfm_api_output_formatter.Partner
//		var address *[]dpfm_api_output_formatter.Address
		for _, fn := range accepter {
			switch fn {
			case "Header":
				func() {
					header = c.Header(mtx, input, output, errs, log)
				}()
			case "Item":
				func() {
					item = c.Item(mtx, input, output, errs, log)
				}()
			case "Items":
				func() {
					item = c.Items(mtx, input, output, errs, log)
				}()
//			case "Partner":
//				func() {
//					partner = c.Partner(mtx, input, output, errs, log)
//				}()
//			case "Address":
//				func() {
//					address = c.Address(mtx, input, output, errs, log)
//				}()
			default:
			}
		}
	data := &dpfm_api_output_formatter.Message{
		Header:  header,
		Item:    item,
		Partner: partner,
		Address: address,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	if input.Header.PurchaseRequisition == nil | {
		err := xerrors.New("入力ファイルのPurchaseRequisitionがnullです。")
		*errs = append(*errs, err)
		return nil
	}
	purchaseRequisition := *input.Header.PurchaseRequisition

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_purchase_requisition_header_data
		WHERE (PurchaseRequisition) = (?);`, purchaseRequisition,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Item(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	var args []interface{}
	purchaseRequisition := input.Header.PurchaseRequisition
	item := input.Header.Item

	cnt := 0
	for _, v := range item {
		args = append(args, purchaseRequisition, v.PurchaseRequisitionItem)
		cnt++
	}

	repeat := strings.Repeat("(?,?),", cnt-1) + "(?,?)"
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_purchase_requisition_item_data
		WHERE (PurchaseRequisition, PurchaseRequisitionItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Items(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	item := &dpfm_api_input_reader.Item{}
	if len(input.Header.Item) > 0 {
		item = &input.Header.Item[0]
	}

	where := fmt.Sprintf("WHERE PurchaseRequisition = %v", input.Header.PurchaseRequisition)
	if item != nil {
		if item.IsMarkedForDeletion != nil {
			where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %v", where, *item.IsMarkedForDeletion)
		}
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_purchase_requisition_item_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToItem(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
