package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-sales-order-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		SalesOrder:                     data.SalesOrder,
		SalesOrderType:                 data.SalesOrderType,
		SalesOrganization:              data.SalesOrganization,
		DistributionChannel:            data.DistributionChannel,
		OrganizationDivision:           data.OrganizationDivision,
		SalesGroup:                     data.SalesGroup,
		SalesOffice:                    data.SalesOffice,
		SalesDistrict:                  data.SalesDistrict,
		SoldToParty:                    data.SoldToParty,
		CreationDate:                   data.CreationDate,
		LastChangeDate:                 data.LastChangeDate,
		ExternalDocumentID:             data.ExternalDocumentID,
		LastChangeDateTime:             data.LastChangeDateTime,
		PurchaseOrderByCustomer:        data.PurchaseOrderByCustomer,
		CustomerPurchaseOrderDate:      data.CustomerPurchaseOrderDate,
		SalesOrderDate:                 data.SalesOrderDate,
		TotalNetAmount:                 data.TotalNetAmount,
		OverallDeliveryStatus:          data.OverallDeliveryStatus,
		TotalBlockStatus:               data.TotalBlockStatus,
		OverallOrdReltdBillgStatus:     data.OverallOrdReltdBillgStatus,
		OverallSDDocReferenceStatus:    data.OverallSDDocReferenceStatus,
		TransactionCurrency:            data.TransactionCurrency,
		SDDocumentReason:               data.SDDocumentReason,
		PricingDate:                    data.PricingDate,
		PriceDetnExchangeRate:          data.PriceDetnExchangeRate,
		RequestedDeliveryDate:          data.RequestedDeliveryDate,
		ShippingCondition:              data.ShippingCondition,
		CompleteDeliveryIsDefined:      data.CompleteDeliveryIsDefined,
		ShippingType:                   data.ShippingType,
		HeaderBillingBlockReason:       data.HeaderBillingBlockReason,
		DeliveryBlockReason:            data.DeliveryBlockReason,
		IncotermsClassification:        data.IncotermsClassification,
		CustomerPriceGroup:             data.CustomerPriceGroup,
		PriceListType:                  data.PriceListType,
		CustomerPaymentTerms:           data.CustomerPaymentTerms,
		PaymentMethod:                  data.PaymentMethod,
		ReferenceSDDocument:            data.ReferenceSDDocument,
		ReferenceSDDocumentCategory:    data.ReferenceSDDocumentCategory,
		CustomerAccountAssignmentGroup: data.CustomerAccountAssignmentGroup,
		AccountingExchangeRate:         data.AccountingExchangeRate,
		CustomerGroup:                  data.CustomerGroup,
		AdditionalCustomerGroup1:       data.AdditionalCustomerGroup1,
		AdditionalCustomerGroup2:       data.AdditionalCustomerGroup2,
		AdditionalCustomerGroup3:       data.AdditionalCustomerGroup3,
		AdditionalCustomerGroup4:       data.AdditionalCustomerGroup4,
		AdditionalCustomerGroup5:       data.AdditionalCustomerGroup5,
		CustomerTaxClassification1:     data.CustomerTaxClassification1,
		TotalCreditCheckStatus:         data.TotalCreditCheckStatus,
		BillingDocumentDate:            data.BillingDocumentDate,
		ToHeaderPartner:                data.ToHeaderPartner.ToHeaderPartnerResults[0],
        ToItem:                         data.ToItem.ToItemResults[0],
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
		SalesOrder:                  data.SalesOrder,
		SalesOrderItem:              data.SalesOrderItem,
		SalesOrderItemCategory:      data.SalesOrderItemCategory,
		SalesOrderItemText:          data.SalesOrderItemText,
		PurchaseOrderByCustomer:     data.PurchaseOrderByCustomer,
		Material:                    data.Material,
		MaterialByCustomer:          data.MaterialByCustomer,
		PricingDate:                 data.PricingDate,
		BillingPlan:                 data.BillingPlan,
		RequestedQuantity:           data.RequestedQuantity,
		RequestedQuantityUnit:       data.RequestedQuantityUnit,
		OrderQuantityUnit:           data.OrderQuantityUnit,
		ConfdDelivQtyInOrderQtyUnit: data.ConfdDelivQtyInOrderQtyUnit,
		ItemGrossWeight:             data.ItemGrossWeight,
		ItemNetWeight:               data.ItemNetWeight,
		ItemWeightUnit:              data.ItemWeightUnit,
		ItemVolume:                  data.ItemVolume,
		ItemVolumeUnit:              data.ItemVolumeUnit,
		TransactionCurrency:         data.TransactionCurrency,
		NetAmount:                   data.NetAmount,
		MaterialGroup:               data.MaterialGroup,
		MaterialPricingGroup:        data.MaterialPricingGroup,
		BillingDocumentDate:         data.BillingDocumentDate,
		Batch:                       data.Batch,
		ProductionPlant:             data.ProductionPlant,
		StorageLocation:             data.StorageLocation,
		DeliveryGroup:               data.DeliveryGroup,
		ShippingPoint:               data.ShippingPoint,
		ShippingType:                data.ShippingType,
		DeliveryPriority:            data.DeliveryPriority,
		IncotermsClassification:     data.IncotermsClassification,
		TaxAmount:                   data.TaxAmount,
		ProductTaxClassification1:   data.ProductTaxClassification1,
		MatlAccountAssignmentGroup:  data.MatlAccountAssignmentGroup,
		CostAmount:                  data.CostAmount,
		CustomerPaymentTerms:        data.CustomerPaymentTerms,
		CustomerGroup:               data.CustomerGroup,
		SalesDocumentRjcnReason:     data.SalesDocumentRjcnReason,
		ItemBillingBlockReason:      data.ItemBillingBlockReason,
		WBSElement:                  data.WBSElement,
		ProfitCenter:                data.ProfitCenter,
		AccountingExchangeRate:      data.AccountingExchangeRate,
		ReferenceSDDocument:         data.ReferenceSDDocument,
		ReferenceSDDocumentItem:     data.ReferenceSDDocumentItem,
		SDProcessStatus:             data.SDProcessStatus,
		DeliveryStatus:              data.DeliveryStatus,
		OrderRelatedBillingStatus:   data.OrderRelatedBillingStatus,
		ToItemPricingElement:        data.ToItemPricingElement.ToItemPricingElementResults[0],
        ToItemScheduleLine:          data.ToItemScheduleLine.ToItemScheduleLineResults[0],
	}

	return item, nil
}

