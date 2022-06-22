package people

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/caiosousaf/go-gin-api-medium/pkg/common/models"
)

type UpdatePersonRequestBody struct {
	Name    	string `json:"name"`
    Profissao 	string `json:"profissao"`
    Team		models.Team    `gorm:"foreignkey:TeamID:constraint:onUpadate:CASCADE,onDelete:CASCADE" json:"team"`
    Task        models.Task    `gorm:"foreignkey:TaskID:constraint:onUpadate:CASCADE,onDelete:CASCADE" json:"task"`
}

func (h handler) UpdatePerson(c *gin.Context) {
    id := c.Param("id")
    body := UpdatePersonRequestBody{}

    // getting request's body
    if err := c.BindJSON(&body); err != nil {
        c.AbortWithError(http.StatusBadRequest, err)
        return
    }

    var person models.Person

    if result := h.DB.First(&person, id); result.Error != nil {
        c.AbortWithError(http.StatusNotFound, result.Error)
        return
    }

    person.Name = body.Name
    person.Profissao = body.Profissao
    person.Team = body.Team
    person.Task = body.Task

    h.DB.Save(&person)

    c.JSON(http.StatusOK, &person)
}