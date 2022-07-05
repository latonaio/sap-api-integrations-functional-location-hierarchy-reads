package responses

type FunctionalLocationStrucList struct {
	Value []struct {
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
	} `json:"value"`
}
