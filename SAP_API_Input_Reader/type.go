package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey                        string `json:"connection_key"`
	Result                               bool   `json:"result"`
	RedisKey                             string `json:"redis_key"`
	Filepath                             string `json:"filepath"`
	FunctionalLocationEquipmentHierarchy struct {
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
		FunctionalLocation            struct {
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
		} `json:"FunctionalLocation"`
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
	} `json:"FunctionalLocationEquipmentHierarchy"`
	APISchema              string   `json:"api_schema"`
	Accepter               []string `json:"accepter"`
	FunctionalLocationCode string   `json:"functional_location_code"`
	Deleted                bool     `json:"deleted"`
}
