package concurrency_control

// alias
type Timestamp = int64
type DataKey = int
type DataValue = int

// Definition for Transaction status
const (
	INIT = iota
	COMMIT
	ROLLBACK
)

type TransactionContext struct {
	nextTs Timestamp
}

func (context *TransactionContext) getNextTs() Timestamp {
	// TODO:add mutex
	ts := context.nextTs
	context.nextTs++
	return ts
}

func (context *TransactionContext) createTx() *Transaction {
	return &Transaction{context.getNextTs(), INIT}
}

type Transaction struct {
	ts Timestamp
	// INIT,COMMIT,ROLLBACK
	status int
}

func (tx *Transaction) read(manager *MultiVersionManager, key DataKey) (DataValue, error) {
	return 0, nil
}

func (tx *Transaction) write(manager *MultiVersionManager, key DataKey, value DataValue) error {
	return nil
}

type MultiVersionManager struct {
	data map[DataKey]*DataNode
}

func (manager *MultiVersionManager) getDataNode(ts Timestamp, key DataKey) *DataNode {
	if _, exist := manager.data[key]; !exist {
		return nil
	}
	node := manager.data[key]
	for node != nil {
		if node.writeTs <= ts {
			break
		}
		node = node.next
	}
	return node
}

type DataNode struct {
	value   DataValue
	readTs  Timestamp
	writeTs Timestamp
	next    *DataNode
	pre     *DataNode
}
