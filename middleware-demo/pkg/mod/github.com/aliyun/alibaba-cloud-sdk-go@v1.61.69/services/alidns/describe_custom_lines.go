package alidns

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

// DescribeCustomLines invokes the alidns.DescribeCustomLines API synchronously
// api document: https://help.aliyun.com/api/alidns/describecustomlines.html
func (client *Client) DescribeCustomLines(request *DescribeCustomLinesRequest) (response *DescribeCustomLinesResponse, err error) {
	response = CreateDescribeCustomLinesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCustomLinesWithChan invokes the alidns.DescribeCustomLines API asynchronously
// api document: https://help.aliyun.com/api/alidns/describecustomlines.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomLinesWithChan(request *DescribeCustomLinesRequest) (<-chan *DescribeCustomLinesResponse, <-chan error) {
	responseChan := make(chan *DescribeCustomLinesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCustomLines(request)
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

// DescribeCustomLinesWithCallback invokes the alidns.DescribeCustomLines API asynchronously
// api document: https://help.aliyun.com/api/alidns/describecustomlines.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCustomLinesWithCallback(request *DescribeCustomLinesRequest, callback func(response *DescribeCustomLinesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCustomLinesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCustomLines(request)
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

// DescribeCustomLinesRequest is the request struct for api DescribeCustomLines
type DescribeCustomLinesRequest struct {
	*requests.RpcRequest
	DomainName   string           `position:"Query" name:"DomainName"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	UserClientIp string           `position:"Query" name:"UserClientIp"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
	Lang         string           `position:"Query" name:"Lang"`
}

// DescribeCustomLinesResponse is the response struct for api DescribeCustomLines
type DescribeCustomLinesResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	TotalItems  int          `json:"TotalItems" xml:"TotalItems"`
	PageNumber  int          `json:"PageNumber" xml:"PageNumber"`
	PageSize    int          `json:"PageSize" xml:"PageSize"`
	TotalPages  int          `json:"TotalPages" xml:"TotalPages"`
	CustomLines []CustomLine `json:"CustomLines" xml:"CustomLines"`
}

// CreateDescribeCustomLinesRequest creates a request to invoke DescribeCustomLines API
func CreateDescribeCustomLinesRequest() (request *DescribeCustomLinesRequest) {
	request = &DescribeCustomLinesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "DescribeCustomLines", "alidns", "openAPI")
	return
}

// CreateDescribeCustomLinesResponse creates a response to parse from DescribeCustomLines response
func CreateDescribeCustomLinesResponse() (response *DescribeCustomLinesResponse) {
	response = &DescribeCustomLinesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
