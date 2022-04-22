package forms

import (
	"bookshelf/models"
	"fmt"

	valid "github.com/asaskevich/govalidator"
)

// IForm form interface
type IForm interface {
	ToModel() (models.IModel, error)
}

// ToModel ...
func ToModel(form IForm) (models.IModel, error) {
	isValid, err := valid.ValidateStruct(form)
	if !isValid {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("all the fields are required")
	}

	return form.ToModel()
}
