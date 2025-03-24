package constant

const (
	DatabaseSchema   = "public"
	TableOrder       = DatabaseSchema + "." + "order"
	TableOrderDetail = DatabaseSchema + "." + "order_detail"
)

type dbTxKey string

const (
	DbTxTransactionKey dbTxKey = "dbtx-transaction"
)

func (t dbTxKey) String() string {
	return string(t)
}
