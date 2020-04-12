package dyplsapi

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

// AddAxnTrackNo invokes the dyplsapi.AddAxnTrackNo API synchronously
// api document: https://help.aliyun.com/api/dyplsapi/addaxntrackno.html
func (client *Client) AddAxnTrackNo(request *AddAxnTrackNoRequest) (response *AddAxnTrackNoResponse, err error) {
	response = CreateAddAxnTrackNoResponse()
	err = client.DoAction(request, response)
	return
}

// AddAxnTrackNoWithChan invokes the dyplsapi.AddAxnTrackNo API asynchronously
// api document: https://help.aliyun.com/api/dyplsapi/addaxntrackno.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddAxnTrackNoWithChan(request *AddAxnTrackNoRequest) (<-chan *AddAxnTrackNoResponse, <-chan error) {
	responseChan := make(chan *AddAxnTrackNoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAxnTrackNo(request)
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

// AddAxnTrackNoWithCallback invokes the dyplsapi.AddAxnTrackNo API asynchronously
// api document: https://help.aliyun.com/api/dyplsapi/addaxntrackno.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddAxnTrackNoWithCallback(request *AddAxnTrackNoRequest, callback func(response *AddAxnTrackNoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAxnTrackNoResponse
		var err error
		defer close(result)
		response, err = client.AddAxnTrackNo(request)
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

// AddAxnTrackNoRequest is the request struct for api AddAxnTrackNo
type AddAxnTrackNoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SubsId               string           `position:"Query" name:"SubsId"`
	PhoneNoX             string           `position:"Query" name:"PhoneNoX"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TrackNo              string           `position:"Query" name:"trackNo"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
}

// AddAxnTrackNoResponse is the response struct for api AddAxnTrackNo
type AddAxnTrackNoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateAddAxnTrackNoRequest creates a request to invoke AddAxnTrackNo API
func CreateAddAxnTrackNoRequest() (request *AddAxnTrackNoRequest) {
	request = &AddAxnTrackNoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "AddAxnTrackNo", "dypls", "openAPI")
	return
}

// CreateAddAxnTrackNoResponse creates a response to parse from AddAxnTrackNo response
func CreateAddAxnTrackNoResponse() (response *AddAxnTrackNoResponse) {
	response = &AddAxnTrackNoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
