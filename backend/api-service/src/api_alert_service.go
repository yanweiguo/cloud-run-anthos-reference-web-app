// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
/*
 * Inventory API
 *
 * Inventory API for the Cloud Run for Anthos Reference Web App
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 *
 * NOTE: The AlertApiService is not yet used in the app - functionality
 * coming soon.
 */

package service

import (
	"context"
	"net/http"
	"os"
)

// AlertApiService is a service that implents the logic for the AlertApiServicer
// This service should implement the business logic for every endpoint for the AlertApi API.
// Include any external packages or services that will be required by this service.
type AlertApiService struct {
	db DatabaseBackend
}

// NewAlertApiService creates a default api service
func NewAlertApiService() AlertApiServicer {
	projectID := os.Getenv("PROJECT_ID")
	backend := NewFirestoreBackend(projectID)
	return &AlertApiService{backend}
}

// DeleteAlert - Delete Alert by ID
func (s *AlertApiService) DeleteAlert(id string, w http.ResponseWriter) error {
	ctx := context.Background()
	err := s.db.DeleteAlert(ctx, id)
	if err != nil {
		return err
	}

	return EncodeJSONStatus(http.StatusOK, "alert deleted", w)
}

// ListAlerts - List all Alerts
func (s *AlertApiService) ListAlerts(w http.ResponseWriter) error {
	ctx := context.Background()
	l, err := s.db.ListAlerts(ctx)
	if err != nil {
		return err
	}

	return EncodeJSONResponse(l, nil, w)
}

// NewAlert - Create a new Alert
func (s *AlertApiService) NewAlert(alert Alert, w http.ResponseWriter) error {
	ctx := context.Background()
	r, err := s.db.NewAlert(ctx, &alert)
	if err != nil {
		return err
	}

	status := http.StatusCreated
	return EncodeJSONResponse(r, &status, w)
}
