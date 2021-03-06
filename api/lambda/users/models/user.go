package models

/**
 * Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

// Group is a struct for Panther Group containing employees.
type Group struct {
	Description *string `json:"description"`
	Name        *string `json:"name"`
}

// User is a struct describing a Panther User.
type User struct {
	CreatedAt   *int64  `json:"createdAt"`
	Email       *string `json:"email"`
	FamilyName  *string `json:"familyName"`
	GivenName   *string `json:"givenName"`
	ID          *string `json:"id"`
	PhoneNumber *string `json:"phoneNumber"`
	Role        *string `json:"role"` // Roles are group name
	Status      *string `json:"status"`
}
