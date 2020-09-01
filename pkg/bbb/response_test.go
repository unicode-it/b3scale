package bbb

import (
	"io/ioutil"
	"path"
	"testing"
)

func readTestResponse(name string) []byte {
	filename := path.Join(
		"../../test/data/responses/",
		name)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return data
}

func TestUnmarshalCreateResponse(t *testing.T) {
	data := readTestResponse("createSuccess.xml")
	response, err := UnmarshalCreateResponse(data)
	if err != nil {
		t.Error(err)
	}
	if response.MeetingID != "Test" {
		t.Error("Unexpected meeting id")
	}
}

func TestMarshalCreateResponse(t *testing.T) {
	data := readTestResponse("createSuccess.xml")
	response, err := UnmarshalCreateResponse(data)
	if err != nil {
		t.Error(err)
	}

	data1, err := response.Marshal()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data1))
}

func TestUnmarshalJoinResponse(t *testing.T) {
	data := readTestResponse("joinSuccess.xml")
	response, err := UnmarshalJoinResponse(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(response)
}

func TestUnmarshalIsMeetingRunningResponse(t *testing.T) {
	data := readTestResponse("isMeetingRunningSuccess.xml")
	response, err := UnmarshalIsMeetingRunningResponse(data)
	if err != nil {
		t.Error(err)
	}

	if response.XMLResponse.Returncode != "SUCCESS" {
		t.Error("Unexpected returncode:", response.XMLResponse.Returncode)
	}
	if response.Running != true {
		t.Error("Expected running to be true")
	}
}

func TestMarshalIsMeetingRunningResponse(t *testing.T) {
	data := readTestResponse("isMeetingRunningSuccess.xml")
	response, err := UnmarshalIsMeetingRunningResponse(data)
	if err != nil {
		t.Error(err)
	}
	data1, err := response.Marshal()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data1))
}

func TestUnmarshalEndResponse(t *testing.T) {
	data := readTestResponse("endSuccess.xml")
	response, err := UnmarshalEndResponse(data)
	if err != nil {
		t.Error(err)
	}

	if response.XMLResponse.MessageKey != "sentEndMeetingRequest" {
		t.Error("Unexpected MessageKey:", response.XMLResponse.MessageKey)
	}
}

func TestUnmarshalGetMeetingInfoRespons(t *testing.T) {
	data := readTestResponse("getMeetingInfoSuccess.xml")
	response, err := UnmarshalGetMeetingInfoResponse(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(response)

	if response.Meeting.MeetingID != "Demo Meeting" {
		t.Error("Unexpected MeetingID:", response.Meeting.MeetingID)
	}

	if response.Meeting.AttendeesCollection == nil {
		t.Error("AttendeesCollection is nil")
	}

	if len(response.Meeting.AttendeesCollection.Attendees) != 2 {
		t.Error("Unexpected attendees length:",
			len(response.Meeting.AttendeesCollection.Attendees))
	}

	t.Log(response.Meeting.AttendeesCollection.Attendees)
	t.Log(response.Meeting.Metadata)
}

func TestUnmarshalGetMeetingInfoResponsBreakout(t *testing.T) {
	// We have a breakout room response
	data := readTestResponse("getMeetingInfoSuccess-breakout.xml")
	response, err := UnmarshalGetMeetingInfoResponse(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(response)

	if response.Meeting.MeetingID != "Demo Meeting" {
		t.Error("Unexpected MeetingID:", response.Meeting.MeetingID)
	}

	if response.Meeting.IsBreakout == false {
		t.Error("Should be a breakout room")
	}

	if response.Meeting.Breakout.ParentMeetingID != "ParentMeetingId" {
		t.Error("Unexpected parentMeetingId:",
			response.Meeting.Breakout.ParentMeetingID)
	}
}

func TestUnmarshalGetMeetingInfoResponsBreakoutParent(t *testing.T) {
	// We have a breakout parent response
	data := readTestResponse("getMeetingInfoSuccess-breakoutParent.xml")
	response, err := UnmarshalGetMeetingInfoResponse(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(response)

	if response.Meeting.MeetingID != "Demo Meeting" {
		t.Error("Unexpected MeetingID:", response.Meeting.MeetingID)
	}

	if response.Meeting.IsBreakout == true {
		t.Error("Should not be a breakout room")
	}

	// Check breakout info
	rooms := response.Meeting.BreakoutRooms
	if len(rooms.BreakoutIDs) != 3 {
		t.Error("Expected 3 breakout ids. got:", len(rooms.BreakoutIDs))
	}

	if rooms.BreakoutIDs[0] != "breakout-room-id-1" {
		t.Error("Unexpected breakout ID:",
			rooms.BreakoutIDs[0])
	}
}

func TestMarshalGetMeetingInfoRequest(t *testing.T) {
	data := readTestResponse("getMeetingInfoSuccess.xml")
	response, err := UnmarshalGetMeetingInfoResponse(data)
	if err != nil {
		t.Error(err)
	}
	// Serialize
	data1, err := response.Marshal()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(data1))
}
