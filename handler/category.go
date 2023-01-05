package handler

import (
	"context"
	"go_category/domain/db"
	"go_category/domain/model"
	"go_category/domain/repository"
	"go_category/pkg"
	pb "go_category/proto/pb/category"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// Category TODO
type Category struct{}

// CreateCategory TODO
func (s *Category) CreateCategory(ctx context.Context,
	in *pb.CreateCategoryReq, out *pb.CreateCategoryRsp) (err error) {
	c := repository.NewCategoryRepository(db.PgSqlDB)
	category := &model.Category{
		ID:          pkg.GetSnowFlake(),
		Name:        in.Name,
		Level:       in.Level,
		Parent:      cast.ToInt64(in.Parent),
		Image:       in.Image,
		Description: in.Description,
	}
	id, err := c.CreateCategory(category)
	if err != nil {
		log.Error(err)
		return err
	}
	out.Id = cast.ToString(id)
	return nil
}

// UpdateCategory TODO
func (s *Category) UpdateCategory(ctx context.Context,
	in *pb.UpdateCategoryReq, out *pb.UpdateCategoryRsp) (err error) {
	c := repository.NewCategoryRepository(db.PgSqlDB)
	category := &model.Category{
		ID:          cast.ToInt64(in.Id),
		Name:        in.Name,
		Level:       in.Level,
		Parent:      cast.ToInt64(in.Parent),
		Image:       in.Image,
		Description: in.Description,
	}
	err = c.UpdateCategory(category)
	if err != nil {
		return errors.Wrap(err, "update category fail")
	}
	return nil
}

// DeleteCategory TODO
func (s *Category) DeleteCategory(ctx context.Context,
	in *pb.DeleteCategoryReq, out *pb.DeleteCategoryRsp) (err error) {
	c := repository.NewCategoryRepository(db.PgSqlDB)
	err = c.DeleteCategoryById(cast.ToInt64(in.Id))
	if err != nil {
		return errors.Wrap(err, "delete category fail")
	}
	return nil
}

// FindAllCategory TODO
func (s *Category) FindAllCategory(ctx context.Context,
	in *pb.FindAllCategoryReq, out *pb.FindAllCategoryRsp) (err error) {
	c := repository.NewCategoryRepository(db.PgSqlDB)
	datas, err := c.FindAll()
	if err != nil {
		return errors.Wrap(err, "find category fail")
	}
	out.Total = int64(len(datas))
	for _, v := range datas {
		out.Datas = append(out.Datas, &pb.CategoryInfo{
			Id:          cast.ToString(v.ID),
			Name:        v.Name,
			Level:       v.Level,
			Parent:      cast.ToString(v.Parent),
			Image:       v.Image,
			Description: v.Description,
		})
	}
	return nil
}

// FindCategoryByName TODO
func (s *Category) FindCategoryByName(ctx context.Context,
	in *pb.FindCategoryByNameReq, out *pb.FindCategoryByNameRsp) (err error) {
	c := repository.NewCategoryRepository(db.PgSqlDB)
	data, err := c.FindCategoryByName(in.Name)
	if err != nil {
		return errors.Wrap(err, "find category by name fail")
	}
	out.Data = &pb.CategoryInfo{
		Id:          cast.ToString(data.ID),
		Name:        data.Name,
		Level:       data.Level,
		Parent:      cast.ToString(data.Parent),
		Image:       data.Image,
		Description: data.Description,
	}
	return nil
}
