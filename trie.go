package axisapi

import "strings"

// node using prefix tree to match routes
type node struct {
	pattern  string  //route to be matched(要匹配的目标路由)
	part     string  //part of route（路由的一部分）
	children []*node //children node
	isWild   bool    // if part is contain * or ：that is universal matching（泛匹配）
}

// 获取第一个成功匹配的节点 get the first successful matched node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 获取所有匹配成功的节点 get the all successful matched node
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// 构造动态路由 Constructing dynamic routes
func (n *node) insert(pattern string, parts []string, height int) {
	// 只有匹配到最终点时pattern才会赋值，祖先节点如果不是路由规则则都是""
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	// 深度一样而且同名或者泛匹配的不生成新节点
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		// 证明没有以parts[height]结束的路由规则
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	// 获取前缀树当前一层所有符合条件的孩子节点
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
