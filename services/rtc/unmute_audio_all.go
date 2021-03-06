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

package rtc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UnmuteAudioAll invokes the rtc.UnmuteAudioAll API synchronously
// api document: https://help.aliyun.com/api/rtc/unmuteaudioall.html
func (client *Client) UnmuteAudioAll(request *UnmuteAudioAllRequest) (response *UnmuteAudioAllResponse, err error) {
	response = CreateUnmuteAudioAllResponse()
	err = client.DoAction(request, response)
	return
}

// UnmuteAudioAllWithChan invokes the rtc.UnmuteAudioAll API asynchronously
// api document: https://help.aliyun.com/api/rtc/unmuteaudioall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnmuteAudioAllWithChan(request *UnmuteAudioAllRequest) (<-chan *UnmuteAudioAllResponse, <-chan error) {
	responseChan := make(chan *UnmuteAudioAllResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnmuteAudioAll(request)
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

// UnmuteAudioAllWithCallback invokes the rtc.UnmuteAudioAll API asynchronously
// api document: https://help.aliyun.com/api/rtc/unmuteaudioall.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnmuteAudioAllWithCallback(request *UnmuteAudioAllRequest, callback func(response *UnmuteAudioAllResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnmuteAudioAllResponse
		var err error
		defer close(result)
		response, err = client.UnmuteAudioAll(request)
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

// UnmuteAudioAllRequest is the request struct for api UnmuteAudioAll
type UnmuteAudioAllRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	AppId         string           `position:"Query" name:"AppId"`
	ConferenceId  string           `position:"Query" name:"ConferenceId"`
	ParticipantId string           `position:"Query" name:"ParticipantId"`
}

// UnmuteAudioAllResponse is the response struct for api UnmuteAudioAll
type UnmuteAudioAllResponse struct {
	*responses.BaseResponse
	RequestId    string                      `json:"RequestId" xml:"RequestId"`
	ConferenceId string                      `json:"ConferenceId" xml:"ConferenceId"`
	Participants UnmuteAudioAllParticipants0 `json:"Participants" xml:"Participants"`
}

type UnmuteAudioAllParticipants0 struct {
	Participant []UnmuteAudioAllParticipant1 `json:"Participant" xml:"Participant"`
}

type UnmuteAudioAllParticipant1 struct {
	Id      string `json:"Id" xml:"Id"`
	Code    string `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
}

// CreateUnmuteAudioAllRequest creates a request to invoke UnmuteAudioAll API
func CreateUnmuteAudioAllRequest() (request *UnmuteAudioAllRequest) {
	request = &UnmuteAudioAllRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "UnmuteAudioAll", "rtc", "openAPI")
	return
}

// CreateUnmuteAudioAllResponse creates a response to parse from UnmuteAudioAll response
func CreateUnmuteAudioAllResponse() (response *UnmuteAudioAllResponse) {
	response = &UnmuteAudioAllResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
