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

// ResetFabricOrganizationUserPassword invokes the baas.ResetFabricOrganizationUserPassword API synchronously
// api document: https://help.aliyun.com/api/baas/resetfabricorganizationuserpassword.html
func (client *Client) ResetFabricOrganizationUserPassword(request *ResetFabricOrganizationUserPasswordRequest) (response *ResetFabricOrganizationUserPasswordResponse, err error) {
	response = CreateResetFabricOrganizationUserPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ResetFabricOrganizationUserPasswordWithChan invokes the baas.ResetFabricOrganizationUserPassword API asynchronously
// api document: https://help.aliyun.com/api/baas/resetfabricorganizationuserpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetFabricOrganizationUserPasswordWithChan(request *ResetFabricOrganizationUserPasswordRequest) (<-chan *ResetFabricOrganizationUserPasswordResponse, <-chan error) {
	responseChan := make(chan *ResetFabricOrganizationUserPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetFabricOrganizationUserPassword(request)
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

// ResetFabricOrganizationUserPasswordWithCallback invokes the baas.ResetFabricOrganizationUserPassword API asynchronously
// api document: https://help.aliyun.com/api/baas/resetfabricorganizationuserpassword.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResetFabricOrganizationUserPasswordWithCallback(request *ResetFabricOrganizationUserPasswordRequest, callback func(response *ResetFabricOrganizationUserPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetFabricOrganizationUserPasswordResponse
		var err error
		defer close(result)
		response, err = client.ResetFabricOrganizationUserPassword(request)
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

// ResetFabricOrganizationUserPasswordRequest is the request struct for api ResetFabricOrganizationUserPassword
type ResetFabricOrganizationUserPasswordRequest struct {
	*requests.RpcRequest
	Password       string `position:"Body" name:"Password"`
	OrganizationId string `position:"Body" name:"OrganizationId"`
	Location       string `position:"Body" name:"Location"`
	Username       string `position:"Body" name:"Username"`
}

// ResetFabricOrganizationUserPasswordResponse is the response struct for api ResetFabricOrganizationUserPassword
type ResetFabricOrganizationUserPasswordResponse struct {
	*responses.BaseResponse
	RequestId string                                      `json:"RequestId" xml:"RequestId"`
	Success   bool                                        `json:"Success" xml:"Success"`
	ErrorCode int                                         `json:"ErrorCode" xml:"ErrorCode"`
	Result    ResultInResetFabricOrganizationUserPassword `json:"Result" xml:"Result"`
}

// CreateResetFabricOrganizationUserPasswordRequest creates a request to invoke ResetFabricOrganizationUserPassword API
func CreateResetFabricOrganizationUserPasswordRequest() (request *ResetFabricOrganizationUserPasswordRequest) {
	request = &ResetFabricOrganizationUserPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "ResetFabricOrganizationUserPassword", "baas", "openAPI")
	return
}

// CreateResetFabricOrganizationUserPasswordResponse creates a response to parse from ResetFabricOrganizationUserPassword response
func CreateResetFabricOrganizationUserPasswordResponse() (response *ResetFabricOrganizationUserPasswordResponse) {
	response = &ResetFabricOrganizationUserPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
