// Copyright 2014 Canonical Ltd.

package v1_test

import (
	"net/http"
	"time"

	"github.com/juju/testing/httptesting"
	"github.com/juju/utils/debugstatus"
	gc "gopkg.in/check.v1"

	buildver "github.com/CanonicalLtd/blues-identity/version"
)

type debugSuite struct {
	apiSuite
}

var _ = gc.Suite(&debugSuite{})

func (s *debugSuite) patchStartTime() time.Time {
	startTime := time.Now()
	s.PatchValue(&debugstatus.StartTime, startTime)
	return startTime
}

func (s *debugSuite) TestServeDebugStatus(c *gc.C) {
	startTime := s.patchStartTime()
	httptesting.AssertJSONCall(c, httptesting.JSONCallParams{
		Handler: s.srv,
		URL:     apiURL("debug/status"),
		ExpectBody: map[string]debugstatus.CheckResult{
			"server_started": {
				Name:   "Server started",
				Value:  startTime.String(),
				Passed: true,
			},
			"mongo_connected": {
				Name:   "MongoDB is connected",
				Value:  "Connected",
				Passed: true,
			},
			"mongo_collections": {
				Name:   "MongoDB collections",
				Value:  "All required collections exist",
				Passed: true,
			},
		},
	})
}

func (s *debugSuite) TestServeDebugInfo(c *gc.C) {
	httptesting.AssertJSONCall(c, httptesting.JSONCallParams{
		Handler:      s.srv,
		URL:          apiURL("debug/info"),
		ExpectStatus: http.StatusOK,
		ExpectBody:   buildver.VersionInfo,
	})
}