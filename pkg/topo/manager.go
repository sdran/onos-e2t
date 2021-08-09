// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"time"

	"github.com/onosproject/onos-lib-go/pkg/uri"

	"github.com/google/uuid"

	"github.com/cenkalti/backoff/v4"
	gogotypes "github.com/gogo/protobuf/types"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-e2t/pkg/store/rnib"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"golang.org/x/net/context"
)

var log = logging.GetLogger("topo", "manager")

type Rnib struct {
	store rnib.Store
}

func (r *Rnib) DeleteE2Relation(ctx context.Context, relationID topoapi.ID) error {
	return r.store.Delete(ctx, relationID)
}

func (r *Rnib) GetE2Relation(ctx context.Context, e2NodeID topoapi.ID) (topoapi.ID, error) {
	objects, err := r.store.List(ctx, &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.CONTROLS,
				},
			},
		},
	})
	if err != nil {
		return "", err
	}

	podID := getPodID()
	for _, object := range objects {
		val := object.Obj.(*topoapi.Object_Relation)
		srcEntity := val.Relation.GetSrcEntityID()
		dstEntity := val.Relation.GetTgtEntityID()
		if srcEntity == topoapi.ID(podID) && dstEntity == e2NodeID {
			return object.ID, nil
		}
	}

	return "", errors.NewNotFound("E2 relation ID is not found")
}

