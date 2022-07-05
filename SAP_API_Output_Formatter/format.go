package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-functional-location-hierarchy-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToFuncnlLocEquipStrucList(raw []byte, l *logger.Logger) ([]FuncnlLocEquipStrucList, error) {
	pm := &responses.FuncnlLocEquipStrucList{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	funcnlLocEquipStrucList := make([]FuncnlLocEquipStrucList, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		funcnlLocEquipStrucList = append(funcnlLocEquipStrucList, FuncnlLocEquipStrucList{
			Equipment:                     data.Equipment,
			ValidityEndDate:               data.ValidityEndDate,
			ValidityStartDate:             data.ValidityStartDate,
			EquipmentName:                 data.EquipmentName,
			EquipmentCategory:             data.EquipmentCategory,
			EquipmentCategoryDesc:         data.EquipmentCategoryDesc,
			TechnicalObjectType:           data.TechnicalObjectType,
			TechnicalObjectTypeDesc:       data.TechnicalObjectTypeDesc,
			SuperordinateEquipment:        data.SuperordinateEquipment,
			SuperordinateEquipmentName:    data.SuperordinateEquipmentName,
			EquipInstallationPositionNmbr: data.EquipInstallationPositionNmbr,
			ConstructionMaterial:          data.ConstructionMaterial,
			FunctionalLocation:            data.FunctionalLocation,
			FunctionalLocationLabelName:   data.FunctionalLocationLabelName,
			EquipmentIsAtCustomer:         data.EquipmentIsAtCustomer,
			EquipmentIsAvailable:          data.EquipmentIsAvailable,
			EquipmentIsInWarehouse:        data.EquipmentIsInWarehouse,
			EquipmentIsAssignedToDelivery: data.EquipmentIsAssignedToDelivery,
			EquipmentIsMarkedForDeletion:  data.EquipmentIsMarkedForDeletion,
			EquipmentIsInstalled:          data.EquipmentIsInstalled,
			EquipIsAllocToSuperiorEquip:   data.EquipIsAllocToSuperiorEquip,
			EquipmentIsInactive:           data.EquipmentIsInactive,
			HierarchyNode:                 data.HierarchyNode,
			HierarchyParentNode:           data.HierarchyParentNode,
			TechObjHierarchyLevel:         data.TechObjHierarchyLevel,
		})

	}

	return funcnlLocEquipStrucList, nil
}
func ConvertToFunctionalLocationStrucList(raw []byte, l *logger.Logger) ([]FunctionalLocationStrucList, error) {
	pm := &responses.FunctionalLocationStrucList{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.Value) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.Value) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.Value))
	}

	functionalLocationStrucList := make([]FunctionalLocationStrucList, 0, 10)
	for i := 0; i < 10 && i < len(pm.Value); i++ {
		data := pm.Value[i]
		functionalLocationStrucList = append(functionalLocationStrucList, FunctionalLocationStrucList{
			FunctionalLocation:             data.FunctionalLocation,
			ValidityStartDate:              data.ValidityStartDate,
			FunctionalLocationName:         data.FunctionalLocationName,
			FunctionalLocationLabelName:    data.FunctionalLocationLabelName,
			SuperiorFunctionalLocation:     data.SuperiorFunctionalLocation,
			SuperiorFunctionalLocationName: data.SuperiorFunctionalLocationName,
			SuperiorFuncnlLocLabelName:     data.SuperiorFuncnlLocLabelName,
			FunctionalLocationCategory:     data.FunctionalLocationCategory,
			FunctionalLocationCategoryDesc: data.FunctionalLocationCategoryDesc,
			TechnicalObjectType:            data.TechnicalObjectType,
			TechnicalObjectTypeDesc:        data.TechnicalObjectTypeDesc,
			FuncnlLocPosInSuperiorTechObj:  data.FuncnlLocPosInSuperiorTechObj,
			ConstructionMaterial:           data.ConstructionMaterial,
			FuncnlLocIsMarkedForDeletion:   data.FuncnlLocIsMarkedForDeletion,
			FuncnlLocIsDeleted:             data.FuncnlLocIsDeleted,
			FunctionalLocationIsActive:     data.FunctionalLocationIsActive,
			FuncnlLocIsDeactivated:         data.FuncnlLocIsDeactivated,
			HierarchyNode:                  data.HierarchyNode,
			HierarchyParentNode:            data.HierarchyParentNode,
			TechObjHierarchyLevel:          data.TechObjHierarchyLevel,
		})

	}

	return functionalLocationStrucList, nil
}
