package waf_openapi

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

// DescribeInstanceSpecInfo invokes the waf_openapi.DescribeInstanceSpecInfo API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/describeinstancespecinfo.html
func (client *Client) DescribeInstanceSpecInfo(request *DescribeInstanceSpecInfoRequest) (response *DescribeInstanceSpecInfoResponse, err error) {
	response = CreateDescribeInstanceSpecInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSpecInfoWithChan invokes the waf_openapi.DescribeInstanceSpecInfo API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/describeinstancespecinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSpecInfoWithChan(request *DescribeInstanceSpecInfoRequest) (<-chan *DescribeInstanceSpecInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSpecInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSpecInfo(request)
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

// DescribeInstanceSpecInfoWithCallback invokes the waf_openapi.DescribeInstanceSpecInfo API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/describeinstancespecinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceSpecInfoWithCallback(request *DescribeInstanceSpecInfoRequest, callback func(response *DescribeInstanceSpecInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSpecInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSpecInfo(request)
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

// DescribeInstanceSpecInfoRequest is the request struct for api DescribeInstanceSpecInfo
type DescribeInstanceSpecInfoRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
	Region     string `position:"Query" name:"Region"`
}

// DescribeInstanceSpecInfoResponse is the response struct for api DescribeInstanceSpecInfo
type DescribeInstanceSpecInfoResponse struct {
	*responses.BaseResponse
	RequestId        string                 `json:"RequestId" xml:"RequestId"`
	InstanceId       string                 `json:"InstanceId" xml:"InstanceId"`
	Version          string                 `json:"Version" xml:"Version"`
	ExpireTime       int64                  `json:"ExpireTime" xml:"ExpireTime"`
	InstanceSpecInfo []InstanceSpecInfoItem `json:"InstanceSpecInfo" xml:"InstanceSpecInfo"`
}

// CreateDescribeInstanceSpecInfoRequest creates a request to invoke DescribeInstanceSpecInfo API
func CreateDescribeInstanceSpecInfoRequest() (request *DescribeInstanceSpecInfoRequest) {
	request = &DescribeInstanceSpecInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeInstanceSpecInfo", "waf", "openAPI")
	return
}

// CreateDescribeInstanceSpecInfoResponse creates a response to parse from DescribeInstanceSpecInfo response
func CreateDescribeInstanceSpecInfoResponse() (response *DescribeInstanceSpecInfoResponse) {
	response = &DescribeInstanceSpecInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
