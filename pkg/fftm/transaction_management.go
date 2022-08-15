// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fftm

import (
	"context"
	"strings"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/i18n"
	"github.com/hyperledger/firefly-transaction-manager/internal/persistence"
	"github.com/hyperledger/firefly-transaction-manager/internal/tmmsgs"
	"github.com/hyperledger/firefly-transaction-manager/pkg/apitypes"
)

func (m *manager) getTransactionByID(ctx context.Context, txID string) (transaction *apitypes.ManagedTX, err error) {
	return m.persistence.GetTransactionByID(ctx, txID)
}

func (m *manager) getTransactions(ctx context.Context, afterStr, limitStr, signer string, pending bool, dirString string) (transactions []*apitypes.ManagedTX, err error) {
	limit, err := m.parseLimit(ctx, limitStr)
	if err != nil {
		return nil, err
	}
	var dir persistence.SortDirection
	switch strings.ToLower(dirString) {
	case "", "desc", "descending":
		dir = persistence.SortDirectionDescending // descending is default
	case "asc", "ascending":
		dir = persistence.SortDirectionAscending
	default:
		return nil, i18n.NewError(ctx, tmmsgs.MsgInvalidSortDirection, dirString)
	}
	var afterTx *apitypes.ManagedTX
	if afterStr != "" {
		// Get the transaction, as we need this to exist to pick the right field depending on the index that's been chosen
		afterTx, err = m.persistence.GetTransactionByID(ctx, afterStr)
		if err != nil {
			return nil, err
		}
		if afterTx == nil {
			return nil, i18n.NewError(ctx, tmmsgs.MsgPaginationErrTxNotFound, afterStr)
		}
	}
	switch {
	case signer != "" && pending:
		return nil, i18n.NewError(ctx, tmmsgs.MsgTXConflictSignerPending)
	case signer != "":
		var afterNonce *fftypes.FFBigInt
		if afterTx != nil {
			afterNonce = afterTx.Nonce
		}
		return m.persistence.ListTransactionsByNonce(ctx, signer, afterNonce, limit, dir)
	case pending:
		var afterSequence *fftypes.UUID
		if afterTx != nil {
			afterSequence = afterTx.SequenceID
		}
		return m.persistence.ListTransactionsPending(ctx, afterSequence, limit, dir)
	default:
		return m.persistence.ListTransactionsByCreateTime(ctx, afterTx, limit, dir)
	}

}