package service

import (
	"fmt"
	"go-template/app/db"
	"go-template/app/model"
)

type NovelService struct{}

func (NovelService) SetNovel(novel *model.Novel) (string, error) {
	_, err := db.DbEngine.Insert(novel)
	if err != nil {
		return "0", err
	}
	return novel.Id, nil
}

func (NovelService) GetNovelById(id int) *model.Novel {
	var novel model.Novel
	has, err := db.DbEngine.ID(id).Get(&novel)
	if err != nil {
		panic(err)
	}
	if !has {
		fmt.Printf("No novel found for id %d\n", id)
		return nil
	}
	return &novel
}

func (NovelService) GetNovelList() []model.Novel {
	novels := make([]model.Novel, 0)
	err := db.DbEngine.Find(&novels)
	if err != nil {
		panic(err)
	}
	return novels
}

func (NovelService) UpdateNovel(newNovel *model.Novel) error {
	_, err := db.DbEngine.ID(newNovel.Id).Update(newNovel)
	if err != nil {
		return err
	}
	return nil
}

func (NovelService) DeleteNovel(id int64) error {
	novel := new(model.Novel)
	_, err := db.DbEngine.ID(id).Delete(novel)
	if err != nil {
		return err
	}
	return nil
}
