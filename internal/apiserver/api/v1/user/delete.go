// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"github.com/marmotedu/component-base/pkg/core"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/apiserver/store"
	"github.com/marmotedu/iam/internal/pkg/code"
	"github.com/marmotedu/log"
)

// Delete delete an user by the user identifier.
// Only administrator can call this function.
func Delete(c *gin.Context) {
	log.Info("delete user function called.", log.String("X-Request-Id", requestid.Get(c)))

	if err := store.Client().Users().Delete(c.Param("name"), metav1.DeleteOptions{Unscoped: true}); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrDatabase, err.Error()), nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}