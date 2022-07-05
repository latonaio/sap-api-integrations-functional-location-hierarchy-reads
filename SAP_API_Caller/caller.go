package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-functional-location-hierarchy-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetFunctionalLocationHierarcy(equipment, functionalLocation string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "FuncnlLocEquipStrucList":
			func() {
				c.FuncnlLocEquipStrucList(equipment)
				wg.Done()
			}()
		case "FunctionalLocationStrucList":
			func() {
				c.FunctionalLocationStrucList(functionalLocation)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) FuncnlLocEquipStrucList(equipment string) {
	data, err := c.callFunctionalLocationHierarcySrvAPIRequirementFuncnlLocEquipStrucList("FuncnlLocEquipStrucList", equipment)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callFunctionalLocationHierarcySrvAPIRequirementFuncnlLocEquipStrucList(api, equipment string) ([]sap_api_output_formatter.FuncnlLocEquipStrucList, error) {
	url := strings.Join([]string{c.baseURL, "api_funcnlloc_struclist/srvd_a2x/sap/funcnllocstruclist/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setFuncnlLocEquipStrucListAPIKeyAccept(req)
	c.getQueryWithFuncnlLocEquipStrucList(req, equipment)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToFuncnlLocEquipStrucList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) FunctionalLocationStrucList(functionalLocation string) {
	data, err := c.callFunctionalLocationHierarcySrvAPIRequirementFunctionalLocationStrucList("FunctionalLocationStrucList", functionalLocation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)

}

func (c *SAPAPICaller) callFunctionalLocationHierarcySrvAPIRequirementFunctionalLocationStrucList(api, functionalLocation string) ([]sap_api_output_formatter.FuncnlLocEquipStrucList, error) {
	url := strings.Join([]string{c.baseURL, "api_funcnlloc_struclist/srvd_a2x/sap/funcnllocstruclist/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setFuncnlLocEquipStrucListAPIKeyAccept(req)
	c.getQueryWithFunctionalLocationStrucList(req, functionalLocation)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToFuncnlLocEquipStrucList(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setFuncnlLocEquipStrucListAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithFuncnlLocEquipStrucList(req *http.Request, equipment string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("Equipment eq '%s'", equipment))
	req.URL.RawQuery = params.Encode()
}

func (c *SAPAPICaller) getQueryWithFunctionalLocationStrucList(req *http.Request, functionalLocation string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("FunctionalLocation eq '%s'", functionalLocation))
	req.URL.RawQuery = params.Encode()
}
