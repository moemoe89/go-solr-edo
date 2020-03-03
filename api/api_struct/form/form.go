//
//  Practicing Solr
//
//  Copyright © 2020. All rights reserved.
//

package form

type UserForm struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Validate represent the validation method from UserForm
func (v *UserForm) Validate() []string {
	errs := []string{}
	if v.ID < 1 {
		errs = append(errs, "ID can't be empty")
	}

	if len(v.Name) < 1 {
		errs = append(errs, "Name can't be empty")
	}

	if len(v.Address) < 1 {
		errs = append(errs, "Address can't be empty")
	}

	return errs
}
