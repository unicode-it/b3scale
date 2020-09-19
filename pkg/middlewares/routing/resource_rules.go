package routing

import (
	"gitlab.com/infra.run/public/b3scale/pkg/bbb"
	"gitlab.com/infra.run/public/b3scale/pkg/cluster"
)

// ResourceRules creates the rule enforcing
// route handler.
func ResourceRules() cluster.RouterMiddleware {
	return func(next cluster.RouterHandler) cluster.RouterHandler {
		return func(backends []*cluster.Backend, req *bbb.Request) ([]*cluster.Backend, error) {
			// Apply rules for resources
			backends = applyResourceRule(backends, req.Resource)
			return next(backends, req)
		}
	}
}

// Apply rule based on resource
func applyResourceRule(
	backends []*cluster.Backend, res string,
) []*cluster.Backend {
	switch res {
	case bbb.ResJoin:
		return selectFirst(backends)
	case bbb.ResCreate:
		return selectFirst(backends)
	case bbb.ResIsMeetingRunning:
		return selectFirst(backends)
	case bbb.ResEnd:
		return selectFirst(backends)
	case bbb.ResGetMeetingInfo:
		return selectFirst(backends)
	case bbb.ResGetMeetings:
		return backends
	case bbb.ResGetRecordings:
		return backends
	case bbb.ResPublishRecordings:
		return backends
	case bbb.ResDeleteRecordings:
		return backends
	case bbb.ResUpdateRecordings:
		return selectFirst(backends)
	case bbb.ResGetDefaultConfigXML:
		return selectFirst(backends)
	case bbb.ResSetConfigXML:
		return selectFirst(backends)
	case bbb.ResGetRecordingTextTracks:
		return selectFirst(backends)
	case bbb.ResPutRecordingTextTrack:
		return selectFirst(backends)
	}

	return backends
}

// Keep only first backend
func selectFirst(backends []*cluster.Backend) []*cluster.Backend {
	// The following slice operation with empty slices
	if len(backends) == 0 {
		return backends
	}
	return backends[:1]
}