func ConvertToHeaderPartner(raw []byte, l *logger.Logger) (*HeaderPartner, error) {
	pm := &responses.HeaderPartner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to HeaderPartner. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	headerPartner := &HeaderPartner{
		SalesOrder:      data.SalesOrder,
		PartnerFunction: data.PartnerFunction,
		Customer:        data.Customer,
		Supplier:        data.Supplier,
	}

	return headerPartner, nil
}

func ConvertToItemPricingElement(raw []byte, l *logger.Logger) (*ItemPricingElement, error) {
	pm := &responses.ItemPricingElement{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemPricingElement. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	itemPricingElement := &ItemPricingElement{
		SalesOrder:                     data.SalesOrder,
		SalesOrderItem:                 data.SalesOrderItem,
		PricingProcedureStep:           data.PricingProcedureStep,
		PricingProcedureCounter:        data.PricingProcedureCounter,
		ConditionType:                  data.ConditionType,
		PriceConditionDeterminationDte: data.PriceConditionDeterminationDte,
		ConditionCalculationType:       data.ConditionCalculationType,
		ConditionBaseValue:             data.ConditionBaseValue,
		ConditionRateValue:             data.ConditionRateValue,
		ConditionCurrency:              data.ConditionCurrency,
		ConditionQuantity:              data.ConditionQuantity,
		ConditionQuantityUnit:          data.ConditionQuantityUnit,
		ConditionCategory:              data.ConditionCategory,
		PricingScaleType:               data.PricingScaleType,
		ConditionRecord:                data.ConditionRecord,
		ConditionSequentialNumber:      data.ConditionSequentialNumber,
		TaxCode:                        data.TaxCode,
		ConditionAmount:                data.ConditionAmount,
		TransactionCurrency:            data.TransactionCurrency,
		PricingScaleBasis:              data.PricingScaleBasis,
		ConditionScaleBasisValue:       data.ConditionScaleBasisValue,
		ConditionScaleBasisUnit:        data.ConditionScaleBasisUnit,
		ConditionScaleBasisCurrency:    data.ConditionScaleBasisCurrency,
		ConditionIsManuallyChanged:     data.ConditionIsManuallyChanged,
	}

	return itemPricingElement, nil
}

func ConvertToItemScheduleLine(raw []byte, l *logger.Logger) (*ItemScheduleLine, error) {
	pm := &responses.ItemScheduleLine{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemScheduleLine. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	itemScheduleLine := &ItemScheduleLine{
		SalesOrder:                    data.SalesOrder,
		SalesOrderItem:                data.SalesOrderItem,
		ScheduleLine:                  data.ScheduleLine,
		RequestedDeliveryDate:         data.RequestedDeliveryDate,
		ConfirmedDeliveryDate:         data.ConfirmedDeliveryDate,
		OrderQuantityUnit:             data.OrderQuantityUnit,
		ScheduleLineOrderQuantity:     data.ScheduleLineOrderQuantity,
		ConfdOrderQtyByMatlAvailCheck: data.ConfdOrderQtyByMatlAvailCheck,
		DeliveredQtyInOrderQtyUnit:    data.DeliveredQtyInOrderQtyUnit,
		OpenConfdDelivQtyInOrdQtyUnit: data.OpenConfdDelivQtyInOrdQtyUnit,
		CorrectedQtyInOrderQtyUnit:    data.CorrectedQtyInOrderQtyUnit,
		DelivBlockReasonForSchedLine:  data.DelivBlockReasonForSchedLine,
	}

	return itemScheduleLine, nil
}
