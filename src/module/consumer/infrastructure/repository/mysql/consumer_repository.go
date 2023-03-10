package mysql

import (
	"context"
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	consumerEvent "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer/event"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/repository/mysql/table"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/repository"
	"gorm.io/gorm"
)

func NewConsumerRepository(db *gorm.DB) repository.ConsumerRepository {
	return &ConsumerRepository{
		db: db,
	}
}

type ConsumerRepository struct {
	db *gorm.DB
}

func (repo ConsumerRepository) Save(ctx context.Context, ag *consumer.ConsumerAggregate) (err error) {
	if ag == nil {
		return
	}
	err = repo.db.Transaction(func(tx *gorm.DB) (txErr error) {
		events := ag.GetEvents()
		for _, aggregateEvent := range events {

			if event, valid := aggregateEvent.(consumerEvent.InsertNewConsumerEvent); valid {
				txErr = repo.InsertNewConsumerEventHandler(ctx, tx, event)
			} else if event, valid := aggregateEvent.(consumerEvent.AddTenorLimitEvent); valid {
				txErr = repo.AddTenorLimitEventHandler(ctx, tx, event)
			} else if event, valid := aggregateEvent.(consumerEvent.UpdateTenorLimitEvent); valid {
				txErr = repo.UpdateTenorLimitEventHandler(ctx, tx, event)
			} else if event, valid := aggregateEvent.(consumerEvent.DeleteTenorLimitEvent); valid {
				txErr = repo.DeleteTenorLimitEventHandler(ctx, tx, event)
			} else if event, valid := aggregateEvent.(consumerEvent.RequestLoanEvent); valid {
				txErr = repo.RequestLoanEventHandler(ctx, tx, event)
			}

			if txErr != nil {
				return
			}
		}
		return
	})
	return
}

func (repo ConsumerRepository) FindTenorLimitByConsumerId(ctx context.Context, consumerId string) (*consumer.ConsumerAggregate, error) {
	tableData := table.ConsumerTable{}
	err := repo.db.Model(tableData).
		Where("id = ? AND deleted_at <= 0", consumerId).
		Preload("TenorLimits", "deleted_at <= 0").
		Find(&tableData).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
		return nil, nil
	}
	ag := consumer.BuildConsumerAggregate(tableData.TransformToEntity())
	return &ag, nil
}

func (repo ConsumerRepository) FindRequestLoanByConsumerId(ctx context.Context, consumerId string) (*consumer.ConsumerAggregate, error) {
	tableData := table.ConsumerTable{}
	err := repo.db.Model(tableData).
		Where("id = ? AND deleted_at <= 0", consumerId).
		Preload("ListRequestLoan", "deleted_at <= 0").
		Find(&tableData).Error

	if err == gorm.ErrRecordNotFound {
		err = nil
		return nil, nil
	}
	ag := consumer.BuildConsumerAggregate(tableData.TransformToEntity())
	return &ag, nil
}

//=============================== EVENT HANDLER ===============================
func (repo ConsumerRepository) InsertNewConsumerEventHandler(ctx context.Context, tx *gorm.DB, event consumerEvent.InsertNewConsumerEvent) (err error) {
	eventData := event.GetData()
	rowData := table.SetupConsumerTable(*eventData)
	err = tx.Omit("TenorLimits,ListRequestLoan").Save(&rowData).Error
	return
}

func (repo ConsumerRepository) AddTenorLimitEventHandler(ctx context.Context, tx *gorm.DB, event consumerEvent.AddTenorLimitEvent) (err error) {
	eventData := event.GetData()
	rowDatas := []table.TenorLimitTable{}
	for _, tl := range eventData {
		rowDatas = append(rowDatas, table.SetupTenorLimitTable(*tl))
	}
	err = tx.Omit("UpdatedAt,DeletedAt").CreateInBatches(&rowDatas, 10).Error
	return
}

func (repo ConsumerRepository) UpdateTenorLimitEventHandler(ctx context.Context, tx *gorm.DB, event consumerEvent.UpdateTenorLimitEvent) (err error) {
	eventData := event.GetData()
	rowDatas := []table.TenorLimitTable{}
	for _, tl := range eventData {
		rowDatas = append(rowDatas, table.SetupTenorLimitTable(*tl))
	}
	err = tx.Omit("CreatedAt,DeletedAt").Save(&rowDatas).Error
	return
}

func (repo ConsumerRepository) DeleteTenorLimitEventHandler(ctx context.Context, tx *gorm.DB, event consumerEvent.DeleteTenorLimitEvent) (err error) {
	eventData := event.GetData()
	rowDatas := []table.TenorLimitTable{}
	deleteTime := time.Now()
	for _, tl := range eventData {
		tl.DeletedAt = &deleteTime
		rowDatas = append(rowDatas, table.SetupTenorLimitTable(*tl))
	}
	err = tx.Omit("CreatedAt,UpdatedAt").Save(&rowDatas).Error
	if err != nil {
		return
	}
	err = tx.Where("deleted_at > 0").Delete(&rowDatas).Error
	return
}

func (repo ConsumerRepository) RequestLoanEventHandler(ctx context.Context, tx *gorm.DB, event consumerEvent.RequestLoanEvent) (err error) {
	eventData := event.GetData()
	rowData := table.SetupRequestLoanTable(*eventData)
	err = tx.Omit("DeletedAt,UpdatedAt").Save(&rowData).Error
	return
}

//==============================================================================
