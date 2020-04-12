package reid

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

// ListMaskDetectionResults invokes the reid.ListMaskDetectionResults API synchronously
// api document: https://help.aliyun.com/api/reid/listmaskdetectionresults.html
func (client *Client) ListMaskDetectionResults(request *ListMaskDetectionResultsRequest) (response *ListMaskDetectionResultsResponse, err error) {
	response = CreateListMaskDetectionResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMaskDetectionResultsWithChan invokes the reid.ListMaskDetectionResults API asynchronously
// api document: https://help.aliyun.com/api/reid/listmaskdetectionresults.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListMaskDetectionResultsWithChan(request *ListMaskDetectionResultsRequest) (<-chan *ListMaskDetectionResultsResponse, <-chan error) {
	responseChan := make(chan *ListMaskDetectionResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMaskDetectionResults(request)
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

// ListMaskDetectionResultsWithCallback invokes the reid.ListMaskDetectionResults API asynchronously
// api document: https://help.aliyun.com/api/reid/listmaskdetectionresults.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListMaskDetectionResultsWithCallback(request *ListMaskDetectionResultsRequest, callback func(response *ListMaskDetectionResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMaskDetectionResultsResponse
		var err error
		defer close(result)
		response, err = client.ListMaskDetectionResults(request)
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

// ListMaskDetectionResultsRequest is the request struct for api ListMaskDetectionResults
type ListMaskDetectionResultsRequest struct {
	*requests.RpcRequest
	EndTime    requests.Integer `position:"Body" name:"EndTime"`
	StartTime  requests.Integer `position:"Body" name:"StartTime"`
	StoreId    requests.Integer `position:"Body" name:"StoreId"`
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
}

// ListMaskDetectionResultsResponse is the response struct for api ListMaskDetectionResults
type ListMaskDetectionResultsResponse struct {
	*responses.BaseResponse
	ErrorCode            string            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage         string            `json:"ErrorMessage" xml:"ErrorMessage"`
	PageNumber           int               `json:"PageNumber" xml:"PageNumber"`
	Message              string            `json:"Message" xml:"Message"`
	Code                 string            `json:"Code" xml:"Code"`
	TotalCount           int64             `json:"TotalCount" xml:"TotalCount"`
	DynamicCode          string            `json:"DynamicCode" xml:"DynamicCode"`
	PageSize             int               `json:"PageSize" xml:"PageSize"`
	RequestId            string            `json:"RequestId" xml:"RequestId"`
	Success              bool              `json:"Success" xml:"Success"`
	DynamicMessage       string            `json:"DynamicMessage" xml:"DynamicMessage"`
	MaskDetectionResults []DetectionResult `json:"MaskDetectionResults" xml:"MaskDetectionResults"`
}

// CreateListMaskDetectionResultsRequest creates a request to invoke ListMaskDetectionResults API
func CreateListMaskDetectionResultsRequest() (request *ListMaskDetectionResultsRequest) {
	request = &ListMaskDetectionResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "ListMaskDetectionResults", "1.1.2", "openAPI")
	return
}

// CreateListMaskDetectionResultsResponse creates a response to parse from ListMaskDetectionResults response
func CreateListMaskDetectionResultsResponse() (response *ListMaskDetectionResultsResponse) {
	response = &ListMaskDetectionResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
