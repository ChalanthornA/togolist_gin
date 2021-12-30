package todoservice

import (
	"fmt"
	"togolist_gin/config"
	"togolist_gin/models"
)

func InputTodo(newTodo *models.TodoModel){
	db := config.Db
	db.Create(newTodo)
}

func GetUserTodo(id int)([]models.TodoModel, error){
	db := config.Db
	userTodo := []models.TodoModel{}
	if err := db.Where("user_refer = ?", id).Find(&userTodo).Error; err != nil{
		return userTodo, err
	}
	return userTodo, nil
}

func DeleteUserTodo(user_id, todo_id int)(string, error){
	db := config.Db
	targetTodo := models.TodoModel{}
	if err := db.First(&targetTodo, todo_id).Error; err != nil {
		return "Cannot find Todo by id.", err
	}
	if targetTodo.UserRefer != user_id{
		return "User id invalid.", fmt.Errorf("user id invalid")
	}
	if err := db.Delete(&models.TodoModel{}, todo_id).Error; err != nil{
		return "There is an error when try to delete todo.", err
	}
	return "Delete successful", nil
}

func UpdateUserTodo(updateTodo models.TodoModel, user_id int)(string, error){
	db := config.Db
	targetTodo := models.TodoModel{}
	if err := db.First(&targetTodo, updateTodo.ID).Error; err != nil {
		return "Cannot find Todo by id.", err
	}
	if targetTodo.UserRefer != user_id{
		return "User id invalid.", fmt.Errorf("user id invalid")
	}
	if updateTodo.Todo != "" && updateTodo.Description != "" {
		if err := db.Model(&models.TodoModel{}).Where("id = ?", updateTodo.ID).Updates(models.TodoModel{Todo: updateTodo.Todo, Description: updateTodo.Description}).Error; err != nil{
			return "There is an error when try to update todo.", err
		}
	}else if updateTodo.Todo != ""{
		if err := db.Model(&models.TodoModel{}).Where("id = ?", updateTodo.ID).Updates(models.TodoModel{Todo: updateTodo.Todo}).Error; err != nil{
			return "There is an error when try to update todo.", err
		}
	}else if updateTodo.Description != ""{
		if err := db.Model(&models.TodoModel{}).Where("id = ?", updateTodo.ID).Updates(models.TodoModel{Description: updateTodo.Description}).Error; err != nil{
			return "There is an error when try to update todo.", err
		}
	}else if updateTodo.Todo == "" && updateTodo.Description == ""{
		return "There are no Todo and Description to update.", fmt.Errorf("no todo and description")
	}
	return "Update successful.", nil
}