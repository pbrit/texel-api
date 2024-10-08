/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package project

type Project struct {
	ID string `uri:"project_id" binding:"required,uuid"`
}

type GenericApiError struct {
}

func (e *GenericApiError) Error() string {
	return "Something went wrong"
}
