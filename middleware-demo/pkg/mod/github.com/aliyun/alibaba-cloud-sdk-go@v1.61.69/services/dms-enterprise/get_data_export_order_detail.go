package dms_enterprise

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

// GetDataExportOrderDetail invokes the dms_enterprise.GetDataExportOrderDetail API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdataexportorderdetail.html
func (client *Client) GetDataExportOrderDetail(request *GetDataExportOrderDetailRequest) (response *GetDataExportOrderDetailResponse, err error) {
	response = CreateGetDataExportOrderDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetDataExportOrderDetailWithChan invokes the dms_enterprise.GetDataExportOrderDetail API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdataexportorderdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataExportOrderDetailWithChan(request *GetDataExportOrderDetailRequest) (<-chan *GetDataExportOrderDetailResponse, <-chan error) {
	responseChan := make(chan *GetDataExportOrderDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDataExportOrderDetail(request)
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

// GetDataExportOrderDetailWithCallback invokes the dms_enterprise.GetDataExportOrderDetail API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/getdataexportorderdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDataExportOrderDetailWithCallback(request *GetDataExportOrderDetailRequest, callback func(response *GetDataExportOrderDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDataExportOrderDetailResponse
		var err error
		defer close(result)
		response, err = client.GetDataExportOrderDetail(request)
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

// GetDataExportOrderDetailRequest is the request struct for api GetDataExportOrderDetail
type GetDataExportOrderDetailRequest struct {
	*requests.RpcRequest
	OrderId requests.Integer `position:"Body" name:"OrderId"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
}

// GetDataExportOrderDetailResponse is the response struct for api GetDataExportOrderDetail
type GetDataExportOrderDetailResponse struct {
	*responses.BaseResponse
	RequestId             string                `json:"RequestId" xml:"RequestId"`
	Success               bool                  `json:"Success" xml:"Success"`
	ErrorMessage          string                `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode             string                `json:"ErrorCode" xml:"ErrorCode"`
	DataExportOrderDetail DataExportOrderDetail `json:"DataExportOrderDetail" xml:"DataExportOrderDetail"`
}

// CreateGetDataExportOrderDetailRequest creates a request to invoke GetDataExportOrderDetail API
func CreateGetDataExportOrderDetailRequest() (request *GetDataExportOrderDetailRequest) {
	request = &GetDataExportOrderDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetDataExportOrderDetail", "", "")
	return
}

// CreateGetDataExportOrderDetailResponse creates a response to parse from GetDataExportOrderDetail response
func CreateGetDataExportOrderDetailResponse() (response *GetDataExportOrderDetailResponse) {
	response = &GetDataExportOrderDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
