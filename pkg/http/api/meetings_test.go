package v1

import (
	"strings"
	"testing"

	"github.com/google/uuid"

	"github.com/b3scale/b3scale/pkg/bbb"
	"github.com/b3scale/b3scale/pkg/store"
)

func createTestMeeting(
	api *APIContext,
	backend *store.BackendState,
) *store.MeetingState {
	ctx := api.Ctx()
	tx, err := api.Conn.Begin(ctx)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback(ctx)

	m := store.InitMeetingState(&store.MeetingState{
		BackendID: &backend.ID,
		Meeting: &bbb.Meeting{
			MeetingID:         uuid.New().String(),
			InternalMeetingID: uuid.New().String(),
			AttendeePW:        "foo42",
			DialNumber:        "+12 345 666",
		},
	})

	if err := m.Save(ctx, tx); err != nil {
		panic(err)
	}
	if err := tx.Commit(ctx); err != nil {
		panic(err)
	}
	return m
}

func TestBackendMeetingsList(t *testing.T) {
	api, _ := NewTestRequest().Context()
	defer api.Release()

	backend := createTestBackend(api)
	meeting := createTestMeeting(api, backend)

	api, res := NewTestRequest().
		KeepState().
		Query("backend_host="+backend.Backend.Host).
		Authorize("admin42", ScopeAdmin).
		Context()
	defer api.Release()

	if err := api.Handle(APIResourceMeetings.List); err != nil {
		t.Fatal(err)
	}
	if err := res.StatusOK(); err != nil {
		t.Error(err)
	}
	body := res.Body()
	if !strings.Contains(body, meeting.ID) {
		t.Error("meeting ID", meeting.ID, "not found in response body", body)
	}
}

func TestMeetingShow(t *testing.T) {
	api, res := NewTestRequest().
		Authorize("test-agent-2000", ScopeNode).
		Context()
	defer api.Release()

	backend := createTestBackend(api)
	meeting := createTestMeeting(api, backend)

	api.SetParamNames("id")
	api.SetParamValues("internal:" + meeting.Meeting.InternalMeetingID)

	if err := api.Handle(APIResourceMeetings.Show); err != nil {
		t.Fatal(err)
	}

	if err := res.StatusOK(); err != nil {
		t.Error(err)
	}

	body := res.JSON()
	t.Log(body)
	meetingRes := body["meeting"].(map[string]interface{})
	if meetingRes["attendee_pw"].(string) != "foo42" {
		t.Error("unexpected meeting:", body)
	}
}

func TestMeetingDestroy(t *testing.T) {
	api, res := NewTestRequest().
		Authorize("test-agent-2000", ScopeNode).
		Context()
	defer api.Release()

	backend := createTestBackend(api)
	meeting := createTestMeeting(api, backend)

	api.SetParamNames("id")
	api.SetParamValues(meeting.Meeting.MeetingID)

	if err := api.Handle(APIResourceMeetings.Destroy); err != nil {
		t.Fatal(err)
	}

	if err := res.StatusOK(); err != nil {
		t.Error(err)
	}

	// Query the meeting again, this should fail.
	api, res = NewTestRequest().
		Authorize("test-agent-2000", ScopeNode).
		KeepState().
		Context()
	defer api.Release()

	api.SetParamNames("id")
	api.SetParamValues(meeting.Meeting.MeetingID)

	if err := api.Handle(APIResourceMeetings.Show); err == nil {
		t.Error("should raise an error")
	}
}

func TestMeetingUpdate(t *testing.T) {
	api, res := NewTestRequest().
		Authorize("test-agent-2000", ScopeNode).
		JSON(map[string]interface{}{
			"meeting": map[string]interface{}{
				"attendees": []map[string]interface{}{
					{
						"user_id":          "user123",
						"internal_user_id": "internal-user-123",
						"full_name":        "Jen Test",
						"role":             "admin",
						"is_presenter":     true,
					},
					{
						"user_id":          "user42",
						"internal_user_id": "internal-user-42",
						"full_name":        "Kate Test",
						"role":             "user",
						"is_presenter":     false,
					},
				},
			},
		}).
		Context()
	defer api.Release()

	backend := createTestBackend(api)
	meeting := createTestMeeting(api, backend)

	api.SetParamNames("id")
	api.SetParamValues(meeting.Meeting.MeetingID)

	if err := api.Handle(APIResourceMeetings.Update); err != nil {
		t.Fatal(err)
	}

	if err := res.StatusOK(); err != nil {
		t.Error(err)
	}

	body := res.JSON()
	t.Log(body)
	meetingRes := body["meeting"].(map[string]interface{})
	attendeesRes := meetingRes["attendees"].([]interface{})
	if len(attendeesRes) != 2 {
		t.Error("unexpected attendees", attendeesRes)
	}
	if meetingRes["dial_number"].(string) != "+12 345 666" {
		t.Error("partial update should not have touched other props", body)
	}

}