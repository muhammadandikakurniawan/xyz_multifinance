package consumer

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer/event"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/abstraction"
	sharedError "github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/error"
)

func CreateNewConsumerAggregate(aggregateRoot entity.ConsumerEntity) (ag ConsumerAggregate, err error) {
	ag = ConsumerAggregate{
		aggregateRoot: aggregateRoot,
	}
	err = ag.CreateNew()
	return
}

func BuildConsumerAggregate(aggregateRoot entity.ConsumerEntity) ConsumerAggregate {
	ag := ConsumerAggregate{
		aggregateRoot: aggregateRoot,
	}
	if ag.aggregateRoot.MapTenorLimitByMonth == nil {
		ag.aggregateRoot.MapTenorLimitByMonth = map[int]*entity.TenorLimitEntity{}
	}
	return ag
}

type ConsumerAggregate struct {
	abstraction.BaseAggregate
	aggregateRoot entity.ConsumerEntity
}

func (ag ConsumerAggregate) GetAggregateRoot() entity.ConsumerEntity {
	return ag.aggregateRoot
}

func (ag *ConsumerAggregate) CreateNew() (err error) {
	if err = ag.aggregateRoot.ValidateNewConsumer(); err != nil {
		return
	}
	ag.aggregateRoot.Id = uuid.NewString()
	ag.aggregateRoot.CreatedAt = time.Now()
	ag.AddEvent(event.NewInsertNewConsumerEvent(&ag.aggregateRoot))
	return
}

func (ag *ConsumerAggregate) AddTenorLimit(tenorlimits ...*entity.TenorLimitEntity) (err error) {

	if len(tenorlimits) <= 0 {
		return
	}

	newLimit := []*entity.TenorLimitEntity{}
	updatedLimit := []*entity.TenorLimitEntity{}
	deletedLimit := []*entity.TenorLimitEntity{}

	monthSet := map[int]bool{}

	for _, tl := range tenorlimits {
		_, alredyProcesses := monthSet[tl.Month]
		if alredyProcesses {
			continue
		}
		monthSet[tl.Month] = true
		existingLimit, alreadyExists := ag.aggregateRoot.MapTenorLimitByMonth[tl.Month]
		if alreadyExists {

			isNotUpdated := existingLimit.LimitValue == tl.LimitValue
			if isNotUpdated {
				continue
			}

			if err = ag.UpdateTenorLimitProcess(tl); err != nil {
				return
			}

			limitIsEmpty := tl.LimitValue <= 0
			if limitIsEmpty {
				deletedLimit = append(deletedLimit, tl)
				continue
			}

			updatedLimit = append(updatedLimit, tl)
			continue
		}
		if err = ag.addTenorLimitProcess(tl); err != nil {
			return
		}
		limitIsEmpty := tl.LimitValue <= 0
		if limitIsEmpty {
			continue
		}
		newLimit = append(newLimit, tl)
	}

	if len(newLimit) > 0 {
		ag.AddEvent(event.NewAddTenorLimitEvent(newLimit))
	}
	if len(updatedLimit) > 0 {
		ag.AddEvent(event.NewUpdateTenorLimitEvent(updatedLimit))
	}
	if len(deletedLimit) > 0 {
		ag.AddEvent(event.NewDeleteTenorLimitEvent(deletedLimit))
	}

	return
}

func (ag *ConsumerAggregate) ValidateTenorLimit(tenorlimit *entity.TenorLimitEntity) (err error) {
	if tenorlimit == nil {
		err = sharedError.NewValidationError("limit cannot be null")
		return
	}

	if tenorlimit.Month <= 0 {
		err = sharedError.NewValidationError("invalid month")
		return
	}
	if tenorlimit.LimitValue < 0 {
		err = sharedError.NewValidationError("invalid limit")
		return
	}
	return
}

func (ag *ConsumerAggregate) addTenorLimitProcess(tenorlimit *entity.TenorLimitEntity) (err error) {

	if err = ag.ValidateTenorLimit(tenorlimit); err != nil {
		return
	}

	tenorlimit.ConsumerId = ag.aggregateRoot.Id
	tenorlimit.CreatedAt = time.Now()
	ag.aggregateRoot.TenorLimits = append(ag.aggregateRoot.TenorLimits, *tenorlimit)
	return
}

func (ag *ConsumerAggregate) UpdateTenorLimitProcess(tenorlimit *entity.TenorLimitEntity) (err error) {

	if err = ag.ValidateTenorLimit(tenorlimit); err != nil {
		return
	}

	tenorlimit.ConsumerId = ag.aggregateRoot.Id
	updatedAt := time.Now()
	tenorlimit.UpdatedAt = &updatedAt
	currentData := ag.aggregateRoot.MapTenorLimitByMonth[tenorlimit.Month]
	currentData.LimitValue = tenorlimit.LimitValue
	ag.aggregateRoot.MapTenorLimitByMonth[tenorlimit.Month] = currentData
	return
}

func (ag *ConsumerAggregate) ValidateNewRequestLoan(req *entity.RequestLoanEntity) (err error) {
	if req == nil {
		err = sharedError.NewValidationError("request cannot be null")
		return
	}

	errorValidationMessages := []string{}

	tenorLimitNotfound := len(ag.aggregateRoot.TenorLimits) <= 0
	if tenorLimitNotfound {
		err = sharedError.NewValidationError("tenor limit not found")
		return
	}

	req.AssetName = strings.ReplaceAll(req.AssetName, " ", "")
	if req.AssetName == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid asset name")
	}

	req.ConsumerId = strings.ReplaceAll(req.ConsumerId, " ", "")
	if req.ConsumerId == "" {
		errorValidationMessages = append(errorValidationMessages, "invalid consumer")
	}

	if len(errorValidationMessages) > 0 {
		err = sharedError.NewValidationError(strings.Join(errorValidationMessages, ", "))
	}

	return
}

func (ag *ConsumerAggregate) AddRequestLoan(req *entity.RequestLoanEntity) (err error) {

	if err = ag.ValidateNewRequestLoan(req); err != nil {
		return
	}

	req.ContractNumber = fmt.Sprintf("REQ-LOAN-%s", uuid.NewString())
	req.CreatedAt = time.Now()
	ag.AddEvent(event.NewRequestLoanEvent(req))
	return
}
