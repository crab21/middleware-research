package cbn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateCenBandwidthPackage invokes the cbn.CreateCenBandwidthPackage API synchronously
// api document: https://help.aliyun.com/api/cbn/createcenbandwidthpackage.html
func (client *Client) CreateCenBandwidthPackage(request *CreateCenBandwidthPackageRequest) (response *CreateCenBandwidthPackageResponse, err error) {
	response = CreateCreateCenBandwidthPackageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCenBandwidthPackageWithChan invokes the cbn.CreateCenBandwidthPackage API asynchronously
// api document: https://help.aliyun.com/api/cbn/createcenbandwidthpackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCenBandwidthPackageWithChan(request *CreateCenBandwidthPackageRequest) (<-chan *CreateCenBandwidthPackageResponse, <-chan error) {
	responseChan := make(chan *CreateCenBandwidthPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCenBandwidthPackage(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateCenBandwidthPackageWithCallback invokes the cbn.CreateCenBandwidthPackage API asynchronously
// api document: https://help.aliyun.com/api/cbn/createcenbandwidthpackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateCenBandwidthPackageWithCallback(request *CreateCenBandwidthPackageRequest, callback func(response *CreateCenBandwidthPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCenBandwidthPackageResponse
		var err error
		defer close(result)
		response, err = client.CreateCenBandwidthPackage(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateCenBandwidthPackageRequest is the request struct for api CreateCenBandwidthPackage
type CreateCenBandwidthPackageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId            requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                string           `position:"Query" name:"ClientToken"`
	Description                string           `position:"Query" name:"Description"`
	BandwidthPackageChargeType string           `position:"Query" name:"BandwidthPackageChargeType"`
	GeographicRegionBId        string           `position:"Query" name:"GeographicRegionBId"`
	Period                     requests.Integer `position:"Query" name:"Period"`
	GeographicRegionAId        string           `position:"Query" name:"GeographicRegionAId"`
	AutoPay                    requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount       string           `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth                  requests.Integer `position:"Query" name:"Bandwidth"`
	OwnerAccount               string           `position:"Query" name:"OwnerAccount"`
	OwnerId                    requests.Integer `position:"Query" name:"OwnerId"`
	Name                       string           `position:"Query" name:"Name"`
	PricingCycle               string           `position:"Query" name:"PricingCycle"`
}

// CreateCenBandwidthPackageResponse is the response struct for api CreateCenBandwidthPackage
type CreateCenBandwidthPackageResponse struct {
	*responses.BaseResponse
	RequestId                  string `json:"RequestId" xml:"RequestId"`
	CenBandwidthPackageId      string `json:"CenBandwidthPackageId" xml:"CenBandwidthPackageId"`
	CenBandwidthPackageOrderId string `json:"CenBandwidthPackageOrderId" xml:"CenBandwidthPackageOrderId"`
}

// CreateCreateCenBandwidthPackageRequest creates a request to invoke CreateCenBandwidthPackage API
func CreateCreateCenBandwidthPackageRequest() (request *CreateCenBandwidthPackageRequest) {
	request = &CreateCenBandwidthPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateCenBandwidthPackage", "Cbn", "openAPI")
	return
}

// CreateCreateCenBandwidthPackageResponse creates a response to parse from CreateCenBandwidthPackage response
func CreateCreateCenBandwidthPackageResponse() (response *CreateCenBandwidthPackageResponse) {
	response = &CreateCenBandwidthPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}