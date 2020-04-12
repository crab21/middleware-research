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

// CreateImageLib invokes the green.CreateImageLib API synchronously
// api document: https://help.aliyun.com/api/green/createimagelib.html
func (client *Client) CreateImageLib(request *CreateImageLibRequest) (response *CreateImageLibResponse, err error) {
	response = CreateCreateImageLibResponse()
	err = client.DoAction(request, response)
	return
}

// CreateImageLibWithChan invokes the green.CreateImageLib API asynchronously
// api document: https://help.aliyun.com/api/green/createimagelib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateImageLibWithChan(request *CreateImageLibRequest) (<-chan *CreateImageLibResponse, <-chan error) {
	responseChan := make(chan *CreateImageLibResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateImageLib(request)
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

// CreateImageLibWithCallback invokes the green.CreateImageLib API asynchronously
// api document: https://help.aliyun.com/api/green/createimagelib.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateImageLibWithCallback(request *CreateImageLibRequest, callback func(response *CreateImageLibResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateImageLibResponse
		var err error
		defer close(result)
		response, err = client.CreateImageLib(request)
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

// CreateImageLibRequest is the request struct for api CreateImageLib
type CreateImageLibRequest struct {
	*requests.RpcRequest
	ServiceModule string           `position:"Query" name:"ServiceModule"`
	Scene         string           `position:"Query" name:"Scene"`
	SourceIp      string           `position:"Query" name:"SourceIp"`
	Enable        requests.Boolean `position:"Query" name:"Enable"`
	BizTypes      string           `position:"Query" name:"BizTypes"`
	Name          string           `position:"Query" name:"Name"`
	Category      string           `position:"Query" name:"Category"`
}

// CreateImageLibResponse is the response struct for api CreateImageLib
type CreateImageLibResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateImageLibRequest creates a request to invoke CreateImageLib API
func CreateCreateImageLibRequest() (request *CreateImageLibRequest) {
	request = &CreateImageLibRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "CreateImageLib", "green", "openAPI")
	return
}

// CreateCreateImageLibResponse creates a response to parse from CreateImageLib response
func CreateCreateImageLibResponse() (response *CreateImageLibResponse) {
	response = &CreateImageLibResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
