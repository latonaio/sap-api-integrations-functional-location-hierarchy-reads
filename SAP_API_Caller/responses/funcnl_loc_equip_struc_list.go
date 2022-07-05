package responses

type FuncnlLocEquipStrucList struct {
	Value []struct {
		Equipment                     string `json:"Equipment"`
		ValidityEndDate               string `json:"ValidityEndDate"`
		ValidityStartDate             string `json:"ValidityStartDate"`
		EquipmentName                 string `json:"EquipmentName"`
		EquipmentCategory             string `json:"EquipmentCategory"`
		EquipmentCategoryDesc         string `json:"EquipmentCategoryDesc"`
		TechnicalObjectType           string `json:"TechnicalObjectType"`
		TechnicalObjectTypeDesc       string `json:"TechnicalObjectTypeDesc"`
		SuperordinateEquipment        string `json:"SuperordinateEquipment"`
		SuperordinateEquipmentName    string `json:"SuperordinateEquipmentName"`
		EquipInstallationPositionNmbr string `json:"EquipInstallationPositionNmbr"`
		ConstructionMaterial          string `json:"ConstructionMaterial"`
		FunctionalLocation            string `json:"FunctionalLocation"`
		FunctionalLocationLabelName   string `json:"FunctionalLocationLabelName"`
		EquipmentIsAtCustomer         bool   `json:"EquipmentIsAtCustomer"`
		EquipmentIsAvailable          bool   `json:"EquipmentIsAvailable"`
		EquipmentIsInWarehouse        bool   `json:"EquipmentIsInWarehouse"`
		EquipmentIsAssignedToDelivery bool   `json:"EquipmentIsAssignedToDelivery"`
		EquipmentIsMarkedForDeletion  bool   `json:"EquipmentIsMarkedForDeletion"`
		EquipmentIsInstalled          bool   `json:"EquipmentIsInstalled"`
		EquipIsAllocToSuperiorEquip   bool   `json:"EquipIsAllocToSuperiorEquip"`
		EquipmentIsInactive           bool   `json:"EquipmentIsInactive"`
		HierarchyNode                 string `json:"HierarchyNode"`
		HierarchyParentNode           string `json:"HierarchyParentNode"`
		TechObjHierarchyLevel         int    `json:"TechObjHierarchyLevel"`
	} `json:"value"`
}