func (r *Rnib) CreateOrUpdateE2Relation(ctx context.Context, e2NodeID topoapi.ID, relationID topoapi.ID) error {
	return backoff.Retry(func() error {
		_, err := r.store.Get(ctx, e2NodeID)
		if err != nil {
			return backoff.Permanent(err)
		}

		currentRelationObject, err := r.store.Get(ctx, relationID)
		if err != nil {
			if !errors.IsNotFound(err) {
				return backoff.Permanent(err)
			}
			e2Relation := &topoapi.Object{
				ID:   relationID,
				Type: topoapi.Object_RELATION,
				Obj: &topoapi.Object_Relation{
					Relation: &topoapi.Relation{
						KindID:      topoapi.ID(topoapi.CONTROLS),
						SrcEntityID: GetE2TID(),
						TgtEntityID: e2NodeID,
					},
				},
			}
			err = r.store.Create(ctx, e2Relation)
			if err != nil {
				if !errors.IsAlreadyExists(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		} else {
			currentRelationObject.Obj.(*topoapi.Object_Relation).Relation.SrcEntityID = topoapi.ID(getPodID())
			currentRelationObject.Obj.(*topoapi.Object_Relation).Relation.TgtEntityID = e2NodeID
			currentRelationObject.Obj.(*topoapi.Object_Relation).Relation.KindID = topoapi.ID(topoapi.RANRelationKinds_CONTROLS.String())
			err = r.store.Update(ctx, currentRelationObject)
			if err != nil {
				if !errors.IsNotFound(err) && !errors.IsConflict(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		}
		return nil
	}, newExpBackoff())
}

func (r *Rnib) CreateOrUpdateE2CellRelation(ctx context.Context, e2NodeID topoapi.ID, cellID topoapi.ID) error {
	return backoff.Retry(func() error {
		cellRelationID := topoapi.ID(uri.NewURI(
			uri.WithScheme("uuid"),
			uri.WithOpaque(uuid.New().String())).String())
		currentCellRelation, err := r.store.Get(ctx, cellRelationID)
		if err != nil {
			if !errors.IsNotFound(err) {
				return backoff.Permanent(err)
			}
			cellRelation := &topoapi.Object{
				ID:   cellRelationID,
				Type: topoapi.Object_RELATION,
				Obj: &topoapi.Object_Relation{
					Relation: &topoapi.Relation{
						KindID:      topoapi.ID(topoapi.CONTAINS),
						SrcEntityID: e2NodeID,
						TgtEntityID: cellID,
					},
				},
			}
			err = r.store.Create(ctx, cellRelation)
			if err != nil {
				if !errors.IsAlreadyExists(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		} else {
			err := r.store.Update(ctx, currentCellRelation)
			if err != nil {
				if !errors.IsNotFound(err) && !errors.IsConflict(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		}
		return nil
	}, newExpBackoff())
}

// CreateOrUpdateE2Cells creates or update E2 cells entities and relations
func (r *Rnib) CreateOrUpdateE2Cells(ctx context.Context, e2NodeID topoapi.ID, e2Cells []*topoapi.E2Cell) error {
	return backoff.Retry(func() error {
		_, err := r.store.Get(ctx, e2NodeID)
		if err != nil {
			return backoff.Permanent(err)
		}
		for _, e2Cell := range e2Cells {
			cellID := GetCellID(e2NodeID, e2Cell.GetCellGlobalID().GetValue())
			currentCellObject, err := r.store.Get(ctx, cellID)
			if err != nil {
				if !errors.IsNotFound(err) {
					return backoff.Permanent(err)
				}
				cellObject := &topoapi.Object{
					ID:   cellID,
					Type: topoapi.Object_ENTITY,
					Obj: &topoapi.Object_Entity{
						Entity: &topoapi.Entity{
							KindID: topoapi.ID(topoapi.E2CELL),
						},
					},
					Aspects: make(map[string]*gogotypes.Any),
					Labels:  map[string]string{},
				}

				err := cellObject.SetAspect(e2Cell)
				if err != nil {
					log.Warn(err)
					return backoff.Permanent(err)
				}
				err = r.store.Create(ctx, cellObject)
				if err != nil {
					if !errors.IsAlreadyExists(err) {
						return backoff.Permanent(err)
					}
					return err
				}

				err = r.CreateOrUpdateE2CellRelation(ctx, e2NodeID, cellID)
				if err != nil {
					return backoff.Permanent(err)
				}
			} else {
				err := currentCellObject.SetAspect(e2Cell)
				if err != nil {
					return backoff.Permanent(err)
				}

				err = r.store.Update(ctx, currentCellObject)
				if err != nil {
					if !errors.IsNotFound(err) && !errors.IsConflict(err) {
						return backoff.Permanent(err)
					}
					return err
				}

				err = r.CreateOrUpdateE2CellRelation(ctx, e2NodeID, cellID)
				if err != nil {
					return backoff.Permanent(err)
				}
			}
		}
		return nil
	}, newExpBackoff())
}

func (r *Rnib) CreateOrUpdateE2T(ctx context.Context) error {
	return backoff.Retry(func() error {
		e2tID := GetE2TID()
		currentE2TObject, err := r.store.Get(ctx, e2tID)
		if err != nil {
			if !errors.IsNotFound(err) {
				return backoff.Permanent(err)
			}
			e2tObject := &topoapi.Object{
				ID:   e2tID,
				Type: topoapi.Object_ENTITY,
				Obj: &topoapi.Object_Entity{
					Entity: &topoapi.Entity{
						KindID: topoapi.ID(topoapi.E2T),
					},
				},
				Aspects: make(map[string]*gogotypes.Any),
				Labels:  map[string]string{},
			}
			err = r.store.Create(ctx, e2tObject)
			if err != nil {
				if !errors.IsAlreadyExists(err) {
					return backoff.Permanent(err)
				}
				return err
			}

		} else {
			err = r.store.Update(ctx, currentE2TObject)
			if err != nil {
				if !errors.IsNotFound(err) && !errors.IsConflict(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		}
		return nil
	}, newExpBackoff())
}

// CreateOrUpdateE2Node creates or updates E2 entities
func (r *Rnib) CreateOrUpdateE2Node(ctx context.Context, deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error {
	return backoff.Retry(func() error {
		e2NodeAspects := &topoapi.E2Node{
			ServiceModels: serviceModels,
		}
		currentE2NodeObject, err := r.store.Get(ctx, deviceID)
		if err != nil {
			if !errors.IsNotFound(err) {
				return backoff.Permanent(err)
			}
			e2NodeObject := &topoapi.Object{
				ID:   deviceID,
				Type: topoapi.Object_ENTITY,
				Obj: &topoapi.Object_Entity{
					Entity: &topoapi.Entity{
						KindID: topoapi.ID(topoapi.E2NODE),
					},
				},
				Aspects: make(map[string]*gogotypes.Any),
				Labels:  map[string]string{},
			}

			err = e2NodeObject.SetAspect(e2NodeAspects)
			if err != nil {
				return backoff.Permanent(err)
			}
			err = r.store.Create(ctx, e2NodeObject)
			if err != nil {
				if !errors.IsAlreadyExists(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		} else {
			err := currentE2NodeObject.SetAspect(e2NodeAspects)
			if err != nil {
				return backoff.Permanent(err)
			}

			err = r.store.Update(ctx, currentE2NodeObject)
			if err != nil {
				if !errors.IsNotFound(err) && !errors.IsConflict(err) {
					return backoff.Permanent(err)
				}
				return err
			}
		}
		return nil
	}, newExpBackoff())
}

func (r *Rnib) WatchE2Relations(ctx context.Context, ch chan<- topoapi.ID) error {
	eventCh := make(chan topoapi.Event)
	go func() {
		podID := getPodID()
		defer close(ch)
		for event := range eventCh {
			switch o := event.Object.Obj.(type) {
			case *topoapi.Object_Relation:
				if o.Relation.SrcEntityID == topoapi.ID(podID) {
					ch <- o.Relation.TgtEntityID
				}
			}
		}
	}()
	return r.store.Watch(ctx, eventCh, &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.CONTROLS,
				},
			},
		},
	})
}

// Manager topology manager
type Manager interface {
	CreateOrUpdateE2Cells(ctx context.Context, deviceID topoapi.ID, e2Cells []*topoapi.E2Cell) error
	CreateOrUpdateE2CellRelation(ctx context.Context, deviceID topoapi.ID, cellID topoapi.ID) error
	CreateOrUpdateE2T(ctx context.Context) error
	CreateOrUpdateE2Node(ctx context.Context, deviceID topoapi.ID, serviceModels map[string]*topoapi.ServiceModelInfo) error
	CreateOrUpdateE2Relation(ctx context.Context, deviceID topoapi.ID, relationID topoapi.ID) error
	DeleteE2Relation(ctx context.Context, relationID topoapi.ID) error
	GetE2Relation(ctx context.Context, deviceID topoapi.ID) (topoapi.ID, error)
	WatchE2Relations(ctx context.Context, ch chan<- topoapi.ID) error
}

// NewManager creates topology manager
func NewManager(store rnib.Store) *Rnib {
	return &Rnib{
		store: store,
	}
}

var _ Manager = &Rnib{}

const (
	backoffInterval = 10 * time.Millisecond
	maxBackoffTime  = 5 * time.Second
)

func newExpBackoff() *backoff.ExponentialBackOff {
	b := backoff.NewExponentialBackOff()
	b.InitialInterval = backoffInterval
	// MaxInterval caps the RetryInterval
	b.MaxInterval = maxBackoffTime
	// Never stops retrying
	b.MaxElapsedTime = 0
	return b
}
