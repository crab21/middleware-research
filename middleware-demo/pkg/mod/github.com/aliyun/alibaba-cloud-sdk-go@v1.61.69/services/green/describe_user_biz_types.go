package green

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

// DescribeUserBizTypes invokes the green.DescribeUserBizTypes API synchronously
// api document: https://help.aliyun.com/api/green/describeuserbiztypes.html
func (client *Client) DescribeUserBizTypes(request *DescribeUserBizTypesRequest) (response *DescribeUserBizTypesResponse, err error) {
	response = CreateDescribeUserBizTypesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserBizTypesWithChan invokes the green.DescribeUserBizTypes API asynchronously
// api document: https://help.aliyun.com/api/green/describeuserbiztypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserBizTypesWithChan(request *DescribeUserBizTypesRequest) (<-chan *DescribeUserBizTypesResponse, <-chan error) {
	responseChan := make(chan *DescribeUserBizTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserBizTypes(request)
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

// DescribeUserBizTypesWithCallback invokes the green.DescribeUserBizTypes API asynchronously
// api document: https://help.aliyun.com/api/green/describeuserbiztypes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserBizTypesWithCallback(request *DescribeUserBizTypesRequest, callback func(response *DescribeUserBizTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserBizTypesResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserBizTypes(request)
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

// DescribeUserBizTypesRequest is the request struct for api DescribeUserBizTypes
type DescribeUserBizTypesRequest struct {
	*requests.RpcRequest
	Customized requests.Boolean `position:"Query" name:"Customized"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
}

// DescribeUserBizTypesResponse is the response struct for api DescribeUserBizTypes
type DescribeUserBizTypesResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	BizTypeList       []Item `json:"BizTypeList" xml:"BizTypeList"`
	BizTypeListImport []Item `json:"BizTypeListImport" xml:"BizTypeListImport"`
}

// CreateDescribeUserBizTypesRequest creates a request to invoke DescribeUserBizTypes API
func CreateDescribeUserBizTypesRequest() (request *DescribeUserBizTypesRequest) {
	request = &DescribeUserBizTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "DescribeUserBizTypes", "green", "openAPI")
	return
}

// CreateDescribeUserBizTypesResponse creates a response to parse from DescribeUserBizTypes response
func CreateDescribeUserBizTypesResponse() (response *DescribeUserBizTypesResponse) {
	response = &DescribeUserBizTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
