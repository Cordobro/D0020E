package Orchestrator

import (
	forms "arrowhead/Orchestrator/forms"
	"encoding/json"
	"fmt"
)

type request struct {
	data           ServiceData
	serviceRegAdrs string
	authAdrs       string
}

func newRequest(data ServiceData, serviceRegAdrs string, authAdrs string) *request {
	r := request{data: data,
		serviceRegAdrs: serviceRegAdrs,
		authAdrs:       authAdrs}
	return &r
}

func sendServiceRequest(r *request) {

	serviceRegClient := newClient(r.serviceRegAdrs)
	serviceQueryListJson := exchangeJson(serviceRegClient, r.data.Discover.ServiceQueryForm)
	var serviceQueryList forms.ServiceQueryList
	err := json.Unmarshal(serviceQueryListJson, &serviceQueryList)
	errorHandler(err)
	r.data.Discover.ServiceQueryList = serviceQueryList
}

/*
func sendAuthQuery(r *request) {
	authClient := NewClient(r.authAdrs)
	intraCloudResultJson := ExchangeJson(authClient, r.data.IntraCloudRule)
	var intraCloudResult forms.IntraCloudResult
	err := json.Unmarshal(intraCloudResultJson, &intraCloudResult)
	errorHandler(err)
	r.data.IntraCloudResult = intraCloudResult
}


func sendTokenQuery(r *request) {
	authClient := NewClient(r.authAdrs)
	tokenResultJson := ExchangeJson(authClient, r.data.TokenRule)
	var tokenResult forms.TokenResult
	err := json.Unmarshal(tokenResultJson, &tokenResult)
	errorHandler(err)
	r.data.TokenResult = tokenResult
}
*/
func matchmaker(r *request) {

	var sql = r.data.Discover.ServiceQueryList

	matchmade := new(forms.Discover)

	// Checks if the serviceRegistry only returned 1 (or less) service

	if len(sql.ServiceQueryData) < 2 {

		matchmade.ServiceQueryList = sql

	} else {

		var requirements = r.data.ServiceRequestForm.RequestedService.MetadataRequirements
		var templist [][]string
		var bestIndex int

		for i := 0; i < len(sql.ServiceQueryData); i++ {
			fmt.Println("METADATA: ", i)
			fmt.Println(sql.ServiceQueryData[i].Metadata)
			fmt.Println("")
		}

		//Run intersection on all indexes and save result in tempList

		for i := 0; i < len(sql.ServiceQueryData); i++ {

			temp := intersection(sql.ServiceQueryData[i].Metadata, requirements)
			templist = append(templist, temp)

		}

		//Find the best index (most matches)

		for i := 0; i < len(templist); i++ {
			if len(templist[i]) > bestIndex {
				bestIndex = i
			}
		}

		fmt.Println("BEST INDEX: ", bestIndex)

		temp := sql.ServiceQueryData[bestIndex]

		fmt.Println("TEMP :")
		fmt.Println(temp)

		matchmade.ServiceQueryList.ServiceQueryData = append(matchmade.ServiceQueryList.ServiceQueryData, temp)

	}

	forms.ConstructOrchestrationResponse(&matchmade.ServiceQueryList, &r.data.OrchestrationResponse)
}

//https://stackoverflow.com/questions/44956031/how-to-get-intersection-of-two-slice-in-golang

func intersection(s1, s2 []string) (inter []string) {
	hash := make(map[string]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

//Remove dups from slice.
func removeDups(elements []string) (nodups []string) {
	encountered := make(map[string]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}
