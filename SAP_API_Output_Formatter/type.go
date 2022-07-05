package sap_api_output_formatter

type FunctionalLocation struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	FunctionalLocation string `json:"functional_location"`
	Deleted            bool   `json:"deleted"`
}

type FuncnlLocEquipStrucList struct {
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
}
type FunctionalLocationStrucList struct {
	FunctionalLocation             string `json:"FunctionalLocation"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	FunctionalLocationName         string `json:"FunctionalLocationName"`
	FunctionalLocationLabelName    string `json:"FunctionalLocationLabelName"`
	SuperiorFunctionalLocation     string `json:"SuperiorFunctionalLocation"`
	SuperiorFunctionalLocationName string `json:"SuperiorFunctionalLocationName"`
	SuperiorFuncnlLocLabelName     string `json:"SuperiorFuncnlLocLabelName"`
	FunctionalLocationCategory     string `json:"FunctionalLocationCategory"`
	FunctionalLocationCategoryDesc string `json:"FunctionalLocationCategoryDesc"`
	TechnicalObjectType            string `json:"TechnicalObjectType"`
	TechnicalObjectTypeDesc        string `json:"TechnicalObjectTypeDesc"`
	FuncnlLocPosInSuperiorTechObj  string `json:"FuncnlLocPosInSuperiorTechObj"`
	ConstructionMaterial           string `json:"ConstructionMaterial"`
	FuncnlLocIsMarkedForDeletion   bool   `json:"FuncnlLocIsMarkedForDeletion"`
	FuncnlLocIsDeleted             bool   `json:"FuncnlLocIsDeleted"`
	FunctionalLocationIsActive     bool   `json:"FunctionalLocationIsActive"`
	FuncnlLocIsDeactivated         bool   `json:"FuncnlLocIsDeactivated"`
	HierarchyNode                  string `json:"HierarchyNode"`
	HierarchyParentNode            string `json:"HierarchyParentNode"`
	TechObjHierarchyLevel          int    `json:"TechObjHierarchyLevel"`
}
