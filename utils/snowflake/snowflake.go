package snowflake

import (
	"errors"
	"goloveu/utils"
	"sync"
)

const (
	NODE_BITS  uint8 = 10                     // 节点ID的位数
	STEP_BITS  uint8 = 12                     // 序列号的位数
	NODE_MAX   int64 = -1 ^ (-1 << NODE_BITS) // 节点ID的最大值，用于检测溢出
	STEP_MAX   int64 = -1 ^ (-1 << STEP_BITS) //序列号的最大值，用于检测溢出
	TIME_SHIFT uint8 = NODE_BITS + STEP_BITS  // 时间戳向左的偏移量
	NODE_SHITF uint8 = STEP_BITS              //节点ID向左的偏移量
)

var epoch int64 = 1288834974657 // timestamp 2006-03-21 20:50:14 GMT

// ID represents
type ID int64

// Node 存储一个节点(机器)上的基础数据
type Node struct {
	mu        sync.Mutex //添加互斥锁，保证并发安全
	timestamp int64      // 时间戳
	node      int64      // 节点ID部分
	step      int64      // 序列号ID部分
}

// NewNode 生成一个Node实例
func NewNode(node int64) (*Node, error) {
	// 如果超出节点的最大范围，产生一个error
	if node < 0 || node > NODE_MAX {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	// 生成并返回节点实列的指针
	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

// Generate 产生一个雪花ID
func (n *Node) Generate() ID {
	n.mu.Lock()         // 保证并发安全，加锁
	defer n.mu.Unlock() // 方法运行完毕后解锁

	// 获取当前时间的时间戳，毫秒
	nowMillisecondstemp := utils.NowTimestamp()

	if n.timestamp == nowMillisecondstemp {
		// step步进+1
		n.step++

		// 当前step用完
		if n.step > STEP_MAX {
			for nowMillisecondstemp <= n.timestamp {
				nowMillisecondstemp = utils.NowTimestamp()
			}
		}
	} else {
		// 本毫秒内step用完
		n.step = 0
	}

	n.timestamp = nowMillisecondstemp
	// 位移运算，生成最终ID
	result := ID((nowMillisecondstemp-epoch)<<TIME_SHIFT | (n.node << NODE_SHITF) | (n.step))
	return result
}
