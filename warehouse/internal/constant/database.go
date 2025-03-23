package constant

const (
	DatabaseSchema = "public"
	TableWarehouse = DatabaseSchema + "." + "warehouse"
	TableStock     = DatabaseSchema + "." + "stock"
)

type dbTxKey string

const (
	DbTxTransactionKey dbTxKey = "dbtx-transaction"
)

func (t dbTxKey) String() string {
	return string(t)
}
