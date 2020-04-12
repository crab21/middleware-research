package iot

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

// GetDataAPIServiceDetail invokes the iot.GetDataAPIServiceDetail API synchronously
// api document: https://help.aliyun.com/api/iot/getdataapiservicedetail.html
func (client *Client) GetDataAPIServiceDetail(request *GetDataAPIServiceDetailRequest) (response *GetDataAPIServiceDetailResponse, err error) {
	response = CreateGetDataAPIServiceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataAPIServiceDetailWithChan invokes the iot.GetDataAPIServiceDetail API asynchronously
// api document: https://help.aliyun.com/api/iot/getdataapiservicedetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataAPIServiceDetailWithChan(request *GetDataAPIServiceDetailRequest) (<-chan *GetDataAPIServiceDetailResponse, <-chan error) {
	responseChan := make(chan *GetDataAPIServiceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataAPIServiceDetail(request)
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

// GetDataAPIServiceDetailWithCallback invokes the iot.GetDataAPIServiceDetail API asynchronously
// api document: https://help.aliyun.com/api/iot/getdataapiservicedetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataAPIServiceDetailWithCallback(request *GetDataAPIServiceDetailRequest, callback func(response *GetDataAPIServiceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataAPIServiceDetailResponse
		var err error
		defer close(result)
		response, err = client.GetDataAPIServiceDetail(request)
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

// GetDataAPIServiceDetailRequest is the request struct for api GetDataAPIServiceDetail
type GetDataAPIServiceDetailRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Body" name:"IotInstanceId"`
	ApiSrn        string `position:"Body" name:"ApiSrn"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GetDataAPIServiceDetailResponse is the response struct for api GetDataAPIServiceDetail
type GetDataAPIServiceDetailResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGetDataAPIServiceDetailRequest creates a request to invoke GetDataAPIServiceDetail API
func CreateGetDataAPIServiceDetailRequest() (request *GetDataAPIServiceDetailRequest) {
	request = &GetDataAPIServiceDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetDataAPIServiceDetail", "Iot", "openAPI")
	return
}

// CreateGetDataAPIServiceDetailResponse creates a response to parse from GetDataAPIServiceDetail response
func CreateGetDataAPIServiceDetailResponse() (response *GetDataAPIServiceDetailResponse) {
	response = &GetDataAPIServiceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
