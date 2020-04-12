package outboundbot

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

// ConversationDetail is a nested struct in outboundbot response
type ConversationDetail struct {
	NodeId         string        `json:"NodeId" xml:"NodeId"`
	TaskId         string        `json:"TaskId" xml:"TaskId"`
	Timestamp      int64         `json:"Timestamp" xml:"Timestamp"`
	Script         string        `json:"Script" xml:"Script"`
	Id             string        `json:"Id" xml:"Id"`
	Speaker        string        `json:"Speaker" xml:"Speaker"`
	ConversationId string        `json:"ConversationId" xml:"ConversationId"`
	Summary        []SummaryItem `json:"Summary" xml:"Summary"`
}
