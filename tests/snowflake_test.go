package tests

import (
	"fmt"
	"goloveu/utils/snowflake"
	"testing"
)

func TestNode_Generate(t *testing.T) {
	type args struct {
		node int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test node from 1",
			args: args{node: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 生成节点实例
			node, err := snowflake.NewNode(tt.args.node)
			if err != nil {
				t.Errorf("snowflake.NewNode failed : %v", err.Error())
				fmt.Println(err)
				return
			}

			ch := make(chan snowflake.ID)
			count := 10000
			// 并发count个goroutine进行snowflake ID生成
			for i := 0; i < count; i++ {
				go func() {
					id := node.Generate()
					ch <- id
				}()
			}

			defer close(ch)

			m := make(map[snowflake.ID]int)
			for i := 0; i < count; i++ {
				id := <-ch
				// 如果map中存在为id的KEY，说明生成的snowflake ID有重复
				_, ok := m[id]
				if ok {
					t.Errorf("ID is not unique!\n")
					return
				}
				// 将 id 作为 key 存入 map
				m[id] = i
			}

			// 成功生成 snowflake ID
			fmt.Printf("All IDs generated, Count = %v\n", len(m))
			for k,v := range m {
				fmt.Printf("[%v] = %v\n",v,k)
			}
		})
	}
}
