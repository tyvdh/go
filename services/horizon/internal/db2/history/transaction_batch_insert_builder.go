package history

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/guregu/null"
	"github.com/stellar/go/exp/ingest/io"
	"github.com/stellar/go/services/horizon/internal/db2/sqx"
	"github.com/stellar/go/services/horizon/internal/toid"
	"github.com/stellar/go/services/horizon/internal/utf8"
	"github.com/stellar/go/support/db"
	"github.com/stellar/go/xdr"
)

// TransactionBatchInsertBuilder is used to insert transactions into the
// history_transactions table
type TransactionBatchInsertBuilder interface {
	Add(transaction io.LedgerTransaction, sequence uint32) error
	Exec() error
}

// transactionBatchInsertBuilder is a simple wrapper around db.BatchInsertBuilder
type transactionBatchInsertBuilder struct {
	builder db.BatchInsertBuilder
}

// NewTransactionBatchInsertBuilder constructs a new TransactionBatchInsertBuilder instance
func (q *Q) NewTransactionBatchInsertBuilder(maxBatchSize int) TransactionBatchInsertBuilder {
	return &transactionBatchInsertBuilder{
		builder: db.BatchInsertBuilder{
			Table:        q.GetTable("history_transactions"),
			MaxBatchSize: maxBatchSize,
		},
	}
}

// Add adds a new transaction to the batch
func (i *transactionBatchInsertBuilder) Add(transaction io.LedgerTransaction, sequence uint32) error {
	m, err := transactionToMap(transaction, sequence)
	if err != nil {
		return err
	}

	return i.builder.Row(m)
}

func (i *transactionBatchInsertBuilder) Exec() error {
	return i.builder.Exec()
}

func formatTimeBounds(transaction io.LedgerTransaction) interface{} {
	if transaction.Envelope.Tx.TimeBounds == nil {
		return nil
	}

	if transaction.Envelope.Tx.TimeBounds.MaxTime == 0 {
		return sq.Expr("int8range(?,?)", transaction.Envelope.Tx.TimeBounds.MinTime, nil)
	}

	maxTime := transaction.Envelope.Tx.TimeBounds.MaxTime
	if maxTime > math.MaxInt64 {
		maxTime = math.MaxInt64
	}

	return sq.Expr("int8range(?,?)", transaction.Envelope.Tx.TimeBounds.MinTime, maxTime)
}

func signatures(transaction io.LedgerTransaction) []string {
	signatures := make([]string, len(transaction.Envelope.Signatures))
	for i, sig := range transaction.Envelope.Signatures {
		signatures[i] = base64.StdEncoding.EncodeToString(sig.Signature)
	}
	return signatures
}

func memoType(transaction io.LedgerTransaction) string {
	switch transaction.Envelope.Tx.Memo.Type {
	case xdr.MemoTypeMemoNone:
		return "none"
	case xdr.MemoTypeMemoText:
		return "text"
	case xdr.MemoTypeMemoId:
		return "id"
	case xdr.MemoTypeMemoHash:
		return "hash"
	case xdr.MemoTypeMemoReturn:
		return "return"
	default:
		panic(fmt.Errorf("invalid memo type: %v", transaction.Envelope.Tx.Memo.Type))
	}
}

func memo(transaction io.LedgerTransaction) null.String {
	var (
		value string
		valid bool
	)
	switch transaction.Envelope.Tx.Memo.Type {
	case xdr.MemoTypeMemoNone:
		value, valid = "", false
	case xdr.MemoTypeMemoText:
		scrubbed := utf8.Scrub(transaction.Envelope.Tx.Memo.MustText())
		notnull := strings.Join(strings.Split(scrubbed, "\x00"), "")
		value, valid = notnull, true
	case xdr.MemoTypeMemoId:
		value, valid = fmt.Sprintf("%d", transaction.Envelope.Tx.Memo.MustId()), true
	case xdr.MemoTypeMemoHash:
		hash := transaction.Envelope.Tx.Memo.MustHash()
		value, valid =
			base64.StdEncoding.EncodeToString(hash[:]),
			true
	case xdr.MemoTypeMemoReturn:
		hash := transaction.Envelope.Tx.Memo.MustRetHash()
		value, valid =
			base64.StdEncoding.EncodeToString(hash[:]),
			true
	default:
		panic(fmt.Errorf("invalid memo type: %v", transaction.Envelope.Tx.Memo.Type))
	}

	return null.NewString(value, valid)
}

func transactionToMap(transaction io.LedgerTransaction, sequence uint32) (map[string]interface{}, error) {
	envelopeBase64, err := xdr.MarshalBase64(transaction.Envelope)
	if err != nil {
		return nil, err
	}
	resultBase64, err := xdr.MarshalBase64(transaction.Result.Result)
	if err != nil {
		return nil, err
	}
	metaBase64, err := xdr.MarshalBase64(transaction.Meta)
	if err != nil {
		return nil, err
	}
	feeMetaBase64, err := xdr.MarshalBase64(transaction.FeeChanges)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":                toid.New(int32(sequence), int32(transaction.Index), 0).ToInt64(),
		"transaction_hash":  hex.EncodeToString(transaction.Result.TransactionHash[:]),
		"ledger_sequence":   sequence,
		"application_order": int32(transaction.Index),
		"account":           transaction.Envelope.Tx.SourceAccount.Address(),
		"account_sequence":  strconv.FormatInt(int64(transaction.Envelope.Tx.SeqNum), 10),
		"max_fee":           int32(transaction.Envelope.Tx.Fee),
		"fee_charged":       int32(transaction.Result.Result.FeeCharged),
		"operation_count":   int32(len(transaction.Envelope.Tx.Operations)),
		"tx_envelope":       envelopeBase64,
		"tx_result":         resultBase64,
		"tx_meta":           metaBase64,
		"tx_fee_meta":       feeMetaBase64,
		"signatures":        sqx.StringArray(signatures(transaction)),
		"time_bounds":       formatTimeBounds(transaction),
		"memo_type":         memoType(transaction),
		"memo":              memo(transaction),
		"created_at":        time.Now().UTC(),
		"updated_at":        time.Now().UTC(),
		"successful":        transaction.Successful(),
	}, nil
}
