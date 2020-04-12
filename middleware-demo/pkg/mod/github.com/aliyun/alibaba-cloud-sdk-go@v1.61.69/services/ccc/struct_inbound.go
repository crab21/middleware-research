package ccc

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

// Inbound is a nested struct in ccc response
type Inbound struct {
	CallsIncomingIVR                          int64   `json:"CallsIncomingIVR" xml:"CallsIncomingIVR"`
	GiveUpByAgentOfQueueCount                 int64   `json:"GiveUpByAgentOfQueueCount" xml:"GiveUpByAgentOfQueueCount"`
	CallsIncomingQueue                        int64   `json:"CallsIncomingQueue" xml:"CallsIncomingQueue"`
	TotalRingTime                             int64   `json:"TotalRingTime" xml:"TotalRingTime"`
	CallsIncomingLine                         int64   `json:"CallsIncomingLine" xml:"CallsIncomingLine"`
	AnsweredByAgentOfQueueCount               int64   `json:"AnsweredByAgentOfQueueCount" xml:"AnsweredByAgentOfQueueCount"`
	InComingQueueOfQueueCount                 int64   `json:"InComingQueueOfQueueCount" xml:"InComingQueueOfQueueCount"`
	MaxRingTime                               int64   `json:"MaxRingTime" xml:"MaxRingTime"`
	AverageRingTime                           int64   `json:"AverageRingTime" xml:"AverageRingTime"`
	AnsweredByAgentOfQueueWaitTimeDuration    int64   `json:"AnsweredByAgentOfQueueWaitTimeDuration" xml:"AnsweredByAgentOfQueueWaitTimeDuration"`
	CallsAbandonedInQueue                     int64   `json:"CallsAbandonedInQueue" xml:"CallsAbandonedInQueue"`
	MaxTalkTime                               string  `json:"MaxTalkTime" xml:"MaxTalkTime"`
	MaxWorkTime                               int64   `json:"MaxWorkTime" xml:"MaxWorkTime"`
	TotalWorkTime                             int64   `json:"TotalWorkTime" xml:"TotalWorkTime"`
	SatisfactionSurveysOffered                int64   `json:"SatisfactionSurveysOffered" xml:"SatisfactionSurveysOffered"`
	ServiceLevel20                            float64 `json:"ServiceLevel20" xml:"ServiceLevel20"`
	SatisfactionSurveysResponded              int64   `json:"SatisfactionSurveysResponded" xml:"SatisfactionSurveysResponded"`
	HandleRate                                float64 `json:"HandleRate" xml:"HandleRate"`
	AverageWorkTime                           int64   `json:"AverageWorkTime" xml:"AverageWorkTime"`
	SatisfactionIndex                         float64 `json:"SatisfactionIndex" xml:"SatisfactionIndex"`
	QueueMaxWaitTimeDuration                  int64   `json:"QueueMaxWaitTimeDuration" xml:"QueueMaxWaitTimeDuration"`
	AverageTalkTime                           int64   `json:"AverageTalkTime" xml:"AverageTalkTime"`
	AbandonedInQueueOfQueueCount              int64   `json:"AbandonedInQueueOfQueueCount" xml:"AbandonedInQueueOfQueueCount"`
	OverFlowInQueueOfQueueCount               int64   `json:"OverFlowInQueueOfQueueCount" xml:"OverFlowInQueueOfQueueCount"`
	TotalTalkTime                             int64   `json:"TotalTalkTime" xml:"TotalTalkTime"`
	CallsAbandonedInIVR                       int64   `json:"CallsAbandonedInIVR" xml:"CallsAbandonedInIVR"`
	QueueWaitTimeDuration                     int64   `json:"QueueWaitTimeDuration" xml:"QueueWaitTimeDuration"`
	CallsOffered                              int64   `json:"CallsOffered" xml:"CallsOffered"`
	CallsHandled                              int64   `json:"CallsHandled" xml:"CallsHandled"`
	AnsweredByAgentOfQueueMaxWaitTimeDuration int64   `json:"AnsweredByAgentOfQueueMaxWaitTimeDuration" xml:"AnsweredByAgentOfQueueMaxWaitTimeDuration"`
}
