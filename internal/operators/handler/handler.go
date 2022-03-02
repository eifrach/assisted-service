package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/openshift/assisted-service/internal/cluster"
	"github.com/openshift/assisted-service/internal/common"
	eventgen "github.com/openshift/assisted-service/internal/common/events"
	eventsapi "github.com/openshift/assisted-service/internal/events/api"
	"github.com/openshift/assisted-service/internal/operators"
	"github.com/openshift/assisted-service/models"
	logutil "github.com/openshift/assisted-service/pkg/log"
	restoperators "github.com/openshift/assisted-service/restapi/operations/operators"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Handler implements REST API interface and deals with HTTP objects and transport data model.
type Handler struct {
	// operatorsAPI is responsible for executing the actual logic related to the operators
	operatorsAPI       operators.API
	db                 *gorm.DB
	log                logrus.FieldLogger
	eventsHandler      eventsapi.Handler
	clusterProgressAPI cluster.ProgressAPI
}

// NewHandler creates new handler
func NewHandler(operatorsAPI operators.API, log logrus.FieldLogger, db *gorm.DB, eventsHandler eventsapi.Handler, clusterProgressAPI cluster.ProgressAPI) *Handler {
	return &Handler{operatorsAPI: operatorsAPI, log: log, db: db, eventsHandler: eventsHandler, clusterProgressAPI: clusterProgressAPI}
}

// ListOperatorProperties Lists properties for an operator name.
func (h *Handler) ListOperatorProperties(ctx context.Context, params restoperators.ListOperatorPropertiesParams) middleware.Responder {
	return common.NewApiError(http.StatusNotFound, errors.New(common.APINotFound))
}

// ListSupportedOperators Retrieves the list of supported operators.
func (h *Handler) ListSupportedOperators(_ context.Context, _ restoperators.ListSupportedOperatorsParams) middleware.Responder {
	return common.NewApiError(http.StatusNotFound, errors.New(common.APINotFound))
}

// ListOfClusterOperators Lists operators to be monitored for a cluster.
func (h *Handler) ListOfClusterOperators(ctx context.Context, params restoperators.ListOfClusterOperatorsParams) middleware.Responder {
	return common.NewApiError(http.StatusNotFound, errors.New(common.APINotFound))
}

// ReportMonitoredOperatorStatus Controller API to report of monitored operators.
func (h *Handler) ReportMonitoredOperatorStatus(ctx context.Context, params restoperators.ReportMonitoredOperatorStatusParams) middleware.Responder {
	return common.NewApiError(http.StatusNotFound, errors.New(common.APINotFound))
}

// ReportMonitoredOperatorStatus Controller API to report of monitored operators.
func (h *Handler) V2ReportMonitoredOperatorStatus(ctx context.Context, params restoperators.V2ReportMonitoredOperatorStatusParams) middleware.Responder {

	log := logutil.FromContext(ctx, h.log)

	txSuccess := false
	tx := h.db.Begin()
	defer func() {
		if !txSuccess {
			log.Error("update monitored operator failed")
			tx.Rollback()
		}
		if r := recover(); r != nil {
			log.Error("update monitored operator failed")
			tx.Rollback()
		}
	}()

	if err := h.UpdateMonitoredOperatorStatus(ctx, params.ClusterID, params.ReportParams.Name, params.ReportParams.Status, params.ReportParams.StatusInfo, tx); err != nil {
		return common.GenerateErrorResponder(err)
	}

	if err := h.clusterProgressAPI.UpdateFinalizingProgress(ctx, tx, params.ClusterID); err != nil {
		return common.GenerateErrorResponder(err)
	}

	if err := tx.Commit().Error; err != nil {
		log.WithError(err).Error("DB error, failed to commit")
		return common.NewApiError(http.StatusInternalServerError, err)
	}

	txSuccess = true

	return restoperators.NewV2ReportMonitoredOperatorStatusOK()
}

// GetMonitoredOperators retrieves list of monitored operators for a cluster
func (h *Handler) GetMonitoredOperators(ctx context.Context, clusterID strfmt.UUID, operatorName *string, db *gorm.DB) (models.MonitoredOperatorsList, error) {
	log := logutil.FromContext(ctx, h.log)
	if operatorName != nil && *operatorName != "" {
		operator, err := h.FindMonitoredOperator(ctx, clusterID, *operatorName, db)
		if err != nil {
			return nil, err
		}
		return models.MonitoredOperatorsList{operator}, nil
	}

	var operatorsList = models.MonitoredOperatorsList{}
	if err := db.Find(&operatorsList, "cluster_id = ?", clusterID).Error; err != nil {
		log.WithError(err).Errorf("failed to find monitored operators")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewApiError(http.StatusNotFound, err)
		}
	}

	return operatorsList, nil
}

// FindMonitoredOperator retrieves monitored operator identified by given cluster ID and non-empty name
func (h *Handler) FindMonitoredOperator(ctx context.Context, clusterID strfmt.UUID, operatorName string, db *gorm.DB) (*models.MonitoredOperator, error) {
	log := logutil.FromContext(ctx, h.log)
	if operatorName == "" {
		return nil, common.NewApiError(http.StatusBadRequest, errors.New("empty operator name is not allowed"))
	}
	var operator models.MonitoredOperator
	if err := db.First(&operator, "cluster_id = ? and name = ?", clusterID, operatorName).Error; err != nil {
		log.WithError(err).Errorf("failed to find monitored operator")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewApiError(http.StatusNotFound, err)
		}
	}
	return &operator, nil
}

// UpdateMonitoredOperatorStatus updates status and status info of a monitored operator for a cluster
func (h *Handler) UpdateMonitoredOperatorStatus(ctx context.Context, clusterID strfmt.UUID, monitoredOperatorName string,
	status models.OperatorStatus, statusInfo string, db *gorm.DB) error {

	log := logutil.FromContext(ctx, h.log)

	operator, err := h.FindMonitoredOperator(ctx, clusterID, monitoredOperatorName, db)
	if err != nil {
		return err
	}

	operator.Status = status
	operator.StatusInfo = statusInfo
	operator.StatusUpdatedAt = strfmt.DateTime(time.Now())

	if err = db.Save(operator).Error; err != nil {
		err = errors.Wrapf(err, "failed to update operator %s of cluster %s", operator.Name, clusterID)
		log.Error(err)
		return common.NewApiError(http.StatusInternalServerError, err)
	}

	eventgen.SendClusterOperatorStatusEvent(ctx, h.eventsHandler, clusterID, operator.Name, string(status), statusInfo)
	return nil
}