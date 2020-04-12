package vod

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

// GetUploadDetails invokes the vod.GetUploadDetails API synchronously
// api document: https://help.aliyun.com/api/vod/getuploaddetails.html
func (client *Client) GetUploadDetails(request *GetUploadDetailsRequest) (response *GetUploadDetailsResponse, err error) {
	response = CreateGetUploadDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// GetUploadDetailsWithChan invokes the vod.GetUploadDetails API asynchronously
// api document: https://help.aliyun.com/api/vod/getuploaddetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUploadDetailsWithChan(request *GetUploadDetailsRequest) (<-chan *GetUploadDetailsResponse, <-chan error) {
	responseChan := make(chan *GetUploadDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUploadDetails(request)
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

// GetUploadDetailsWithCallback invokes the vod.GetUploadDetails API asynchronously
// api document: https://help.aliyun.com/api/vod/getuploaddetails.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUploadDetailsWithCallback(request *GetUploadDetailsRequest, callback func(response *GetUploadDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUploadDetailsResponse
		var err error
		defer close(result)
		response, err = client.GetUploadDetails(request)
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

// GetUploadDetailsRequest is the request struct for api GetUploadDetails
type GetUploadDetailsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceRealOwnerId  requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	MediaIds             string           `position:"Query" name:"MediaIds"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	MediaType            string           `position:"Query" name:"MediaType"`
}

// GetUploadDetailsResponse is the response struct for api GetUploadDetails
type GetUploadDetailsResponse struct {
	*responses.BaseResponse
	RequestId         string         `json:"RequestId" xml:"RequestId"`
	NonExistMediaIds  []string       `json:"NonExistMediaIds" xml:"NonExistMediaIds"`
	ForbiddenMediaIds []string       `json:"ForbiddenMediaIds" xml:"ForbiddenMediaIds"`
	UploadDetails     []UploadDetail `json:"UploadDetails" xml:"UploadDetails"`
}

// CreateGetUploadDetailsRequest creates a request to invoke GetUploadDetails API
func CreateGetUploadDetailsRequest() (request *GetUploadDetailsRequest) {
	request = &GetUploadDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetUploadDetails", "vod", "openAPI")
	return
}

// CreateGetUploadDetailsResponse creates a response to parse from GetUploadDetails response
func CreateGetUploadDetailsResponse() (response *GetUploadDetailsResponse) {
	response = &GetUploadDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
