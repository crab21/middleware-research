package baas

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

// ConfirmFabricConsortiumMember invokes the baas.ConfirmFabricConsortiumMember API synchronously
// api document: https://help.aliyun.com/api/baas/confirmfabricconsortiummember.html
func (client *Client) ConfirmFabricConsortiumMember(request *ConfirmFabricConsortiumMemberRequest) (response *ConfirmFabricConsortiumMemberResponse, err error) {
	response = CreateConfirmFabricConsortiumMemberResponse()
	err = client.DoAction(request, response)
	return
}

// ConfirmFabricConsortiumMemberWithChan invokes the baas.ConfirmFabricConsortiumMember API asynchronously
// api document: https://help.aliyun.com/api/baas/confirmfabricconsortiummember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfirmFabricConsortiumMemberWithChan(request *ConfirmFabricConsortiumMemberRequest) (<-chan *ConfirmFabricConsortiumMemberResponse, <-chan error) {
	responseChan := make(chan *ConfirmFabricConsortiumMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfirmFabricConsortiumMember(request)
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

// ConfirmFabricConsortiumMemberWithCallback invokes the baas.ConfirmFabricConsortiumMember API asynchronously
// api document: https://help.aliyun.com/api/baas/confirmfabricconsortiummember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ConfirmFabricConsortiumMemberWithCallback(request *ConfirmFabricConsortiumMemberRequest, callback func(response *ConfirmFabricConsortiumMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfirmFabricConsortiumMemberResponse
		var err error
		defer close(result)
		response, err = client.ConfirmFabricConsortiumMember(request)
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

// ConfirmFabricConsortiumMemberRequest is the request struct for api ConfirmFabricConsortiumMember
type ConfirmFabricConsortiumMemberRequest struct {
	*requests.RpcRequest
	Organization *[]ConfirmFabricConsortiumMemberOrganization `position:"Query" name:"Organization"  type:"Repeated"`
	ConsortiumId string                                       `position:"Query" name:"ConsortiumId"`
}

// ConfirmFabricConsortiumMemberOrganization is a repeated param struct in ConfirmFabricConsortiumMemberRequest
type ConfirmFabricConsortiumMemberOrganization struct {
	OrganizationId string `name:"OrganizationId"`
}

// ConfirmFabricConsortiumMemberResponse is the response struct for api ConfirmFabricConsortiumMember
type ConfirmFabricConsortiumMemberResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateConfirmFabricConsortiumMemberRequest creates a request to invoke ConfirmFabricConsortiumMember API
func CreateConfirmFabricConsortiumMemberRequest() (request *ConfirmFabricConsortiumMemberRequest) {
	request = &ConfirmFabricConsortiumMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "ConfirmFabricConsortiumMember", "baas", "openAPI")
	return
}

// CreateConfirmFabricConsortiumMemberResponse creates a response to parse from ConfirmFabricConsortiumMember response
func CreateConfirmFabricConsortiumMemberResponse() (response *ConfirmFabricConsortiumMemberResponse) {
	response = &ConfirmFabricConsortiumMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
