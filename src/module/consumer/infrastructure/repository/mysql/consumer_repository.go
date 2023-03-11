package mysql

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer"
	consumerEvent "github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/aggregate/consumer/event"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/entity"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/infrastructure/repository/mysql/table"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/module/consumer/repository"
	"github.com/muhammadandikakurniawan/xyz_multifinance/src/shared/model"
	"github.com/spf13/cast"
	"golang.org/x/sync/errgroup"
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
			} else if event, valid := aggregateEvent.(consumerEvent.ApproveRequestLoanEvent); valid {
				txErr = repo.ApproveRequestLoanEventHandler(ctx, tx, event)
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

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	if tableData.Id != consumerId {
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
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	ag := consumer.BuildConsumerAggregate(tableData.TransformToEntity())
	return &ag, nil
}

func (repo ConsumerRepository) SearchListRequestLoan(ctx context.Context, paginationReq model.PaginationRequestModel, filter entity.SearchRequestLoanFilterModel) (res []entity.ConsumerEntity, paginationRes model.PaginationResponseModel, err error) {

	eg := errgroup.Group{}

	eg.Go(func() (egErr error) {
		res, egErr = repo.SearchRequestLoan(ctx, paginationReq, filter)
		return
	})

	eg.Go(func() (egErr error) {
		total, egErr := repo.CountSearchRequestLoan(ctx, paginationReq, filter)
		paginationRes.TotalData = total
		paginationRes.TotalPage = uint64(math.Ceil(cast.ToFloat64(total) / float64(paginationReq.TotalData)))
		paginationRes.CurrentPage = paginationReq.Page
		return
	})

	err = eg.Wait()

	return
}

func (repo ConsumerRepository) SearchRequestLoan(ctx context.Context, paginationreq model.PaginationRequestModel, filter entity.SearchRequestLoanFilterModel) (res []entity.ConsumerEntity, err error) {
	queryStrBuilder := strings.Builder{}
	queryStrBuilder.WriteString(`
		SELECT request_loan.*, consumer.legal_name consumer_name, consumer.salary consumer_salary FROM request_loan
		JOIN consumer ON request_loan.consumer_id = consumer.id
		WHERE (request_loan.deleted_at <= 0 AND consumer.deleted_at <= 0)
	`)

	filterStr, filterValue := repo.SetupSearchRequestLoanFiter(paginationreq, filter)
	if filterStr != "" {
		queryStrBuilder.WriteString(" " + filterStr)
	}

	page := paginationreq.Page
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * paginationreq.TotalData
	queryStrBuilder.WriteString(fmt.Sprintf(" LIMIT %d OFFSET %d", paginationreq.TotalData, offset))

	queryStr := queryStrBuilder.String()
	rowDatas := []map[string]interface{}{}
	if len(filterValue) > 0 {
		err = repo.db.Raw(queryStr, filterValue).Find(&rowDatas).Error
	} else {
		err = repo.db.Raw(queryStr).Find(&rowDatas).Error
	}
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	mapConsumerById := map[string]*entity.ConsumerEntity{}

	for i := range rowDatas {
		row := rowDatas[i]

		consumerId := cast.ToString(row["consumer_id"])
		consumerName := cast.ToString(row["consumer_name"])
		salary := cast.ToFloat64(row["consumer_salary"])
		consumerData, exists := mapConsumerById[consumerId]
		if !exists {
			res = append(res, entity.ConsumerEntity{Id: consumerId, Legalname: consumerName, Salary: salary})
			consumerData = &res[len(res)-1]
			mapConsumerById[consumerId] = consumerData
		}

		requestLoanData := entity.RequestLoanEntity{
			Id:             cast.ToInt64(row["id"]),
			ConsumerId:     consumerId,
			ContractNumber: cast.ToString(row["contract_number"]),
			AssetName:      cast.ToString(row["asset_name"]),
			OTR:            cast.ToFloat64(row["otr"]),
			AdminFee:       cast.ToFloat64(row["admin_fee"]),
			Installment:    cast.ToFloat64(row["installment"]),
			Interest:       cast.ToFloat64(row["interest"]),
			CreatedAt:      time.UnixMilli(cast.ToInt64(row["created_at"])),
		}
		if updatedAtUnix := cast.ToInt64(row["updated_at"]); updatedAtUnix > 0 {
			v := time.UnixMilli(updatedAtUnix)
			requestLoanData.UpdatedAt = &v
		}
		if isApproved := row["is_approved"]; isApproved != nil {
			v := cast.ToBool(isApproved)
			requestLoanData.IsApproved = &v
		}
		consumerData.AddRequestLoan(requestLoanData)
	}

	return
}

func (repo ConsumerRepository) CountSearchRequestLoan(ctx context.Context, paginationreq model.PaginationRequestModel, filter entity.SearchRequestLoanFilterModel) (res uint64, err error) {
	queryStrBuilder := strings.Builder{}
	queryStrBuilder.WriteString(`
		SELECT count(1) FROM request_loan
		JOIN consumer ON request_loan.consumer_id = consumer.id
		WHERE (request_loan.deleted_at <= 0 AND consumer.deleted_at <= 0)
	`)

	filterStr, filterValue := repo.SetupSearchRequestLoanFiter(paginationreq, filter)
	queryStrBuilder.WriteString(" " + filterStr)

	queryStr := queryStrBuilder.String()

	if len(filterValue) > 0 {
		err = repo.db.Raw(queryStr, filterValue).Row().Scan(&res)
	} else {
		err = repo.db.Raw(queryStr).Row().Scan(&res)
	}
	return
}

func (repo ConsumerRepository) SetupSearchRequestLoanFiter(paginationreq model.PaginationRequestModel, filter entity.SearchRequestLoanFilterModel) (filterStr string, filterValue map[string]interface{}) {
	filtersBuilder := []string{}
	filterValue = map[string]interface{}{}
	if val := strings.ReplaceAll(filter.ConsumerId, " ", ""); val != "" {
		filtersBuilder = append(filtersBuilder, "request_loan.consumer_id = @consumerId")
		filterValue["consumerId"] = val
	}
	if val := strings.TrimSpace(filter.ConsumerName); val != "" {
		filtersBuilder = append(filtersBuilder, "(consumer.full_name LIKE @consumerName OR consumer.legal_name LIKE @consumerName)")
		filterValue["consumerName"] = "%" + val + "%"
	}
	if val := strings.TrimSpace(filter.AssetName); val != "" {
		filtersBuilder = append(filtersBuilder, "asset_name LIKE @assetName")
		filterValue["assetName"] = "%" + val + "%"
	}
	if val := strings.ReplaceAll(filter.ContractNumber, " ", ""); val != "" {
		filtersBuilder = append(filtersBuilder, "contract_number = @contractNumber")
		filterValue["contractNumber"] = val
	}

	switch filter.ApprovalStatus {
	case entity.RequestLoanApprovalStatus_APPROVED:
		filtersBuilder = append(filtersBuilder, "is_approved = true")
	case entity.RequestLoanApprovalStatus_REJECTED:
		filtersBuilder = append(filtersBuilder, "is_approved = false")
	case entity.RequestLoanApprovalStatus_PENDING:
		filtersBuilder = append(filtersBuilder, "is_approved IS NULL")
	}

	if len(filtersBuilder) > 0 {
		filterStr = fmt.Sprintf("AND (%s)", strings.Join(filtersBuilder, " OR "))
	}
	return
}

func (repo ConsumerRepository) FindRequestLoanById(ctx context.Context, requestLoanId int64) (ag *consumer.ConsumerAggregate, err error) {
	requestLoanRowData := table.RequestLoanTable{}
	err = repo.db.Where("id = ? AND deleted_at <= 0", requestLoanId).First(&requestLoanRowData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	consumerData := table.ConsumerTable{}
	err = repo.db.Where("id = ? AND deleted_at <= 0", requestLoanRowData.ConsumerId).First(&consumerData).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	consumerData.ListRequestLoan = append(consumerData.ListRequestLoan, requestLoanRowData)

	consumerEntity := consumerData.TransformToEntity()
	res := consumer.BuildConsumerAggregate(consumerEntity)
	ag = &res
	return
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

func (repo ConsumerRepository) ApproveRequestLoanEventHandler(ctx context.Context, tx *gorm.DB, event consumerEvent.ApproveRequestLoanEvent) (err error) {
	eventData := event.GetData()
	rowData := table.SetupRequestLoanTable(*eventData)
	updatedAt := time.Now().UnixMilli()
	if eventData.UpdatedAt != nil {
		updatedAt = eventData.UpdatedAt.UnixMilli()
	}
	err = tx.Model(rowData).Where("id = ?", rowData.Id).Updates(map[string]interface{}{"is_approved": rowData.IsApproved, "updated_at": updatedAt}).Error
	return
}

//==============================================================================
